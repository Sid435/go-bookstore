package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
Creating a variable called DB which helps to interact with the database (something like
repository in springboot)
*/
var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "siddharth:siddharth@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
