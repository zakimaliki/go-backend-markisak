package main

import (
	"log"
	"markisak/src/configs"
	"markisak/src/helpers"
	"markisak/src/routes"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "myapp",
		Usage: "A simple application",
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "Run database migrations",
				Action: func(c *cli.Context) error {
					helpers.Migration()
					return nil
				},
			},
			{
				Name:  "seed",
				Usage: "Run database seeder",
				Action: func(c *cli.Context) error {
					helpers.Seeder()
					return nil
				},
			},
		},
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	// Membuat instance dari aplikasi Gofiber
	fiberApp := fiber.New()

	configs.InitDB()
	helpers.Seeder()

	// mendenifisikan path
	routes.Router(fiberApp)

	// Menjalankan server pada port 3000
	fiberApp.Listen(":3000")

	// Menjalankan aplikasi CLI
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
