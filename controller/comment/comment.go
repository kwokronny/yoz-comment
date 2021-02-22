package comment

import (
	"KBCommentAPI/helper"
	"KBCommentAPI/model"
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

var commentModel = model.Comment{}
var resp = helper.Response{}

type getPageRequestQuery struct {
	NickName string `form:"nickName"`
	Mail     string `form:"mail"`
	Content  string `form:"content"`
}

// GetPage 获取所有问答并分页
func GetPage(c *gin.Context) {
	data := &getPageRequestQuery{}
	c.BindQuery(&data)
	page := helper.GetPagination(c)
	comments := commentModel.GetPage(data.NickName, data.Mail, data.Content, page)
	resp.Success(c, comments)
}

// GetComment 通过token分页获取多级评论
func GetComment(c *gin.Context) {
	var token string = c.DefaultQuery("token", "")
	if token == "" {
		helper.Logger().Warning("token值为空")
		resp.Error(c, helper.ResponseParamError, "入参错误")
		return
	}
	page := helper.GetPagination(c)

	comments := commentModel.GetCommentByArticle(token, page)
	resp.Success(c, comments)
}

type deleteRequestJSON struct {
	ID uint `json:"id" binding:"required"`
}

// Delete 删除评论
func Delete(c *gin.Context) {
	data := &deleteRequestJSON{}
	err := c.Bind(&data)
	if err != nil {
		helper.Logger().Warning(err.Error())
		resp.Error(c, helper.ResponseParamError, "入参错误")
		return
	}
	commentModel.Delete(data.ID)
	resp.Success(c, true)
}

// Save 保存评论
func Save(c *gin.Context) {
	var data model.Comment
	err := c.Bind(&data)
	if err != nil {
		helper.Logger().Warning(err.Error())
		resp.Error(c, helper.ResponseParamError, "入参错误")
		return
	}
	data.IP = c.ClientIP()
	notBlockWord, _ := helper.SensitiveValid(data.Content)
	if !notBlockWord {
		resp.Error(c, helper.ResponseParamError, "提交内容含敏感内容")
		return
	}
	data.Content = template.HTMLEscapeString(data.Content)
	commentModel.Save(data)
	resp.Success(c, true)
	if helper.Config.SMTPEnabled == true {
		sendEmail(data)
	}
}

func sendEmail(data model.Comment) {
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(helper.Config.SMTPUsername, "KBComment"))
	m.SetHeader("To", helper.Config.SMTPTo)
	m.SetHeader("Subject", "[KB-Comment]你有一条新的留言")
	m.SetBody("text/html", "["+fmt.Sprint(data.ID)+"]<a href=\""+data.PageLink+"\">"+data.ArticleToken+"</a><br/>"+data.NickName+"&lt;"+data.Mail+"&gt;:"+data.Content)

	d := gomail.NewDialer(helper.Config.SMTPHost, helper.Config.SMTPPort, helper.Config.SMTPUsername, helper.Config.SMTPPassword)
	err := d.DialAndSend(m)
	if err != nil {
		helper.Logger().Error(err.Error())
	}
}
