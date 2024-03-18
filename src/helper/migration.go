package helper

import (
	"be_food_recipe/src/config"
	likermodel "be_food_recipe/src/models/LikerModel"
	recipemodel "be_food_recipe/src/models/RecipeModel"
	savedmodel "be_food_recipe/src/models/SavedModel"
	usermodel "be_food_recipe/src/models/UserModel"
	videomodel "be_food_recipe/src/models/VideoModel"
	commentmodel "be_food_recipe/src/models/CommentModel"
)

func Migration() {
	config.DB.AutoMigrate(&usermodel.User{})
	config.DB.AutoMigrate(&recipemodel.Recipe{})
	config.DB.AutoMigrate(&videomodel.Video{})
	config.DB.AutoMigrate(&likermodel.Liker{})
	config.DB.AutoMigrate(&savedmodel.Saved{})
	config.DB.AutoMigrate(&commentmodel.Comment{})
}
