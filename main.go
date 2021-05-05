package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 加载CSS文件
	r.Static("/static", "static")

	// 加载html文件
	r.LoadHTMLGlob("templates/*")
	// 创建路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run()

}
