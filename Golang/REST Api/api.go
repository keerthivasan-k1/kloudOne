/*package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "POwer Of Now", Desc: "Eckhart Tolle", Content: "Spritual"},
	}
	fmt.Println("Viewing Article Page")
	json.NewEncoder(w).Encode(articles)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	writer := html.EscapeString(r.URL.Path)
	fmt.Fprintf(w, "Hello, %q", writer)
	fmt.Println("Viewing HomePage",writer)
}
func page2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "To pursue a highly challenging and a creative career, where I can apply my knowledge and creativity, acquire new skills and contribute effectively to the organization and to achieve expertise.\n\n Have completed BSC (Visual Communication) in Patrician college of arts and science")
	fmt.Println("Viewing viki Page")
}

func handleReq() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/article", allArticles)
	http.HandleFunc("/viki", page2)
	http.ListenAndServe(":10000", nil)

}

func main() {
	handleReq()

}
