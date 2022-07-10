package main

import (
	"fmt"
	"io"
	"os"
)

type User struct {
	Name    string
	Address string
}

func main() {
	DumpUser(&User{
		Name:    "KEY",
		Address: "Japan",
	})
}

func DumpUser(u *User) {
	DumpUserTo(os.Stdout, u)
}

func DumpUserTo(w io.Writer, u *User) {
	if u.Address == "" {
		fmt.Fprintf(w, "%s(住所不定)", u.Name)
	} else {
		fmt.Fprintf(w, "%s@%s", u.Name, u.Address)
	}
}
