package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "todo/docs"
	v1 "todo/router/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/todos", v1.GetAllTodo)
		apiV1.GET("/todos/:id", v1.GetTodoByID)
		apiV1.POST("/todos", v1.AddTodo)
		apiV1.PUT("/todos/:id", v1.UpdateTodo)
		apiV1.DELETE("/todos/:id", v1.DeleteTodo)
	}

	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
