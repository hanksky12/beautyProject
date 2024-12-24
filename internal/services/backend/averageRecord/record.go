package averageRecord

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/pkg/util/web/requestConversion"
	log "github.com/sirupsen/logrus"
)

type AverageRecord struct {
	UserId     uint64
	Repo       *repository.StatusRecordAverage
	Conditions map[string]any
	Paging     *requestConversion.PagingSchema
}

func (r *AverageRecord) Query() dto.Table {
	result, err := r.Repo.FindByUserId(r.UserId, r.Conditions, r.Paging)
	if err != nil {
		return dto.Table{}
	}
	log.Info(result)

	// table column 硬體名稱, 時間, percent, 建立時間
	return dto.Table{Success: true, Message: "查詢成功", DataList: []map[string]any{}, Total: 0}

}
