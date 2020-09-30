package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		//else http.Error(w, "Forbidden", http.StatusForbidden)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)

	})
}
func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	postRouter := r.PathPrefix("/posts").Subrouter()
	/*	postRouter.HandleFunc("/", listPosts).Methods("GET").Name("posts.index")
		postRouter.HandleFunc("/create", createPost).Methods("POST").Name("posts.create")
		postRouter.HandleFunc("/update", updatePost).Methods("PUT")
		postRouter.HandleFunc("/delete", deletePost).Methods("DELETE")
		postRouter.HandleFunc("/show/{id:[0-9]+}", showPost).Methods("GET").Name("posts.show")*/
	log.Fatal(http.ListenAndServe(":8082", postRouter))
}
