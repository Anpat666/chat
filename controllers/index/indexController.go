package index

import "github.com/gin-gonic/gin"

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "登录页面",
	})
}

func (con IndexController) IndexOk(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "测试",
	})
}
