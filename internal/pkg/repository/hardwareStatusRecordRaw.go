package repository

import (
	"beautyProject/internal/pkg/entity"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
	"beautyProject/internal/pkg/util/web/requestConversion"
)

type HardwareStatusRecordRaw struct {
}

func (s *HardwareStatusRecordRaw) Add(record *model.HardwareStatusRecordRaw) error {
	return sql.Db.Create(record).Error
}

func (s *HardwareStatusRecordRaw) FindByUserId(userId uint64, cond map[string]any, paging *requestConversion.PagingSchema) ([]entity.HardwareRecord, int, error) {
	query := sql.Db.Model(&model.HardwareStatusRecordRaw{}).
		Select("hardware_status_record_raw.id, hardware_status_record_raw.time, hardware_status_record_raw.percent, hardware.name as hardware_name").
		Joins("left join hardware ON hardware.id = hardware_status_record_raw.hardware_id") // 修改為使用 status_record_raw
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
