package main

import (
	"fmt"
	"regexp"
)

const text = `My email is 815050565@163.com
815050565@163.com 815050565@163.com
815050565@1634.com
`

func main() {
	var regex = `([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`
	/*
			. 就是匹配任何一个字符
			+ 匹配一个或者多个
			.* 0个或者多个
		 	\  一个斜杠会被go语言认为是转义字符 所以一般是 \\ 或者是  ` xxxxx\xxxxx `
			注意斜杠要放在前面
			[a-zA-Z0-9]设置字符  取代.  防止取到 email之前的字符

	*/
	//设置正则表达式
	compile := regexp.MustCompile(regex)              //这是我们自己写的时候 可以看出错的，如果是用户输入的可以使用Mustcompile
	match := compile.FindAllString(text, -1)          //查询一个是 findstring  findallstring就是所有结果 第二个参数表示找多少个 -1表示所有
	match2 := compile.FindAllStringSubmatch(text, -1) //提取
	fmt.Println(match, match2)
	fmt.Println()
	for _, i := range match2 {
		fmt.Println(i) //[815050565@163.com 815050565 163 com]
		for _, j := range i {
			fmt.Println(j)
		}
	}
}
