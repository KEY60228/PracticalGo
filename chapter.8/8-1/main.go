package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type user struct {
	UserID   int64
	UserName string
}

func main() {
	var b bytes.Buffer
	u := user{
		UserID:   001,
		UserName: "gopher",
	}
	_ = json.NewEncoder(&b).Encode(u)
	fmt.Printf("%v\n", b.String())
}
