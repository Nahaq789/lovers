package group

import (
	"context"
	"lovers/internal/domain/models/aggregates/group"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/group/groupname"
	domainUserId "lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/updatedat"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	"lovers/internal/shared/infrastructure/security/userid"
	groupDto "lovers/internal/usecases/dto/group"
)

type GroupCreate struct {
	groupRepository repositories.GroupRepository
}

func NewGroupCreate(g repositories.GroupRepository) *GroupCreate {
	return &GroupCreate{
		groupRepository: g,
	}
}

func (gc *GroupCreate) Execute(ctx context.Context, d *groupDto.GroupCreateDto) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "グループ作成処理を開始します。")

	userId, err := domainUserId.NewUserIdFromString(userid.FromContext(ctx))
	if err != nil {
		l.ErrorContext(ctx, "ユーザーIDの取得に失敗しました。", "error", err)
		return err
	}

	groupId, err := groupid.NewGroupId()
	if err != nil {
		l.ErrorContext(ctx, "グループIDの生成に失敗しました。", "error", err)
		return err
	}

	groupName, err := groupname.NewGroupName(d.GroupName)
	if err != nil {
		l.ErrorContext(ctx, "グループ名の検証に失敗しました。", "error", err)
		return err
	}

	createdAt := createdat.NewCreatedAt()
	updatedAt := updatedat.NewUpdatedAt()

	group := group.NewGroupAggregate(
		groupId,
		userId,
		groupName,
		createdAt,
		updatedAt,
	)

	// グループ作成者をメンバーに追加する。
	group.AddMember(userId)

	dbErr := gc.groupRepository.Create(ctx, *group)
	if dbErr != nil {
		l.ErrorContext(ctx, "データベース保存に失敗しました。", "error", dbErr)
		return dbErr
	}

	l.InfoContext(ctx, "グループ作成処理が正常に完了しました。")

	return nil
}
