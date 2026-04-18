package password

import (
	"testing"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なパスワード",
			input:   "abc123",
			wantErr: false,
		},
		{
			name:    "有効なパスワード（長い）",
			input:   "test1234",
			wantErr: false,
		},
		{
			name:    "有効なパスワード（大文字含む）",
			input:   "MyPass123",
			wantErr: false,
		},
		{
			name:    "境界値：ちょうど6文字",
			input:   "abc123",
			wantErr: false,
		},
		{
			name:    "エラー：6文字未満",
			input:   "abc12",
			wantErr: true,
		},
		{
			name:    "エラー：小文字なし",
			input:   "123456",
			wantErr: true,
		},
		{
			name:    "エラー：数字なし",
			input:   "abcdef",
			wantErr: true,
		},
		{
			name:    "エラー：空文字",
			input:   "",
			wantErr: true,
		},
		{
			name:    "エラー：大文字と記号のみ",
			input:   "ABC!@#",
			wantErr: true,
		},
		{
			name:    "エラー：スペースを含む",
			input:   "abc 123",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := NewPassword(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewPassword() error = nil, wantErr %v", tt.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("NewPassword() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if password.value != tt.input {
					t.Errorf("NewPassword().value = %v, want %v", password.value, tt.input)
				}
			}
		})
	}
}