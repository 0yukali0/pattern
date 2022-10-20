package main

import (
	"fmt"
	"time"
	"main/problem"
	"main/solution"
)

func CannotReturnProblem() {
	allNews := []string {
                "John is here",
                "Tim is here",
                "Sam is here",
        }

        start := time.Now()
        for _, news := range allNews {
                problem.PushNews(news, start)
        }
        fmt.Printf("Cost %s\n", time.Since(start))

        time.Sleep(problem.THREE_SECOND)

        start = time.Now()
        for _, news := range allNews {
                go problem.PushNews(news, start)
        }
        time.Sleep(problem.FIVE_SECOND)
        fmt.Printf("Cost %s\n", time.Since(start))
}

func Solution() {
	allNews := []string {
                "John is here",
                "Tim is here",
                "Sam is here",
        }
	newsChs := []<-chan time.Time{}
	start := time.Now()
	for _, news := range allNews {
		newsChs = append(newsChs, solution.PushNews(news, start))
	}

	for index, newsCh := range newsChs {
		fmt.Printf("News %d is sent at %s\n", index, <-newsCh)
	}
}

func main() {
	CannotReturnProblem()
	Solution()
}
