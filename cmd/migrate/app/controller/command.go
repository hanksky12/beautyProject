package controller

import (
	"beautyProject/internal/pkg/util/app/initApp/initBeforeApp"
	"beautyProject/internal/pkg/util/db/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Command struct{}

func (j *Command) Run(command string, stepsOrVersion string) {
	initBeforeApp.Run("log/migrate.log", false)
	driver, err := mysql.WithInstance(sql.SqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create MySQL driver for migrations: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migration",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	switch command {
	case "up":
		log.Info("Up~~~~~~~~")
		err = m.Up()
	case "down":
		log.Info("Down~~~~~~~~")
		err = m.Down()
	case "steps":
		uint64Steps, err := strconv.ParseInt(stepsOrVersion, 10, 64)
		if err != nil {
			log.Info("Failed !")
			panic(err)
		}
		intSteps := int(uint64Steps)
		log.Infof("Steps~~~~~~~~ %v", intSteps)
		err = m.Steps(intSteps) //有bug
	case "goto":
		uint64Version, err := strconv.ParseFloat(stepsOrVersion, 64)
		if err != nil {
			log.Info("Failed !")
			panic(err)
		}
		uintVersion := uint(uint64Version)
		log.Infof("Goto~~~~~~~~ %v", uintVersion)
		err = m.Migrate(uintVersion)
	default:
		log.Info("Nothing !")
	}
	if err != nil {
		log.Info("Failed !")
		panic(err)
	}
	log.Info("Success !")
}
