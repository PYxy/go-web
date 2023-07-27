package main

import (
	svc "github.com/PYxy/go-web/internal/customer-app"
	//logrus 初始化
	_ "github.com/PYxy/go-web/pkg/logger"
)

func main() {
	svc.NewApp("web-1", ":8080").RUN()
}
