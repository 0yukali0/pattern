package main

import (
	"fmt"
	"time"
	"main/problem"
	"main/solution"
)

const (
    PARRALISM_LEVEL = 4
)

func SequenceGetDataFromDistributedServer() {
	start := time.Now()
	hosts_IP := map[string]string{
		"host1": "192.168.0.1",
		"host2": "192.168.0.2",
		"host3": "192.168.0.3",
	}
	responses := make([]string, 0)
	for hostName, ip := range hosts_IP {
		fmt.Printf("Access data from %s\n", hostName)
		responses = append(responses, problem.GetServerData(ip))
	}
	problem.ShowNews(responses)
	fmt.Printf("Cost %s", time.Since(start))
}

func fanInFanOut() {
    start := time.Now()
    hosts_IP := []string{
        "192.168.0.1",
        "192.168.0.2",
        "192.168.0.3",
    }
    producerCh := solution.Producer(hosts_IP...)
    tasks := make([]<-chan string, 0)
    for id := 0; id < PARRALISM_LEVEL; id++ {
        tasks = append(tasks, solution.Task(producerCh, id))
    }

    consumerCh := solution.Consumer(tasks...)

    for new := range consumerCh {
        problem.ShowNews(new)
    }
    fmt.Printf("cost %s", time.Since(start))
}

func main() {
	// SequenceGetDataFromDistributedServer()
	fanInFanOut()
}
