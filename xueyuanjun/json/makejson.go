package main

import (
	"encoding/json"
	"fmt"
)

type Language struct {
	Golang int
	Java   int
	Php    int
}
type Person struct {
	Name  string
	Hobby string
	Lang  Language
}

func main2() {
	p := Person{"5lmh.com", "女", Language{1, 0, 0}}
	// 编码json
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))

	// 格式化输出
	b, err = json.MarshalIndent(p, "", "     ")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}
