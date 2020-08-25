package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func main() {
	variableZeroValue()
	variableInitValue()
	variableShorter()
	euler()
	consts()
	enums()
	fmt.Println(calcTriangle(3, 4))
	fmt.Println(aa, ss, bb)
}

//没有初始值的声明变量方式
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //%q 打印字符串
}

//强类型语言 声明变量必须带类型
func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//直接赋值的声明变量方式
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//欧拉公式 e的πi 次幂 +1 = 0 ,i的出现使用1i 不然会以为是变量
func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

//勾股定理
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //因为sqrt方法需要传入一个float类型的值 所以需要强转
	return c
}

//常量
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	const number = "xxx"
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, number, c)
}

//枚举
func enums() {
	const (
		cpp = iota //一个递增的种子 下面就都是 0 1 2 3
		_
		_
		_
		python
		golang
		js
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(cpp, js, python, golang)
	fmt.Println(b, kb, mb, gb, tb)
}
