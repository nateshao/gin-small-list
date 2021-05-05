package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Todo model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

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
		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		// 查看单个待办事项
		v1Group.GET(("/todo:id", func(c *gin.Context) {

		}))
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			
			
		})
		// 修改
		v1Group.PUT("/todo:id", func(c *gin.Context) {
			
		})
		// 删除
		v1Group.DELETE("/todo:id", func(c *gin.Context) {
			
		})
	}

	r.Run()

}
