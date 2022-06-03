package main

import (
	"fmt"
)

type CarType int

const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)

type CarOption uint64

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSet
	DriverAssist
)

func main() {
	suv := SUV
	awd := AWD

	fmt.Printf("%T, %d, %s\n", suv, suv, suv)
	fmt.Printf("%T, %d, %s\n", awd, awd, awd)
}
