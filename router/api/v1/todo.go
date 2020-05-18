package v1

import (
	"github.com/gin-gonic/gin"
	"todo/pkg/resp"
	"todo/services"
)

// @Summary Get All
// @Produce json
// @Success 200 {object} resp.Response
// @Failure 400 {object} resp.Response
// @Failure 500 {object} resp.Response
// @Router /todos [get]
func GetAllTodo(c *gin.Context) {
	service := services.TodoService{}
	todos, err := service.GetAll()
	if err != nil {
		resp.InternalServerError(c, nil, err.Error())
		return
	}
	resp.OK(c, gin.H{"todos": todos}, "get all todo")
}

// @Summary Get By ID
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} resp.Response
// @Failure 400 {object} resp.Response
// @Failure 500 {object} resp.Response
// @Router /todos/{id} [get]
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

// @Summary Add
// @Produce json
// @Param title body string true "Title"
// @Success 200 {object} resp.Response
// @Failure 400 {object} resp.Response
// @Failure 500 {object} resp.Response
// @Router /todos [post]
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

// @Summary Delete
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} resp.Response
// @Failure 400 {object} resp.Response
// @Failure 500 {object} resp.Response
// @Router /todos/{id} [delete]
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

// @Summary Update
// @Produce json
// @Param id path int true "ID"
// @Param title body string true "Title"
// @Success 200 {object} resp.Response
// @Failure 400 {object} resp.Response
// @Failure 500 {object} resp.Response
// @Router /todos/{id} [put]
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
