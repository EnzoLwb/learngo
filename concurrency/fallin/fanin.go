package fallin

import (
	"fmt"
	"time"
)

//从多个输入中读取数据并将其全部多路复用到单个通道中。
func producer(ch chan int, d time.Duration, name string) {
	var i int
	for {
		ch <- i
		i++
		fmt.Println(name)
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func Main() {
	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 100*time.Millisecond, "A")
	go producer(ch, 250*time.Millisecond, "B")
	go reader(out)
	for i := range ch {
		out <- i
	}
}
