package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Post struct {
	title string `json:"json"`
	body string `json:"body"`
}

type Posts []Post


func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func posts(w http.ResponseWriter, r *http.Request) {
	posts_list := Posts{
		Post{title:"post 1", body:"body of post 1"},
		Post{title:"post 2", body:"body of post 2"},
		Post{title:"post 3", body:"body of post 3"},
	}

	fmt.Println("request: get posts")
	json.NewEncoder(w).Encode(posts_list)
}

func requestHandler() {

	new_router := mux.NewRouter().StrictSlash(true)

	new_router.HandleFunc("/", hello)
	new_router.HandleFunc("/posts", posts)
	log.Fatal(http.ListenAndServe(":8080", new_router))
}


func main() {
	requestHandler()
}