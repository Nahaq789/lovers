package group

import (
	"context"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/updatedat"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	"lovers/internal/usecases/dto/group"
)

type GroupCreate struct {
	groupRepository repositories.GroupRepository
}

func NewGroupCreate(g repositories.GroupRepository) *GroupCreate {
	return &GroupCreate{
		groupRepository: g,
	}
}

func (gc *GroupCreate) Execute(ctx context.Context, d *group.GroupCreateDto) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "グループ作成処理を開始します。")

	groupId, err := groupid.NewGroupId()
	if err != nil {
		l.ErrorContext(ctx, "グループ作成処理でエラーが発生しました。", "error", err)
		return err
	}

	created_by, err := userid.NewUserIdFromString(d.CreatedBy)
	if err != nil {
		l.ErrorContext(ctx, "グループ作成処理でエラーが発生しました。", "error", err)
		return err
	}

	createdAt := createdat.NewCreatedAt()
	updatedAt := updatedat.NewUpdatedAt()

	return nil
}
