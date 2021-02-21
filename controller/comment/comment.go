package comment

import (
	"KBCommentAPI/helper"
	"KBCommentAPI/model"

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

// GetComment 获取所有问答并分页
func GetComment(c *gin.Context) {
	var token string = c.DefaultQuery("token", "")
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
	if c.Bind(&data) == nil {
		commentModel.Delete(data.ID)
		resp.Success(c, true)
	}
	resp.Error(c, helper.ResponseParamError, "入参错误")
}

// Save 保存评论
func Save(c *gin.Context) {
	var data model.Comment
	if c.Bind(&data) != nil {
		resp.Error(c, helper.ResponseParamError, "入参错误")
		return
	}
	data.IP = c.ClientIP()
	notBlockWord, _ := helper.SensitiveValid(data.Content)
	if !notBlockWord {
		resp.Error(c, helper.ResponseParamError, "提交内容含敏感内容")
		return
	}
	commentModel.Save(data)
	resp.Success(c, true)
	if helper.Config.SMTPEnabled == true {
		sendEmail(data.Content)
	}
}

func sendEmail(content string) {
	m := gomail.NewMessage()
	m.SetHeader("From", helper.Config.SMTPUsername)
	m.SetHeader("To", helper.Config.SMTPTo)
	m.SetHeader("Subject", "[KB-Comment]你有一条新的留言")
	m.SetBody("text/html", content)

	d := gomail.NewDialer(helper.Config.SMTPHost, helper.Config.SMTPPort, helper.Config.SMTPUsername, helper.Config.SMTPPassword)
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}
