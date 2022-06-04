package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCategoryByName(name string) int {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ErrorCategoryNameUsed
	}
	return errmsg.SUCCESS
}

func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ErrorCategoryCreateFail
	}
	return errmsg.SUCCESS
}

func GetCategoryInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

func GetCategoryList(pageSize int, pageNum int) ([]Category, int, int64) {
	var categoryList []Category
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err := db.Limit(pageSize).Offset(offset).Find(&categoryList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ErrorCategoryQueryFail, total
	}
	return categoryList, errmsg.SUCCESS, total
}

func EditCategory(id uint, data *Category) int {
	var category Category
	var cateMap = make(map[string]interface{})
	cateMap["name"] = data.Name
	err := db.Model(&category).Where("id = ?", id).Updates(cateMap).Error
	if err != nil {
		return errmsg.ErrorCategoryQueryFail
	}
	return errmsg.SUCCESS
}

func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ErrorCategoryDeleteFail
	}
	return errmsg.SUCCESS
}
