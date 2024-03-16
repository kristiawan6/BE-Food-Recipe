package bookmarkercontroller

import (
	models "be_food_recipe/src/models/BookmarkerModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllBookmark(c *fiber.Ctx) error {
	bookmark := models.SelectAllBookmarker()
	res, err := json.Marshal(bookmark)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func GetBookmarkById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectBookmarkerById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostBookmark(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var Bookmarker models.Bookmarker
		if err := c.BodyParser(&Bookmarker); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Bookmarker{
			UserId:   Bookmarker.UserId,
			RecipeId: Bookmarker.RecipeId,
		}
		models.PostBookmarker(&item)

		return c.JSON(fiber.Map{
			"Message": "Bookmark Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateBookmark(c *fiber.Ctx) error {

	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var Bookmarker models.Bookmarker
		if err := c.BodyParser(&Bookmarker); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newBookmark := models.Bookmarker{
			UserId:   Bookmarker.UserId,
			RecipeId: Bookmarker.RecipeId,
		}
		models.UpdateBookmarker(id, &newBookmark)

		return c.JSON(fiber.Map{
			"Message": "Bookmark Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteBookmark(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteBookmarker(id)

	return c.JSON(fiber.Map{
		"Message": "Bookmark Deleted",
	})

}
