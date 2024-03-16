package usermodel

import (
	"be_food_recipe/src/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string
	Email            string
	Password         string
	Picture          string
	PhoneNumber      string
	Role             string
	VerificationCode string
}

func SelectAllUser() []*User {
	var items []*User
	config.DB.Find(&items)
	return items
}

func SelectUserById(id string) *User {
	var item User
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostUser(item *User) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateUser(id int, newUser *User) error {
	var item User
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newUser)
	return result.Error
}

func DeleteUser(id int) error {
	var item User
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*User {
	var items []*User
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*User {
	var items []*User
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&User{}).Count(&count)
	return count
}
