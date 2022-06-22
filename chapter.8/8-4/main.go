package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ianlopshire/go-fixedwidth"
)

type Book struct {
	ISBN        string `fixed:"1,17"`
	PublishDate string `fixed:"18,25"`
	Price       int    `fixed:"26,29"`
	PDF         string `fixed:"30,34,left"`
	EPUB        string `fixed:"35,39,left"`
	EbookPrice  int    `fixed:"40,44"`
}

func main() {
	s := `978-4-87311-865-9201909174620true true 3696
978-4-87311-924-3202010102750falsefalse0000
978-4-87311-878-9201903120000true true 0000`

	for _, line := range strings.Split(s, "\n") {
		var b Book
		if err := fixedwidth.Unmarshal([]byte(line), &b); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", b)
	}
}
