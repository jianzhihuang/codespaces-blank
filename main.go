package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/", "./public")

	// 其他路由和处理器
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
