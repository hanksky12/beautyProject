package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type GoroutineIDHook struct{}

// Fire 是 Hook 的核心邏輯，用於添加 Goroutine ID 到日誌中
func (hook *GoroutineIDHook) Fire(entry *log.Entry) error {
	entry.Data["go_id"] = getGoroutineID()
	return nil
}

// Levels 返回 Hook 適用的日誌級別
func (hook *GoroutineIDHook) Levels() []log.Level {
	return log.AllLevels
}

// getGoroutineID 獲取當前 Goroutine 的 ID
func getGoroutineID() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stackLine := strings.Split(string(buf[:n]), "\n")[0]
	var gID string
	fmt.Sscanf(stackLine, "goroutine %s ", &gID)
	return gID
}
