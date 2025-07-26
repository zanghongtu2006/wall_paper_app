package main

import (
	"github.com/zanghongtu2006/wallet_paper_backend/config"
	"github.com/zanghongtu2006/wallet_paper_backend/model"
	"github.com/zanghongtu2006/wallet_paper_backend/router"
)

func main() {
    
    // 初始化数据库
    config.InitDB()

    // 自动迁移：注册所有模型（按顺序避免外键依赖错误）
    config.DB.AutoMigrate(
        &model.User{},
        &model.Category{},
        &model.Wallpaper{},
    )

    r := router.InitRouter()
    r.Run()
}
