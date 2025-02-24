package record

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/pkg/util/web/gin/handler/response"
	"beautyProject/internal/pkg/util/web/requestConversion"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/user/mouseActionRecord/rawRecord"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) RawRecordInfo(c *gin.Context, req request.MouseActionRecordReq) {
	conditions, paging, err := requestConversion.Map(req)
	if err != nil {
		tableDto := dto.Table{Success: false, Message: "轉換失敗"}
		response.ProcessTableDto(c, tableDto)
		return
	}
	userId := uint64(c.GetUint("userId"))
	repo := repository.MouseActionStatusRecordRaw{}
	r := rawRecord.Info{
		UserId:     userId,
		Repo:       &repo,
		Conditions: conditions,
		Paging:     paging,
	}
	tableDto := r.Query()
	response.ProcessTableDto(c, tableDto)
}
