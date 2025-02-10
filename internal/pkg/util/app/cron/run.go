package cron

import (
	"beautyProject/internal/pkg/interfaces"
	"github.com/go-co-op/gocron"
	"os"
	"time"
)

func Run(addJob interfaces.IEnvironment, poolSize int) {
	scheduler := gocron.NewScheduler(time.UTC)

	// 配置並行池大小，控制同時執行的任務數量
	scheduler.SetMaxConcurrentJobs(poolSize, gocron.RescheduleMode)

	// 根據環境配置不同的調度邏輯
	envConfig := os.Getenv("GO_ENV")
	if envConfig == "production" {
		addJob.InProduction(scheduler)
	} else {
		addJob.InOther(scheduler)
	}

	// 啟動調度器
	//scheduler.StartAsync()
	scheduler.StartBlocking()
	// 捕捉停止信號
	select {}
}
