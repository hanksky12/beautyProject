package controller

import (
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/services/cron/server"
	log "github.com/sirupsen/logrus"
)

type Job struct{}

func (j *Job) AnalyzeTask() {
	recordRepo := &repository.HardwareStatusRecord{}
	recordAverageRepo := &repository.HardwareStatusRecordAverage{}
	recordQueryRepo := &repository.HardwareStatusRecordQuery{}
	h := server.Hardware{
		AverageRow:        5,
		RecordRepo:        recordRepo,
		RecordAverageRepo: recordAverageRepo,
		RecordQueryRepo:   recordQueryRepo,
	}
	h.Analyze()
}

func (j *Job) Test() {
	log.Info("Test:")
}
