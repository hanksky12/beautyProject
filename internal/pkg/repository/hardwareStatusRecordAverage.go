package repository

import (
	"beautyProject/internal/pkg/entity"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
	"beautyProject/internal/pkg/util/web/requestConversion"
	"gorm.io/gorm"
)

type HardwareStatusRecordAverage struct {
}

func (s *HardwareStatusRecordAverage) Add(tx *gorm.DB, record *model.HardwareStatusRecordAverage) error {
	return tx.Create(record).Error
}

func (s *HardwareStatusRecordAverage) FindByUserId(userId uint64, cond map[string]any, paging *requestConversion.PagingSchema) ([]entity.HardwareRecord, int, error) {
	query := sql.Db.Model(&model.HardwareStatusRecordAverage{}).
		Select("hardware_status_record_average.id, hardware_status_record_average.time, hardware_status_record_average.percent, hardware.name as hardware_name").
		Joins("left join hardware ON hardware.id = hardware_status_record_average.hardware_id")
	query = query.Where("user_id = ?", userId)
	sqlConds := []sql.Cond{
		{Name: "Hardware", Type: "string", Query: "hardware.name = ?"},
		{Name: "MinDateTime", Type: "int64", Query: "time >= ?"},
		{Name: "MaxDateTime", Type: "int64", Query: "time <= ?"},
		{Name: "MinPercent", Type: "float64", Query: "percent >= ?"},
		{Name: "MaxPercent", Type: "float64", Query: "percent <= ?"},
	}
	sqlOrders := []sql.Order{
		{Name: "hardware_name", Query: "hardware_id"},
		{Name: "time", Query: "time"},
		{Name: "percent", Query: "percent"},
	}
	query = sql.ApplyConditions(query, cond, sqlConds)
	query = sql.ApplyOrder(query, paging, sqlOrders)
	return sql.DynamicQuery[entity.HardwareRecord](query, paging)
}
