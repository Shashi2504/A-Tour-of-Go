package main

import (
	"fmt"
	"math"
)

type I interface {
	Make()
}

type Take struct {
	Stake string
}

func (t *Take) Make() {
	fmt.Println(t.Stake)
}

type Fake float64

func (f Fake) Make() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &Take{"Hello"}
	describe(i)
	i.Make()

	i = Fake(math.Pi)
	describe(i)
	i.Make()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

