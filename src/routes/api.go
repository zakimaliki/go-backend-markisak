package routes

import (
	"markisak/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	// Mendefinisikan handler untuk route GET "/"
	app.Get("/", controllers.GetHello)
}
