package config

import "os"

type ClientId string
type ClientSecret string

type CognitoConfig struct {
	ClientId     string
	ClientSecret string
}

func LoadCognitoConfig() *CognitoConfig {
	return &CognitoConfig{
		ClientId:     os.Getenv("COGNITO_CLIENT_ID"),
		ClientSecret: os.Getenv("COGNITO_CLIENT_SECRET"),
	}
}
