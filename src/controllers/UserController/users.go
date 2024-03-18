package usercontroller

import (
	"be_food_recipe/src/helper"
	models "be_food_recipe/src/models/UserModel"
	"encoding/json"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAdmin(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := models.User{
			Name:             user.Name,
			Email:            user.Email,
			Password:         Password,
			PhoneNumber:      user.PhoneNumber,
			Picture:          user.Picture,
			Role:             "Admin",
			VerificationCode: "-",
		}
		models.PostUser(&item)

		return c.JSON(fiber.Map{
			"Message": "Admin Registered",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func RegisterUser(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := models.User{
			Name:             user.Name,
			Email:            user.Email,
			Password:         Password,
			PhoneNumber:      user.PhoneNumber,
			Picture:          user.Picture,
			Role:             "User",
			VerificationCode: "-",
		}
		models.PostUser(&item)

		return c.JSON(fiber.Map{
			"Message": "User Registered",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func Login(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		users, err := models.FindEmail(user.Email)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Email Not Found")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password)); err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Password Not found")
		}

		jwtKey := os.Getenv("SECRETKEY")
		token, err := helper.GenerateToken(jwtKey, user.Email, user.Role)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed To Generate Tokens")
		}

		payload := fiber.Map{
			"Message": "HI, " + users.Name + " as a " + users.Role,
			"Email":   user.Email,
			"Role":    users.Role,
			"Token":   token,
		}
		return c.JSON(payload)
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func GetAllUsers(c *fiber.Ctx) error {
	users := models.SelectAllUser()
	res, err := json.Marshal(users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal Konversi Json")
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetUserById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectUserById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		newUser := models.User{
			Name:        user.Name,
			Email:       user.Email,
			Password:    Password,
			PhoneNumber: user.PhoneNumber,
			Picture:     user.Picture,
		}
		models.UpdateUser(id, &newUser)

		return c.JSON(fiber.Map{
			"Message": "User Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteUser(id)

	return c.JSON(fiber.Map{
		"Message": "User Deleted",
	})

}
