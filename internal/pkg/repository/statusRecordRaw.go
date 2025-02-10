package repository

import (
	"beautyProject/internal/pkg/entity"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
	"beautyProject/internal/pkg/util/web/requestConversion"
)

type StatusRecordRaw struct {
}

func (s *StatusRecordRaw) Add(record *model.StatusRecordRaw) error {
	return sql.Db.Create(record).Error
}

func (s *StatusRecordRaw) FindByUserId(userId uint64, cond map[string]any, paging *requestConversion.PagingSchema) ([]entity.Record, int, error) {
	query := sql.Db.Model(&model.StatusRecordRaw{}). // 修改為使用 StatusRecordRaw
								Select("status_record_raw.id, status_record_raw.time, status_record_raw.percent, hardware.name as hardware_name").
								Joins("left join hardware ON hardware.id = status_record_raw.hardware_id") // 修改為使用 status_record_raw
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
	return sql.DynamicQuery[entity.Record](query, paging)
}
