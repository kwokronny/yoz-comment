package comment

import (
	"YozComment/model"
	"YozComment/plugins"
	"YozComment/util"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/importcjj/sensitive"
	log "github.com/sirupsen/logrus"
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
		err := plugins.SendEmail(data)
		if err != nil {
			log.Errorf("邮件通知 %s", err.Error())
		}
	}

	if util.Config.SendCloudEnabled == true && data.RID > 0 {
		beComment := commentModel.GetComment(data.RID)
		err := plugins.SendCloud(beComment, data)
		if err != nil {
			log.Errorf("邮件通知留言者 %s", err.Error())
		}
	}
}

// sensitiveValid 敏感字验证
func sensitiveValid(content string) (bool, string) {
	filter := sensitive.New()
	filter.LoadWordDict(util.Config.SensitivePath)
	return filter.Validate(content)
}
