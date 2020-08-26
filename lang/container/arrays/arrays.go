package main

import "fmt"

//索引式数组的声明
func main() {
	var arr1 [5]int
	arr2 := [5]int{1, 23, 4, 5, 1}
	arr3 := [...]int{1, 23, 4, 5, 1}
	var grid [4][5]int
	var d = [...]struct { //struct 创建出来的类型
		name string
		age  int
	}{
		{"xxxx", 14},
		{"xxxx1", 24}, //换行的话 就需要逗号
	}
	fmt.Println(d)
	fmt.Println("arr1")
	printArray(arr1)
	fmt.Println("arr2")
	printArray(arr2)
	fmt.Println("arr3")
	printArray(arr3)
	fmt.Println("grid")
	fmt.Println(grid)
}

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
