package likermodel

import (
	"be_food_recipe/src/config"

	"gorm.io/gorm"
)

type Liker struct {
	gorm.Model
	UserId   uint
	User     User
	RecipeId uint
	Recipe   Recipe
}

type User struct {
	gorm.Model
	Name  string
	Email string
	Role  string
}

type Recipe struct {
	gorm.Model
	Title      string
	Ingredient string
	Thumbnail  string
	VideoUrl   string
}

func SelectAllLiker() []*Liker {
	var items []*Liker
	config.DB.Preload("User").Preload("Recipe").Find(&items)
	return items
}

func SelectLikerById(id string) *Liker {
	var item Liker
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostLiker(item *Liker) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateLiker(id int, newLiker *Liker) error {
	var item Liker
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newLiker)
	return result.Error
}

func DeleteLiker(id int) error {
	var item Liker
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Liker {
	var items []*Liker
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Liker {
	var items []*Liker
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Liker{}).Count(&count)
	return count
}
