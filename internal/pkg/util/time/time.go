package time

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type DefaultParams struct {
	Layout      string
	LocationEtc string
}

func GetDefaultParams() DefaultParams {
	return DefaultParams{
		Layout:      "2006-01-02 15:04:05",
		LocationEtc: "Etc/GMT-8", // 和直覺相反=> GMT+8
	}
}

func TrackExecutionTime(name string, start time.Time) {
	duration := time.Since(start).Seconds()
	log.Infof("%s 執行時間: %v 秒", name, duration)
}

func GetNowTimeStamp() int64 {
	param := GetDefaultParams()
	location, err := time.LoadLocation(param.LocationEtc)
	if err != nil {
		return 0
	}
	now := time.Now().In(location)
	return now.Unix()
}

func GetFormatTime(timestamp int64, param DefaultParams) string {
	location, err := time.LoadLocation(param.LocationEtc)
	if err != nil {
		return fmt.Sprintf("invalid location: %s", err)
	}
	return time.Unix(timestamp, 0).In(location).Format(param.Layout)
}

func GetTimeStamp(timeStr string, param DefaultParams) int64 {
	location, err := time.LoadLocation(param.LocationEtc)
	if err != nil {
		fmt.Println("Error loading location:", err)
		return 0
	}
	t, err := time.ParseInLocation(param.Layout, timeStr, location)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}
	return t.Unix()

}

type key struct {
	Date     string
	Time     string
	DateTime string
}

var keyArray = [2]key{
	{Date: "MaxDate", Time: "MaxTime", DateTime: "MaxDateTime"},
	{Date: "MinDate", Time: "MinTime", DateTime: "MinDateTime"},
}

func CombineKeyDateTime(conditions map[string]any) {
	for _, key := range keyArray {
		if _, exists := conditions[key.Date]; exists {
			strDate := conditions[key.Date].(string) + " " + conditions[key.Time].(string)
			conditions[key.DateTime] = GetTimeStamp(strDate, GetDefaultParams())
			delete(conditions, key.Date)
			delete(conditions, key.Time)
		}
	}
}
