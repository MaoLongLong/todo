package router

import (
	"github.com/gin-gonic/gin"
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

	return r
}
