package comment

import (
	"kwok-comment/helper"
	"kwok-comment/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var commentModel = model.Comment{}
var resp = helper.Response{}

// GetPage 获取所有问答并分页
func GetPage(c *gin.Context) {
	page := helper.GetPagination(c)
	comments := commentModel.GetPage(page)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"data": comments,
	})
}

// GetComment 获取所有问答并分页
func GetComment(c *gin.Context) {
	var token string = c.DefaultQuery("token", "")
	page := helper.GetPagination(c)

	comments := commentModel.GetCommentByArticle(token, page)
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
	notBlockWord, text := helper.SensitiveValid(data.Content)
	println(text)
	if !notBlockWord {
		resp.Error(c, helper.ResponseParamError, "提交内容含敏感内容")
		return
	}
	commentModel.Save(data)
	resp.Success(c, true)
}
