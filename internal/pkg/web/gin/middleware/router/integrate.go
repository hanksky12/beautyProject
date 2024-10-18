package router

import (
	"github.com/gin-gonic/gin"
)

func Integrate[T any](handlerFunc func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer respRecover(c)
		validator(handlerFunc)(c)
	}
}
