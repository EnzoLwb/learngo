package main

import "fmt"

//需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。
func main() {
	//1.声明切片
	var s11 []int
	if s11 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	slice22 := []int{}
	// 3.make()
	var slice3 []int = make([]int, 0)
	fmt.Println("声明数组:", slice22, slice3)
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	//data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8
	fmt.Println("声明数组:", s)
	/*
		数组or切片转字符串：
		    strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
	*/
	return
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	fmt.Println(arr[2:6])
	fmt.Println(arr[2:]) //包含 2这个位置
	fmt.Println(arr[:2]) //不包含 2这个位置
	s2 := arr[:]
	fmt.Println(s2)
	fmt.Println("updateSlice 后")
	updateSlice(s2)          //切片就是会引用传递
	fmt.Printf("s2地址%p", s2) //s2 是切片类型了 而不是arr类型
	fmt.Println()
	fmt.Printf("arr地址%p", arr) //arr 是数组 所以不能这样打印出地址
	fmt.Println()
	slice1 := new([]int)
	fmt.Println(slice1) //输出的是一个地址  &[]

	//使用make创建切片
	slice2 := make([]int, 5)
	fmt.Println(slice2) //输出初始值都为0的数组， [0 0 0 0 0]
	//fmt.Println(slice1[0]) 结果出错 slice1是一个空指针 invalid operation: slice1[0] (type *[]int does not support indexing)
	fmt.Println(slice2[0]) //结果为 0 因为已经初始化了

	s1 := make([]int, 0, 3)
	fmt.Println(s1)
	fmt.Printf("地址%p,长度%d,容量%d\n", s1, len(s1), cap(s1))
	var a = append(s1, 1, 2)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Printf("地址%p,长度%d,容量%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 3, 4, 5)
	s1 = append(s1, 3, 4, 5)
	fmt.Printf("地址%p,长度%d,容量%d\n", s1, len(s1), cap(s1)) //扩容后地址也会随之改变
}

func updateSlice(s []int) {
	s[0] = 100
}
