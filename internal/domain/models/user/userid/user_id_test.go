package userid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserId(t *testing.T) {
	t.Run("新しいUserIdの生成", func(t *testing.T) {
		userID, err := NewUserId()
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, userID.value)
		assert.NotEmpty(t, userID.GetValue())
	})
}

func TestNewUserIdFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なUUID文字列からUserIdを生成",
			input:   "550e8400-e29b-41d4-a716-446655440000",
			wantErr: false,
		},
		{
			name:    "UUID_v7文字列からUserIdを生成",
			input:   "018d8a3a-3e5e-7000-8000-000000000000",
			wantErr: false,
		},
		{
			name:    "無効なUUID文字列",
			input:   "invalid-uuid",
			wantErr: true,
		},
		{
			name:    "無効なUUID文字列/空文字列",
			input:   "",
			wantErr: true,
		},
		{
			name:    "無効なUUID文字列/不正な形式",
			input:   "12345",
			wantErr: true,
		},
		{
			name:    "無効なUUID文字列/短すぎる",
			input:   "123e4567-e89b-12d3-a456",
			wantErr: true,
		},
		{
			name:    "無効なUUID文字列/不正な文字を含む",
			input:   "123e4567-e89b-12d3-a456-00000000000g",
			wantErr: true,
		},
		{
			name:    "無効なUUID文字列/長すぎる",
			input:   "123e4567-e89b-12d3-a456-0000000000000",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userID, err := NewUserIdFromString(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewUserIdFromString() error = nil, wantErr %v", tt.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("NewUserIdFromString() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if userID.GetValue() != tt.input {
					t.Errorf("NewUserIdFromString().GetValue() = %v, want %v", userID.GetValue(), tt.input)
				}
			}
		})
	}
}

func TestUserIdGetValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "GetValueが正しい文字列を返す",
			input: "550e8400-e29b-41d4-a716-446655440000",
		},
		{
			name:  "ゼロ値のUUID",
			input: "00000000-0000-0000-0000-000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userID, _ := NewUserIdFromString(tt.input)
			if userID.GetValue() != tt.input {
				t.Errorf("UserId.GetValue() = %v, want %v", userID.GetValue(), tt.input)
			}
		})
	}
}

func TestUserIdEqual(t *testing.T) {
	t.Run("UserIdが正しく比較される", func(t *testing.T) {
		userID1, _ := NewUserIdFromString("550e8400-e29b-41d4-a716-446655440000")
		userID2, _ := NewUserIdFromString("550e8400-e29b-41d4-a716-446655440000")
		userID3, _ := NewUserIdFromString("550e8400-e29b-41d4-a716-446655440001")

		assert.True(t, userID1.Equal(userID2))
		assert.False(t, userID1.Equal(userID3))
	})
}