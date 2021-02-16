package controller

import (
	"github.com/gustavopergola/go-rest-study/entity"
	"gorm.io/gorm"
)

type StudentController struct {}

func (s *StudentController) Retrieve(db *gorm.DB) (*entity.Student, error) {
	student := entity.Student{}
	err := db.First(&student, 1).Error
	return &student, err
}

func (s *StudentController) Create(db *gorm.DB) (*entity.Student, error) {
	student := entity.Student{Name: "Gustavo"}
	err := db.Create(&student).Error
	return &student, err
}