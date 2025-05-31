package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tkdnbb/bookofben-api/internal/routes"
)

func main() {
	r := routes.SetupRoutes()

	// 设置优雅关闭
	go func() {
		fmt.Println("Bible API Server starting on :8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	// 关闭数据库连接
	if err := routes.CloseDatabase(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	fmt.Println("Server exited")
}
