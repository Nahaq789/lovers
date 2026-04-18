package amount

import (
	"testing"
)

func TestNewAmount(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		wantErr bool
	}{
		{
			name:    "有効な金額 - プロダクション値",
			input:   1000,
			wantErr: false,
		},
		{
			name:    "有効な金額 - ゼロ",
			input:   0,
			wantErr: false,
		},
		{
			name:    "有効な金額 - 大きな値",
			input:   9223372036854775807,
			wantErr: false,
		},
		{
			name:    "無効な金額 - 負の値",
			input:   -100,
			wantErr: true,
		},
		{
			name:    "無効な金額 - 負のゼロ",
			input:   -0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, err := NewAmount(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewAmount() error = nil, wantErr %v", tt.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("NewAmount() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if amount.GetValue() != tt.input {
					t.Errorf("NewAmount().GetValue() = %v, want %v", amount.GetValue(), tt.input)
				}
			}
		})
	}
}

func TestNewAmountZero(t *testing.T) {
	amount := NewAmountZero()
	if amount.GetValue() != 0 {
		t.Errorf("NewAmountZero() = %v, want 0", amount.GetValue())
	}
}

func TestAmountGetValue(t *testing.T) {
	tests := []struct {
		name  string
		input int64
	}{
		{
			name:  "ゼロ金額",
			input: 0,
		},
		{
			name:  "小さな金額",
			input: 100,
		},
		{
			name:  "大きな金額",
			input: 9223372036854775807,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, _ := NewAmount(tt.input)
			if amount.GetValue() != tt.input {
				t.Errorf("Amount.GetValue() = %v, want %v", amount.GetValue(), tt.input)
			}
		})
	}
}

func TestAmountAdd(t *testing.T) {
	tests := []struct {
		name    string
		a       int64
		b       int64
		want    int64
		wantErr bool
	}{
		{
			name:    "正常な加算 - ゼロとゼロ",
			a:       0,
			b:       0,
			want:    0,
			wantErr: false,
		},
		{
			name:    "正常な加算 - 小さな金額",
			a:       100,
			b:       200,
			want:    300,
			wantErr: false,
		},
		{
			name:    "正常な加算 - 通常の金額",
			a:       1000,
			b:       2000,
			want:    3000,
			wantErr: false,
		},
		{
			name:    "正常な加算 - 大きな金額",
			a:       9223372036854775800,
			b:       5,
			want:    9223372036854775805,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amountA, _ := NewAmount(tt.a)
			amountB, _ := NewAmount(tt.b)

			result, err := amountA.Add(amountB)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Amount.Add() error = nil, wantErr %v", tt.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("Amount.Add() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if result.GetValue() != tt.want {
					t.Errorf("Amount.Add() = %v, want %v", result.GetValue(), tt.want)
				}
			}
		})
	}
}

func TestAmountSubtract(t *testing.T) {
	tests := []struct {
		name    string
		a       int64
		b       int64
		want    int64
		wantErr bool
	}{
		{
			name:    "正常な減算 - ゼロからゼロを引く",
			a:       0,
			b:       0,
			want:    0,
			wantErr: false,
		},
		{
			name:    "正常な減算 - 小さな金額",
			a:       300,
			b:       100,
			want:    200,
			wantErr: false,
		},
		{
			name:    "正常な減算 - 同じ金額",
			a:       100,
			b:       100,
			want:    0,
			wantErr: false,
		},
		{
			name:    "減算エラー - 結果が負数",
			a:       100,
			b:       150,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amountA, _ := NewAmount(tt.a)
			amountB, _ := NewAmount(tt.b)

			result, err := amountA.Subtract(amountB)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Amount.Subtract() error = nil, wantErr %v", tt.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("Amount.Subtract() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if result.GetValue() != tt.want {
					t.Errorf("Amount.Subtract() = %v, want %v", result.GetValue(), tt.want)
				}
			}
		})
	}
}