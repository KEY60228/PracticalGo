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

type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(opt Option) *Udon {
	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}

func main() {
	opt := Option{
		men:      Large,
		aburaage: false,
		ebiten:   3,
	}
	tempuraUdon := NewUdon(opt)
	fmt.Println(tempuraUdon)
}
