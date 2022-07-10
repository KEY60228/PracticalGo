package main

import "fmt"

type User struct {
	UserID    string
	UserName  string
	Languages []string
}

func main() {
	fmt.Println(getTom())
	fmt.Println(getTom2())
}

func getTom() User {
	return User{
		UserID:    "001",
		UserName:  "Tom",
		Languages: []string{"Java", "Go", "Rust"},
	}
}

func getTom2() User {
	return User{
		UserID:    "001",
		UserName:  "Tom",
		Languages: []string{"Java", "Go"},
	}
}
