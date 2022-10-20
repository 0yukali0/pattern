package main

import (
	"fmt"
	"time"
)

const (
	THREE_SECOND = 3 * time.Second
	FIVE_SECOND = 5 * time.Second
)

func PushNews(news string, startTime time.Time) {
	time.Sleep(time.Duration(THREE_SECOND))
	fmt.Printf("%s cost %s\n", news, time.Since(startTime))
}

func main() {
	allNews := []string {
		"John is here",
		"Tim is here",
		"Sam is here",
	}
	
	start := time.Now()
	for _, news := range allNews {
		PushNews(news, start)
	}
	fmt.Printf("Cost %s\n", time.Since(start))
	
	time.Sleep(THREE_SECOND)
	fmt.Print("Solution\n")
	start = time.Now()
	for _, news := range allNews {
		go PushNews(news, start)
	}
	time.Sleep(FIVE_SECOND)
	fmt.Printf("Cost %s\n", time.Since(start))
}
