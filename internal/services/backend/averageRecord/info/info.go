package info

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/repository"
	structUtil "beautyProject/internal/pkg/util/struct"
	timeUtil "beautyProject/internal/pkg/util/time"
	"beautyProject/internal/pkg/util/web/requestConversion"
	log "github.com/sirupsen/logrus"
)

type AverageRecordInfo struct {
	UserId     uint64
	Repo       *repository.StatusRecordAverage
	Conditions map[string]any
	Paging     *requestConversion.PagingSchema
}

func (info *AverageRecordInfo) Query() dto.Table {
	//log.Info(info.Conditions)
	//log.Info(info.Paging)
	info.combineKeyDateTime()
	result, total, err := info.Repo.FindByUserId(info.UserId, info.Conditions, info.Paging)
	if err != nil {
		return dto.Table{Success: false, Message: "資料查詢失敗"}
	}
	dataArray := structUtil.ToDataArray(result,
		func(value repository.FindByUserId) map[string]any {
			return map[string]any{
				"id":            value.ID,
				"time":          timeUtil.GetFormatTime(value.Time, timeUtil.GetDefaultParams()),
				"percent":       value.Percent,
				"hardware_name": value.HardwareName}
		})
	log.Info(total)
	return dto.Table{Success: true, Message: "查詢成功", DataArray: dataArray, Total: total}
}

func (info *AverageRecordInfo) combineKeyDateTime() {
	type Key struct {
		Date     string
		Time     string
		DateTime string
	}
	keyArray := [2]Key{
		{Date: "MaxDate", Time: "MaxTime", DateTime: "MaxDateTime"},
		{Date: "MinDate", Time: "MinTime", DateTime: "MinDateTime"}}
	for _, key := range keyArray {
		//驗證rule1 日期與時間必同時存在
		if _, exists := info.Conditions[key.Date]; exists {
			strDate := info.Conditions[key.Date].(string) + " " + info.Conditions[key.Time].(string)
			info.Conditions[key.DateTime] = timeUtil.GetTimeStamp(strDate, timeUtil.GetDefaultParams())
			delete(info.Conditions, key.Date)
			delete(info.Conditions, key.Time)
		}
	}
	//log.Info(info.Conditions)
}
