package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	t := time.Now()
	for range 10 {
		go getHttpCode()
	}
	time.Sleep(time.Millisecond * 1100)
	fmt.Println(time.Since(t))
}

func getHttpCode() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode, len(body))
}
