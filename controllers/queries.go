package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetQuerys(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": false,
		"msg":   "test",
	})
}
