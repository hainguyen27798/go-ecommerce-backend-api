package main

import "github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/router"

func main() {
	r := router.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
