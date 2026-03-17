package initialize

import (
	"context"
	"lovers/cmd/di/template"
	"lovers/internal/shared/infrastructure/db"
)

func InitTemplate(ctx context.Context, d *db.DbClient) *template.TemplateSet {
	templateSet := template.Initialize(d)
	return templateSet
}
