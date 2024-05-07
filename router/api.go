package router

import (
	controllers "ErgoGo/internal/http/controllers/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// 路由组
	v1 := r.Group("/v1")

	v1.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	user := new(controllers.UserController)
	// 获取用户列表
	v1.GET("user", user.Index)
}
