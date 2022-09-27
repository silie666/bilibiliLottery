// mysql db drives
package mysql

import (
	"bilibili/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB
var DbErr error

func init() {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		config.Env.GetString("mysql.user"),
		config.Env.GetString("mysql.pwd"),
		config.Env.GetString("mysql.host"),
		config.Env.GetString("mysql.port"),
		config.Env.GetString("mysql.name"),
		config.Env.GetString("mysql.charset"),
	)
	Db, DbErr = gorm.Open(mysql.Open(dbDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Env.GetString("mysql.prefix"),
			SingularTable: true,
		},
	})
	sqlDB, _ := Db.DB()
	sqlDB.SetMaxOpenConns(config.Env.GetInt("mysql.max_open_connections"))
	sqlDB.SetMaxIdleConns(config.Env.GetInt("mysql.max_idle_connections"))
	sqlDB.SetConnMaxLifetime(config.Env.GetDuration("mysql.max_lifetime_connections"))

	if DbErr != nil {
		panic("database data source name error: " + DbErr.Error())
	}
}
