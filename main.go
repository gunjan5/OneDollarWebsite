package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Everything is awesome!!!"))
// }
