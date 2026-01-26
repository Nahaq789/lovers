package template

import (
	"context"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/template/templateid"
	"lovers/internal/domain/models/template/templatename"
	domainUserId "lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
	"lovers/internal/domain/models/valueobjects/updatedat"
	"lovers/internal/domain/repositories"
	"lovers/internal/shared/infrastructure/logger"
	"lovers/internal/shared/infrastructure/security/userid"
	templateDto "lovers/internal/usecases/dto/template"
)

type TemplateCreate struct {
	templateRepository repositories.TemplateRepository
}

func NewTemplateCreate(t repositories.TemplateRepository) *TemplateCreate {
	return &TemplateCreate{
		templateRepository: t,
	}
}

func (tc *TemplateCreate) Execute(ctx context.Context, d templateDto.TemplateCreateDto) error {
	l := logger.FromContext(ctx)
	l.InfoContext(ctx, "テンプレート作成処理を開始します。")

	userId, err := domainUserId.NewUserIdFromString(userid.FromContext(ctx))
	if err != nil {
		l.ErrorContext(ctx, "ユーザーIDの取得に失敗しました。", "error", err)
		return err
	}

	templateId, err := templateid.NewTemplateId()
	if err != nil {
		l.ErrorContext(ctx, "テンプレートIDの生成に失敗しました。", "error", err)
		return err
	}

	groupId, err := groupid.NewGroupIdFromString(d.GroupId)
	if err != nil {
		l.ErrorContext(ctx, "グループIDの生成に失敗しました。", "error", err)
		return err
	}

	templateName, err := templatename.NewTemplateName(d.TemplateName)
	if err != nil {
		l.ErrorContext(ctx, "テンプレート名の検証に失敗しました。", "error", err)
		return err
	}

	createAt := createdat.NewCreatedAt()
	updatedAt := updatedat.NewUpdatedAt()
	return nil
}
