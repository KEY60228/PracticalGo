package main

import (
	"fmt"
)

type Portion int

const (
	Regular Portion = iota // 普通
	Small                  // 小盛り
	Large                  // 大盛り
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

type OptFunc func(r *Udon)

func NewUdon(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) {
		r.men = p
	}
}

func OptAburaage() OptFunc {
	return func(r *Udon) {
		r.aburaage = true
	}
}

func OptEbiten(n uint) OptFunc {
	return func(r *Udon) {
		r.ebiten = n
	}
}

func main() {
	tempuraUdon := NewUdon(OptMen(Large), OptEbiten(3))
	fmt.Println(tempuraUdon)
}
