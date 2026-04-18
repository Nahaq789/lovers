package username

import (
	"testing"
)

func TestNewUserName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		errorMsg  string
	}{
		{
			name:      "正常系：有効なユーザー名",
			input:     "testuser",
			wantError: false,
		},
		{
			name:      "正常系：空文字列",
			input:     "",
			wantError: false,
		},
		{
			name:      "正常系：20文字のユーザー名",
			input:     "12345678901234567890",
			wantError: false,
		},
		{
			name:      "異常系：21文字のユーザー名",
			input:     "123456789012345678901",
			wantError: true,
			errorMsg:  "ユーザー名は20文字以内にしてください。",
		},
		{
			name:      "異常系：30文字のユーザー名",
			input:     "123456789012345678901234567890",
			wantError: true,
			errorMsg:  "ユーザー名は20文字以内にしてください。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUserName(tt.input)

			if tt.wantError {
				if err == nil {
					t.Errorf("NewUserName(%q) expected error, but got nil", tt.input)
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("NewUserName(%q) error message = %q, want %q", tt.input, err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("NewUserName(%q) unexpected error: %v", tt.input, err)
					return
				}
			}
		})
	}
}

func TestUserNameGetValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "正常系：空文字列",
			input: "",
			want:  "",
		},
		{
			name:  "正常系：通常のユーザー名",
			input: "testuser",
			want:  "testuser",
		},
		{
			name:  "正常系：数字を含むユーザー名",
			input: "user123",
			want:  "user123",
		},
		{
			name:  "正常系：特殊文字を含むユーザー名",
			input: "user_name",
			want:  "user_name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userName, err := NewUserName(tt.input)
			if err != nil {
				t.Fatalf("NewUserName(%q) unexpected error: %v", tt.input, err)
			}

			value := userName.GetValue()
			if value != tt.want {
				t.Errorf("UserName.GetValue() = %q, want %q", value, tt.want)
			}
		})
	}
}

// ベンチマークテスト
func BenchmarkNewUserName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewUserName("testuser")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetValue(b *testing.B) {
	userName, err := NewUserName("testuser")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = userName.GetValue()
	}
}