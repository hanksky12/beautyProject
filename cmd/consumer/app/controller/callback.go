package controller

import (
	"beautyProject/internal/pkg/enum/hardware"
	"beautyProject/internal/pkg/enum/mouseAction"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/services/consumer/server"
	"beautyProject/internal/services/consumer/user"
	log "github.com/sirupsen/logrus"
)

type Callback struct {
	IsRaw bool
}

func (c *Callback) RecordHardwareTask(hardware *hardware.Hardware, key string, value string, headers map[string]string) {
	log.Info("RecordHardwareTask")
	log.Info(hardware.Name)
	repo := &repository.HardwareStatusRecord{}
	repoRaw := &repository.HardwareStatusRecordRaw{}
	pc := server.Hardware{RecordRepo: repo, RecordRepoRaw: repoRaw}
	pc.Record(hardware, value, headers, c.IsRaw)
}

func (c *Callback) RecordMouseTask(action *mouseAction.MouseAction, key string, value string, headers map[string]string) {
	log.Info("RecordMouseTask")
	log.Info(action.Name)
	repoRaw := &repository.MouseActionStatusRecordRaw{}
	act := user.MouseAction{RecordRepoRaw: repoRaw}
	act.Record(action, value, headers)
}
