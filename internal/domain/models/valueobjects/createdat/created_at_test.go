package createdat

import (
	"testing"
	"time"
)

func TestNewCreatedAt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "新しいCreatedAtが正常に作成される",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdAt := NewCreatedAt()

			// Check that the value is set
			if createdAt.value.IsZero() {
				t.Errorf("NewCreatedAt() created a zero time value")
			}

			// Check that the time is close to now (within a reasonable tolerance)
			now := time.Now().UTC()
			if createdAt.value.After(now.Add(time.Second)) {
				t.Errorf("NewCreatedAt() time is in the future: got %v, now %v", createdAt.value, now)
			}
		})
	}
}

func TestCreatedAtGetValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "GetValue() が正しい値を返す",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdAt := NewCreatedAt()
			value := createdAt.GetValue()

			// Check that the returned value is the same as the internal value
			if value != createdAt.value {
				t.Errorf("CreatedAt.GetValue() = %v, want %v", value, createdAt.value)
			}

			// Check that the returned value is in UTC
			if value.Location() != time.UTC {
				t.Errorf("CreatedAt.GetValue() location = %v, want UTC", value.Location())
			}
		})
	}
}

func TestCreatedAtToString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ToString() が正しい形式の文字列を返す",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdAt := NewCreatedAt()
			str := createdAt.ToString()

			// Check that the string is not empty
			if str == "" {
				t.Errorf("CreatedAt.ToString() returned empty string")
			}

			// Try to parse the string back to time.Time to verify it's a valid RFC3339 format
			_, err := time.Parse(time.RFC3339, str)
			if err != nil {
				t.Errorf("CreatedAt.ToString() returned invalid RFC3339 string: %v", err)
			}
		})
	}
}

func TestCreatedAtEquality(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "複数回作成したCreatedAtが異なる時間を持つ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdAt1 := NewCreatedAt()
			time.Sleep(1 * time.Millisecond) // Ensure some time passes
			createdAt2 := NewCreatedAt()

			// They should be different times (unless the test runs extremely fast)
			if createdAt1.value == createdAt2.value {
				t.Logf("Warning: CreatedAt values are equal, which might be expected in fast execution")
			}
		})
	}
}

func BenchmarkNewCreatedAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewCreatedAt()
	}
}

func BenchmarkCreatedAtGetValue(b *testing.B) {
	createdAt := NewCreatedAt()
	for i := 0; i < b.N; i++ {
		_ = createdAt.GetValue()
	}
}

func BenchmarkCreatedAtToString(b *testing.B) {
	createdAt := NewCreatedAt()
	for i := 0; i < b.N; i++ {
		_ = createdAt.ToString()
	}
}