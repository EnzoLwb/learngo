package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//r 代表请求对象 w代表响应对象
func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                              //解析参数
	fmt.Println(strings.Join(r.Form["a"], "")) //http://xxxxx?a=2&b=sb => 2
	fmt.Println("url:", r.URL.Path)            //请求url
	fmt.Println("scheme:", r.URL.Scheme)       //请求url
	for k, v := range r.Form {
		fmt.Println(k + ":" + strings.Join(v, ""))
	}
	fmt.Fprintf(w, "你好！Enzo~") //发送响应到客户端 也就是return
}
func main() {
	//注册路由
	http.HandleFunc("/", sayHelloWorld)
	//设立服务器 监听端口
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
