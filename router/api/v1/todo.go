package v1

import (
	"github.com/gin-gonic/gin"
	"todo/pkg/resp"
	"todo/services"
)

func GetAllTodo(c *gin.Context) {
	service := services.TodoService{}
	todos, err := service.GetAll()
	if err != nil {
		resp.InternalServerError(c, nil, err.Error())
		return
	}
	resp.OK(c, gin.H{"todos": todos}, "get all todo")
}

func GetTodoByID(c *gin.Context) {
	service := services.TodoService{}
	err := c.ShouldBindUri(&service)
	if err != nil {
		resp.BadRequest(c, nil, err.Error())
		return
	}
	todo, err := service.GetByID()
	if err != nil {
		resp.InternalServerError(c, nil, err.Error())
		return
	}
	resp.OK(c, gin.H{"todo": todo}, "get todo by id")
}

func AddTodo(c *gin.Context) {
	service := services.TodoService{}
	err := c.ShouldBind(&service)
	if err != nil {
		resp.BadRequest(c, nil, err.Error())
		return
	}
	todo, err := service.CreateByTitle()
	if err != nil {
		resp.InternalServerError(c, nil, err.Error())
		return
	}
	resp.OK(c, gin.H{"todo": todo}, "created")
}

func DeleteTodo(c *gin.Context) {
	service := services.TodoService{}
	err := c.ShouldBindUri(&service)
	if err != nil {
		resp.BadRequest(c, nil, err.Error())
		return
	}
	err = service.DeleteByID()
	if err != nil {
		resp.InternalServerError(c, nil, err.Error())
		return
	}
	resp.OK(c, nil, "deleted")
}

func UpdateTodo(c *gin.Context) {
	service := services.TodoService{}
	err := c.ShouldBindUri(&service)
	if err != nil {
		resp.BadRequest(c, nil, err.Error())
		return
	}
	err = c.ShouldBind(&service)
	if err != nil {
		resp.BadRequest(c, nil, err.Error())
		return
	}
	todo, err := service.UpdateByIdAndTitle()
	if err != nil {
		resp.InternalServerError(c, nil, err.Error())
		return
	}
	resp.OK(c, gin.H{"todo": todo}, "updated")
}
