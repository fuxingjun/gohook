package route

import (
	"github.com/fuxingjun/hookgo/app/service"
	"github.com/gofiber/fiber/v2"
)

func HookRoute(router fiber.Router) {
	// 分组前缀 /hook
	hookGroup := router.Group("/hook")

	hookGroup.Post("/lark", service.HandleLark)
}
