package pingpong

import (
	"fmt"
	"time"
)

func Main() {
	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	table <- Ball

	time.Sleep(10 * time.Second)

}

func player(table chan int) {
	for {
		ball := <-table
		ball++
		time.Sleep(1 * time.Second)
		table <- ball
		fmt.Println(ball)
	}
}
