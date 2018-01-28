package gormdb

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //init mysql dia
)

var (
	GORDB *gorm.DB //全局变量不能用：= 否则会被覆盖范围
)

func init() {
	connStr := "root:123@tcp(localhost:3306)/login?charset=utf8"
	var err error
	GORDB, err = gorm.Open("mysql", connStr)
	if err != nil {
		defer GORDB.Close()
		panic(err)
	}
	GORDB.SingularTable(true)
	log.Printf("mysql successful", GORDB)
}

func Close() {
	GORDB.Close()
}
