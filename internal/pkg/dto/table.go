package dto

type Table struct {
	Success   bool
	Message   string
	DataArray []map[string]any
	Total     int
}
