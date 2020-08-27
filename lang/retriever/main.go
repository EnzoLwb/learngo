package main

import (
	"fmt"
	"learngo/lang/retriever/mock"
)

//接口
type RetrieverInf interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r RetrieverInf) {
	r.Get(url)
}
func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

//多继承
type RetrieverPoster interface {
	RetrieverInf
	Poster
}

//使用多继承的方法
func session(s RetrieverPoster) string {
	s.Post(url,
		map[string]string{
			"contents": "another faked imooc.com",
		})
	return s.Get(url)
}

func main() {
	//我们要new 一个爬虫类
	var r RetrieverInf
	mockRetriever := mock.Retriever{
		Contents: "this is a fake imooc.com",
	}
	r = &mockRetriever

	/*	r = &real2.Retriever{ // 因为 接口中 Get方法的接收者是 指针类型
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}*/
	fmt.Println(r)
	fmt.Println(mockRetriever)
	fmt.Println(r.(*mock.Retriever))
	// type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not mock.Retriever")
	}
}
