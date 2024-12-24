package request

type RecordReq struct {
	PagingReq
	Hardware   string  `json:"hardware" binding:"hardwareEnum"`
	MinTime    int64   `json:"min_time" `
	MaxTime    int64   `json:"max_time" `
	MinPercent float64 `json:"min_percent" `
	MaxPercent float64 `json:"max_percent" `
}
