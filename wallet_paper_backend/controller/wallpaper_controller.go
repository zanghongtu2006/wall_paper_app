package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zanghongtu2006/wallet_paper_backend/service"
)

func GetWallpaperList(c *gin.Context) {
    data := service.GetWallpaperList()
    c.JSON(http.StatusOK, gin.H{
        "code": 0,
        "msg": "ok",
        "data": data,
    })
}
