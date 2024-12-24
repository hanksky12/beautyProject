package requestConversion

import (
	"fmt"
	"reflect"
)

func Map(src any) (map[string]any, *PagingSchema, error) {
	conditions := make(map[string]any)
	paging := &PagingSchema{}

	srcVal := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
		srcType = srcType.Elem()
	}

	if srcVal.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("source must be a struct or a pointer to struct")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		field := srcType.Field(i)
		fieldName := field.Name
		fieldValue := srcVal.Field(i).Interface()

		// 動態設置 PagingSchema 的字段值
		switch fieldName {
		case "PerPage":
			paging.PerPage = fieldValue.(int)
		case "Page":
			paging.Page = fieldValue.(int)
		case "Sort":
			paging.Sort = fieldValue.(string)
		case "SortOrder":
			paging.SortOrder = fieldValue.(string)
		default:
			conditions[fieldName] = fieldValue
		}
	}
	return conditions, paging, nil
}
