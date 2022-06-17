package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	UserID   int64
	UserName string
}

func main() {
	u := user{
		UserID:   001,
		UserName: "gopher",
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
