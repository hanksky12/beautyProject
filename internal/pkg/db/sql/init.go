package sql

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	env "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func Init() {
	envMysql := env.GetStringMapString("MySQL")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		envMysql["user"],
		envMysql["password"],
		envMysql["host"],
		envMysql["port"],
		envMysql["db_name"],
	)
	log.Info("dsn: ", dsn)
	DB = connectDB(dsn)
}

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Info("Sql Init Success")
	return db
}
