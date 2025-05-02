package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/hello"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listenning on port 8081")
	server.ListenAndServe()
}
