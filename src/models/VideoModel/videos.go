package videomodel

import (
	"be_food_recipe/src/config"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title     string
	Url       string
	Thumbnail string
	RecipeId  uint
	Recipe    Recipe
}

type Recipe struct {
	gorm.Model
	Title      string
	Ingredient string
	Thumbnail  string
	VideoUrl   string
}

func SelectAllVideo() []*Video {
	var items []*Video
	config.DB.Preload("Recipe").Find(&items)
	return items
}

func SelectVideoById(id string) *Video {
	var item Video
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostVideo(item *Video) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateVideo(id int, newVideo *Video) error {
	var item Video
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newVideo)
	return result.Error
}

func DeleteVideo(id int) error {
	var item Video
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Video {
	var items []*Video
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Video {
	var items []*Video
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Video{}).Count(&count)
	return count
}
