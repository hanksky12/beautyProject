package main

import (
	"beautyProject/cmd/backend/app"
	"beautyProject/internal/pkg/util/app/initAfterApp"
	"beautyProject/internal/pkg/util/app/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/router.log", true)
	app.Run()
	initAfterApp.Run(true)
}

//sql.Db.AutoMigrate(&model.User{}) //todo delete
