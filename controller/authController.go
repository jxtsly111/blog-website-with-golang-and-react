package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jxtsly111/blog/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	// check if password is less than 6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 character",
		})
	}

}
