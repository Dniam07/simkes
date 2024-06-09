package authcontroller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
	"simkes-go/config/database"
	"simkes-go/helpers"
	"simkes-go/models/request"
	"simkes-go/models/response"
	"time"
)

func Register(c *fiber.Ctx) error {

	db := database.DB
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	usersReq := request.Users{}
	if err := c.BodyParser(&usersReq); err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// Validasi kustom untuk username
	if usersReq.Username == "" {
		return c.JSON(fiber.Map{
			"status": 401,
			"error":  "pastikan username sudah benar",
		})
	}
	var count int64
	err := db.WithContext(ctx).Model(&response.Users{}).Where("username = ?", usersReq.Username).Find(&response.Users{}).Count(&count).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	if count > 0 {
		return c.JSON(fiber.Map{
			"status": 409,
			"error":  "username telah digunakan",
		})
	}

	password, err := helpers.HashPassword(usersReq.Password, helpers.GenerateSalt())
	if err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	insertUser := request.Users{
		Username:   usersReq.Username,
		Password:   password,
		UserroleId: usersReq.UserroleId,
		SiswaId:    usersReq.SiswaId,
		GuruId:     usersReq.GuruId,
		Status:     false,
		CreatedAt:  time.Now(),
	}

	errCreateUser := db.WithContext(ctx).Create(&insertUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed To Create Data",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Data Created Successfully",
		"Data":    insertUser,
	})
}
