package problem

import (
	"fmt"
)

type MessageBoard struct {
	messages []string
}

func NewMessageBoard() *MessageBoard {
	return &MessageBoard{messages: make([]string, 0)}
}

func (m *MessageBoard) SendMessage(message string) {
	m.messages = append(m.messages, message)
}

func (m *MessageBoard) Pop() string {
	message := m.messages[0]
        m.messages = m.messages[1:]
	return message
}

func (m *MessageBoard) Show() {
	for true {
		fmt.Println(m.Pop())
	}
}
