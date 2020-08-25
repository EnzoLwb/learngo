package main

import "fmt"

func main() {
	fmt.Println(print(13, 39, "//"))
	fmt.Println(sum(13, 39, 1))
	var a, b = 3, 4
	fmt.Println("交换前：", a, b)
	//交换变量
	swap(&a, &b)
	fmt.Println("交换后：", a, b)
	fmt.Println(sum(13, 39))
}

//计算器
func print(q, r int, op string) (int, error) {
	if res, err := eval(q, r, op); err != nil {
		return 0, err
	} else {
		return res, nil
	}
}

func eval(a, b int, op string) (int, error) { //返回类型
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("error operation:%s", op) //可以返回error 类型
	}
}

//除法 求模
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//求多个数字的和
func sum(numbers ...int) int { //... 表示多个参数 以数组形式展示
	s := 0
	fmt.Printf("numbers:%d", numbers)
	for _, v := range numbers {
		s += v
	}
	return s
}

// 交换两个变量的值 指针方式改变
func swap(a, b *int) {
	*a, *b = *b, *a

}
