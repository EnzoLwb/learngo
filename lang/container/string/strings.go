package main

import "fmt"

type iAdder func(int) (int, iAdder) //声明一个类型 这里是个函数 函数返回 int 和一个函数
func main() {
	/*
		1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
		2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。*/
	var a *int
	a = new(int) //使用new函数得到的是一个类型的指针 并且该指针对应的值为该类型的零值
	*a = 10
	fmt.Println(a)

	var b = make(map[string]int)
	b["测试"] = 100
	fmt.Println(b)
}
