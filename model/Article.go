package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     uint   `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ErrorArticleCreateFail
	}
	return errmsg.SUCCESS
}

// GetArticleUnderCategory 查询分类下的所有文章
func GetArticleUnderCategory(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	// Count()查询必须在where条件之后，limit,offset 分页 之前
	err := db.Preload("Category").Limit(pageSize).Offset(offset).Where("cid = ?", cid).Find(&articleList).Error
	db.Model(&articleList).Where("cid = ?", cid).Count(&total)
	if err != nil {
		return articleList, errmsg.ErrorArticleQueryUnderCategoryFail, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// GetArticleById 查询单个文章
func GetArticleById(id uint) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ErrorArticleQueryByIdFail
	}
	return article, errmsg.SUCCESS
}

// GetArticleList 文章列表
func GetArticleList(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := db.Preload("Category").Limit(pageSize).Offset(offset).Order("Updated_At DESC").Find(&articleList).Error
	db.Model(&articleList).Count(&total)
	if err != nil {
		return nil, errmsg.ErrorArticleQueryFail, total
	}
	return articleList, errmsg.SUCCESS, total

}

func SearchArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := db.Preload("Category").Where("title LIKE ?", "%"+title+"%").Limit(pageSize).Offset(offset).Order("Updated_At DESC").Find(&articleList).Error
	db.Model(&articleList).Where("title LIKE ?", "%"+title+"%").Count(&total)
	if err != nil {
		return nil, errmsg.ErrorArticleQueryFail, total
	}
	return articleList, errmsg.SUCCESS, total
}

func EditArticle(id uint, data *Article) int {
	var article Article
	var articleMap = make(map[string]interface{})
	articleMap["title"] = data.Title
	articleMap["cid"] = data.Cid
	articleMap["desc"] = data.Desc
	articleMap["content"] = data.Content
	articleMap["img"] = data.Img
	err := db.Model(&article).Where("id = ?", id).Updates(articleMap).Error
	if err != nil {
		return errmsg.ErrorArticleEditFail
	}
	return errmsg.SUCCESS
}

func DeleteArticle(id uint) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ErrorArticleDeleteFail
	}
	return errmsg.SUCCESS
}
