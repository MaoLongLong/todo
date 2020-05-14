package main

import (
	"todo/dao"
	"todo/models"
	"todo/router"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	dao.DB.LogMode(true)

	err = models.Migrate()
	if err != nil {
		panic(err)
	}

	r := router.NewRouter()
	r.Run(":9090")
}
