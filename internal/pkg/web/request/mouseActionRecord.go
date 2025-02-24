package request

import log "github.com/sirupsen/logrus"

type MouseActionRecordReq struct {
	PagingReq
	MouseActionName string `form:"mouse_action_name" binding:"mouseActionValidations"`
	MinX            int64  `form:"min_x" binding:"rangeValidations"`
	MaxX            int64  `form:"max_x" binding:"rangeValidations"`
	MinY            int64  `form:"min_y" binding:"rangeValidations"`
	MaxY            int64  `form:"max_y" binding:"rangeValidations"`
	MaxDate         string `form:"max_date" binding:"datetimeRangeValidations" time_format:"2006-01-02"`
	MaxTime         string `form:"max_time" binding:"datetimeRangeValidations" time_format:"15:04:05"`
	MinDate         string `form:"min_date" binding:"datetimeRangeValidations" time_format:"2006-01-02"`
	MinTime         string `form:"min_time" binding:"datetimeRangeValidations" time_format:"15:04:05"`
}

func (m MouseActionRecordReq) GetAllDateTime() (string, string, string, string) {
	return m.MaxDate, m.MaxTime, m.MinDate, m.MinTime
}
func (m MouseActionRecordReq) GetMaxMinValue(field string) (int64, int64) {
	log.Info("GetMaxMinValue", field)
	switch field {
	case "MaxX", "MinX":
		return m.MaxX, m.MinX
	case "MaxY", "MinY":
		return m.MaxY, m.MinY
	default:
		return 0, 0
	}
}
