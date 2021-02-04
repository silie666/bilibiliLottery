// mysql db drives
package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	"time"
	"bilibili/config"
)

// query need rows.Close to release db ins
// exec will release automatic
var Db *gorm.DB
var DbErr error

func init() {
	// get db config
	dbConfig := config.GetDbConfig()

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
	)
	Db, DbErr = gorm.Open(mysql.Open(dbDSN),&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: dbConfig["DB_PREFIX"].(string),
			SingularTable: true,
		},
	})
	sqlDB,_ := Db.DB()
	maxOpenConns,_ :=strconv.Atoi(dbConfig["DB_MAX_OPEN_CONNS"].(string))
	maxIdleConns,_ :=strconv.Atoi(dbConfig["DB_MAX_IDLE_CONNS"].(string))
	connMaxLifetime := dbConfig["DB_MAX_LIFETIME_CONNS"].(time.Duration)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)


	if DbErr != nil {
		panic("database data source name error: " + DbErr.Error())
	}
}
