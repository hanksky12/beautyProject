package router

import (
	"beautyProject/cmd/backend/router/api"
	"beautyProject/cmd/backend/router/index"
	mid "beautyProject/internal/pkg/web/gin/middleware"
	"github.com/gin-gonic/gin"
)

var router = gin.New()

func Run() {
	addMiddleware()
	addRouter()
	router.Run(":8080")
}

func addMiddleware() {
	mid.ContextLogger(router)
	mid.AddLog(router)
	//mid.RequestID(router)
}

func addRouter() {
	apiRG := router.Group("/api")
	api.AddApiRoutes(apiRG)

	indexRG := router.Group("/")
	index.AddIndexRoutes(indexRG)
}
