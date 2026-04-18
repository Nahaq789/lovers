package paymentdate

import (
	"testing"
	"time"
)

func TestNewPaymentDate(t *testing.T) {
	t.Run("新しい支払日が正常に作成される", func(t *testing.T) {
		paymentDate := NewPaymentDate()
		now := time.Now().UTC()
		value := paymentDate.GetValue()

		// Check that the date is today (UTC)
		if value.Year() != now.Year() || value.Month() != now.Month() || value.Day() != now.Day() {
			t.Errorf("NewPaymentDate() = %v, want today's date", value)
		}
	})
}

func TestNewPaymentDateFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効な日付形式_-_YYYY-MM-DD",
			input:   "2023-12-25",
			wantErr: false,
		},
		{
			name:    "有効な日付形式_-_年が4桁",
			input:   "2024-01-01",
			wantErr: false,
		},
		{
			name:    "無効な日付形式_-_日付が短い",
			input:   "2023-1-1",
			wantErr: true,
		},
		{
			name:    "無効な日付形式_-_形式が異なる",
			input:   "25/12/2023",
			wantErr: true,
		},
		{
			name:    "無効な日付形式_-_文字列が空",
			input:   "",
			wantErr: true,
		},
		{
			name:    "無効な日付形式_-_日付が存在しない",
			input:   "2023-02-29",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paymentDate, err := NewPaymentDateFromString(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewPaymentDateFromString() error = nil, wantErr %v", tt.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("NewPaymentDateFromString() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				// Verify that the date is correctly parsed
				expected, _ := time.Parse(PaymentDateFormat, tt.input)
				if paymentDate.GetValue() != expected {
					t.Errorf("NewPaymentDateFromString().GetValue() = %v, want %v", paymentDate.GetValue(), expected)
				}
			}
		})
	}
}

func TestPaymentDateGetValue(t *testing.T) {
	t.Run("GetValue()が正しい値を返す", func(t *testing.T) {
		paymentDate, _ := NewPaymentDateFromString("2023-12-25")
		value := paymentDate.GetValue()
		expected, _ := time.Parse(PaymentDateFormat, "2023-12-25")
		if value != expected {
			t.Errorf("PaymentDate.GetValue() = %v, want %v", value, expected)
		}
	})
}

func TestPaymentDateToString(t *testing.T) {
	t.Run("ToString()が正しい形式の文字列を返す", func(t *testing.T) {
		paymentDate, _ := NewPaymentDateFromString("2023-12-25")
		str := paymentDate.ToString()
		if str != "2023-12-25" {
			t.Errorf("PaymentDate.ToString() = %v, want 2023-12-25", str)
		}
	})
}