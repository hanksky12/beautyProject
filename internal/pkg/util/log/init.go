package log

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

/*
ex:
log.WithFields(log.Fields{
  "event": event,
  "topic": topic,
  "key": key,
}).Fatal("Failed to send event")
*/

func Init(filePath string) {
	env := os.Getenv("GO_ENV")
	setLevel(env)
	setFormatter(env)
	setOutput(filePath)
	log.SetReportCaller(true)
	log.Info("Log Init Success")
}

func setOutput(filePath string) {
	writer1 := os.Stdout
	writer2, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	log.SetOutput(io.MultiWriter(writer1, writer2))

}

func setFormatter(env string) {
	switch env {
	case "development":
		log.SetFormatter(&log.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) { //自定义Caller的返回
				fileName := path.Base(frame.File)
				return frame.Function, fileName
			}})
	default:
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FieldMap: log.FieldMap{
				log.FieldKeyTime:  "[time]",
				log.FieldKeyLevel: "[lv]",
				log.FieldKeyMsg:   "message",
				log.FieldKeyFunc:  "[caller]",
				log.FieldKeyFile:  "[file]",
			},
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) { //自定义Caller的返回
				fileName := path.Base(frame.File)
				return frame.Function, fileName
			},
			PrettyPrint: true})
	}
}

func setLevel(env string) {
	switch env {
	case "development":
		log.SetLevel(log.DebugLevel)
	case "test", "production":
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.InfoLevel) // 默認使用 info 級別
	}
}
