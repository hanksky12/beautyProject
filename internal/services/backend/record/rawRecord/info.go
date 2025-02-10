package rawRecord

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/entity"
	"beautyProject/internal/pkg/repository"
	structUtil "beautyProject/internal/pkg/util/struct"
	timeUtil "beautyProject/internal/pkg/util/time"
	"beautyProject/internal/pkg/util/web/requestConversion"
)

type Info struct {
	UserId     uint64
	Repo       *repository.StatusRecordRaw
	Conditions map[string]any
	Paging     *requestConversion.PagingSchema
}

func (info *Info) Query() dto.Table {
	timeUtil.CombineKeyDateTime(info.Conditions)
	result, total, err := info.Repo.FindByUserId(info.UserId, info.Conditions, info.Paging)
	if err != nil {
		return dto.Table{Success: false, Message: "資料查詢失敗"}
	}
	dataArray := structUtil.ToDataArray(result,
		func(value entity.Record) map[string]any {
			return map[string]any{
				"id":            value.ID,
				"time":          timeUtil.GetFormatTime(value.Time, timeUtil.GetDefaultParams()),
				"percent":       value.Percent,
				"hardware_name": value.HardwareName}
		})
	return dto.Table{Success: true, Message: "查詢成功", DataArray: dataArray, Total: total}
}
