package HomeController

import "github.com/gin-gonic/gin"

// HomeController 控制器

// Index 首页
func Index(c *gin.Context) {
	c.JSON(200, "This is Index")
}

// Hi 首页
func Hi(c *gin.Context) {
	c.JSON(200, " Welcome to your from my  Blog!")
}
