package main

import (
	"beautyProject/cmd/backend/router"
	"beautyProject/internal/pkg/config"
	"beautyProject/internal/pkg/db/sql"
	logCustom "beautyProject/internal/pkg/log"
	log "github.com/cihub/seelog"
	env "github.com/spf13/viper"
)

func main() {
	config.Init()
	logCustom.Init("log/data.log")
	envMysql := env.GetStringMapString("MySQL")
	sql.Init(
		envMysql["host"],
		envMysql["user"],
		envMysql["password"],
		envMysql["db_name"],
		envMysql["port"],
	)
	defer log.Flush()
	router.Run()
	//atk := env.GetStringMap("MySQL")
	//fmt.Println(atk["host"])
	//
	//log.Debug("MySQL")
	//log.Info("MySQL")
}
