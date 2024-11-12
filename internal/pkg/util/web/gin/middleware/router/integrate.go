package router

import (
	"github.com/gin-gonic/gin"
)

func Integrate[T any](handlerFunc func(*gin.Context, T), isAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer respRecover(c)
		if isAuth && !Auth(c) {
			return
		}
		validator(handlerFunc)(c)
	}
}
