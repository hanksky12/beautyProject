package requestConversion

import (
	"beautyProject/internal/pkg/web/request"
	"fmt"
	"reflect"
)

func Map(src any) (map[string]any, *PagingSchema, error) {
	conditions := make(map[string]any)
	paging := &PagingSchema{}

	srcVal := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)

	// Dereference if src is a pointer
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
		srcType = srcType.Elem()
	}

	// Ensure src is a struct
	if srcVal.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("source must be a struct or a pointer to struct")
	}

	// Iterate through the fields of the struct
	for i := 0; i < srcVal.NumField(); i++ {
		field := srcType.Field(i)
		fieldName := field.Name
		fieldValue := srcVal.Field(i)

		// Skip unexported fields
		if !fieldValue.CanInterface() {
			continue
		}

		// Check if the field belongs to PagingReq
		if field.Anonymous && field.Type == reflect.TypeOf(request.PagingReq{}) {
			// Extract PagingReq fields into PagingSchema
			for j := 0; j < fieldValue.NumField(); j++ {
				subField := field.Type.Field(j)
				subFieldName := subField.Name
				subFieldValue := fieldValue.Field(j).Interface()

				switch subFieldName {
				case "PerPage":
					paging.PerPage = subFieldValue.(int)
				case "Page":
					paging.Page = subFieldValue.(int)
				case "Sort":
					paging.Sort = subFieldValue.(string)
				case "SortOrder":
					paging.SortOrder = subFieldValue.(string)
				}
			}
		} else {
			// Add non-PagingReq fields to conditions
			conditions[fieldName] = fieldValue.Interface()
		}
	}

	return conditions, paging, nil
}
