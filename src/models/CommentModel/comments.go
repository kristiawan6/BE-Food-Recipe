package commentmodel

import (
	"be_food_recipe/src/config"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text     string
	UserId   uint
	RecipeId uint
}

func SelectAllComment() []*Comment {
	var items []*Comment
	config.DB.Find(&items)
	return items
}

func SelectCommentById(id string) *Comment {
	var item Comment
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostComment(item *Comment) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateComment(id int, newComment *Comment) error {
	var item Comment
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newComment)
	return result.Error
}

func DeleteComment(id int) error {
	var item Comment
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Comment {
	var items []*Comment
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Comment {
	var items []*Comment
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Comment{}).Count(&count)
	return count
}
