package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("ゴルーチンを実行します")
	go func() {
		fmt.Println("ゴルーチンが実行しています")
	}()

	fmt.Println("ゴルーチンの終了を待ちます")
	time.Sleep(time.Second)
	fmt.Println("ゴルーチンが終了しました")
}
