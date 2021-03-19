package comment

import (
	"YozComment/model"
	"YozComment/util"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/importcjj/sensitive"
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
	notBlockWord, _ := sensitiveValid(data.Content)
	if !notBlockWord {
		resp.Error(c, util.ResponseParamError, "提交内容含敏感内容")
		return
	}
	data.Content = template.HTMLEscapeString(data.Content)
	commentModel.Save(data)

	resp.Success(c, commentModel.GetComment(data.RID))
	// if util.Config.SMTPEnabled == true {
	// sendEmail(data)
	// }
}

// sensitiveValid 敏感字验证
func sensitiveValid(content string) (bool, string) {
	filter := sensitive.New()
	filter.LoadWordDict(util.Config.SensitivePath)
	return filter.Validate(content)
}

// func sendEmail(data model.Comment) {
// 	m := gomail.NewMessage()

// 	m.SetHeader("From", m.FormatAddress(util.Config.SMTPUsername, "KBComment"))
// 	m.SetHeader("To", util.Config.SMTPTo)
// 	m.SetHeader("Subject", "[KB-Comment]你有一条新的留言")

// const mailTmpl = `<span id="9999" style="display: none !important; font-size:0; line-height:0">你在 {{.BlogName}} 博客上的留言有回复啦</span><div style="background-color:white;border-top:2px solid #12ADDB;box-shadow:0 1px 3px #AAAAAA; line-height:180%; padding:0 15px 12px;width:500px;margin:100px auto;color:#555555;font-family:Century Gothic,Trebuchet MS,Hiragino Sans GB,微软雅黑,Microsoft Yahei,Tahoma,Helvetica,Arial,SimSun,sans-serif;font-size:14px;"><h2 style="border-bottom:1px solid #DDD;font-size:16px;font-weight:normal;padding:13px 0 10px 0;"><span style="color: #12ADDB;font-weight: bold;">&gt; </span>你在 <a href="{{.BlogUrl}}" style="text-decoration:none;color: #12ADDB;" target="_blank">{{.BlogName}}</a> 博客上的留言有回复啦！</h2>	<div style="padding:0 12px 0 12px;margin-top:18px">
// {{.CommentAuth}} 同学，你在文章《<a href="{{.PostUrl}}" style="text-decoration:none; color:#12addb" target="_blank">{{.PostName}}</a>》上的评论:
// <p style="background-color: #f5f5f5;border: 0 solid #DDD;padding: 10px 15px;margin:18px 0">%you_comment%</p>%comment_author% 给你的回复如下:	<p style="background-color: #f5f5f5;border: 0 solid #DDD;padding: 10px 15px;margin:18px 0">{{.Content}}</p>你可以点击 <a href="{{.PostUrl}}" style="text-decoration:none; color:#12addb" target="_blank">查看回复的完整內容 </a>，欢迎再来玩呀~</div></div>`

// tmpl,err:=template.New("webpage").Parse(mailTmpl)
// if err!=nil{
// 	print(err)
// }
// tmpl.Execute(os.Stdout, )
// m.SetBody("text/html", "["+fmt.Sprint(data.ID)+"]<a href=\""+data.pageUrl+"\">"+data.ArticleToken+"</a><br/>"+data.NickName+"&lt;"+data.Mail+"&gt;:"+data.Content)

// d := gomail.NewDialer(util.Config.SMTPHost, util.Config.SMTPPort, util.Config.SMTPUsername, util.Config.SMTPPassword)
// err := d.DialAndSend(m)
// if err != nil {
// middleware.Logger().Error(err.Error())
// 	}
// }
