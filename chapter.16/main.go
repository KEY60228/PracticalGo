package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx := context.Background()
	err := runJobs(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}

func runJobs(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ec := make(chan error)
	done := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			err := exec.CommandContext(ctx, "sleep", "30").Run()
			if err != nil {
				ec <- err
			} else {
				done <- struct{}{}
			}
		}()
	}

	go func() {
		time.Sleep(10 * time.Second)
		ec <- errors.New("accidental error")
	}()

	for i := 0; i < 10; i++ {
		select {
		case err := <-ec:
			return err
		case <-done:
		}
	}

	return nil
}
