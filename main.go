package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 設定一個簡單的路由
	// 當訪問根 URL "/" 時，執行的函數
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// 運行HTTP服務器在本地機器上的 8080 端口
	router.Run() // 預設監聽並服務於 0.0.0.0:8080
}
