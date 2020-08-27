package main

import (
	"fmt"
	"learngo/lang/queue"
)

func main() {
	q := queue.Queue{1} //这其实就是new + 构造方法
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
