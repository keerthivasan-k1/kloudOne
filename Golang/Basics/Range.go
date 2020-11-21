package main

import (
	"fmt"
)

func main() {

	str :=[]string{"s","S","3","d","l"}
	for i ,r := range str{
	fmt.Println("Hello, playground",r,i)
	}
	
	m := map[string]int{"a":1,"d":3}
	for k , v := range m{
	fmt.Println("key is ",k ,"value is ", v)
	}
}