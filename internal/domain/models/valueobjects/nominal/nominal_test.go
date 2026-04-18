package nominal

import (
	"testing"
)

func TestNewNominal(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効な名目_-_15文字以内",
			input:   "テスト名目",
			wantErr: false,
		},
		{
			name:    "境界値_-_ちょうど15文字",
			input:   "123456789012345",
			wantErr: false,
		},
		{
			name:    "無効な名目_-_16文字以上",
			input:   "1234567890123456",
			wantErr: true,
		},
		{
			name:    "空文字",
			input:   "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nominal, err := NewNominal(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewNominal() error = nil, wantErr %v", tt.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("NewNominal() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if nominal.value != tt.input {
					t.Errorf("NewNominal().value = %v, want %v", nominal.value, tt.input)
				}
			}
		})
	}
}

func TestNominalGetValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "基本的な名目",
			input: "テスト名目",
		},
		{
			name:  "空文字",
			input: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nominal, _ := NewNominal(tt.input)
			value := nominal.GetValue()
			if value != tt.input {
				t.Errorf("Nominal.GetValue() = %v, want %v", value, tt.input)
			}
		})
	}
}