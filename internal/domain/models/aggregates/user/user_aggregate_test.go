package user

import (
	"testing"

	"github.com/google/uuid"
	userid "lovers/internal/domain/models/user/userid"
	username "lovers/internal/domain/models/user/username"
	createdat "lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/email"
	updatedat "lovers/internal/domain/models/valueobjects/updatedat"
)

func TestNewUserAggregate(t *testing.T) {
	t.Run("正常系：新しいUserAggregateの生成", func(t *testing.T) {
		// テスト用の値を作成
		userId, err := userid.NewUserId()
		if err != nil {
			t.Fatalf("UserIdの生成エラー: %v", err)
		}

		email, err := email.NewEmail("test@example.com")
		if err != nil {
			t.Fatalf("Emailの生成エラー: %v", err)
		}

		userName, err := username.NewUserName("testuser")
		if err != nil {
			t.Fatalf("UserNameの生成エラー: %v", err)
		}

		createdAt := createdat.NewCreatedAt()
		updatedAt := updatedat.NewUpdatedAt()

		// UserAggregateを作成
		user := NewUserAggregate(userId, email, userName, createdAt, updatedAt)

		// 値が正しく設定されているか確認
		if user.GetUserId().GetValue() != userId.GetValue() {
			t.Error("UserIdが正しく設定されていません")
		}

		if user.GetEmail().GetValue() != email.GetValue() {
			t.Error("Emailが正しく設定されていません")
		}

		if user.GetUserName().GetValue() != userName.GetValue() {
			t.Error("UserNameが正しく設定されていません")
		}

		if user.GetCreatedAt().GetValue() != createdAt.GetValue() {
			t.Error("CreatedAtが正しく設定されていません")
		}

		if user.GetUpdatedAt().GetValue() != updatedAt.GetValue() {
			t.Error("UpdatedAtが正しく設定されていません")
		}
	})
}

func TestUserAggregate_Getters(t *testing.T) {
	t.Run("正常系：各getterメソッドが正しく動作", func(t *testing.T) {
		// テスト用の値を作成
		testUUID := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		userId, err := userid.NewUserIdFromString(testUUID.String())
		if err != nil {
			t.Fatalf("UserIdの生成エラー: %v", err)
		}

		email, err := email.NewEmail("test@example.com")
		if err != nil {
			t.Fatalf("Emailの生成エラー: %v", err)
		}

		userName, err := username.NewUserName("testuser")
		if err != nil {
			t.Fatalf("UserNameの生成エラー: %v", err)
		}

		createdAt := createdat.NewCreatedAt()
		updatedAt := updatedat.NewUpdatedAt()

		// UserAggregateを作成
		user := NewUserAggregate(userId, email, userName, createdAt, updatedAt)

		// 各getterメソッドの結果を確認
		if user.GetUserId() != userId {
			t.Error("GetUserId()が正しく動作していません")
		}

		if user.GetEmail() != email {
			t.Error("GetEmail()が正しく動作していません")
		}

		if user.GetUserName() != userName {
			t.Error("GetUserName()が正しく動作していません")
		}

		if user.GetCreatedAt() != createdAt {
			t.Error("GetCreatedAt()が正しく動作していません")
		}

		if user.GetUpdatedAt() != updatedAt {
			t.Error("GetUpdatedAt()が正しく動作していません")
		}
	})
}