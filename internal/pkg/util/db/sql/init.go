package sql

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	env "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB
var SqlDB *sql.DB

func Init() {
	envMysql := env.GetStringMapString("MySQL")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&multiStatements=true&loc=Local",
		envMysql["user"],
		envMysql["password"],
		envMysql["host"],
		envMysql["port"],
		envMysql["db_name"],
	)
	//log.Info("dsn: ", dsn)
	Db = connectDB(dsn)
	log.Info("Sql Init Success")
}

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
			DryRun:         false},
	)
	if err != nil {
		panic(err)
	}
	SqlDB, err = db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	SqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	SqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	SqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
