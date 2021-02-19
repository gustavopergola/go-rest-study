package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavopergola/go-rest-study/controller"
	"gorm.io/gorm"
	"strconv"
)

type StudentHandler struct {
	controller *controller.StudentController
}

func NewStudentHandler(db *gorm.DB) StudentHandler {
	return StudentHandler{
		controller: &controller.StudentController{
			DB: db,
		},
	}
}

func (s *StudentHandler) Retrieve() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		student, err := s.controller.Retrieve(id)

		if err != nil {
			return fiber.ErrNotFound
		}

		return c.JSON(student)
	}
}

func (s *StudentHandler) Create() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := controller.StudentPost{}
		err := c.BodyParser(&body)

		if err != nil {
			return fiber.ErrUnprocessableEntity
		}

		student, err2 := s.controller.Create(&body)

		if err2 != nil {
			return fiber.ErrInternalServerError
		}

		return c.JSON(student)
	}
}


func (s *StudentHandler) MountRoutes(f *fiber.App) {
	s.controller = &controller.StudentController{}
	studentGroup := f.Group("/students")
	studentGroup.Get("/:id", s.Retrieve())
	studentGroup.Post("", s.Create())
}


