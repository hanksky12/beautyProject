package repository

import (
	"beautyProject/internal/pkg/entity"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
	"beautyProject/internal/pkg/util/web/requestConversion"
)

type MouseActionStatusRecordRaw struct {
}

func (s *MouseActionStatusRecordRaw) Add(record *model.MouseActionStatusRecordRaw) error {
	return sql.Db.Create(record).Error
}

func (s *MouseActionStatusRecordRaw) FindByUserId(userId uint64, cond map[string]any, paging *requestConversion.PagingSchema) ([]entity.MouseActionRecord, int, error) {
	query := sql.Db.Model(&model.MouseActionStatusRecordRaw{}).
		Select("mouse_action_status_record_raw.id, mouse_action_status_record_raw.time, mouse_action_status_record_raw.x,mouse_action_status_record_raw.y, mouse_action.name as mouse_action_name").
		Joins("left join mouse_action ON mouse_action.id = mouse_action_status_record_raw.mouse_action_id")
	query = query.Where("user_id = ?", userId)
	sqlConds := []sql.Cond{
		{Name: "MouseActionName", Type: "string", Query: "mouse_action.name = ?"},
		{Name: "MinDateTime", Type: "int64", Query: "time >= ?"},
		{Name: "MaxDateTime", Type: "int64", Query: "time <= ?"},
		{Name: "MinX", Type: "int64", Query: "x >= ?"},
		{Name: "MaxX", Type: "int64", Query: "x <= ?"},
		{Name: "MinY", Type: "int64", Query: "y >= ?"},
		{Name: "MaxY", Type: "int64", Query: "y <= ?"},
	}
	//SortValidations
	sqlOrders := []sql.Order{
		{Name: "mouse_action_name", Query: "mouse_action_id"},
		{Name: "time", Query: "time"},
		{Name: "x", Query: "x"},
		{Name: "y", Query: "y"},
	}
	query = sql.ApplyConditions(query, cond, sqlConds)
	query = sql.ApplyOrder(query, paging, sqlOrders)
	return sql.DynamicQuery[entity.MouseActionRecord](query, paging)
}
