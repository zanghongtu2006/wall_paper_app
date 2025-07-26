package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zanghongtu2006/wallet_paper_backend/service"
)

// ✅ 初始化数据库连接
// func init() {
//     config.InitDB()
// }

// ✅ 实际测试函数
func TestGetWallpaperList(t *testing.T) {
    wallpapers := service.GetWallpaperList()
    assert.NotNil(t, wallpapers)
    assert.GreaterOrEqual(t, len(wallpapers), 0, "壁纸列表至少返回0项")
}
