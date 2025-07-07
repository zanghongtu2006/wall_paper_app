package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zanghongtu2006/wallet_paper_backend/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		api.GET("/test", controller.TestHandler)
		api.GET("/wallpapers", controller.GetWallpaperList) // ✅ 添加壁纸接口

	}

	return r
}
