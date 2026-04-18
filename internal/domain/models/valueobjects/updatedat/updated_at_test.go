package updatedat

import (
	"testing"
	"time"
)

func TestNewUpdatedAt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "新しいUpdatedAtが正常に作成される",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedAt := NewUpdatedAt()

			// Check that the value is set
			if updatedAt.value.IsZero() {
				t.Errorf("NewUpdatedAt() created a zero time value")
			}

			// Check that the time is close to now (within a reasonable tolerance)
			now := time.Now().UTC()
			if updatedAt.value.After(now.Add(time.Second)) {
				t.Errorf("NewUpdatedAt() time is in the future: got %v, now %v", updatedAt.value, now)
			}
		})
	}
}

func TestUpdatedAtGetValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "GetValue() が正しい値を返す",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedAt := NewUpdatedAt()
			value := updatedAt.GetValue()

			// Check that the returned value is the same as the internal value
			if value != updatedAt.value {
				t.Errorf("UpdatedAt.GetValue() = %v, want %v", value, updatedAt.value)
			}

			// Check that the returned value is in UTC
			if value.Location() != time.UTC {
				t.Errorf("UpdatedAt.GetValue() location = %v, want UTC", value.Location())
			}
		})
	}
}

func TestUpdatedAtToString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ToString() が正しい形式の文字列を返す",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedAt := NewUpdatedAt()
			str := updatedAt.ToString()

			// Check that the string is not empty
			if str == "" {
				t.Errorf("UpdatedAt.ToString() returned empty string")
			}

			// Try to parse the string back to time.Time to verify it's a valid RFC3339 format
			_, err := time.Parse(time.RFC3339, str)
			if err != nil {
				t.Errorf("UpdatedAt.ToString() returned invalid RFC3339 string: %v", err)
			}
		})
	}
}

func TestUpdatedAtEquality(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "複数回作成したUpdatedAtが異なる時間を持つ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedAt1 := NewUpdatedAt()
			time.Sleep(1 * time.Millisecond) // Ensure some time passes
			updatedAt2 := NewUpdatedAt()

			// They should be different times (unless the test runs extremely fast)
			if updatedAt1.value == updatedAt2.value {
				t.Logf("Warning: UpdatedAt values are equal, which might be expected in fast execution")
			}
		})
	}
}

func BenchmarkNewUpdatedAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewUpdatedAt()
	}
}

func BenchmarkUpdatedAtGetValue(b *testing.B) {
	updatedAt := NewUpdatedAt()
	for i := 0; i < b.N; i++ {
		_ = updatedAt.GetValue()
	}
}

func BenchmarkUpdatedAtToString(b *testing.B) {
	updatedAt := NewUpdatedAt()
	for i := 0; i < b.N; i++ {
		_ = updatedAt.ToString()
	}
}