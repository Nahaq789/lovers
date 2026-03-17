//go:build wireinject
// +build wireinject

package template

import (
	domainRepos "lovers/internal/domain/repositories"
	"lovers/internal/presentation/template"

	"lovers/internal/infrastructure/repositories"
	"lovers/internal/shared/infrastructure/db"
	templateCreate "lovers/internal/usecases/template"

	"github.com/google/wire"
)

func ProvideTemplateRepository(d *db.DbClient) *repositories.TemplateRepositoryImpl {
	repository := repositories.NewTemplateRepository(d)
	return repository
}

var templateRepositorySet = wire.NewSet(
	ProvideTemplateRepository,
	wire.Bind(new(domainRepos.TemplateRepository), new(*repositories.TemplateRepositoryImpl)),
)

var createSet = wire.NewSet(templateCreate.NewTemplateCreate)
var templateControllerSet = wire.NewSet(template.NewTemplateController)

type TemplateSet struct {
	TemplateController *template.TemplateController
}

func Initialize(d *db.DbClient) *TemplateSet {
	wire.Build(
		templateRepositorySet,
		createSet,
		templateControllerSet,
		wire.Struct(new(TemplateSet), "*"),
	)
	return nil
}
