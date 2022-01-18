package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa Kabar teman!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " Hello")
	})
	http.HandleFunc("/index", index)
	fmt.Println("Memulai Web Server Pada Localhost:8080")
	http.ListenAndServe(":8080", nil)
}
