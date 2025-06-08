package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
	"github.com/tkdnbb/bookofben-api/internal/routes"
)

// FunctionEvent 定义函数计算的事件结构
type FunctionEvent struct {
	HTTPMethod      string            `json:"httpMethod"`
	Path            string            `json:"path"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	QueryParameters map[string]string `json:"queryParameters"`
}

// FunctionResponse 定义函数计算的响应结构
type FunctionResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

var router http.Handler

func HandleRequest(ctx context.Context, event []byte) (*FunctionResponse, error) {
	fctx, _ := fccontext.FromContext(ctx)
	logger := fctx.GetLogger()

	// 解析事件
	var fcEvent FunctionEvent
	if err := json.Unmarshal(event, &fcEvent); err != nil {
		logger.Errorf("Error parsing event: %v", err)
		return &FunctionResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "Invalid request format"}`,
		}, nil
	}

	logger.Infof("Processing request: %s %s", fcEvent.HTTPMethod, fcEvent.Path)

	// 创建 HTTP 请求
	req, err := createHTTPRequest(fcEvent)
	if err != nil {
		logger.Errorf("Error creating HTTP request: %v", err)
		return &FunctionResponse{
			StatusCode: 500,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "Internal server error"}`,
		}, nil
	}

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 处理请求
	router.ServeHTTP(w, req)

	// 构建函数计算响应
	headers := make(map[string]string)
	for key, values := range w.Header() {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	response := &FunctionResponse{
		StatusCode: w.Code,
		Headers:    headers,
		Body:       w.Body.String(),
	}

	logger.Infof("Response status: %d", w.Code)
	return response, nil
}

func createHTTPRequest(event FunctionEvent) (*http.Request, error) {
	// 构建完整的 URL
	url := event.Path
	if len(event.QueryParameters) > 0 {
		query := make([]string, 0, len(event.QueryParameters))
		for key, value := range event.QueryParameters {
			query = append(query, fmt.Sprintf("%s=%s", key, value))
		}
		url += "?" + strings.Join(query, "&")
	}

	// 创建请求
	req, err := http.NewRequest(event.HTTPMethod, url, strings.NewReader(event.Body))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for key, value := range event.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func initialize(ctx context.Context) {
	fctx, _ := fccontext.FromContext(ctx)
	logger := fctx.GetLogger()

	logger.Info("Initializing Bible API Server...")

	// 初始化路由和数据库连接
	router = routes.SetupRoutes()

	logger.Info("Bible API Server initialized successfully")
}

func preStop(ctx context.Context) {
	fctx, _ := fccontext.FromContext(ctx)
	logger := fctx.GetLogger()

	logger.Info("Shutting down Bible API Server...")

	// 关闭数据库连接
	if err := routes.CloseDatabase(); err != nil {
		logger.Errorf("Error closing database: %v", err)
	} else {
		logger.Info("Database connection closed successfully")
	}

	logger.Info("Bible API Server shutdown completed")
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.RegisterPreStopFunction(preStop)
	fc.Start(HandleRequest)
}
