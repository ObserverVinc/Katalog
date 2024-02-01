package routes

import (
	"github.com/ObserverVinc/Katalog_pusri/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/newkatalog", controllers.NewAppKatalog)
	app.Post("/api/newkategori", controllers.NewKategori)
	app.Get("/api/getkatalog", controllers.GetKatalog)
	app.Get("/api/getkategori", controllers.GetKategori)
	app.Delete("/api/deletekatalog", controllers.DeleteKatalog)
	app.Patch("/api/editkatalog", controllers.EditKatalog)

}
