package router

import (
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/pkg/web/request/validation/base"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	validatorV10 "github.com/go-playground/validator/v10"
	_ "net/http"
	"reflect"
)

/*
Validator 是一個通用的驗證中間件，它會將綁定好的數據直接傳給處理器函數
*/
func validator[T any](handlerFunc func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData T
		// 如果不是 request.EmptyReq 才進行綁定檢查
		if reflect.TypeOf(reqData) != reflect.TypeOf(request.EmptyReq{}) {
			// 用c.ShouldBind 就不用區分是GET還是POST
			if err := c.ShouldBind(&reqData); err != nil {
				var errorStr string
				var validationErrs validatorV10.ValidationErrors
				switch {
				case errors.As(err, &validationErrs):
					errorStr = base.ParseValidationErrors(validationErrs)
				default:
					errorStr = err.Error()
				}
				response.Error(c, fmt.Sprintf("驗證錯誤=> %v", errorStr))
				c.Abort()
				return
			}
		}
		handlerFunc(c, reqData)
	}
}
