package models

import "todo/dao"

func Migrate() error {
	return dao.DB.AutoMigrate(&Todo{}).Error
}
