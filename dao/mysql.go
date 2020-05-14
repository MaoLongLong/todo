package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMySQL() (err error) {
	DB, err = gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	return err
}
