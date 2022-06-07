package main

import "fmt"

func main() {
	s := NewNoCopyStruct("string")
	fmt.Println(s.String())

	// コピーするとpanic
	// s2 := *s
	// fmt.Println(s2.String())

	s3 := s.Copy()
	fmt.Println(s3.String())
}

type NoCopyStruct struct {
	self  *NoCopyStruct
	Value *string
}

func NewNoCopyStruct(value string) *NoCopyStruct {
	r := &NoCopyStruct{
		Value: &value,
	}
	r.self = r
	return r
}

func (n *NoCopyStruct) String() string {
	if n != n.self {
		panic("should not copy NoCopyStruct instance without Copy() method")
	}
	return *n.Value
}

func (n *NoCopyStruct) Copy() *NoCopyStruct {
	str := *n.Value
	p2 := &NoCopyStruct{
		Value: &str,
	}
	p2.self = p2
	return p2
}
