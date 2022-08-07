package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "hamish:password@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

/*
docker run --name dev-mysql \
        -e MYSQL_ROOT_PASSWORD=password \
        -e MYSQL_DATABASE=simplerest \
        -e MYSQL_USER=hamish \
        -e MYSQL_PASSWORD=password \
        -p 3306:3306 -d mysql:latest
*/
