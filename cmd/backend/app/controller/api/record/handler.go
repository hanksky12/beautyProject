package record

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/util/web/requestConversion"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/averageRecord"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler struct{}

func (h *Handler) Record(c *gin.Context, req request.RecordReq) {
	conditions, paging, err := requestConversion.Map(req)
	if err != nil {
		log.Info("Error:", err)
		tableDto := dto.Table{Success: false, Message: "轉換失敗"}
		response.ProcessTableDto(c, tableDto)
		return
	}
	userId := uint64(c.GetUint("userId"))
	repo := repository.StatusRecordAverage{}
	r := averageRecord.AverageRecord{
		UserId:     userId,
		Repo:       &repo,
		Conditions: conditions,
		Paging:     paging,
	}
	tableDto := r.Query()
	response.ProcessTableDto(c, tableDto)
}
