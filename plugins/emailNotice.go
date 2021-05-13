package plugins

import (
	"YozComment/model"
	"YozComment/statics"
	"YozComment/util"
	"bytes"
	"text/template"

	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

// SendEmail 发送评论通知给站长有新的评论
func SendEmail(data model.Comment) (err error) {
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
		SiteURL     string
		CommentUser string
		PostURL     string
		PostTitle   string
		Content     string
	}{
		SiteName:    util.Config.SiteName,
		SiteURL:     util.Config.SiteURL,
		CommentUser: data.NickName,
		PostURL:     data.PageURL,
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
