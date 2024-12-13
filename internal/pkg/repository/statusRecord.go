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

func (s *StatusRecord) FindByUserAndHardware(userId uint, hardwareId int, limit int) ([]model.MiniStatusRecord, error) {
	var records []model.MiniStatusRecord
	err := sql.Db.Model(&model.StatusRecord{}).
		Where("user_id = ?", userId).
		Where("hardware_id = ?", hardwareId).
		Where("processed = ?", false).
		Order("time").
		Limit(limit).
		Find(&records).Error
	return records, err
}

func (s *StatusRecord) UpdateProcessed(tx *gorm.DB, ids []int) error {
	return tx.Model(&model.StatusRecord{}).
		Where("id IN (?)", ids).
		Update("processed", true).Error
}
