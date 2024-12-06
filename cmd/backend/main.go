package main

import (
	"beautyProject/cmd/backend/app"
	"beautyProject/internal/pkg/util/app/initApp/initAfterApp"
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/router.log", true)
	app.Run()
	initAfterApp.Run(true)
}

//sql.Db.AutoMigrate(&model.User{}) //todo delete
