package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
)

type StatusRecord struct {
}

func (s *StatusRecord) Add(record *model.StatusRecord) error {
	return sql.Db.Create(record).Error
}
