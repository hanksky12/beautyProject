package sql

import "gorm.io/gorm"

type Cond struct {
	Name  string // 对应 map 的 key
	Type  string // 数据类型，例如 "string", "int64", "float64"
	Query string // SQL 查询模板
}

func ApplyConditions(query *gorm.DB, cond map[string]any, sqlConds []Cond) *gorm.DB {
	for _, sqlCond := range sqlConds {
		if value, exists := cond[sqlCond.Name]; exists {
			switch sqlCond.Type {
			case "string":
				if v, ok := value.(string); ok && v != "" {
					query = query.Where(sqlCond.Query, v)
				}
			case "int64":
				if v, ok := value.(int64); ok && v != 0 {
					query = query.Where(sqlCond.Query, v)
				}
			case "float64":
				if v, ok := value.(float64); ok && v != 0 {
					query = query.Where(sqlCond.Query, v)
				}
			}
		}
	}
	return query
}
