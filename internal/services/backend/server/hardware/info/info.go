package info

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/enum/hardware"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/repository"
	structUtil "beautyProject/internal/pkg/util/struct"
)

type HardwareInfo struct {
	Repo *repository.Hardware
}

func (r *HardwareInfo) Query() dto.Table {
	result, err := r.Repo.FindAll()
	if err != nil {
		return dto.Table{Success: false, Message: "資料查詢失敗"}
	}
	dataArray := structUtil.ToDataArray(result,
		func(value model.Hardware) map[string]any {
			return map[string]any{
				"id":            value.ID,
				"hardware_name": enum.GetChineseNameByName(value.Name, hardware.Map)}
		})
	return dto.Table{Success: true, Message: "查詢成功", DataArray: dataArray, Total: len(result)}
}
