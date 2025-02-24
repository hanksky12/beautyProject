package api

import (
	"beautyProject/cmd/backend/app/controller/api/server"
	"beautyProject/cmd/backend/app/controller/api/user"
	"github.com/gin-gonic/gin"
)

func AddApiRoutes(rg *gin.RouterGroup) {
	serverRG := rg.Group("/server")
	server.AddRoutes(serverRG)

	userRG := rg.Group("/user")
	user.AddRoutes(userRG)
}
