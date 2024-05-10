package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler 是 Vercel 需要的导出函数
func Handler(w http.ResponseWriter, r *http.Request) {
	// 创建 Gin 路由器
	router := gin.Default()

	// 设置静态文件服务
	// 从 'static' 文件夹中服务文件，URL 路径也是 '/static'
	// router.Static("/static", "./static")

	// 根路径请求处理
	router.GET("/sss", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Gin server!"})
	})

	// 将 Gin 的处理器适配到标准 net/http
	http.Handle("/sss", router)
	router.ServeHTTP(w, r)
}
