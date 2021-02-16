package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavopergola/go-rest-study/controller"
	"gorm.io/gorm"
)

type StudentHandler struct {
	controller *controller.StudentController
}

func (s *StudentHandler) Retrieve(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		student, err := s.controller.Retrieve(db)

		if err != nil {
			return fiber.ErrNotFound
		}

		return c.JSON(student)
	}
}

func (s *StudentHandler) Create(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		student, err := s.controller.Create(db)

		if err != nil {
			return fiber.ErrUnprocessableEntity
		}

		return c.JSON(student)
	}
}


func (s *StudentHandler) MountRoutes(f *fiber.App, db *gorm.DB) {
	s.controller = &controller.StudentController{}
	studentGroup := f.Group("/students")
	studentGroup.Get("/:id", s.Retrieve(db))
	studentGroup.Get("/save/testsave", s.Create(db))
}


