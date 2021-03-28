package comment

import (
	"YozComment/model"
	"YozComment/statics"
	"YozComment/util"
	"bytes"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/importcjj/sensitive"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var commentModel = model.Comment{}
var resp = util.Response{}

// GetComment 通过token分页获取多级评论
func GetComment(c *gin.Context) {
	var token string = c.DefaultQuery("token", "")
	if token == "" {
		resp.Error(c, util.ResponseParamError, "入参错误")
		return
	}
	page := util.GetPagination(c)

	comments := commentModel.GetCommentByArticle(token, page)
	resp.Success(c, comments)
}

// Save 保存评论
func Save(c *gin.Context) {
	var data model.Comment
	err := c.Bind(&data)
	if err != nil {
		resp.Error(c, util.ResponseParamError, "入参错误")
		return
	}
	data.IP = c.ClientIP()
	if util.Config.SensitiveEnabled == true {
		notBlockWord, _ := sensitiveValid(data.Content)
		if !notBlockWord {
			resp.Error(c, util.ResponseParamError, "提交内容含敏感内容")
			return
		}
	}
	data.NickName = template.HTMLEscapeString(data.NickName)
	data.Site = template.HTMLEscapeString(data.Site)
	data.Mail = template.HTMLEscapeString(data.Mail)
	data.Content = template.HTMLEscapeString(data.Content)
	commentModel.Save(data)

	resp.Success(c, true)

	if util.Config.SMTPEnabled == true {
		err := sendEmail(data)
		if err != nil {
			log.Errorf("邮件通知 %s", err.Error())
		}
	}
}

// sensitiveValid 敏感字验证
func sensitiveValid(content string) (bool, string) {
	filter := sensitive.New()
	filter.LoadWordDict(util.Config.SensitivePath)
	return filter.Validate(content)
}

func sendEmail(data model.Comment) (err error) {
	tmpl := template.New("")
	mailTmpl, err := statics.Asset("templates/manage/mail_notice.html")
	if err != nil {
		log.Errorf("加载邮件模板文件错误: %s", err.Error())
		return
	}
	tmpl, err = tmpl.Parse(string(mailTmpl))
	if err != nil {
		log.Errorf("解析邮件模板文件错误: %s", err.Error())
		return
	}
	var body bytes.Buffer
	err = tmpl.Execute(&body, struct {
		SiteName    string
		SiteUrl     string
		CommentUser string
		PostUrl     string
		PostTitle   string
		Content     string
	}{
		SiteName:    util.Config.SiteName,
		SiteUrl:     util.Config.SiteUrl,
		CommentUser: data.NickName,
		PostUrl:     data.PageUrl,
		PostTitle:   data.PageTitle,
		Content:     data.Content,
	})

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(util.Config.SMTPUsername, "YozComment"))
	m.SetHeader("To", util.Config.SMTPTo)
	m.SetHeader("Subject", "你的 "+util.Config.SiteName+" 有一条新留言")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(util.Config.SMTPHost, util.Config.SMTPPort, util.Config.SMTPUsername, util.Config.SMTPPassword)
	err = d.DialAndSend(m)
	return
}
