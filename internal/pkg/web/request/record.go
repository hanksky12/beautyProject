package request

type RecordReq struct {
	PagingReq
	Hardware   string  `form:"hardware_name" binding:"hardwareValidations"`
	MinPercent float64 `form:"min_percent" binding:"percentRangeValidations"`
	MaxPercent float64 `form:"max_percent" binding:"percentRangeValidations"`

	MaxDate string `form:"max_date" binding:"datetimeRangeValidations" time_format:"2006-01-02"`
	MaxTime string `form:"max_time" binding:"datetimeRangeValidations" time_format:"15:04:05"`
	MinDate string `form:"min_date" binding:"datetimeRangeValidations" time_format:"2006-01-02"`
	MinTime string `form:"min_time" binding:"datetimeRangeValidations" time_format:"15:04:05"`
}
