package controller

import (
	"github.com/gustavopergola/go-rest-study/entity"
	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

type StudentPost struct {
	Name string `json:"name"`
}

func (s *StudentController) Retrieve(id int) (*entity.Student, error) {
	student := entity.Student{}
	err := s.DB.First(&student, id).Error
	return &student, err
}

func (s *StudentController) Create(body *StudentPost) (*entity.Student, error) {
	student := entity.Student{
		Name: body.Name,
	}
	err := s.DB.Create(&student).Error
	return &student, err
}