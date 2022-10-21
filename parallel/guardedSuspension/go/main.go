package main

import (
	"main/problem"
	"main/solution"
	"math/rand"
	"time"
)

const (
	UPPERBOUND_UNIT = 10
	TEN_SECOND      = 10 * time.Second
	MESSAGE_CONTEXT = "Hi"
)

func ProblemCauseRuntimeError() {
	messageBoard := problem.NewMessageBoard()
	for index := 0; index < 2; index++ {
		go func() {
			for true {
				time.Sleep(time.Duration(rand.Intn(UPPERBOUND_UNIT)) * time.Millisecond)
				messageBoard.SendMessage(MESSAGE_CONTEXT)
			}
		}()
	}
	go func() {
		messageBoard.Show()
	}()
	time.Sleep(TEN_SECOND)
}

func SolutionWillWaitingTheMessage() {
	messageBoard := solution.NewMessageBoard()
	for index := 0; index < 2; index++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(UPPERBOUND_UNIT)) * time.Millisecond)
			messageBoard.SendMessage(MESSAGE_CONTEXT)
		}()
	}
	go func() {
		messageBoard.Show()
	}()
	time.Sleep(TEN_SECOND)
}

func main() {
	// ProblemCauseRuntimeError()
	SolutionWillWaitingTheMessage()
}
