package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	code := make(chan int)
	go getHttpCode(code)
	<-code
}

func getHttpCode(codeCh chan int) {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	codeCh <- resp.StatusCode
}
