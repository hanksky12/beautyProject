package controller

import (
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/services/consumer/pc"
	log "github.com/sirupsen/logrus"
)

type Callback struct {
	Minutes int
}

func (c *Callback) CpuRecord(key string, value string, headers map[string]string) {
	c.RecordTask(enum.Cpu, key, value, headers)
}

func (c *Callback) DiskRecord(key string, value string, headers map[string]string) {
	c.RecordTask(enum.Disk, key, value, headers)
}

func (c *Callback) MemRecord(key string, value string, headers map[string]string) {
	c.RecordTask(enum.Memory, key, value, headers)
}

func (c *Callback) RecordTask(hardware *enum.Hardware, key string, value string, headers map[string]string) {
	log.Info("RecordTask")
	log.Info(hardware.Name)
	repo := repository.StatusRecord{}
	pc := pc.Hardware{}
	pc.Record(hardware, key, value, headers, c.Minutes, repo)
}
