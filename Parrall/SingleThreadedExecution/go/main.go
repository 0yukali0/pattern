package main

import(
	"fmt"
	"time"
)

const (
	duration = 1 * time.Second
	ids = ["A", "B", "C", "D"]
)

var like Like

func main() {
	// Scenario 1: No guarding case
	like = new(Like)
	for id := range ids {
		go like.AddLikesWithNoLock(id, like)
	}
	time.Sleep(duration)
	fmt.Println(like.GetCount())
	fmt.Println(like.GetConttribution())

	// Scenario 2: Guarding case
	like = new(Like)
	for id := range ids {
		go like.AddLikes(id, like)
	}
	time.Sleep(duration)
	fmt.Println(like.GetCount())
	fmt.Println(like.GetConttribution())
}