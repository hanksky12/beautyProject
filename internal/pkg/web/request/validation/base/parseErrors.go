package base

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func ParseValidationErrors(errs validator.ValidationErrors) string {
	// 自定义错误消息
	var errStr string
	for _, err := range errs {
		log.Info(err.Tag())
		switch err.Tag() {
		case "percentRangeValidations":
			errStr = "運作率最小值不可大于最大值"
			break

		case "datetimeRangeValidations":
			errStr = "日期與時間需填寫(可只填最大or最小)，最晚日期時間不可大于最早日期時間"
			break
		default:
			errStr = err.Error()
			break
		}
	}
	return errStr
}
