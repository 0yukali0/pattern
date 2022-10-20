package solution

import (
	"fmt"
	"sync"
)

// sync.Mutex inheritance
type Like struct {
	sync.Mutex
	count uint64
}

func NewLike() *Like {
	return &Like{count: 0}
}

func (l *Like) GetCount() uint64 {
	return l.count
}

func (l *Like) Add(writerID string) {
	l.Lock()
	defer l.Unlock()
	l.count++
	fmt.Printf("%s change count: %d\n", writerID, l.GetCount())
}

func (l *Like) Show(readerID string) {
	l.Lock()
	defer l.Unlock()
	fmt.Printf("%s read count: %d\n", readerID, l.GetCount())
}

func AddLikes(writerID string, like *Like) {
	for i := 0; i < 100; i++ {
		like.Add(writerID)
	}
}

func ReadLikes(readerID string, like *Like) {
	for i := 0; i < 200; i++ {
		like.Show(readerID)
	}
}
