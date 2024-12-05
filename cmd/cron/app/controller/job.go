package controller

import (
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/services/cron/crawler"
	log "github.com/sirupsen/logrus"
)

type Job struct{}

func (j *Job) Test(lotteryType string) {
	status := enum.Cpu
	// 打印枚舉值
	log.Info("Status:", status)
	log.Info("Status Name:", status.Name)
	log.Info("Status Number:", status.Number)

	// 打印枚舉名稱
	//fmt.Println("Status:", status)
	//fmt.Println("Status Name:", status.String())
}

func (j *Job) AnalyzeTask() {
	// 依照 user id  依照時間做平均
	//user => 同樣 min => 放到 1go

	c := crawler.Crawler{}
	c.Crawl("hk")
	//m := map[string]any{
	//	"name":    "backy",
	//	"species": "dog",
	//}
	//mJson, _ := json.Marshal(m)
	//log.Info(string(mJson))
	//log.Infof("Executing another task without parameters.")
}
