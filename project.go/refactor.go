package main

import (
	"fmt"
	"time"
)

type Job struct {
	id   int
	task func()
}

func main() {
	eventLoop := make(chan Job, 10)
	done := make(chan bool)

	// event loop goroutine
	go func() {
		for {
			select {
			case job := <-eventLoop:
				fmt.Println("Executing job", job.id)
				job.task()
			case <-done:
				return
			}
		}
	}()

	// schedule jobs
	for i := 1; i <= 5; i++ {
		id := i
		time.AfterFunc(time.Duration(i)*300*time.Millisecond, func() {
			eventLoop <- Job{
				id: id,
				task: func() {
					fmt.Println("Job", id, "completed")
				},
			}
		})
	}

	time.Sleep(2 * time.Second)
	done <- true
}
