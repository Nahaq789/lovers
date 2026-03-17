package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/template"
	"lovers/internal/shared/infrastructure/db"
	"lovers/internal/shared/infrastructure/logger"
)

type TemplateRepositoryImpl struct {
	db *db.DbClient
}

func NewTemplateRepository(d *db.DbClient) *TemplateRepositoryImpl {
	return &TemplateRepositoryImpl{db: d}
}

func (t *TemplateRepositoryImpl) Create(ctx context.Context, template template.TemplateAggregate) error {
	l := logger.FromContext(ctx)
	query := `insert into "template" (template_id, group_id, created_by, template_name, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)`
	c := t.db.GetClient()
	_, err := c.ExecContext(ctx, query,
		template.GetTemplateId().GetValue(),
		template.GetGroupId().GetValue(),
		template.GetCreatedBy().GetValue(),
		template.GetTemplateName().GetValue(),
		template.GetCreatedAt().GetValue(),
		template.GetUpdatedAt().GetValue(),
	)
	if err != nil {
		l.ErrorContext(ctx, "failed to create template", "error", err)
		return err
	}

	return nil
}
