package api

import (
	"ErgoGo/internal/http/controllers"
	"ErgoGo/internal/http/requests"
	"ErgoGo/internal/model/user"
	"ErgoGo/pkg/http/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllers.Controller
}

func (ctrl *UserController) Index(c *gin.Context) {
	// 处理请求
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	// 读取数据库数据
	data, pager := user.Paginate(c, 3)

	// 返回数据
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
