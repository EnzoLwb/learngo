package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	grade(11)
	const filename = "lang/branch/abc.txt"
	if content, error := ioutil.ReadFile(filename); error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", content)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score)) //sprintf(终端中不会有显示)
	case score < 60:
		g = "不及格"
	}
	return g
}
