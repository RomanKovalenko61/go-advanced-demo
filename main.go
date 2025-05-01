package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoriutines := 3
	res := make(chan int, numGoriutines)
	partSize := len(arr) / numGoriutines
	for i := range numGoriutines {
		start := i * partSize
		end := start + partSize
		go sumPart(res, arr[start:end])
	}
	totalSum := 0
	for range numGoriutines {
		totalSum += <-res
	}
	fmt.Printf("Итог вычисления: %v", totalSum)
}

func sumPart(ch chan int, slice []int) {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	ch <- sum
}
