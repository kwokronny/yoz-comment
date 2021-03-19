package manage

import (
	"YozComment/model"
	"YozComment/util"

	"github.com/gin-gonic/gin"
)

var commentModel = model.Comment{}
var resp = util.Response{}

type getPageRequestQuery struct {
	NickName string `form:"nickName"`
	Mail     string `form:"mail"`
	Content  string `form:"content"`
}

// GetPage 获取所有问答并分页
func GetPage(c *gin.Context) {
	data := &getPageRequestQuery{}
	c.BindQuery(&data)
	page := util.GetPagination(c)
	comments := commentModel.GetPage(data.NickName, data.Mail, data.Content, page)
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
		resp.Error(c, util.ResponseParamError, "入参错误")
		return
	}
	commentModel.Delete(data.ID)
	resp.Success(c, true)
}
