package main

import (
	"fmt"
	"time"
)

const numRequest = 1000

var count int

func worker(result chan any, work func() int) {
	fmt.Println("Start")
	result <- work()
	fmt.Println("Finish")
}

func networkRequest() int {
	time.Sleep(time.Millisecond)
	count++
	return count
}
func main() {
	numWorkers := 10
	result := make(chan any, numRequest)

	for i := 0; i < numWorkers; i++ {
		go worker(result, networkRequest)
	}

	for i := 0; i < numRequest; i++ {
		<-result
	}

}
