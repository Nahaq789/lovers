package amount

import (
	"testing"
)

func TestNewAmount(t *testing.T) {
	tests := []struct {
		name      string
		input     int64
		wantError bool
		errorMsg  string
	}{
		{
			name:      "有効な金額 - プロダクション値",
			input:     1000,
			wantError: false,
		},
		{
			name:      "有効な金額 - ゼロ",
			input:     0,
			wantError: false,
		},
		{
			name:      "有効な金額 - 大きな値",
			input:     999999999999,
			wantError: false,
		},
		{
			name:      "無効な金額 - 負の値",
			input:     -100,
			wantError: true,
			errorMsg:  "amount must be non-negative, got -100",
		},
		{
			name:      "無効な金額 - 負のゼロ",
			input:     -0,
			wantError: false,
			errorMsg:  "amount must be non-negative, got 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAmount(tt.input)

			if tt.wantError {
				if err == nil {
					t.Errorf("NewAmount(%d) expected error, but got nil", tt.input)
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("NewAmount(%d) error message = %q, want %q", tt.input, err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("NewAmount(%d) unexpected error: %v", tt.input, err)
					return
				}
			}
		})
	}
}

func TestNewAmountZero(t *testing.T) {
	amount := NewAmountZero()

	if amount.GetValue() != 0 {
		t.Errorf("NewAmountZero() returned amount with value %d, want 0", amount.GetValue())
	}
}

func TestAmountGetValue(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  int64
	}{
		{
			name:  "ゼロ金額",
			input: 0,
			want:  0,
		},
		{
			name:  "小さな金額",
			input: 50,
			want:  50,
		},
		{
			name:  "大きな金額",
			input: 1000000,
			want:  1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, err := NewAmount(tt.input)
			if err != nil {
				t.Fatalf("NewAmount(%d) unexpected error: %v", tt.input, err)
			}

			value := amount.GetValue()
			if value != tt.want {
				t.Errorf("Amount.GetValue() = %d, want %d", value, tt.want)
			}
		})
	}
}

func TestAmountAdd(t *testing.T) {
	tests := []struct {
		name       string
		amount1    int64
		amount2    int64
		wantError  bool
		errorMsg   string
		wantResult int64
	}{
		{
			name:       "正常な加算 - ゼロとゼロ",
			amount1:    0,
			amount2:    0,
			wantError:  false,
			wantResult: 0,
		},
		{
			name:       "正常な加算 - 小さな金額",
			amount1:    100,
			amount2:    200,
			wantError:  false,
			wantResult: 300,
		},
		{
			name:       "正常な加算 - 大きな金額",
			amount1:    999999999999,
			amount2:    1,
			wantError:  false,
			wantResult: 1000000000000,
		},
		{
			name:      "加算エラー - 第二引数が負数",
			amount1:   100,
			amount2:   -50,
			wantError: true,
			errorMsg:  "amount must be non-negative, got -50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount1, err1 := NewAmount(tt.amount1)
			if err1 != nil {
				t.Fatalf("NewAmount(%d) unexpected error: %v", tt.amount1, err1)
			}

			amount2, err2 := NewAmount(tt.amount2)
			if err2 != nil {
				if tt.wantError {
					// If we expect an error, check if it matches
					if err2.Error() != tt.errorMsg {
						t.Errorf("NewAmount(%d) error message = %q, want %q", tt.amount2, err2.Error(), tt.errorMsg)
					}
					return
				} else {
					t.Fatalf("NewAmount(%d) unexpected error: %v", tt.amount2, err2)
				}
			}

			result, err := amount1.Add(amount2)
			if tt.wantError {
				if err == nil {
					t.Errorf("Amount.Add() expected error, but got nil")
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("Amount.Add() error message = %q, want %q", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("Amount.Add() unexpected error: %v", err)
					return
				}
				if result.GetValue() != tt.wantResult {
					t.Errorf("Amount.Add() result = %d, want %d", result.GetValue(), tt.wantResult)
				}
			}
		})
	}
}

func TestAmountSubtract(t *testing.T) {
	tests := []struct {
		name       string
		amount1    int64
		amount2    int64
		wantError  bool
		errorMsg   string
		wantResult int64
	}{
		{
			name:       "正常な減算 - ゼロからゼロを引く",
			amount1:    0,
			amount2:    0,
			wantError:  false,
			wantResult: 0,
		},
		{
			name:       "正常な減算 - 小さな金額",
			amount1:    300,
			amount2:    100,
			wantError:  false,
			wantResult: 200,
		},
		{
			name:       "正常な減算 - 同じ金額",
			amount1:    500,
			amount2:    500,
			wantError:  false,
			wantResult: 0,
		},
		{
			name:      "減算エラー - 第二引数が負数",
			amount1:   100,
			amount2:   -50,
			wantError: true,
			errorMsg:  "amount must be non-negative, got -50",
		},
		{
			name:      "減算エラー - 結果が負数",
			amount1:   100,
			amount2:   200,
			wantError: true,
			errorMsg:  "subtraction result is negative: 100 - 200 = -100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount1, err1 := NewAmount(tt.amount1)
			if err1 != nil {
				t.Fatalf("NewAmount(%d) unexpected error: %v", tt.amount1, err1)
			}

			amount2, err2 := NewAmount(tt.amount2)
			if err2 != nil {
				if tt.wantError {
					// If we expect an error, check if it matches
					if err2.Error() != tt.errorMsg {
						t.Errorf("NewAmount(%d) error message = %q, want %q", tt.amount2, err2.Error(), tt.errorMsg)
					}
					return
				} else {
					t.Fatalf("NewAmount(%d) unexpected error: %v", tt.amount2, err2)
				}
			}

			result, err := amount1.Subtract(amount2)
			if tt.wantError {
				if err == nil {
					t.Errorf("Amount.Subtract() expected error, but got nil")
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("Amount.Subtract() error message = %q, want %q", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("Amount.Subtract() unexpected error: %v", err)
					return
				}
				if result.GetValue() != tt.wantResult {
					t.Errorf("Amount.Subtract() result = %d, want %d", result.GetValue(), tt.wantResult)
				}
			}
		})
	}
}

func BenchmarkNewAmount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewAmount(1000)
	}
}

func BenchmarkAmountGetValue(b *testing.B) {
	amount, _ := NewAmount(1000)
	for i := 0; i < b.N; i++ {
		_ = amount.GetValue()
	}
}

func BenchmarkAmountAdd(b *testing.B) {
	amount1, _ := NewAmount(1000)
	amount2, _ := NewAmount(2000)
	for i := 0; i < b.N; i++ {
		_, _ = amount1.Add(amount2)
	}
}

func BenchmarkAmountSubtract(b *testing.B) {
	amount1, _ := NewAmount(3000)
	amount2, _ := NewAmount(1000)
	for i := 0; i < b.N; i++ {
		_, _ = amount1.Subtract(amount2)
	}
}

