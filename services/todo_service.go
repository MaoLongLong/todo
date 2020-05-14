package services

import (
	"todo/dao"
	"todo/models"
)

type TodoService struct {
	ID    uint   `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
}

func (s *TodoService) GetAll() (todos []models.Todo, err error) {
	err = dao.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return
}

func (s *TodoService) GetByID() (todo models.Todo, err error) {
	err = dao.DB.First(&todo, s.ID).Error
	if err != nil {
		return models.Todo{}, err
	}
	return
}

func (s *TodoService) CreateByTitle() (todo models.Todo, err error) {
	todo.Title = s.Title
	err = dao.DB.Create(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}
	return
}

func (s *TodoService) DeleteByID() error {
	return dao.DB.Delete(&models.Todo{ID: s.ID}).Error
}

func (s *TodoService) UpdateByIdAndTitle() (todo models.Todo, err error) {
	err = dao.DB.First(&todo, s.ID).Error
	if err != nil {
		return models.Todo{}, err
	}
	todo.Title = s.Title
	err = dao.DB.Save(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}
	return
}
