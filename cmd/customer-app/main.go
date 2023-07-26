package main

import svc "github.com/PYxy/go-web/internal/customer-app"

func main() {
	svc.NewApp("web-1", ":8080").RUN()
}
