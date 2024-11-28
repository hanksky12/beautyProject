package app

import (
	"beautyProject/cmd/cron/app/controller"
	"beautyProject/internal/pkg/util/app/cron"
	"beautyProject/internal/pkg/util/log"
	"github.com/go-co-op/gocron"
)

type AddJob struct{}

func (j *AddJob) InProduction(scheduler *gocron.Scheduler) {
	job := &controller.Job{}

	_, err := scheduler.CronWithSeconds("0 * * * * *").Do(job.UploadImagesToWebServer, "hk")
	log.PrintCron(err)

	// 上傳圖片，每分鐘的第 20 秒和第 50 秒執行
	_, err = scheduler.CronWithSeconds("20,50 * * * * *").Do(job.UploadImagesToWebServer)
	log.PrintCron(err)
}

func (j *AddJob) InOther(scheduler *gocron.Scheduler) {
	job := &controller.Job{}

	// 上傳圖片，每分鐘的第 0 秒執行
	_, err := scheduler.CronWithSeconds("0 * * * * *").Do(job.UploadImagesToWebServer, "hk")
	log.PrintCron(err)

	_, err = scheduler.CronWithSeconds("20,50 * * * * *").Do(job.UploadImagesToWebServer)
	log.PrintCron(err)
}

func Run() {
	job := &AddJob{}
	cron.Run(job, 20)
}
