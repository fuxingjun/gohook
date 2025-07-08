package util

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type HTTPClient struct {
	client *fasthttp.Client
}

func NewHTTPClient(maxConnsPerHost int) *HTTPClient {
	return &HTTPClient{
		client: &fasthttp.Client{
			MaxConnsPerHost:     maxConnsPerHost,
			MaxIdleConnDuration: 10 * time.Second,
		},
	}
}

func (a *HTTPClient) SendGetRequest(url string, params map[string]any, headers map[string]string) ([]byte, error) {
	// 构造查询字符串，过滤空值
	queryString := ToQueryStrWithoutEncode(params)
	// 构造完整 URL
	fullUrl := url
	if queryString != "" {
		fullUrl += "?" + queryString
	}

	// 创建请求对象
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fullUrl)

	req.Header.Set("Accept", "application/json")
	for k, v := range headers {
		if v != "" {
			req.Header.Set(k, v)
		}
	}

	req.Header.SetMethod("GET")

	// 创建响应对象
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 发送请求
	if err := a.client.Do(req, resp); err != nil {
		return nil, err
	}

	// 检查响应状态码
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, &StructuredError{
			Code:    resp.StatusCode(),
			Message: fmt.Sprintf("request failed with status code: %d, response: %s", resp.StatusCode(), resp.Body()),
			Details: map[string]any{
				"response": string(resp.Body()),
			},
		}
	}

	return resp.Body(), nil
}

func (a *HTTPClient) SendPostRequest(url string, payload, params map[string]any, headers map[string]string) ([]byte, error) {
	// 构造查询字符串，过滤空值
	queryString := ToQueryStrWithoutEncode(params)

	payloadString := ""
	if payload != nil {
		// 过滤 nil 值
		filtered := make(map[string]any)
		for k, v := range payload {
			if v != nil {
				filtered[k] = v
			}
		}

		// 序列化 payload
		payloadBytes, err := json.Marshal(filtered)
		if err != nil {
			return nil, err
		}
		payloadString = string(payloadBytes)
	}

	// 构造完整 URL
	fullUrl := url
	if queryString != "" {
		fullUrl += "?" + queryString
	}

	// 创建请求对象
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fullUrl)

	req.Header.Set("Accept", "application/json")
	for k, v := range headers {
		if v != "" {
			req.Header.Set(k, v)
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBodyString(payloadString)
	req.Header.SetMethod("POST")

	// 创建响应对象
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 发送请求
	if err := a.client.Do(req, resp); err != nil {
		return nil, err
	}

	// 检查响应状态码
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, &StructuredError{
			Code:    resp.StatusCode(),
			Message: fmt.Sprintf("request failed with status code: %d, response: %s", resp.StatusCode(), resp.Body()),
			Details: map[string]any{
				"response": string(resp.Body()),
			},
		}
	}

	return resp.Body(), nil
}

var (
	httpClient     *HTTPClient
	httpClientOnce sync.Once
)

func initHTTPClient(maxConnsPerHost int) {
	httpClientOnce.Do(func() {
		httpClient = NewHTTPClient(maxConnsPerHost)
	})
}

func GetHTTPClient() *HTTPClient {
	if httpClient == nil {
		initHTTPClient(200)
	}
	return httpClient
}
