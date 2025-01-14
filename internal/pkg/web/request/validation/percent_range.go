package validation

import (
	"beautyProject/internal/pkg/web/request"
	"github.com/go-playground/validator/v10"
)

func PercentRangeValidations(fl validator.FieldLevel) bool {
	req := fl.Parent().Interface().(request.RecordReq) // 获取父结构体
	if req.MinPercent == 0 || req.MaxPercent == 0 {
		//不用比较
		return true
	}
	return req.MinPercent <= req.MaxPercent
}
