package main

import (
	"todo/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	//err := dao.InitMySQL()
	//if err != nil {
	//	panic(err)
	//}
	//defer dao.DB.Close()
	//dao.DB.LogMode(true)
	//
	//err = models.Migrate()
	//if err != nil {
	//	panic(err)
	//}

	r := router.NewRouter()
	r.Run(":9090")
}
