package validation

import (
	"beautyProject/internal/pkg/interfaces"
	"github.com/go-playground/validator/v10"
)

func RangeValidations(fl validator.FieldLevel) bool {
	parent := fl.Parent().Interface()
	req, ok := parent.(interfaces.IRangeValidatable)
	if !ok {
		return false // 结构体未实现接口，校验失败
	}
	field := fl.FieldName() // 获取当前字段名
	maxValue, minValue := req.GetMaxMinValue(field)
	if minValue == 0 || maxValue == 0 {
		return true
	}

	return minValue <= maxValue
}
