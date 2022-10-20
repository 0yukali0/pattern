package solution

import (
	"fmt"
	"time"
)

const (
	THREE_SECOND = 3 * time.Second
	FIVE_SECOND = 5 * time.Second
)

func PushNews(news string, startTime time.Time) <-chan time.Time {
	newsCh := make(chan time.Time)
	go func() {
		time.Sleep(THREE_SECOND)
		fmt.Printf("%s cost %s\n", news, time.Since(startTime))
		newsCh <- time.Now()
	}()
	return newsCh
}
