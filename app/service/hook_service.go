package service

import (
	"fmt"

	"github.com/fuxingjun/gohook/app/util"
	"github.com/gofiber/fiber/v2"
)

type larkTextModel struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

func HandleLark(c *fiber.Ctx) error {
	key := c.Query("key", "")
	if key == "" {
		return c.JSON(fiber.Map{
			"msg":  "key is required",
			"code": -1,
			"data": nil,
		})
	}
	// 解析JSON数据
	var msg larkTextModel
	if err := c.BodyParser(&msg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  "无法解析消息",
			"code": -1,
			"data": nil,
		})
	}
	fmt.Println(msg)
	// 发送消息
	config := util.GetAppConfig("")
	count := 0
	for _, v := range *config {
		if v.HookFrom.Platform == "lark" && v.HookFrom.Key == key {
			for _, to := range v.HookToList {
				if to.Platform == "qywx" {
					count = count + 1
					go func() {
						res, err := util.QYWXSendTextMsg(to.Webhook, msg.Content.Text)
						if err != nil {
							fmt.Println(err)
						} else {
							fmt.Println(string(res))
						}
					}()
				}
			}
		}
	}

	return c.JSON(fiber.Map{
		"msg":  "success",
		"code": 0,
		"data": count,
	})
}
