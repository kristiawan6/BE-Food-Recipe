package routes

import (
	commentcontroller "be_food_recipe/src/controllers/CommentController"
	likercontroller "be_food_recipe/src/controllers/LikerController"
	recipecontroller "be_food_recipe/src/controllers/RecipeController"
	savedcontroller "be_food_recipe/src/controllers/SavedController"
	usercontroller "be_food_recipe/src/controllers/UserController"
	videocontroller "be_food_recipe/src/controllers/VideoController"

	// "github.com/goddtriffin/helmet"
	"github.com/gofiber/fiber/v2"
)

func Router(c *fiber.App) {

	// helmet := helmet.Default()

	v1 := c.Group("/api/v1")

	// v1.Use(helmet)
	// c.Use(helmet)

	c.Post("/login", usercontroller.Login)
	c.Post("/register-user", usercontroller.RegisterUser)
	c.Post("/register-admin", usercontroller.RegisterAdmin)

	user := v1.Group("/user")
	{
		user.Get("/data", usercontroller.GetAllUsers)
		user.Get("/:id", usercontroller.GetUserById)
		user.Put("/update/:id", usercontroller.UpdateUser)
		user.Delete("/delete/:id", usercontroller.DeleteUser)
	}

	recipe := v1.Group("/recipe")
	{
		recipe.Get("/data", recipecontroller.GetAllRecipes)
		recipe.Get("/:id", recipecontroller.GetRecipeById)
		recipe.Post("/create", recipecontroller.PostRecipe)
		// recipe.Get("/paginated-data", recipecontroller)
		recipe.Put("/update/:id", recipecontroller.UpdateRecipe)
		recipe.Delete("/delete/:id", recipecontroller.DeleteRecipe)
	}

	video := v1.Group("/video")
	{
		video.Get("/data", videocontroller.GetAllVideos)
		video.Get("/:id", videocontroller.GetVideoById)
		video.Post("/create", videocontroller.PostVideo)
		// video.Get("/paginated-data", videocontroller)
		video.Put("/update/:id", videocontroller.UpdateVideo)
		video.Delete("/delete/:id", videocontroller.DeleteVideo)
	}

	comment := v1.Group("/comment")
	{
		comment.Get("/data", commentcontroller.GetAllComments)
		comment.Get("/:id", commentcontroller.GetCommentById)
		comment.Post("/create", commentcontroller.PostComment)
		// comment.Get("/paginated-data", commentcontroller)
		comment.Put("/update/:id", commentcontroller.UpdateComment)
		comment.Delete("/delete/:id", commentcontroller.DeleteComment)
	}

	saved := v1.Group("/saved")
	{
		saved.Get("/data", savedcontroller.GetAllSaveds)
		saved.Get("/:id", savedcontroller.GetSavedById)
		saved.Post("/create", savedcontroller.PostSaved)
		// saved.Get("/paginated-data", savedcontroller)
		saved.Put("/update/:id", savedcontroller.UpdateSaved)
		saved.Delete("/delete/:id", savedcontroller.DeleteSaved)
	}

	liked := v1.Group("/liked")
	{
		liked.Get("/data", likercontroller.GetAllLikers)
		liked.Get("/:id", likercontroller.GetLikerById)
		liked.Post("/create", likercontroller.PostLiker)
		// liked.Get("/paginated-data", likercontroller)
		liked.Put("/update/:id", likercontroller.UpdateLiker)
		liked.Delete("/delete/:id", likercontroller.DeleteLiker)
	}

}
