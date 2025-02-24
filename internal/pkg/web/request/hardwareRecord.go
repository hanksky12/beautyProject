package request

type HardwareRecordReq struct {
	PagingReq
	Hardware   string  `form:"hardware_name" binding:"hardwareValidations"`
	MinPercent float64 `form:"min_percent" binding:"rangeValidations"`
	MaxPercent float64 `form:"max_percent" binding:"rangeValidations"`
	MaxDate    string  `form:"max_date" binding:"datetimeRangeValidations" time_format:"2006-01-02"`
	MaxTime    string  `form:"max_time" binding:"datetimeRangeValidations" time_format:"15:04:05"`
	MinDate    string  `form:"min_date" binding:"datetimeRangeValidations" time_format:"2006-01-02"`
	MinTime    string  `form:"min_time" binding:"datetimeRangeValidations" time_format:"15:04:05"`
}

func (h HardwareRecordReq) GetAllDateTime() (string, string, string, string) {
	return h.MaxDate, h.MaxTime, h.MinDate, h.MinTime
}
func (h HardwareRecordReq) GetMaxMinValue(field string) (int64, int64) {
	return int64(h.MaxPercent * 100), int64(h.MinPercent * 100)
}
