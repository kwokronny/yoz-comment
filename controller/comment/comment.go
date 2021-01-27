package comment

import (
	"kwok-comment/helper"
	"kwok-comment/model"

	"github.com/gin-gonic/gin"
)

var commentModel = model.Comment{}
var resp = helper.Response{}

// GetPage 获取所有问答并分页
func GetPage(c *gin.Context) {
	// authGet, _ := c.Get("auth")
	// var auth middleware.Auth
	// auth = authGet.(middleware.Auth)

	var token string = c.DefaultQuery("token", "")
	page := helper.GetPagination(c)

	comments := commentModel.GetArticleComment(token, page)
	resp.Success(c, comments)
}

// Delete 删除评论
// func Delete(c *gin.Context) {
// authGet, _ := c.Get("auth")
// var auth middleware.Auth
// auth = authGet.(middleware.Auth)
// user := userModel.Get(auth.OpenID)
// id, err := strconv.Atoi(c.DefaultPostForm("id", ""))
// if user.Role == 1 && err == nil {
// 	questionModel.Delete(uint(id))
// 	resp.Success(c, gin.H{})
// } else {
// 	c.Status(404)
// }
// }

// Save 保存评论
// func Save(c *gin.Context) {
// 	authGet, _ := c.Get("auth")
// 	var auth middleware.Auth
// 	auth = authGet.(middleware.Auth)
// 	user := userModel.Get(auth.OpenID)
// 	var data model.Question
// 	err := c.ShouldBind(&data)
// 	if err != nil {
// 		resp.Error(c, helper.ResponseParamError, "入参错误")
// 		return
// 	}
// 	var checkContent string
// 	if data.ID != 0 && data.Answer != "" {
// 		if user.Role == 0 {
// 			resp.Error(c, helper.ResponseAuthorized, "无权限执行该操作")
// 			return
// 		}
// 		checkContent = data.Answer
// 	} else if data.Ask != "" && auth.OpenID != "" {
// 		checkContent = data.Ask
// 		data.OpenID = auth.OpenID
// 	}
// 	notBlockWord, err := wechat.MsgSecCheck(checkContent)
// 	if err != nil {
// 		resp.Error(c, helper.ResponseServerError, "微信SDK服务请求失败")
// 	} else if !notBlockWord {
// 		resp.Error(c, helper.ResponseNotFound, "提交内容含违法违规内容")
// 	} else {
// 		data := questionModel.Save(data)
// 		resp.Success(c, true)
// 		if data.ID > 0 {
// 			sendSubscribeMsg(data.ID)
// 		}
// 	}
// }
