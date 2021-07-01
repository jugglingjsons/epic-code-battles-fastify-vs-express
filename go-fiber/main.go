package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Answer struct {
	Hello string
}

func test(c *fiber.Ctx) error {
	data := Answer{Hello: "world"}

	// timeoutString := os.Getenv("TIMEOUT")
	// timeout, _ := strconv.Atoi(timeoutString)
	// time.Sleep(time.Duration(timeout) * time.Millisecond)

	return c.JSON(data)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", test)

	app.Listen(":" + os.Getenv("PORT"))
}
