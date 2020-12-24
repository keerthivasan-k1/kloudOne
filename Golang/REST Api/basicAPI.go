package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func homepage(w http.RespondsWriter, r *http.Request) {
	writer := html.EscapeString(r.URL.Path)
	fmt.Printf(w, "Hello, %q", writer)
}
func homepage2(w http.RespondsWriter, r *http.Request) {
	fmt.Printf(w, "HI")
}

func handleReq() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/", homepage2)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func main() {
	handleReq()

}
