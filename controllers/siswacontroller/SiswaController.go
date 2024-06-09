package siswacontroller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
	"simkes-go/config/database"
	"simkes-go/models/request"
	"simkes-go/models/response"
	"time"
)

func GetSiswa(c *fiber.Ctx) error {
	var (
		Siswa []response.Siswa
		count int64
	)

	db := database.DB
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := db.WithContext(ctx).Model(&Siswa).Find(&Siswa).Count(&count).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Data Fetch Successfully",
		"Data":    Siswa,
	})
}

func PostSiswa(c *fiber.Ctx) error {
	db := database.DB
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	siswaReq := request.Siswa{}
	if err := c.BodyParser(&siswaReq); err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	insertSiswa := request.Siswa{
		Nama:          siswaReq.Nama,
		NISN:          siswaReq.NISN,
		Kelas:         siswaReq.Kelas,
		TahunAjaranID: siswaReq.TahunAjaranID,
		JurusanID:     siswaReq.JurusanID,
		NamaOrtu:      siswaReq.NamaOrtu,
		NoTelpOrtu:    siswaReq.NoTelpOrtu,
		NoTelpSiswa:   siswaReq.NoTelpSiswa,
		Point:         siswaReq.Point,
		CreatedAt:     time.Now(),
	}

	err := db.WithContext(ctx).Model(&response.Siswa{}).Create(&insertSiswa).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed To Create Data",
		})
	}
	return c.JSON(fiber.Map{
		"Message": "Data Created Successfully",
		"Data":    insertSiswa,
	})
}
