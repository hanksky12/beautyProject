package app

import (
	"beautyProject/cmd/backend/app/controller/api"
	"beautyProject/internal/pkg/util/web/gin/middleware"
	"beautyProject/internal/pkg/web/request"
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
}

func addMiddleware(router *gin.Engine) {
	middleware.ContextLogger(router)
	middleware.AddLog(router)
}

func addRouter(router *gin.Engine) {
	apiRG := router.Group("/api")
	api.AddApiRoutes(apiRG)
}

func registerValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 统一注册所有自定义验证器
		if err := request.RegisterCustomValidations(v); err != nil {
			panic(err)
		}
	}
}
