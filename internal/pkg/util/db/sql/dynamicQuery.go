package sql

import (
	"beautyProject/internal/pkg/util/web/requestConversion"
	"gorm.io/gorm"
)

func DynamicQuery[T any](query *gorm.DB, paging *requestConversion.PagingSchema) ([]T, int, error) {
	var results []T
	totalQuery := query
	err := totalQuery.Find(&results).Error
	if err != nil {
		return nil, 0, err
	}
	total := len(results)

	if paging != nil {
		query = query.Offset((paging.Page - 1) * paging.PerPage).Limit(paging.PerPage)
	}
	//stmt := query.Find(&results).Statement
	//log.Info(stmt.SQL.String())
	//log.Info(stmt.Vars)

	err = query.Find(&results).Error
	if err != nil {
		return nil, 0, err
	}
	return results, total, err
}
