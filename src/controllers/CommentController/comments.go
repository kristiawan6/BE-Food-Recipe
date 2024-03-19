package commentcontroller

import (
	"be_food_recipe/src/helper"
	models "be_food_recipe/src/models/CommentModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllComments(c *fiber.Ctx) error {
	helper.EnableCors(c)
	comments := models.SelectAllComment()

	response := fiber.Map{
		"Message": "Success",
		"data":    comments,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetCommentById(c *fiber.Ctx) error {
	helper.EnableCors(c)	
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectCommentById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostComment(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPost {
		var comment models.Comment
		if err := c.BodyParser(&comment); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Comment{
			UserId:   comment.UserId,
			RecipeId: comment.RecipeId,
		}
		models.PostComment(&item)

		return c.JSON(fiber.Map{
			"Message": "Comment Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateComment(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var comment models.Comment
		if err := c.BodyParser(&comment); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newComment := models.Comment{
			UserId:   comment.UserId,
			RecipeId: comment.RecipeId,
		}
		models.UpdateComment(id, &newComment)

		return c.JSON(fiber.Map{
			"Message": "Comment Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteComment(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteComment(id)

	return c.JSON(fiber.Map{
		"Message": "Comment Deleted",
	})

}
