package models

import (
	"errors"
	"fmt"
	"loginv/server/gormdb"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       int `gorm:"primary_key"`
	Username string
	Pwd      string
}

func ValidateUser(user User) error {
	var u User
	fmt.Println(gormdb.GORDB)
	fmt.Println(user)
	db := gormdb.GORDB.First(&u, "username=? and pwd=?", user.Username, user.Pwd)

	if db.Error != nil {
		return db.Error
	}

	if u.Username == "" {
		return errors.New("用户名和密码错误")
	}

	// GORDB.First(&user, "code = ?", "L1212")
	// var u User
	// o.Raw("select * from user where username=? and pwd=?", user.Username, user.Pwd).QueryRow(&u)

	return nil
}

func SaveUser(user User) error {
	var db *gorm.DB
	fmt.Println(user)
	db = gormdb.GORDB.Save(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
