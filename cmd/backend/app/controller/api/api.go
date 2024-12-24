package api

import (
	"beautyProject/cmd/backend/app/controller/api/hardware"
	"beautyProject/cmd/backend/app/controller/api/record"
	"beautyProject/cmd/backend/app/controller/api/user"
	"github.com/gin-gonic/gin"
)

func AddApiRoutes(rg *gin.RouterGroup) {

	userRG := rg.Group("/user")
	user.AddRoutes(userRG)

	hardwareRG := rg.Group("/hardware")
	hardware.AddRoutes(hardwareRG)

	recordRG := rg.Group("/record")
	record.AddRoutes(recordRG)
}
