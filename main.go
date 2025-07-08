package main

import (
	"fmt"

	"github.com/fuxingjun/gohook/app/route"
	"github.com/fuxingjun/gohook/app/util"
	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	route.HookRoute(app)
}

var (
	version = "dev"
	date    = "unknown"
)

func main() {
	fmt.Printf("版本: %s, 构建时间: %s\n", version, date)
	util.GetAppConfig("config.json")
	fmt.Println("hello, go-fiber! 🌟")
	// 创建 Fiber 应用实例
	app := fiber.New()

	// 静态文件服务
	app.Static("/assets", "./public")

	// 基本路由
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber! 🚀")
	})

	SetupRoutes(app)

	// 启动服务器在 15492 端口
	app.Listen(":15492")
}
