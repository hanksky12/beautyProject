package app

import (
	"beautyProject/cmd/consumer/app/controller"
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/enum/hardware"
	"beautyProject/internal/pkg/enum/mouseAction"
	log "github.com/sirupsen/logrus"
)

type ConsumeHandler struct {
	Callback *controller.Callback
}

func (c *ConsumeHandler) Handle(key string, value string, headers map[string]string) {
	log.Info("ConsumeHandler Handle")
	topic := headers["topic"]
	switch topic {
	case hardware.Cpu.Name, hardware.Disk.Name, hardware.Memory.Name:
		hw, exist := enum.GetEnumByName(topic, hardware.Map)
		if !exist {
			log.Errorf("Hardware %s not exist", topic)
			return
		}
		c.Callback.RecordHardwareTask(hw, key, value, headers)
	case mouseAction.Click.Name, mouseAction.Move.Name, mouseAction.Scroll.Name:
		action, exist := enum.GetEnumByName(topic, mouseAction.Map)
		if !exist {
			log.Errorf("MouseAction %s not exist", topic)
			return
		}
		c.Callback.RecordMouseTask(action, key, value, headers)
	default:
		log.Errorf("Topic %s not exist", topic)
	}
}
