package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
)

type StatusRecordRaw struct {
}

func (s *StatusRecordRaw) Add(record *model.StatusRecordRaw) error {
	return sql.Db.Create(record).Error
}
