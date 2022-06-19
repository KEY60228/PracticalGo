package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	X        func() `json:"-"`
}

func main() {
	u := user{
		UserID:   001,
		UserName: "gopher",
		X:        func() { fmt.Println("Hello") },
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
