package server

import (
	"beautyProject/cmd/backend/app/controller/api/server/hardware"
	"beautyProject/cmd/backend/app/controller/api/server/hardwareSpecs"
	"beautyProject/cmd/backend/app/controller/api/server/record"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	hardwareRG := rg.Group("/hardware")
	hardware.AddRoutes(hardwareRG)

	hardwareSpecsRG := rg.Group("/hardware-specs")
	hardwareSpecs.AddRoutes(hardwareSpecsRG)

	recordRG := rg.Group("/record")
	record.AddRoutes(recordRG)
}
