package main

import (
	"beautyProject/cmd/migrate/app/controller"
	"beautyProject/internal/pkg/util/app/cmd"
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
)

func main() {
	initBeforeApp.Run("log/cmd.log", false)
	cmd.Execute(&controller.Command{})
	//go run ./cmd/migrate/cmd.go -c Run -p up -p stepsOrVersion
}
