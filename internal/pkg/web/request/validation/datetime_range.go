package validation

import (
	"beautyProject/internal/pkg/web/request"
	"github.com/go-playground/validator/v10"
)

func DatetimeRangeValidations(fl validator.FieldLevel) bool {
	req := fl.Parent().Interface().(request.RecordReq) // 获取父结构体
	//Rule1 有日期就要有時間
	if (req.MinDate == "" && req.MinTime != "") || (req.MinDate != "" && req.MinTime == "") {
		return false
	}
	if (req.MaxDate == "" && req.MaxTime != "") || (req.MaxDate != "" && req.MaxTime == "") {
		return false
	}
	if req.MinDate == "" || req.MaxDate == "" {
		// 只有單邊有日期時間，不用比較最大與最小
		return true
	}
	//Rule2 最大日期時間不可小於最小日期時間
	if req.MinDate > req.MaxDate {
		return false
	}
	if (req.MinDate == req.MaxDate) && (req.MinTime > req.MaxTime) {
		return false
	}
	return true
}

//log.Info("MinDate ", reflect.TypeOf(req.MinDate), ' ', req.MinDate)
//log.Info("MinTime ", reflect.TypeOf(req.MinTime), ' ', req.MinTime)
//log.Info("MaxDate ", reflect.TypeOf(req.MaxDate), ' ', req.MaxDate)
//log.Info("MaxTime ", reflect.TypeOf(req.MaxTime), ' ', req.MaxTime)
