package routes

import (
	"github.com/gofiber/fiber/v2"
	"simkes-go/controllers/authcontroller"
	"simkes-go/controllers/siswacontroller"
)

func RouteInit(r *fiber.App) {
	api := r.Group("/Api")
	auth := api.Group("/Auth")
	siswa := api.Group("/Siswa")

	// AUTH //
	auth.Post("/Register", authcontroller.Register)

	// SISWA //
	siswa.Post("/Get", siswacontroller.GetSiswa)
	siswa.Post("/", siswacontroller.PostSiswa)

}
