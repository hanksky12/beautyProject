package controller

import (
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/services/cron/pc"
	log "github.com/sirupsen/logrus"
)

type Job struct{}

func (j *Job) AnalyzeTask() {
	userRepo := &repository.User{}
	recordRepo := &repository.StatusRecord{}
	recordAverageRepo := &repository.StatusRecordAverage{}
	h := pc.Hardware{
		AverageRow:        5,
		RecordRepo:        recordRepo,
		RecordAverageRepo: recordAverageRepo,
		UserRepo:          userRepo}
	h.Analyze()
}

func (j *Job) Test(test string) {
	status := enum.Cpu
	log.Info("Status:", status)
	log.Info("Status Name:", status.Name)
	log.Info("Status Number:", status.Number)
}
