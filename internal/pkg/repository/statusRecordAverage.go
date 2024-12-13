package repository

import (
	"beautyProject/internal/pkg/model"
	"gorm.io/gorm"
)

type StatusRecordAverage struct {
}

func (s *StatusRecordAverage) Add(tx *gorm.DB, record *model.StatusRecordAverage) error {
	return tx.Create(record).Error
}
