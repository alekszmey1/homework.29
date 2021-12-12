package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Task struct {
	TaskID int
	Number int
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("START")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	queue := make(chan Task, 10)
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		i := 0
		for {
			<-ticker.C
			fmt.Println("new task: ", i)
			wg.Add(1)
			queue <- Task{
				TaskID: i,
				Number: i,
			}
			i++
		}
	}()

	go func() {
		for {
			task := <-queue
			fmt.Println("task start: ", task.TaskID, "v: ", task.Number*task.Number)
			time.Sleep(5 * time.Second)
			fmt.Println("task done: ", task.TaskID)
			wg.Done()
		}
	}()

	go func() {
		<-c
		wg.Done()
		ticker.Stop()
	}()
	wg.Wait()
	fmt.Println("OK")
	os.Exit(1)
}
