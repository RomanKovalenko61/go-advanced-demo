package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	code := make(chan int)
	for range 10 {
		go getHttpCode(code)
	}
	for res := range code {
		fmt.Printf("Код: %d\n", res)
	}
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
	codeCh <- resp.StatusCode
	fmt.Println("Готово!")
}
