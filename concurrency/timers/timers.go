package timers

import (
	"fmt"
	"time"
)

func timer(d time.Duration) chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()

	return c
}
func reader(out chan int, j int) {
	for x := range out {
		fmt.Println(j, x)
	}
}

func Main() {
	for i := 0; i < 24; i++ {
		c := timer(1 * time.Second)
		go reader(c, i)

	}

}
