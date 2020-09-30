package main

import (
	"encoding/json"
	"fmt"
)

type Person2 struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func main() {
	student := make(map[string]interface{})
	student["name"] = "5lmh.com"
	student["age"] = 18
	student["sex"] = "man"
	// b := []byte(`{"age":"18","name":"5lmh.com","marry":false}`)
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	var p Person2
	err = json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
