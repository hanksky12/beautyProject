package repository

import (
	"beautyProject/internal/pkg/entity"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
)

type HardwareStatusRecordQuery struct {
}

func (r *HardwareStatusRecordQuery) FindByNoProcessedWithUser(limit int) ([]entity.HardwareStatusRecordWithUser, error) {
	var results []entity.HardwareStatusRecordWithUser
	err := sql.Db.Model(&model.HardwareStatusRecord{}).
		Select("hardware_status_record.id as hardware_status_record_id, hardware_status_record.user_id, hardware_status_record.hardware_id, hardware_status_record.percent, hardware_status_record.time, hardware_status_record.processed").
		Joins("JOIN user ON hardware_status_record.user_id = user.id").
		Where("hardware_status_record.processed = ?", false).
		Order("hardware_status_record.time").
		Limit(limit).
		Scan(&results).Error
	return results, err
}
