package model

import (
	"YozComment/dao"
	"YozComment/helper"
)

// Comment is Model Type
type Comment struct {
	dao.Model
	ArticleToken string    `gorm:"column:article_token;type:varchar(50);not null;comment:文章token" json:"articleToken" form:"articleToken" binding:"required"`
	ParentID     uint64    `gorm:"column:parent_id;type:int;size:32;not null;comment:楼ID" json:"parentId" form:"parentId"`
	RID          uint      `gorm:"column:r_id;type:int;size:32;not null;comment:回复ID" json:"rId" form:"rId"`
	NickName     string    `gorm:"column:nickname;type:varchar(50);not null;comment:昵称" json:"nickName" form:"nickName" binding:"required"`
	Mail         string    `gorm:"column:mail;type:varchar(50);not null;comment:邮箱" json:"mail" form:"mail" binding:"required,email"`
	Site         string    `gorm:"column:site;type:varchar(50);comment:网站" json:"site" form:"site"`
	Content      string    `gorm:"column:content;type:varchar(255);not null;comment:内容" json:"content" form:"content" binding:"required"`
	pageUrl      string    `gorm:"column:page_url;type:varchar(255);not null;comment:来源页面" json:"pageUrl" binding:"required"`
	PageTitle    string    `gorm:"column:page_title;type:varchar(100);not null;comment:页面标题" json:"PageTitle" binding:"required"`
	IP           string    `gorm:"column:ip;type:varchar(50);not null;comment:IP"`
	Replys       []Comment `gorm:"-" json:"replys"`
}

func init() {
	dao.DB.AutoMigrate(&Comment{})
}

// TableName is Change GORM default TableName
func (q Comment) TableName() string {
	return "yoz-comment"
}

// GetPage is get all comment
func (q Comment) GetPage(nickName string, mail string, content string, page util.Pagination) util.PageData {
	var data util.PageData
	var comments []Comment
	dao.DB.Where("content LIKE ? or nickname = ? or mail = ?", "%%"+content+"%%", nickName, mail).Order("created_at DESC").Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&comments).Offset(-1).Count(&data.Total)
	data.Records = comments
	data.Page = page.GetPage()
	data.PageSize = page.GetPageSize()
	return data
}

// GetCommentByArticle is get comment by article
func (q Comment) GetCommentByArticle(ArticleToken string, page util.Pagination) util.PageData {
	var data util.PageData
	var comments []Comment
	condition := dao.DB
	condition = condition.Where("article_token = ? and r_id = 0", ArticleToken)
	condition.Order("created_at DESC").Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&comments).Offset(-1).Count(&data.Total)
	ids := make([]uint, data.Total)
	for index, comment := range comments {
		ids[index] = comment.ID
	}
	var replys []Comment
	dao.DB.Where("parent_id in (?)", ids).Order("created_at DESC").Find(&replys)
	comments = append(comments, replys...)
	data.Records = genrateTree(comments)
	data.Page = page.GetPage()
	data.PageSize = page.GetPageSize()
	return data
}

func genrateTree(comments []Comment) (trees []Comment) {
	trees = []Comment{}
	var roots, childs []Comment
	for _, comment := range comments {
		if comment.ParentID == 0 {
			// 判断顶层根节点
			roots = append(roots, comment)
		}
		childs = append(childs, comment)
	}

	for _, comment := range roots {
		recursiveTree(&comment, childs)
		trees = append(trees, comment)
	}
	return
}

func recursiveTree(tree *Comment, nodes []Comment) {
	for _, comment := range nodes {
		if comment.ParentID == 0 {
			continue
		}
		if tree.ID == comment.RID {
			recursiveTree(&comment, nodes)
			tree.Replys = append(tree.Replys, comment)
		}
	}
}

// Save is insert or update record
func (q Comment) Save(data Comment) Comment {
	dao.DB.Create(&data)
	return data
}

// Delete is get one record Comment
func (q Comment) Delete(id uint) Comment {
	var model Comment
	dao.DB.First(&model, id)
	dao.DB.Delete(&model)
	return model
}
