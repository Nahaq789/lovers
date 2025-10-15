package userid

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestNewUserId(t *testing.T) {
	t.Run("正常系：新しいUserIdの生成", func(t *testing.T) {
		userId, err := NewUserId()
		if err != nil {
			t.Fatalf("予期しないエラー: %v", err)
		}

		// 値が空でないことを確認
		if userId.value == uuid.Nil {
			t.Error("UserIdの値が空です")
		}

		// GetValue()が有効なUUID文字列を返すことを確認
		value := userId.GetValue()
		if value == "" {
			t.Error("GetValue()が空文字列を返しました")
		}

		// 返された文字列が有効なUUIDであることを確認
		_, err = uuid.Parse(value)
		if err != nil {
			t.Errorf("GetValue()が無効なUUID文字列を返しました: %v", err)
		}
	})

	t.Run("一意性：複数のUserIdが異なる値を持つ", func(t *testing.T) {
		userId1, err := NewUserId()
		if err != nil {
			t.Fatalf("userId1の生成エラー: %v", err)
		}

		userId2, err := NewUserId()
		if err != nil {
			t.Fatalf("userId2の生成エラー: %v", err)
		}

		if userId1.GetValue() == userId2.GetValue() {
			t.Error("異なるUserIdが同じ値を持っています")
		}
	})
}

func TestNewUserIdFromString(t *testing.T) {
	t.Run("正常系：有効なUUID文字列からUserIdを生成", func(t *testing.T) {
		// 有効なUUID v4文字列
		validUUID := "550e8400-e29b-41d4-a716-446655440000"

		userId, err := NewUserIdFromString(validUUID)
		if err != nil {
			t.Fatalf("予期しないエラー: %v", err)
		}

		// GetValue()が元の文字列と一致することを確認
		if userId.GetValue() != validUUID {
			t.Errorf("期待値: %s, 実際値: %s", validUUID, userId.GetValue())
		}
	})

	t.Run("正常系：UUID v7文字列からUserIdを生成", func(t *testing.T) {
		// まず新しいUserIdを生成
		originalUserId, err := NewUserId()
		if err != nil {
			t.Fatalf("元のUserIdの生成エラー: %v", err)
		}

		// その文字列表現を取得
		uuidString := originalUserId.GetValue()

		// 文字列から新しいUserIdを生成
		userId, err := NewUserIdFromString(uuidString)
		if err != nil {
			t.Fatalf("予期しないエラー: %v", err)
		}

		// 値が一致することを確認
		if userId.GetValue() != uuidString {
			t.Errorf("期待値: %s, 実際値: %s", uuidString, userId.GetValue())
		}
	})

	t.Run("異常系：無効なUUID文字列", func(t *testing.T) {
		testCases := []struct {
			name  string
			input string
		}{
			{"空文字列", ""},
			{"不正な形式", "not-a-uuid"},
			{"短すぎる", "550e8400"},
			{"不正な文字を含む", "550e8400-e29b-41d4-a716-44665544000g"},
			{"長すぎる", "550e8400-e29b-41d4-a716-446655440000-extra"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := NewUserIdFromString(tc.input)
				if err == nil {
					t.Errorf("エラーが期待されましたが、nilが返されました: input=%s", tc.input)
				}
			})
		}
	})
}

func TestNewUserIdWithGenerator(t *testing.T) {
	t.Run("正常系：カスタムジェネレータでUserIdを生成", func(t *testing.T) {
		expectedUUID := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")

		// モックジェネレータ
		mockGenerator := func() (uuid.UUID, error) {
			return expectedUUID, nil
		}

		userId, err := newUserIdWithGenerator(mockGenerator)
		if err != nil {
			t.Fatalf("予期しないエラー: %v", err)
		}

		if userId.GetValue() != expectedUUID.String() {
			t.Errorf("期待値: %s, 実際値: %s", expectedUUID.String(), userId.GetValue())
		}
	})

	t.Run("異常系：ジェネレータがエラーを返す", func(t *testing.T) {
		expectedErr := errors.New("UUID生成エラー")

		// エラーを返すモックジェネレータ
		errorGenerator := func() (uuid.UUID, error) {
			return uuid.Nil, expectedErr
		}

		_, err := newUserIdWithGenerator(errorGenerator)
		if err == nil {
			t.Error("エラーが期待されましたが、nilが返されました")
		}

		if err != expectedErr {
			t.Errorf("期待されたエラー: %v, 実際のエラー: %v", expectedErr, err)
		}
	})
}

func TestUserId_GetValue(t *testing.T) {
	t.Run("正常系：GetValueが正しい文字列を返す", func(t *testing.T) {
		expectedUUID := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		userId := UserId{value: expectedUUID}

		result := userId.GetValue()
		expected := "550e8400-e29b-41d4-a716-446655440000"

		if result != expected {
			t.Errorf("期待値: %s, 実際値: %s", expected, result)
		}
	})

	t.Run("エッジケース：ゼロ値のUUID", func(t *testing.T) {
		userId := UserId{value: uuid.Nil}

		result := userId.GetValue()
		expected := "00000000-0000-0000-0000-000000000000"

		if result != expected {
			t.Errorf("期待値: %s, 実際値: %s", expected, result)
		}
	})
}

// ベンチマークテスト
func BenchmarkNewUserId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewUserId()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewUserIdFromString(b *testing.B) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := NewUserIdFromString(validUUID)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetValue(b *testing.B) {
	userId, err := NewUserId()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = userId.GetValue()
	}
}
