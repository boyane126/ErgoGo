package bootstrap

import (
	"ErgoGo/internal/http/middlewares"
	"ErgoGo/pkg/http/middleware"
	"ErgoGo/router"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func SetupRoute(r *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(r)

	// 注册 API 路由
	router.RegisterAPIRoutes(r)

	// 配置 404 路由
	setup404Handler(r)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middleware.Logger(),
		middleware.Recovry(),
		middlewares.ForceUA(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理404请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})
}
