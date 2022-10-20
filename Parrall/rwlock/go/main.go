package main

import (
	"time"
	"main/problem"
	"main/solution"
)

const (
	TEN_SECOND=10 * time.Second 
)

func main() {
        like := new(problem.Like)
        users := [4]string{"A", "B", "C", "D"}
	
        for index, user := range users {
		if index % 2 == 1 {
                	go problem.AddLikes(user, like)
			continue
		}
		go problem.ReadLikes(user, like)
        }
        time.Sleep(TEN_SECOND)

	like2 := solution.NewLike()
	for index, user := range users {
		if index % 2 == 1 {
			go solution.AddLikes(user, like2)
		}
		go solution.ReadLikes(user, like2)
	}
	time.Sleep(TEN_SECOND)
}
