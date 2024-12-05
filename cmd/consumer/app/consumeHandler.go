package app

import (
	"beautyProject/cmd/consumer/app/controller"
	"beautyProject/internal/pkg/enum"
	log "github.com/sirupsen/logrus"
)

type ConsumeHandler struct {
	handle func(string, string, map[string]string)
}

func (c *ConsumeHandler) SetHandle(topic string, callback controller.Callback) {
	log.Infof("Set handle for topic: %s", topic)
	switch topic {
	case enum.Cpu.Name:
		c.handle = callback.CpuRecord
	case enum.Disk.Name:
		c.handle = callback.DiskRecord
	case enum.Memory.Name:
		c.handle = callback.MemRecord
	default:
		c.handle = nil
	}
}

func (c *ConsumeHandler) Handle(key string, value string, headers map[string]string) {
	c.handle(key, value, headers)
}
