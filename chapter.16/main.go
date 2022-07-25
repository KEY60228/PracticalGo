package main

import (
	"context"
	"log"
	"time"
)

func main() {
	recv := make(chan string)

	ctx1, cancel1 := context.WithCancel(context.Background())
	defer cancel1()

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	select {
	case <-recv:
		log.Println("情報があれば受け取りたいが、いつまでもブロックしたくないチャネル")
	case <-ctx1.Done():
		log.Println("cancel1()が呼ばれて終了")
	case <-ctx2.Done():
		log.Println("タイムアウトで終了")
	}
}
