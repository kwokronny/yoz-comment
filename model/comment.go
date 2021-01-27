package model

import (
	"kwok-comment/dao"
	"kwok-comment/helper"
)

// Comment is Model Type
type Comment struct {
	dao.Model
	ArticleToken string     `gorm:"column:article_token" json:"article_token" form:"article_token"`
	ParentID     uint       `gorm:"column:parent_id" json:"parent_id" form:"parent_id"`
	RID          uint       `gorm:"column:r_id" json:"r_id" form:"r_id"`
	Nickname     string     `gorm:"column:nickname" json:"nickname" form:"nickname"`
	Mail         string     `gorm:"column:mail" json:"mail" form:"mail"`
	Site         string     `gorm:"column:site" json:"site" form:"site"`
	Content      string     `gorm:"column:content" json:"content" form:"content"`
	IP           string     `gorm:"column:ip" json:"ip" form:"ip"`
	Replys       []*Comment `sql:"default:null" json:"replys"`
}

// TableName is Change GORM default TableName
func (q Comment) TableName() string {
	return "comment"
}

// GetArticleComment is get all or search all question records
func (q Comment) GetArticleComment(ArticleToken string, page helper.Pagination) helper.PageData {
	var data helper.PageData
	var comments []Comment
	condition := dao.DB
	condition = condition.Where("article_token = ? and r_id = 0", ArticleToken)
	condition.Order("created_at DESC").Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&comments).Offset(-1).Count(&data.Total)
	var ids []uint
	for index, comment := range comments {
		println(index, comment.ID)
	}
	dao.DB.Where("parent_id in (?)", ids).Order("created_at DESC").Find(&comments)
	data.Records = comments
	data.Page = page.GetPage()
	data.PageSize = page.GetPageSize()
	return data
}

// genrateTree is
func genrateTree(comments []Comment) {

}

// Save is insert or update record
// func (q Comment) Save(data Comment) Comment {
// 	var model Comment
// 	dao.DB.First(&model, data.ID)
// 	if model.ID == 0 {
// 		dao.DB.Create(&data)
// 	} else {
// 		dao.DB.Model(&model).Updates(&data)
// 	}
// 	return data
// }

// // Delete is get one record question
// func (q Comment) Delete(id uint) Comment {
// 	var model Comment
// 	dao.DB.First(&model, id)
// 	dao.DB.Delete(&model)
// 	return model
// }
