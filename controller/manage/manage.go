package manage

import (
	"YozComment/middleware"
	"YozComment/model"
	"YozComment/util"

	"github.com/gin-gonic/gin"
)

var commentModel = model.Comment{}
var resp = util.Response{}

type getPageRequestQuery struct {
	NickName  string `form:"nickName"`
	Mail      string `form:"mail"`
	Content   string `form:"content"`
	PageTitle string `form:"pageTitle"`
}

// GetPage 获取所有问答并分页
func GetPage(c *gin.Context) {
	data := model.QueryCommentField{}
	c.BindQuery(&data)
	page := util.GetPagination(c)
	comments := commentModel.GetPage(data, page)
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

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var admin LoginParams
	err := c.ShouldBind(&admin)
	if err != nil {
		resp.Error(c, util.ResponseParamError, "入参错误")
		return
	}
	// 校验用户名和密码是否正确
	if admin.Username == util.Config.AdminRoot && admin.Password == util.Config.AdminPass {
		// 生成Token
		tokenString, _ := middleware.GenerateToken()
		resp.Success(c, tokenString)
		return
	}
	resp.Error(c, util.ResponseParamError, "用户名或密码错误")
	return
}
