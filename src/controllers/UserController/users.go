package usercontroller

import (
	models "be_food_recipe/src/models/UserModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users := models.SelectAllUser()
	res, err := json.Marshal(users)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetUsersById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectUserById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostUser(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.User{
			Name:             user.Name,
			Email:            user.Email,
			Password:         user.Password,
			Picture:          user.Picture,
			Phonenumber:      user.Phonenumber,
			Role:             user.Role,
			Verificationcode: user.Verificationcode,
		}
		models.PostUser(&item)

		return c.JSON(fiber.Map{
			"Message": "User Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}
