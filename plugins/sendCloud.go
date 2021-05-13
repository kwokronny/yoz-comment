package plugins

import (
	"YozComment/model"
	"YozComment/util"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type xsmtpAPI struct {
	To  []string            `json:"to"`
	Sub map[string][]string `json:"sub"`
}

// SendCloud 当评论者收到回复时应用sendCloud平台发信通知评论者在此站的评论有回复，请查看
func SendCloud(beComment model.Comment, reply model.Comment) (err error) {
	xs := xsmtpAPI{
		To: []string{beComment.Mail},
		Sub: map[string][]string{
			"%post%":           {beComment.PageTitle},
			"%post_url%":       {beComment.PageURL},
			"%you%":            {beComment.NickName},
			"%you_comment%":    {beComment.Content},
			"%comment_author%": {reply.NickName},
			"%comment%":        {reply.Content},
		},
	}
	xsmtpAPIValue, _ := json.Marshal(xs)
	RequestURI := "http://api.sendcloud.net/apiv2/mail/sendtemplate"
	PostParams := url.Values{
		"apiUser":            {util.Config.SendCloudAPIUser},
		"apiKey":             {util.Config.SendCloudAPIKey},
		"from":               {util.Config.SendCloudFrom},
		"xsmtpapi":           {string(xsmtpAPIValue)},
		"templateInvokeName": {util.Config.SendCloudTemplateName},
	}
	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		log.Errorf("SendCloud API请求失败: %s", err.Error())
		return
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		log.Errorf("SendCloud API请求转码失败: %s", err.Error())
		return
	}
	log.Infof("SendCloud API调用返回: %s", BodyByte)
	return
}
