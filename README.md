# hookgo
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
nohup ./release/hookgo_linux_arm64 > logs/hookgo.out 2>&1 &

# 支持指定端口, 默认 15492
./release/hookgo_linux_arm64 -port=8082

# 启动程序后测试
curl -X POST -H "Content-Type: application/json" \
-d '{"msg_type":"text","content":{"text":"test message\nhello lark"}}' \
http://localhost:15492/hook/lark?key=your-custom-key

```