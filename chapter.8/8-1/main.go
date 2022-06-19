package main

import (
	"encoding/json"
	"fmt"
)

type FormInput struct {
	Name        string `json:"name"`
	CompanyName string `json:"company_name,omitempty"`
}

func main() {
	in := FormInput{
		Name: "山田太郎",
	}
	b, _ := json.Marshal(in)
	fmt.Println(string(b))
}
