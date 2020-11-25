package main

import (
	"fmt"
	)
func test2(myFunc func(x int) int) {
	fmt.Println(myFunc(7))
}	
	
	
func main(){
	test:= func(x int) int {
		return x -1
	}
fmt.Println(test(5))
test2(test)
}