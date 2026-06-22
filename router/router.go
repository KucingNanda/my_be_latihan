package router

import (
	"be_latihan/config/middleware"
	"be_latihan/handler"
	"be_latihan/model"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.Response{
			Message: "API be_latihan aktif",
		})
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	userGroup := app.Group("/api/user", middleware.JWTProtected(""))
	userGroup.Put("/password", handler.ChangePassword)

	mahasiswa := app.Group("/api/mahasiswa", middleware.JWTProtected("admin"))
	mahasiswa.Get("/", handler.GetAllMahasiswa)
	mahasiswa.Get("/:npm", handler.GetMahasiswaByNPM)
	mahasiswa.Post("/", handler.InsertMahasiswa)
	mahasiswa.Put("/:npm", handler.UpdateMahasiswa)
	mahasiswa.Delete("/:npm", handler.DeleteMahasiswa)

	// jika menggunakan query parameter
	mahasiswa.Get("/search", handler.GetMahasiswaByNPM)
}
