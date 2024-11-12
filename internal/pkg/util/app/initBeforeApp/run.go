package initBeforeApp

import (
	"beautyProject/internal/pkg/util/config"
	"beautyProject/internal/pkg/util/db/sql"
	"beautyProject/internal/pkg/util/log"
)

func Run(filePath string) {
	config.Init()
	log.Init(filePath)
	sql.Init()
}
