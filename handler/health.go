package handler

import (
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
}

func (h *HealthHandler) HealthCheck() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("OK")
	}
}

func (h *HealthHandler) MountRoutes(f *fiber.App) {
	f.Get("/health_check", h.HealthCheck())
}


