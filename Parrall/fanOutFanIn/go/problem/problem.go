package problem

import (
	"time"
	"fmt"
)

const (
	UPPERBOUND_UNIT = 3
)

func GetServerData(serverName string) string {
	time.Sleep(time.Duration(UPPERBOUND_UNIT) * time.Second)
	return fmt.Sprintf("%s server data", serverName)
}

func ShowNews(news ...interface{}) {
	fmt.Println(news...)
}
