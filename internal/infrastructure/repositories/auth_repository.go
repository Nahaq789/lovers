package repositories

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"lovers/internal/domain/models/aggregates/auth"
	sharedAws "lovers/internal/shared/infrastructure/aws"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type AuthRepositoryImpl struct {
	logger                 *slog.Logger
	client                 *sharedAws.CognitoClient
	clientId, clientSecret string
}

func NewAuthRepositoryImpl(l *slog.Logger, c *sharedAws.CognitoClient, clientId, clientSecret string) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		logger:       l,
		client:       c,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (a *AuthRepositoryImpl) SignUp(ctx context.Context, auth auth.AuthAggregate) (*string, error) {
	secretHash := a.generateSecretHash(auth.GetEmail().GetValue())
	c := a.client.GetClient()
	output, err := c.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(a.clientId),
		Password:   aws.String(auth.GetPassword().GetValue()),
		Username:   aws.String(auth.GetEmail().GetValue()),
		SecretHash: aws.String(secretHash),
	})
	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			return nil, fmt.Errorf("%s", *invalidPassword.Message)
		}
		return nil, fmt.Errorf("Failed to sign up user %v. message: %w\n", auth.GetEmail().GetValue(), err)
	}
	return output.UserSub, nil
}

func (a *AuthRepositoryImpl) generateSecretHash(email string) string {
	mac := hmac.New(sha256.New, []byte(a.clientSecret))
	mac.Write([]byte(email + a.clientId))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
