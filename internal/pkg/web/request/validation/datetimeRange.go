package validation

import (
	"beautyProject/internal/pkg/interfaces"
	"github.com/go-playground/validator/v10"
)

func DatetimeRangeValidations(fl validator.FieldLevel) bool {
	parent := fl.Parent().Interface()
	// 嘗試將 parent 轉換為 DatetimeRangeRequest interface
	req, ok := parent.(interfaces.IDatetimeRangeRequest)
	if !ok {
		return false // 如果無法轉換，則驗證失敗
	}
	maxDate, maxTime, minDate, minTime := req.GetAllDateTime()

	// Rule1: 如果有日期，就必須有時間，反之亦然
	if (minDate == "" && minTime != "") || (minDate != "" && minTime == "") {
		return false
	}
	if (maxDate == "" && maxTime != "") || (maxDate != "" && maxTime == "") {
		return false
	}

	// Rule2: 如果有完整的日期時間，則比較大小
	if minDate == "" || maxDate == "" {
		return true // 單邊有日期時間時，不強制比較大小
	}

	// 最大日期時間不可小於最小日期時間
	if minDate > maxDate {
		return false
	}
	if minDate == maxDate && minTime > maxTime {
		return false
	}

	return true
}
