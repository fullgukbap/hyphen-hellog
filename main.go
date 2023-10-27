package main

import (
	"flag"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/controller"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var port string

func init() {
	flagPort := flag.Int("p", 8080, "Enter the port")
	flag.Parse()

	port = ":" + strconv.Itoa(*flagPort)
}

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: cerrors.ErrorHandler})
	app.Use(recover.New())

	log.Fatal(controller.Route(app).Listen(port))

}
