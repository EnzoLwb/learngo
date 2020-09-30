package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func listPosts(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "文章列表")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "发布文章")
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "修改文章")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "删除文章")
}

func showPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "文章详情")
}
func main() {
	// 路由分组（基于子路由+路径前缀）
	r := mux.NewRouter()
	postRouter := r.PathPrefix("/posts").Subrouter()
	postRouter.HandleFunc("/", listPosts).Methods("GET").Name("posts.index")
	postRouter.HandleFunc("/create", createPost).Methods("POST").Name("posts.create")
	postRouter.HandleFunc("/update", updatePost).Methods("PUT")
	postRouter.HandleFunc("/delete", deletePost).Methods("DELETE")
	postRouter.HandleFunc("/show/{id:[0-9]+}", showPost).Methods("GET").Name("posts.show")
	indexUrl, _ := r.Get("posts.index").URL()
	log.Println("文章列表链接：", indexUrl)
	showUrl, _ := r.Get("posts.show").URL("id", "1")
	log.Println("文章详情链接：", showUrl)
	log.Fatal(http.ListenAndServe(":8082", postRouter))
}
