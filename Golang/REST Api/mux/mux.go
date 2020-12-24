package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//type Articles []Article
var Articles = []Article{
	Article{Id: "1", Title: "POwer Of Now", Desc: "Eckhart Tolle", Content: "Spritual"},
	Article{Id: "2", Title: "Harry Potter", Desc: "J.K Rowling", Content: "Fiction"},
}

func allArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Viewing Article Page")
	json.NewEncoder(w).Encode(Articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var articleNew Article

	json.Unmarshal(reqBody, &articleNew)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, articleNew)

	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]

	for _, Article := range Articles {
		if Article.Id == key {
			json.NewEncoder(w).Encode(Article)
		}
	}
}

/*func DeleteSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	for index, article := range Articles {
		if article.Id == id {
			//	json.NewEncoder(w).Encode(Article)
			Articles = append(Articles[:index], Articles[index+1:]...)
			//Articles = append(Articles[:index], Articles[index+1:]...)
			fmt.Println("delete successful")
			//json.NewEncoder(w).Encode(Articles)

		}
	}
}*/
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to Homepage")
}
func testPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST method Successful")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles)
	//	myRouter.HandleFunc("/articles", testPost).Methods("POST")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{Id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {

	handleRequest()
}
