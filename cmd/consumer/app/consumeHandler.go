package app

import "beautyProject/cmd/consumer/app/controller"

type ConsumeHandler struct {
	handle func(string, string, map[string]string)
}

func (c ConsumeHandler) SetHandle(topic string, callback controller.Callback) {
	switch topic {
	case "cpu":
		c.handle = callback.CpuAnalyzeTask
	case "disk":
		c.handle = callback.DiskAnalyzeTask
	case "mem":
		c.handle = callback.MemAnalyzeTask
	default:
		c.handle = nil
	}
}

func (c ConsumeHandler) Handle(key string, value string, headers map[string]string) {
	c.handle(key, value, headers)
}
