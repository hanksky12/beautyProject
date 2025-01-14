package _struct

func ToDataArray[T any](result []T, fun func(T) map[string]any) []map[string]any {
	dataArray := make([]map[string]any, len(result))
	for i, value := range result {
		dataArray[i] = fun(value)
	}
	return dataArray
}
