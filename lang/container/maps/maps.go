package main

import "fmt"

func main() {
	m := map[string]string{ //key value  都是string类型
		"name": "enzo",
		"age":  "14",
		"site": "imooc",
	}
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      // m3 == nil
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
