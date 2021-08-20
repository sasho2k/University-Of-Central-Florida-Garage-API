package client

import (
	"github.com/gofiber/fiber"
	"github.com/sasho2k/University-Of-Central-Florida-Garage-API/internals"
	"github.com/sasho2k/University-Of-Central-Florida-Garage-API/job"
	"net/http"
	"time"
)

// StartService will serve as our router service. It is responsible for capturing traffic to our route and returning
// a response based on the parameter given to :storeNumber. There is also a check beforehand.
func StartService() {
	app := fiber.New(fiber.Config{GETOnly: true})

	app.Get("/", func(c *fiber.Ctx) error {
		garages, err := job.GarageRequest()
		if err != nil {
			return c.SendString(err.Error())
		}

		year, month, day := time.Now().Date()
		hour, min, sec := time.Now().Clock()
		date := internals.ParseDate(year, month, day, hour, min, sec)

		str := "UCF GARAGE INFORMATION AS OF " + date + "\n"

		for _, garage := range garages {
			str += garage.Print() + "\n"
		}

		return c.SendString(str)
	})

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	app.Get("/get-garages/", func(c *fiber.Ctx) error {
		garages, err := job.GarageRequest()
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.JSON(garages)
	})

	app.Get("/get-garages/:garage", func(c *fiber.Ctx) error {
		garageNum, err := ReturnParam(c.Params("garage"))
		if err != nil {
			return c.SendString(err.Error())
		}

		garages, err := job.GarageRequest()
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.JSON(garages[garageNum])
	})

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
