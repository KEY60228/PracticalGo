package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	m.Run()
}

func setup() {
	fmt.Println("tests start")
}

func teardown() {
	fmt.Println("tests end")
}

func TestCalc(t *testing.T) {
	type args struct {
		a        int
		b        int
		operator string
	}
	tests := map[string]struct {
		args   args
		want   int
		hasErr bool
	}{
		"足し算": {
			args: args{
				a:        10,
				b:        2,
				operator: "+",
			},
			want:   12,
			hasErr: false,
		},
		"不正な演算子": {
			args: args{
				a:        10,
				b:        2,
				operator: "?",
			},
			hasErr: true,
		},
	}

	fmt.Println("TestCalc start")
	defer func() {
		fmt.Println("TestCalc done")
	}()

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			fmt.Printf("test %v start\n", name)
			defer func() {
				fmt.Printf("test %v end\n", name)
			}()

			got, err := Calc(tt.args.a, tt.args.b, tt.args.operator)
			if (err != nil) != tt.hasErr {
				t.Errorf("Calc() error = %v, hasErr %v", err, tt.hasErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
