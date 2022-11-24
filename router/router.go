package router

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type messageRespon struct {
	Message string `json:"message"`
	status  bool   `json:"status"`
}

func SetupRoutes() {
	app := fiber.New()

	app.Get("/:id?", func(c *fiber.Ctx) error {
		if c.Params("id") != "" {
			id := c.Params("id")
			return c.SendFile(fmt.Sprintf(`./output/%v.json`, id))
		}
		m := messageRespon{Message: "Not Data", status: false}
		byteRespon, _ := json.Marshal(m)
		return c.SendString(string(byteRespon))
	})

	app.Get("/provinces", func(c *fiber.Ctx) error {
		return c.SendFile("./output/provinces.json")
	})

	log.Fatal(app.Listen(":3003"))
}
