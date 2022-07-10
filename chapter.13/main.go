package main

import "fmt"

type User struct {
	id        string
	UserName  string
	Languages []string
}

func main() {
	fmt.Println(getTom())
	fmt.Println(getTom2())
}

func getTom() User {
	return User{
		id:        "001",
		UserName:  "Tom",
		Languages: []string{"Java", "Go"},
	}
}

func getTom2() User {
	return User{
		id:        "002",
		UserName:  "Tom",
		Languages: []string{"Java", "Go"},
	}
}
