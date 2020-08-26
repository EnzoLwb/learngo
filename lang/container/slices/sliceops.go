package main

import "fmt"

func main2() {
	sliceOps()
}
func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}
func sliceOps() {
	fmt.Println("Creating slice")
	/*var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)*/
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	s2 := make([]int, 16) //默认 cap和len相同都是 这里就是16喽
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) //cope s1给 s2
	fmt.Println(s2)
	fmt.Println("Delete elements from slice") //删除就是重新拼接
	s2 = append(s2[:3], s2[4:]...)            //这不就是删除 位置为3的元素
	fmt.Println(s2)
	fmt.Println("Popping from front")
	front := s2[0] //从前面删除
	s2 = s2[1:]
	fmt.Println(front)
	fmt.Println("Popping from back")
	//从后面删除
	s2 = s2[:len(s2)-1]
	fmt.Println(s2)
}
