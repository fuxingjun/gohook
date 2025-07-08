# gohook
一个简单的webhook转发工具


没有配置文件时执行将生成示例配置

```json
[
  {
    "from": {
      "platform": "lark",
      "key": "your-custom-key"
    },
    "to": [
      {
        "platform": "qywx",
        "webhook": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx"
      }
    ]
  }
]

```

```bash
# 后台运行程序
nohup ./release/gohook_linux_arm64 > logs/gohook.out 2>&1 &

# 启动程序后测试
curl -X POST -H "Content-Type: application/json" \
-d '{"msg_type":"text","content":{"text":"test message\nhello lark"}}' \
http://localhost:15492/hook/lark?key=your-custom-key

```