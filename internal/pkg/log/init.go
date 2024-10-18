package log

import (
	"bytes"
	"fmt"
	"github.com/cihub/seelog"
	"os"
	"text/template"
)

func Init(filePath string) {
	minLevel := getMinLevel()
	appConfig, err := loadSeeLogConfig(minLevel, filePath)
	if err != nil {
		fmt.Println("Failed to load logger configuration:", err)
		return
	}
	logger, err2 := seelog.LoggerFromConfigAsBytes(appConfig)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	err = seelog.ReplaceLogger(logger)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Log Init Success")
}
func getMinLevel() string {
	env := os.Getenv("GO_ENV")
	switch env {
	case "development":
		return "debug"
	case "test", "production":
		return "info"
	default:
		return "info" // 默認使用 info 級別
	}
}

func loadSeeLogConfig(minLevel, filePath string) ([]byte, error) {
	tmpl, err := template.ParseFiles("internal/pkg/log/seelog.xml.template")
	if err != nil {
		return nil, err
	}

	// 使用模板生成最終的 Seelog 配置
	var configBuffer bytes.Buffer
	err = tmpl.Execute(&configBuffer,
		struct {
			MinLevel string
			FilePath string
		}{MinLevel: minLevel, FilePath: filePath})
	if err != nil {
		return nil, err
	}
	return configBuffer.Bytes(), nil
}
