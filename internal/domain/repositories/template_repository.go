package repositories

import (
	"context"
	"lovers/internal/domain/models/aggregates/template"
)

type TemplateRepository interface {
	Create(ctx context.Context, template template.TemplateAggregate) error
}
