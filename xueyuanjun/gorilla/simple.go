package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)                      //设置状态码为200
	fmt.Fprintf(w, "Hello,World! %s", params["name"]) // return
}
func tokenQuery(w http.ResponseWriter, r *http.Request) {
	query := "token"
	fmt.Fprintf(w, "包含指定查询字符串[%s=%s]", query, r.FormValue(query))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Hello", sayHelloWorld)
	r.HandleFunc("/token", tokenQuery).Queries("token") //限制请求参数中必须带有token字段
	r.HandleFunc("/header", sayHelloWorld).Headers("X-Requested-With", "XMLHttpRequest")
	r.HandleFunc("/Hello/{name}", sayHelloWorld).Methods("GET", "POST")
	r.HandleFunc("/Hello/{name:[a-z]+}", sayHelloWorld).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", r))
}
