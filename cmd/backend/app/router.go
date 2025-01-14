package app

import (
	"beautyProject/cmd/backend/app/controller/api"
	"beautyProject/internal/pkg/util/web/gin/middleware"
	"beautyProject/internal/pkg/web/request/validation/base"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Run() {
	router := gin.New()
	addMiddleware(router)
	addRouter(router)
	registerValidator()
	router.Run(":8070")
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
}

func addMiddleware(router *gin.Engine) {
	middleware.ContextLogger(router)
	middleware.AddLog(router)
	middleware.AddCors(router)
}

func addRouter(router *gin.Engine) {
	apiRG := router.Group("/api")
	api.AddApiRoutes(apiRG)
}

func registerValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 统一注册所有自定义验证器
		if err := base.RegisterCustomValidations(v); err != nil {
			panic(err)
		}
	}
}
