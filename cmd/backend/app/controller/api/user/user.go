package user

import (
	"beautyProject/cmd/backend/app/controller/api/user/authentication"
	"beautyProject/cmd/backend/app/controller/api/user/mouse"
	"beautyProject/cmd/backend/app/controller/api/user/record"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	authenticationRG := rg.Group("/authentication")
	authentication.AddRoutes(authenticationRG)

	mouseRG := rg.Group("/mouse")
	mouse.AddRoutes(mouseRG)

	recordRG := rg.Group("/record")
	record.AddRoutes(recordRG)
}
