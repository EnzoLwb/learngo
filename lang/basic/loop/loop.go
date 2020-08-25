package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	convertToBin(8)
	printFile("lang/branch/abc.txt")
	//读取每行文字 原样输出
	printFileContents(strings.NewReader(
		`abc"d"
	kkkk
	123

	p`))
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader) //读取每行文字
	//fmt.Println(scanner.Scan()) 指针阅读
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func convertToBin(n int) {
	result := ""
	for ; n > 0; n /= 2 {
		//对2取模 然后反转
		lsb := n % 2
		//strconv.Itoa将数字转换成对应的字符串类型的数字 而string方法 返回的是一个ascII码
		result = strconv.Itoa(lsb) + result
	}
	fmt.Println(result)
}
