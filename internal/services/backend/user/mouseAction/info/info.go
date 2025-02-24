package info

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/enum/mouseAction"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/repository"
	structUtil "beautyProject/internal/pkg/util/struct"
)

type ActionInfo struct {
	Repo *repository.MouseAction
}

func (a *ActionInfo) Query() dto.Table {
	result, err := a.Repo.FindAll()
	if err != nil {
		return dto.Table{Success: false, Message: "資料查詢失敗"}
	}
	dataArray := structUtil.ToDataArray(result,
		func(value model.MouseAction) map[string]any {
			return map[string]any{
				"id":                value.ID,
				"mouse_action_name": enum.GetChineseNameByName(value.Name, mouseAction.Map)}
		})
	return dto.Table{Success: true, Message: "查詢成功", DataArray: dataArray, Total: len(result)}
}
