package app

import (
	"beautyProject/cmd/backend/app/controller/api"
	"beautyProject/cmd/backend/app/controller/index"
	"beautyProject/internal/pkg/util/web/gin/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()
	addMiddleware(router)
	addRouter(router)
	router.Run(":8080")
}

func addMiddleware(router *gin.Engine) {
	middleware.ContextLogger(router)
	middleware.AddLog(router)
}

func addRouter(router *gin.Engine) {
	apiRG := router.Group("/api")
	api.AddApiRoutes(apiRG)

	indexRG := router.Group("/")
	index.AddIndexRoutes(indexRG)
}
