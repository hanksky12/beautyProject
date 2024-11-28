package controller

import log "github.com/sirupsen/logrus"

type Callback struct {
	Minutes int
}

func (c *Callback) CpuAnalyzeTask(key string, value string, headers map[string]string) {
	log.Info(c.Minutes)
	log.Info(key, value)
	log.Info(headers)

}

func (c *Callback) DiskAnalyzeTask(key string, value string, headers map[string]string) {
	log.Info(c.Minutes)
	log.Info(key, value)
	log.Info(headers)
}

func (c *Callback) MemAnalyzeTask(key string, value string, headers map[string]string) {
	log.Info(c.Minutes)
	log.Info(key, value)
	log.Info(headers)
}
