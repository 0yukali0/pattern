package main

import (
	"sync"
)

const (
	likeCountPerPersion = 10000
)

type Like struct {
	sync.Mutex
	count           uint16
	IDcontributions map[string]uint16
}

func (l *Like) Add(writerID string) {
	l.Lock()
	defer l.Unlock()
	l.count++
	l.IDcontributions[writerID]++
}

func (l *Like) AddLikes(writerID string) {
	for i := 0; i < likeCountPerPersion; i++ {
		l.Add(writerID)
	}
}

func (l *Like) GetCount() {
	return l.count
}

func (l *Like) GetConttribution() {
	return l.IDcontributions
}
