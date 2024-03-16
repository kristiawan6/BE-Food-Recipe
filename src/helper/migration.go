package helper

import (
	"be_food_recipe/src/config"
	usermodel "be_food_recipe/src/models/UserModel"
	recipemodel "be_food_recipe/src/models/RecipeModel"
	videomodel "be_food_recipe/src/models/VideoModel"
	likermodel "be_food_recipe/src/models/LikerModel"
	bookmarkermodel "be_food_recipe/src/models/BookmarkerModel"
)

func Migration() {
	config.DB.AutoMigrate(&usermodel.User{})
	config.DB.AutoMigrate(&recipemodel.Recipe{})
	config.DB.AutoMigrate(&videomodel.Video{})
	config.DB.AutoMigrate(&likermodel.Liker{})
	config.DB.AutoMigrate(&bookmarkermodel.Bookmarker{})
}
