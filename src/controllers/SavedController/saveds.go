package savedcontroller

import (
	models "be_food_recipe/src/models/SavedModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllSaveds(c *fiber.Ctx) error {
	saved := models.SelectAllSaved()
	res, err := json.Marshal(saved)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func GetSavedById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectSavedById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostSaved(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var Saved models.Saved
		if err := c.BodyParser(&Saved); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Saved{
			UserId:   Saved.UserId,
			RecipeId: Saved.RecipeId,
		}
		models.PostSaved(&item)

		return c.JSON(fiber.Map{
			"Message": "Saved Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateSaved(c *fiber.Ctx) error {

	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var Saved models.Saved
		if err := c.BodyParser(&Saved); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newSaved := models.Saved{
			UserId:   Saved.UserId,
			RecipeId: Saved.RecipeId,
		}
		models.UpdateSaved(id, &newSaved)

		return c.JSON(fiber.Map{
			"Message": "Saved Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteSaved(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteSaved(id)

	return c.JSON(fiber.Map{
		"Message": "Saved Deleted",
	})

}
