package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/", "./public")

	r.Run() // listen and serve on 0.0.0.0:8080
}
