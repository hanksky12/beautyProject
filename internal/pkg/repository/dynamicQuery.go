package repository

import (
	"beautyProject/internal/pkg/util/web/requestConversion"
	"gorm.io/gorm"
)

func DynamicQuery(cond map[string]any, query *gorm.DB, paging *requestConversion.PagingSchema) ([]map[string]any, int, error) {
	for key, value := range cond {
		query = query.Where(key+" = ?", value)
	}
	var results []map[string]any

	totalQuery := query
	err := totalQuery.Find(&results).Error
	if err != nil {
		return nil, 0, err
	}
	total := len(results)

	if paging != nil {
		query = query.Offset((paging.Page - 1) * paging.PerPage).Limit(paging.PerPage)
		query = query.Order(paging.Sort + " " + paging.SortOrder)
	}
	err = query.Find(&results).Error
	if err != nil {
		return nil, 0, err
	}
	return results, total, err
}
