package rawRecord

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/entity"
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/enum/mouseAction"
	"beautyProject/internal/pkg/repository"
	structUtil "beautyProject/internal/pkg/util/struct"
	timeUtil "beautyProject/internal/pkg/util/time"
	"beautyProject/internal/pkg/util/web/requestConversion"
)

type Info struct {
	UserId     uint64
	Repo       *repository.MouseActionStatusRecordRaw
	Conditions map[string]any
	Paging     *requestConversion.PagingSchema
}

func (info *Info) Query() dto.Table {
	//log.Info("Query raw record")
	timeUtil.CombineKeyDateTime(info.Conditions)
	//log.Info("conditions:", info.Conditions)
	enum.ReplaceChineseNameToName(info.Conditions, mouseAction.Map)
	//log.Info("conditions:", info.Conditions)
	result, total, err := info.Repo.FindByUserId(info.UserId, info.Conditions, info.Paging)
	if err != nil {
		return dto.Table{Success: false, Message: "資料查詢失敗"}
	}
	//log.Info("result:", result, "total:", total)
	dataArray := structUtil.ToDataArray(result,
		func(value entity.MouseActionRecord) map[string]any {
			return map[string]any{
				"id":                value.ID,
				"time":              timeUtil.GetFormatTime(value.Time, timeUtil.GetDefaultParams()),
				"x":                 value.X,
				"y":                 value.Y,
				"mouse_action_name": enum.GetChineseNameByName(value.MouseActionName, mouseAction.Map)}
		})
	return dto.Table{Success: true, Message: "查詢成功", DataArray: dataArray, Total: total}
}
