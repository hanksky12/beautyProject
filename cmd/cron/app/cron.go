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

	_, err := scheduler.CronWithSeconds("20,50 * * * * *").Do(job.AnalyzeTask)
	log.PrintCron(err)
}

func (j *AddJob) InOther(scheduler *gocron.Scheduler) {
	job := &controller.Job{}
	// 每5秒執行一次
	_, err := scheduler.CronWithSeconds("*/5 * * * * *").Do(job.AnalyzeTask)
	log.PrintCron(err)

	//_, err = scheduler.CronWithSeconds("0 * * * * *").Do(job.Test, "test")
	//log.PrintCron(err)

}

func Run() {
	job := &AddJob{}
	cron.Run(job, 20)
}
