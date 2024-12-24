package dto

type Table struct {
	Success  bool
	Message  string
	DataList []map[string]any
	Total    int
}
