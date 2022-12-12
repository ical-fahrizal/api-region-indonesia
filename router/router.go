package router

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type messageRespon struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func SetupRoutes() {
	app := fiber.New()

	app.Get("/:id?", func(c *fiber.Ctx) (err error) {
		if c.Params("id") != "" {
			id := c.Params("id")
			if err = c.SendFile(fmt.Sprintf(`./output/%v.json`, id)); err != nil {
				m := messageRespon{Message: "Not Data", Status: false}
				byteRespon, _ := json.Marshal(m)
				return c.SendString(string(byteRespon))
			}
			return c.SendFile(fmt.Sprintf(`./output/%v.json`, id))
		}
		m := messageRespon{Message: "Not Data", Status: false}
		byteRespon, _ := json.Marshal(m)
		return c.SendString(string(byteRespon))
	})

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3003"))
}
