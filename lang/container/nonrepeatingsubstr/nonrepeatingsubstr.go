package main

import "fmt"

/**
寻找最长不含有重复子串的数量
1.lastOccurred[x] 不存在 或者  <start 时  =>> 无需操作
2.lastOccurred[x] >= start =>> 更新start
3.更新lastOccurred[x]时，更新maxLength
*/
func lengthOfNonRepeatingSubstr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) { //中文字符串  []rune()强转rune类型
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start { //有重复的字符串了 重置 字符串的start位置
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
func main() {
	fmt.Println(lengthOfNonRepeatingSubstr("aaabc"))
	fmt.Println("中文")
	traversalString()
	fmt.Println(lengthOfNonRepeatingSubstr("红旗手红旗手手"))
}
