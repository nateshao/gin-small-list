package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {

	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()        // 程序退出关闭数据库
	DB.AutoMigrate(&Todo{}) // 创建表 todos
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
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项 点击提交，会发送请求到这里
			// 1. 从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 2. 存入数据库
			//err = DB.Create(&todo).Error
			//if err != nil {
			//	panic(err)
			//}
			// 3. 返回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//// 查看单个待办事项
		//v1Group.POST("/todo:id", func(c *gin.Context) {
		//
		//})
		//// 添加
		//v1Group.POST("/todo", func(c *gin.Context) {
		//
		//
		//})
		//// 修改
		//v1Group.PUT("/todo:id", func(c *gin.Context) {
		//
		//})
		//// 删除
		//v1Group.DELETE("/todo:id", func(c *gin.Context) {
		//
		//})
	}

	r.Run()

}
