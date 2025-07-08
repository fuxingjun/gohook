package main

import (
	"flag"
	"fmt"

	"github.com/fuxingjun/hookgo/app/route"
	"github.com/fuxingjun/hookgo/app/util"
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

	port := flag.Int("port", 15492, "port")
	flag.Parse() // ✅ 必须调用以解析命令行参数

	addr := fmt.Sprintf(":%d", *port)
	// 启动服务器在 指定 端口
	fmt.Printf("Listening on %s\n", addr)
	if err := app.Listen(addr); err != nil {
		fmt.Printf("listen failed: %v\n", err)
	}
}
