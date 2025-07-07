package main

import (
	"github.com/zanghongtu2006/wallet_paper_backend/config"
	"github.com/zanghongtu2006/wallet_paper_backend/model"
	"github.com/zanghongtu2006/wallet_paper_backend/router"
)

func main() {
    config.InitDB()

    // 自动建表
    config.DB.AutoMigrate(&model.Wallpaper{})

    r := router.InitRouter()
    r.Run()
}
