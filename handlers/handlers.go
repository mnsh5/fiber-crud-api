package handlers

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.SendString("{ping: pong}")
}
