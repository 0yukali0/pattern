package solution

import (
	"fmt"
)

type MessageBoard struct {
	messages chan string
}

func NewMessageBoard() *MessageBoard {
	return &MessageBoard{messages: make(chan string, 100)}
}

func (m *MessageBoard) SendMessage(message string) {
	m.messages <- message
}

func (m *MessageBoard) Show() {
	for true {
		fmt.Println(<-m.messages)
	}
}
