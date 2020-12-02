package main

import (
	"fmt"
	"encoding/json"
)

type books struct {
	Title string `json:"title"`
	Author string `json:"author"`
} 
func main() {
	
	sp := books{Title :"Power of now ",Author :"Eckhart tolle"}
	//fmt.Println(sp)
	
	json1, err := json.Marshal(sp)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(string(json1))
	
}
