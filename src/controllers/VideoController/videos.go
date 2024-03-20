package videocontroller

import (
	"be_food_recipe/src/helper"
	models "be_food_recipe/src/models/VideoModel"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllVideos(c *fiber.Ctx) error {
	helper.EnableCors(c)
	video := models.SelectAllVideo()

	response := fiber.Map{
		"Message": "Success",
		"data":    video,
	}

	res, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Gagal Konversi Json")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(res)
}

func GetVideoById(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectVideoById(strconv.Itoa(id))
	return c.JSON(fiber.Map{
		"Message": "Success",
		"data":    res,
	})
}

func PostVideo(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPost {
		var Video models.Video
		if err := c.BodyParser(&Video); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		item := models.Video{
			Title:     Video.Title,
			Url:       Video.Url,
			Thumbnail: Video.Thumbnail,
			RecipeId:  Video.RecipeId,
		}
		models.PostVideo(&item)

		return c.JSON(fiber.Map{
			"Message": "Video Posted",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func UpdateVideo(c *fiber.Ctx) error {
	helper.EnableCors(c)
	if c.Method() == fiber.MethodPut {
		idParam := c.Params("id")
		id, _ := strconv.Atoi(idParam)
		var Video models.Video
		if err := c.BodyParser(&Video); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		newVideo := models.Video{
			Title:     Video.Title,
			Url:       Video.Url,
			Thumbnail: Video.Thumbnail,
			RecipeId:  Video.RecipeId,
		}
		models.UpdateVideo(id, &newVideo)

		return c.JSON(fiber.Map{
			"Message": "Video Updated",
		})
	} else {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method tidak diizinkan")
	}
}

func DeleteVideo(c *fiber.Ctx) error {
	helper.EnableCors(c)
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteVideo(id)

	return c.JSON(fiber.Map{
		"Message": "Video Deleted",
	})

}
