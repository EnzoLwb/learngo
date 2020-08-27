package main

import (
	"fmt"
	"math"
)

// 1, 1, 2, 3, 5, 8, 13, ...
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送。
func main() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))

	// --- function variable ---
	fn := func() { println("Hello, World!") }
	fn()

	// --- function collection ---
	fns := [](func(x, y int) int){
		func(x, y int) int { return x + 1 },
		func(x, y int) int { return y + 2 },
		func(x, y int) int { return x + 3 },
	}
	println(fns[0](100, 200))

	// --- function as field ---
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	println(d.fn())
}
