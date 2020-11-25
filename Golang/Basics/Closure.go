package main

import (
	"fmt"
	)
	
func  returnFunc(x int) func() int{
	x = 0
	return func() int {
	x++
	return x
	}
}	
	
func main(){
	v :=returnFunc(4)
	fmt.Println(v())
	fmt.Println(v())
	
}