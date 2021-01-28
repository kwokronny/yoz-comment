package comment

import (
	"kwok-comment/helper"
	"kwok-comment/model"
	"strconv"

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
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultPostForm("id", ""))
	if err == nil {
		commentModel.Delete(uint(id))
		resp.Success(c, true)
	} else {
		c.Status(404)
	}
}

// Save 保存评论
func Save(c *gin.Context) {
	var data model.Comment
	err := c.ShouldBind(&data)
	if err != nil {
		resp.Error(c, helper.ResponseParamError, "入参错误")
		return
	}
	notBlockWord, _ := helper.Filter.Validate(data.Content)
	if !notBlockWord {
		resp.Error(c, helper.ResponseNotFound, "提交内容含违法违规内容")
	} else {
		commentModel.Save(data)
		resp.Success(c, true)
	}
}
