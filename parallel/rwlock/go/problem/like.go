package problem

import (
	"fmt"
)

type Like struct {
	count uint64
}

func (l *Like) Add(writerID string) {
	l.count++
	fmt.Printf("%s change count: %d\n", writerID, l.count)
}

func (l *Like) Show(readerID string) {
	fmt.Printf("%s read count: %d\n", readerID, l.count)
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
