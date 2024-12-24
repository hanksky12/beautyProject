package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
	"gorm.io/gorm"
)

type StatusRecord struct {
}

func (s *StatusRecord) Add(record *model.StatusRecord) error {
	return sql.Db.Create(record).Error
}

func (s *StatusRecord) UpdateProcessed(tx *gorm.DB, ids []int) error {
	return tx.Model(&model.StatusRecord{}).
		Where("id IN (?)", ids).
		Update("processed", true).Error
}
