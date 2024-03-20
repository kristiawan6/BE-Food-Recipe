package recipecontroller

import (
	models "be_food_recipe/src/models/RecipeModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllRecipes(c *fiber.Ctx) error {
	recipes := models.SelectAllRecipe()

	response := fiber.Map{
		"Message": "Success",
		"data":    recipes,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to convert to JSON")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetRecipeById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectRecipeById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostRecipe(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var recipe models.Recipe
		if err := c.BodyParser(&recipe); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Recipe{
			Title:      recipe.Title,
			Ingredient: recipe.Ingredient,
			Thumbnail:  recipe.Thumbnail,
			VideoUrl:   recipe.VideoUrl,
			UserId:     recipe.UserId,
		}
		models.PostRecipe(&item)

		return c.JSON(fiber.Map{
			"Message": "Recipe Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateRecipe(c *fiber.Ctx) error {

	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var recipe models.Recipe
		if err := c.BodyParser(&recipe); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newRecipe := models.Recipe{
			Title:      recipe.Title,
			Ingredient: recipe.Ingredient,
			Thumbnail:  recipe.Thumbnail,
			VideoUrl:   recipe.VideoUrl,
			UserId:     recipe.UserId,
		}
		models.UpdateRecipe(id, &newRecipe)

		return c.JSON(fiber.Map{
			"Message": "Recipe Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteRecipe(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteRecipe(id)

	return c.JSON(fiber.Map{
		"Message": "Recipe Deleted",
	})

}
