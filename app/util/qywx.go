package util

func QYWXSendTextMsg(webhook, msg string) ([]byte, error) {
	httpClient = GetHTTPClient()
	data := map[string]any{
		"msgtype": "text",
		"text": map[string]string{
			"content": msg,
		},
	}
	return httpClient.SendPostRequest(webhook, data, nil, nil)
}
