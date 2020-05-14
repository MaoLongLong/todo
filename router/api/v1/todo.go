package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.BadRequest(c, nil, "param invalid")
		return
	}
	service := services.TodoService{ID: uint(id)}
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.BadRequest(c, nil, "param invalid")
		return
	}
	service := services.TodoService{ID: uint(id)}
	err = service.DeleteByID()
	if err != nil {
		resp.InternalServerError(c, nil, err.Error())
		return
	}
	resp.OK(c, nil, "deleted")
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.BadRequest(c, nil, "param invalid")
		return
	}
	service := services.TodoService{}
	err = c.ShouldBind(&service)
	if err != nil {
		resp.BadRequest(c, nil, err.Error())
		return
	}
	service.ID = uint(id)
	todo, err := service.UpdateByIdAndTitle()
	if err != nil {
		resp.InternalServerError(c, nil, err.Error())
		return
	}
	resp.OK(c, gin.H{"todo": todo}, "updated")
}
