package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://fb.com",
	"https://amazon.com",
}

func fetchstatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%v \n", err)
			}
			fmt.Fprintf(w, "%v \n", resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/", fetchstatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
