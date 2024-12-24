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
		Select("status_records.id as status_record_id, status_records.user_id, status_records.hardware_id, status_records.percent, status_records.time, status_records.processed").
		Joins("JOIN users ON status_records.user_id = users.id").
		Where("status_records.processed = ?", false).
		Order("status_records.time").
		Limit(limit).
		Scan(&results).Error
	return results, err
}
