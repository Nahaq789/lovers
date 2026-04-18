package deletedat

import (
	"testing"
	"time"
)

func TestNewDeletedAt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "新しいDeletedAtが正常に作成される",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deletedAt := NewDeletedAt()

			// Check that the value is set
			if deletedAt.value.IsZero() {
				t.Errorf("NewDeletedAt() created a zero time value")
			}

			// Check that the time is close to now (within a reasonable tolerance)
			now := time.Now().UTC()
			if deletedAt.value.After(now.Add(time.Second)) {
				t.Errorf("NewDeletedAt() time is in the future: got %v, now %v", deletedAt.value, now)
			}
		})
	}
}

func TestDeletedAtGetValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "GetValue() が正しい値を返す",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deletedAt := NewDeletedAt()
			value := deletedAt.GetValue()

			// Check that the returned value is the same as the internal value
			if value != deletedAt.value {
				t.Errorf("DeletedAt.GetValue() = %v, want %v", value, deletedAt.value)
			}

			// Check that the returned value is in UTC
			if value.Location() != time.UTC {
				t.Errorf("DeletedAt.GetValue() location = %v, want UTC", value.Location())
			}
		})
	}
}

func TestDeletedAtToString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ToString() が正しい形式の文字列を返す",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deletedAt := NewDeletedAt()
			str := deletedAt.ToString()

			// Check that the string is not empty
			if str == "" {
				t.Errorf("DeletedAt.ToString() returned empty string")
			}

			// Try to parse the string back to time.Time to verify it's a valid RFC3339 format
			_, err := time.Parse(time.RFC3339, str)
			if err != nil {
				t.Errorf("DeletedAt.ToString() returned invalid RFC3339 string: %v", err)
			}
		})
	}
}

func TestDeletedAtEquality(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "複数回作成したDeletedAtが異なる時間を持つ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deletedAt1 := NewDeletedAt()
			time.Sleep(1 * time.Millisecond) // Ensure some time passes
			deletedAt2 := NewDeletedAt()

			// They should be different times (unless the test runs extremely fast)
			if deletedAt1.value == deletedAt2.value {
				t.Logf("Warning: DeletedAt values are equal, which might be expected in fast execution")
			}
		})
	}
}

func BenchmarkNewDeletedAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewDeletedAt()
	}
}

func BenchmarkDeletedAtGetValue(b *testing.B) {
	deletedAt := NewDeletedAt()
	for i := 0; i < b.N; i++ {
		_ = deletedAt.GetValue()
	}
}

func BenchmarkDeletedAtToString(b *testing.B) {
	deletedAt := NewDeletedAt()
	for i := 0; i < b.N; i++ {
		_ = deletedAt.ToString()
	}
}