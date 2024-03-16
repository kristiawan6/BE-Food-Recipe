package bookmarkermodel

import (
	"be_food_recipe/src/config"

	"gorm.io/gorm"
)

type Bookmarker struct {
	gorm.Model
	UserId   string
	RecipeId string
}

func SelectAllBookmarker() []*Bookmarker {
	var items []*Bookmarker
	config.DB.Find(&items)
	return items
}

func SelectBookmarkerById(id string) *Bookmarker {
	var item Bookmarker
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostBookmarker(item *Bookmarker) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateBookmarker(id int, newBookmarker *Bookmarker) error {
	var item Bookmarker
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newBookmarker)
	return result.Error
}

func DeleteBookmarker(id int) error {
	var item Bookmarker
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Bookmarker {
	var items []*Bookmarker
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Bookmarker {
	var items []*Bookmarker
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Bookmarker{}).Count(&count)
	return count
}
