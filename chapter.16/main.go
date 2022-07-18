package main

import "fmt"

type Account struct {
	balance  int
	transfer chan int
}

func NewAccount() *Account {
	transfer := make(chan int)
	r := &Account{0, transfer}
	go func() {
		for {
			amount := <-transfer
			r.balance += amount
		}
	}()
	return r
}

func (a Account) GetBalance() int {
	return a.balance
}

func (a Account) Transfer(amount int) {
	a.transfer <- amount
}

func main() {
	a := NewAccount()

	for {
		var s string
		fmt.Print("enter command: ")
		fmt.Scan(&s)

		switch s {
		case "get":
			fmt.Println(a.GetBalance())
		case "add":
			var i int
			fmt.Print("enter amount: ")
			fmt.Scan(&i)
			a.Transfer(i)
		}
	}
}
