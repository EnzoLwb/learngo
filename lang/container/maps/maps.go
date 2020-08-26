package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"unicode/utf8"
)

//按照顺序遍历map
func order() {
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	/*str := "我和你我和你我和你我和你"
	randStr := []rune(str)[rand.Intn(9)]    //转成rune 不然就是byte的unicode
	fmt.Printf("%v(%c) ", randStr, randStr) //[97 98 99 100 97 100 119 49 50 51 52]([a b c d a d w 1 2 3 4])
	*/
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //带有前置0的字符串数字
		value := rand.Intn(99)
		scoreMap[key] = value
	}
	//取出key 存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	//排序
	sort.Strings(keys)
	fmt.Println("排序后")
	fmt.Println(keys)
	fmt.Println("原Map")
	fmt.Println(scoreMap)
	fmt.Println("按照顺序遍历Map")
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func main() {
	order()
	return
	s := "Yes我爱慕课网!" // UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) //59 65 73 E6 88 91 E7 88 B1 E6 85 95 E8 AF BE E7 BD 91 21
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch) //对应的 Unicode 代码点表示的字符
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
	return
	m := map[string]string{ //key value  都是string类型
		"name": "enzo",
		"age":  "14",
		"site": "imooc",
	}
	m2 := make(map[string]int, 6) //m2 == empty map
	var m3 map[string]int         // m3 == nil
	fmt.Println(m, m2, m3)
	fmt.Println("traverse")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting Value :NAME is", m["name"])
	if causeName, ok := m["cc"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key not exist")
	}
	fmt.Println("Deleteing Values")
	//delete(m, "name")
	delete(m, "case")
	fmt.Println(m)
	name, ok := m["name"] //ok 返回 true 证明键存在
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
}
