package index

import (
	"github.com/gin-gonic/gin"
)

func AddIndexRoutes(rg *gin.RouterGroup) {
	rg.GET("/index", (&Handler{}).Events)
}
