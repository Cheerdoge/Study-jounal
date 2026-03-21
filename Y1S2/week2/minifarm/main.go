package main

import (
	"log"
	"minifarm/farm"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化农场管理器 (全局单例存储)
	fm := farm.NewFarmManager()

	// 初始化 Gin 框架
	r := gin.Default()

	// 托管静态前端页面
	r.StaticFile("/", "./static/index.html")

	// 注册路由
	handlers := farm.NewFarmHandlers(fm)
	handlers.RegisterRoutes(r)

	// 打印启动日志
	log.Println("迷你农场系统正在启动... 运行在 http://localhost:8080")
	log.Println("1. 种植：POST /api/farm/plant (body: {\"type\": \"白菜\"})")
	log.Println("2. 状态：GET /api/farm/status")
	log.Println("3. 互动：POST /api/farm/action (body: {\"id\": \"xxxx\", \"action\": \"浇水/除草虫/施肥/收获\"})")

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("启动失败: %v", err)
	}
}
