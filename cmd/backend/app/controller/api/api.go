package api

import (
	"beautyProject/cmd/backend/app/controller/api/user"
	"github.com/gin-gonic/gin"
)

func AddApiRoutes(rg *gin.RouterGroup) {

	userRG := rg.Group("/user")
	user.AddUserRoutes(userRG)

	//todo 等待發展
	//indexRG := rg.Group("/")
	//index.AddIndexRoutes(indexRG)

}
