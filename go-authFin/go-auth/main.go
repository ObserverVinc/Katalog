package main

import (
	"github.com/ObserverVinc/Katalog_pusri/database"
	"github.com/ObserverVinc/Katalog_pusri/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	routes.Setup(app)

	app.Listen(":8000")
}

//func initDB() {
//	var err error
//	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3307)/pusri-go")
//	if err != nil {
//		log.Fatal(err)
//	}

// Ensure the connection to the database
//	if err = db.Ping(); err != nil {
//		log.Fatal(err)
//	}
