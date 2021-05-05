package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 连接数据库
	r := gin.Default()

	// 加载CSS文件
	r.Static("/static", "static")

	// 加载html文件
	r.LoadHTMLGlob("templates/*")
	// 创建路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		// 查看
		// 删除
	}

	r.Run()

}
