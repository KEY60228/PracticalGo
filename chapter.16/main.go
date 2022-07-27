package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done: 1")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Done: 2")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("Done: 3")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done all tasks")
}
