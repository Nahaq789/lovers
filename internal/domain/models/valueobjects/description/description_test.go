package description

import (
	"testing"
)

func TestNewDescription(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "基本的な説明文",
			input: "テスト説明文",
		},
		{
			name:  "空文字",
			input: "",
		},
		{
			name:  "長い説明文",
			input: "これは非常に長い説明文です。これは非常に長い説明文です。これは非常に長い説明文です。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			description := NewDescription(tt.input)
			if description.value != tt.input {
				t.Errorf("NewDescription().value = %v, want %v", description.value, tt.input)
			}
		})
	}
}

func TestDescriptionGetValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "基本的な説明文",
			input: "テスト説明文",
		},
		{
			name:  "空文字",
			input: "",
		},
		{
			name:  "特殊文字を含む",
			input: "テスト@#$%説明文",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			description := NewDescription(tt.input)
			value := description.GetValue()
			if value != tt.input {
				t.Errorf("Description.GetValue() = %v, want %v", value, tt.input)
			}
		})
	}
}