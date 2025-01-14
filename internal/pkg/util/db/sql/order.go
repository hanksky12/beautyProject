package sql

import (
	"beautyProject/internal/pkg/util/web/requestConversion"
	"gorm.io/gorm"
)

type Order struct {
	Name  string // 对应 map 的 key
	Query string // SQL 查询模板
}

func ApplyOrder(query *gorm.DB, paging *requestConversion.PagingSchema, sqlOrders []Order) *gorm.DB {
	for _, order := range sqlOrders {
		if paging.Sort == order.Name {
			query = query.Order(order.Query + " " + paging.SortOrder)
		}
	}
	return query
}
