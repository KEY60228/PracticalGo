package main

import (
	"fmt"
	"math/big"
)

type Currency string

const (
	USD Currency = "USD"
	JPY Currency = "JPY"
)

func main() {
	m := NewMutableMoney(USD, big.NewInt(10))
	fmt.Printf("m's currency: %s\n", m.Currency())
	m.SetCurrency(JPY)
	fmt.Println("=== SetCurrency ===")
	fmt.Printf("m's currency: %s\n", m.Currency())
	fmt.Println()

	im := NewImmutableMoney(USD, big.NewInt(10))
	fmt.Printf("im's currency: %s\n", im.Currency())
	im2 := im.SetCurrency(JPY)
	fmt.Println("=== SetCurrency ===")
	fmt.Printf("im's currency: %s\n", im.Currency())
	fmt.Printf("im2's currency: %s\n", im2.Currency())
}

// ミュータブルな構造体
type MutableMoney struct {
	currency Currency
	amount   *big.Int
}

func NewMutableMoney(currency Currency, amount *big.Int) *MutableMoney {
	return &MutableMoney{
		currency: currency,
		amount:   amount,
	}
}

func (m MutableMoney) Currency() Currency {
	return m.currency
}

func (m *MutableMoney) SetCurrency(c Currency) {
	m.currency = c
}

// イミュータブルな構造体
type ImmutableMoney struct {
	currency Currency
	amount   *big.Int
}

func NewImmutableMoney(currency Currency, amount *big.Int) *ImmutableMoney {
	return &ImmutableMoney{
		currency: currency,
		amount:   amount,
	}
}

func (im ImmutableMoney) Currency() Currency {
	return im.currency
}

func (im ImmutableMoney) SetCurrency(c Currency) ImmutableMoney {
	return ImmutableMoney{
		currency: c,
		amount:   im.amount,
	}
}
