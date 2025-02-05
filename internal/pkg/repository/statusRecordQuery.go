package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
)

type StatusRecordQuery struct {
}

type StatusRecordWithUser struct {
	StatusRecordID uint
	UserId         uint
	HardwareId     uint
	Percent        float64
	Time           int64
	Processed      bool
}

func (r *StatusRecordQuery) FindByNoProcessedWithUser(limit int) ([]StatusRecordWithUser, error) {
	var results []StatusRecordWithUser
	err := sql.Db.Model(&model.StatusRecord{}).
		Select("status_record.id as status_record_id, status_record.user_id, status_record.hardware_id, status_record.percent, status_record.time, status_record.processed").
		Joins("JOIN user ON status_record.user_id = user.id").
		Where("status_record.processed = ?", false).
		Order("status_record.time").
		Limit(limit).
		Scan(&results).Error
	return results, err
}
