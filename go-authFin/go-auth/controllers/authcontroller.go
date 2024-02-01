package controllers

import (
	"strconv"
	"time"

	"github.com/ObserverVinc/Katalog_pusri/database"
	"github.com/ObserverVinc/Katalog_pusri/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

//	type Register_request struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//		Email    string `json:"email"`
//	}
const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string
	//data := new(Register_request)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func NewAppKatalog(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	katalog := models.Katalog{
		Id_kategori:   data["id_katalog"],
		Nama_aplikasi: data["nama_aplikasi"],
		Deskripsi:     data["deskripsi"],
		Link:          data["link"],
		Gambar:        data["gambar"],
	}
	database.DB.Create(&katalog)
	return c.JSON(katalog)
}

func GetKatalog(c *fiber.Ctx) error {
	var katalogs []models.Katalog
	database.DB.Find(&katalogs)
	return c.JSON(katalogs)
}

func NewKategori(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	kategori := models.Kategori{
		Id_kategori: data["id_katalog"],
		Deskripsi_K: data["deskripsiK"],
	}
	database.DB.Create(&kategori)
	return c.JSON(kategori)
}

func GetKategori(c *fiber.Ctx) error {
	var kategoris []models.Kategori
	database.DB.Find(&kategoris)
	return c.JSON(kategoris)
}

// this function delete data from mysql database (id is the one taken to delete)
func DeleteKatalog(c *fiber.Ctx) error {
	// Parse request body to get the ID of the Katalog to be deleted
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Get the ID of the Katalog to be deleted from the request body
	id := data["id"]

	// Delete the Katalog from the database based on the provided ID
	if err := database.DB.Where("id_aplikasi = ?", id).Delete(&models.Katalog{}).Error; err != nil {
		return err
	}

	// Return a success message or response
	return c.SendString("Katalog deleted successfully")
}

// this function edit the data in database
func EditKatalog(c *fiber.Ctx) error {
	// Parse request body to get the updated Katalog data
	var updatedKatalog models.Katalog
	if err := c.BodyParser(&updatedKatalog); err != nil {
		return err
	}

	// Update the Katalog in the database
	if err := database.DB.Model(&models.Katalog{}).Where("id_aplikasi = ?", updatedKatalog.Id_aplikasi).Updates(updatedKatalog).Error; err != nil {
		return err
	}

	// Return a success message or response
	return c.SendString("Katalog updated successfully")
}
