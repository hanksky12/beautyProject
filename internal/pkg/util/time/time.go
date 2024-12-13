package time

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func TrackExecutionTime(name string, start time.Time) {
	duration := time.Since(start).Seconds()
	log.Infof("%s 執行時間: %v 秒", name, duration)
}
