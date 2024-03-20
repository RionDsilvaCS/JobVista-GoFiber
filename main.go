package main

import (
	"os"

	"github.com/RionDsilvaCS/configs"
	"github.com/RionDsilvaCS/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":6000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// app.Get("/", func(c *fiber.Ctx) {
	// 	c.Send("Jobvista server !!!")
	// })

	routes.UserRoute(app)
	routes.JobRoute(app)
	app.Listen(getPort())
}
