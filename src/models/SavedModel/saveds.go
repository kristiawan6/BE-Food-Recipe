package savedmodel

import (
	"be_food_recipe/src/config"

	"gorm.io/gorm"
)

type Saved struct {
	gorm.Model
	UserId   uint
	RecipeId uint
}

func SelectAllSaved() []*Saved {
	var items []*Saved
	config.DB.Find(&items)
	return items
}

func SelectSavedById(id string) *Saved {
	var item Saved
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostSaved(item *Saved) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateSaved(id int, newSaved *Saved) error {
	var item Saved
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newSaved)
	return result.Error
}

func DeleteSaved(id int) error {
	var item Saved
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Saved {
	var items []*Saved
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Saved {
	var items []*Saved
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Saved{}).Count(&count)
	return count
}
