package likercontroller

import (
	models "be_food_recipe/src/models/LikerModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllLikers(c *fiber.Ctx) error {
	liker := models.SelectAllLiker()
	res, err := json.Marshal(liker)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func GetLikerById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectLikerById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostLiker(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var liker models.Liker
		if err := c.BodyParser(&liker); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Liker{
			UserId:   liker.UserId,
			RecipeId: liker.RecipeId,
		}
		models.PostLiker(&item)

		return c.JSON(fiber.Map{
			"Message": "Liker Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateLiker(c *fiber.Ctx) error {

	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var liker models.Liker
		if err := c.BodyParser(&liker); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newLiker := models.Liker{
			UserId:   liker.UserId,
			RecipeId: liker.RecipeId,
		}
		models.UpdateLiker(id, &newLiker)

		return c.JSON(fiber.Map{
			"Message": "Liker Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteLiker(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteLiker(id)

	return c.JSON(fiber.Map{
		"Message": "Liker Deleted",
	})

}
