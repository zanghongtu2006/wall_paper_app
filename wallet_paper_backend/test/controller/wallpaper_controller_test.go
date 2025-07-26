package controller_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zanghongtu2006/wallet_paper_backend/controller"
	"github.com/zanghongtu2006/wallet_paper_backend/test"
	_ "github.com/zanghongtu2006/wallet_paper_backend/test"
)

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/api/wallpapers", controller.GetWallpaperList)
    return r
}

func TestGetWallpaperList(t *testing.T) {
    router := setupRouter()
    resp := test.PerformRequest(router, "GET", "/api/wallpapers", nil)

    assert.Equal(t, http.StatusOK, resp.Code)
    assert.Contains(t, resp.Body.String(), `"code":0`)
}
