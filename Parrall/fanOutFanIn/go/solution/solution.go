package solution

import (
    "fmt"
	"sync"
	"main/problem"
)

func Producer(hostNames ...string) <-chan string {
	producerCh := make(chan string, len(hostNames))
	go func() {
		defer close(producerCh)
		for _, hostName := range hostNames {
			producerCh <- hostName	
		}
	}()
	return producerCh
}

func Task(producerCh <-chan string, id int) <-chan string {
	taskCh := make(chan string)
	go func() {
		defer close(taskCh)
        for hostName := range producerCh {
            fmt.Printf("go %d got a job %s\n", id, hostName)
			taskCh <- problem.GetServerData(hostName)
		}
	}()
	return taskCh
}

func Consumer(taskChs ...<-chan string) <-chan string {
	consumerCh := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(taskChs))
	go func(){
		defer close(consumerCh)
		wg.Wait()
	}()

	for _, task := range taskChs {
		go func(task <-chan string) {
			defer wg.Done()
			for new := range task {
				consumerCh <- new
			}
		}(task)
	}
	return consumerCh
}
