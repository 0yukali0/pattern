package problem

import (
	"time"
	"fmt"
)

const (
	THREE_SECOND = 3 * time.Second
	FIVE_SECOND = 5 * time.Second
)

func PushNews(news string, startTime time.Time) {
	time.Sleep(time.Duration(THREE_SECOND))
	fmt.Printf("%s cost %s\n", news, time.Since(startTime))
}
