package Routers

import (
	"github.com/gin-gonic/gin"
	"golang-blog/Controller/HomeController"
)

// 路由器

func Init(router *gin.Engine) {
	// 获取组
	home := router.Group("Home")
	// 不区分大小写，并重定向到对应的handler
	router.RedirectFixedPath = true
	{
		home.GET("/", HomeController.Index)
		home.GET("/hi", HomeController.Hi)
	}
	// 启动服务
	_ = router.Run(":8888")
}
