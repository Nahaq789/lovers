package services

import "lovers/internal/shared/infrastructure/db"

type GroupQueryServiceImpl struct {
	db *db.DbClient
}

func NewGroupQueryService(d *db.DbClient) *GroupQueryServiceImpl {
	return &GroupQueryServiceImpl{db: d}
}
