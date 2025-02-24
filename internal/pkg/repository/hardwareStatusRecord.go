package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
	"gorm.io/gorm"
)

type HardwareStatusRecord struct {
}

func (s *HardwareStatusRecord) Add(record *model.HardwareStatusRecord) error {
	return sql.Db.Create(record).Error
}

func (s *HardwareStatusRecord) UpdateProcessed(tx *gorm.DB, ids []int) error {
	return tx.Model(&model.HardwareStatusRecord{}).
		Where("id IN (?)", ids).
		Update("processed", true).Error
}
