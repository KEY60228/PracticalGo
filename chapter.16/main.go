package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"go.uber.org/ratelimit"
)

type Task string

type Result struct {
	Value int64
	Task  Task
	Err   error
}

func worker(id int, tasks <-chan Task, results chan<- Result, rt ratelimit.Limiter) {
	for t := range tasks {
		rt.Take()

		fmt.Printf("worker: %d task: %s\n", id, t)
		s, err := os.Stat(string(t))
		if err == nil && s.IsDir() {
			err = fmt.Errorf("worker: %d err: %s is dir", id, string(t))
		}
		result := Result{
			Task: t,
		}
		if err != nil {
			result.Err = err
		} else {
			fmt.Printf("worker: %d path: %s size: %d\n", id, string(t), s.Size())
			result.Value = s.Size()
		}
		results <- result
	}
}

func fixedTasks(taskSrcs []Task) int64 {
	tasks := make(chan Task, len(taskSrcs))
	results := make(chan Result)
	for _, src := range taskSrcs {
		tasks <- src
	}
	close(tasks)

	rl := ratelimit.New(100)
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, tasks, results, rl)
	}

	var count int
	var size int64
	for {
		result := <-results
		count += 1
		if result.Err != nil {
			fmt.Printf("err %v for %s\n", result.Err, result.Task)
		} else {
			size += result.Value
		}
		if count == len(taskSrcs) {
			break
		}
	}
	return size
}

func main() {
	tasks := make([]Task, 0, 10000)
	filepath.Walk(runtime.GOROOT(), func(path string, info os.FileInfo, err error) error {
		tasks = append(tasks, Task(path))
		return nil
	})
	fmt.Printf("total file size is %dMB\n", fixedTasks(tasks)/1000/1000)
}
