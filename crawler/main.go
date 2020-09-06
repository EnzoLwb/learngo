package main

import (
	"bufio"
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
	"os"
	"regexp"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}

func printCityList(contents []byte) {
	//正则表达式 （ 贴过来手动改写 比较方便）.*> 这里是不行的 因为会匹配到最后的>  所以要修改成[^>]*
	//这里的中文处理就是 [^<] 表示除了< 一直向后匹配 如果匹配到< 就停
	//<a href="http://www.zhenai.com/zhenghun/gannan1" data-v-2cb5b6a2>甘南</a>
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" [^>]+>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1) //返回 [][]string
	for _, i := range matches {
		//for _, subMatch := range i {
		//fmt.Printf("%s ", subMatch) //0 号元素是匹配的整个字符串
		fmt.Printf("City:%s Url:%s", i[2], i[1])
		//}
		fmt.Println()
		//fmt.Println(i) //这打印的是byte
		//fmt.Printf("%s\n", i)
	}
	fmt.Printf("Matches Count: %d", len(matches))

}

//写入内容到文件
func writeStringToFile(contents []byte, filename string) {
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		//f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		f, err1 = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0600) //打开文件 写入会覆盖
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err3 := w.WriteString(string(contents))
	check(err3)
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
