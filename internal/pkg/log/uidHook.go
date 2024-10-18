package log

import (
	"beautyProject/internal/pkg/str"
	log "github.com/sirupsen/logrus"
)

type UidHook struct {
}

// 设置所有的日志等级都走这个钩子
func (hook *UidHook) Levels() []log.Level {
	return log.AllLevels
}

// 修改其中的数据，或者进行其他操作
func (hook *UidHook) Fire(entry *log.Entry) error {
	entry.Data["request_id"] = str.GetNewUid()
	return nil
}
