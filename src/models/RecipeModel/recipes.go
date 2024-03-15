package recipemodel

import (
	"be_food_recipe/src/config"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	title              string
}

func SelectAllRecipe() []*Recipe {
	var items []*Recipe
	config.DB.Find(&items)
	return items
}

func SelectRecipeById(id string) *Recipe {
	var item Recipe
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostRecipe(item *Recipe) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateRecipe(id int, newRecipe *Recipe) error {
	var item Recipe
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newRecipe)
	return result.Error
}

func DeleteRecipe(id int) error {
	var item Recipe
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Recipe {
	var items []*Recipe
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Recipe {
	var items []*Recipe
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Recipe{}).Count(&count)
	return count
}
