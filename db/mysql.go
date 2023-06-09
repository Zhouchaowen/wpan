package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"wpan/config"
)

var DB *gorm.DB

func InitDB(conf config.Mysql) {
	var err error
	dbParams := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
	)
	DB, err = gorm.Open("mysql", dbParams)
	if err != nil {
		log.Fatal(2, err)
	}

	// 全局禁用表名复数
	DB.SingularTable(true)
	DB.LogMode(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	fmt.Println("database init on port ", conf.Host)
}
