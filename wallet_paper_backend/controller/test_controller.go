package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zanghongtu2006/wallet_paper_backend/service"
)

func TestHandler(c *gin.Context) {
	msg := service.GetTestMessage()
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
