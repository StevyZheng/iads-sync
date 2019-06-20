package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

type Connection struct {
	Eloquent *gorm.DB
}

func NewConnection() (conn *Connection) {
	var err error
	Eloquent, err := gorm.Open("mysql", "root:000000@tcp(127.0.0.1:3306)/iads?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Eloquent.Error != nil {
		fmt.Printf("common error %v", Eloquent.Error)
	}
	Eloquent.DB().SetMaxIdleConns(10)
	Eloquent.DB().SetMaxOpenConns(100)
	conn = &Connection{Eloquent}
	return
}
