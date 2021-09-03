package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//creating a new instance of fiber.
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "Fiber",
	})

	app.Use(cors.New())

	//handling routes for app.
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("templates/index.gohtml", nil)
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		city := ctx.FormValue("city")
		url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=imperial&appid=1d004cd295137a8caad925f150098a25"
		resp, err := http.Get(url)
		if err != nil {
			return fiber.DefaultErrorHandler(ctx, err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fiber.DefaultErrorHandler(ctx, err)
		}
		return ctx.SendString(string(body))
	})

	//setting up server.
	app.Listen(":8080")
}
