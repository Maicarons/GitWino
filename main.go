package main

import (
	handler "gitwino/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 提供前端静态文件
	r.Static("/frontend", "./frontend")

	// API路由
	r.GET("/api", handler.GinHandler)

	// 根路径重定向到前端
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/frontend/index.html")
	})

	// 启动服务器
	r.Run(":8080")
}
