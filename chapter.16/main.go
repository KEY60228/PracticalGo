package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, _ := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("Done: 1")
		return nil
	})

	eg.Go(func() error {
		time.Sleep(time.Second * 2)
		fmt.Println("Done: 2")
		return nil
	})

	eg.Go(func() error {
		time.Sleep(time.Second * 3)
		fmt.Println("Done: 3")
		return nil
	})

	err := eg.Wait()
	fmt.Println("Done all tasks", err)
}
