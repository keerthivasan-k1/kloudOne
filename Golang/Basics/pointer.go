package main

import "fmt"

func main() {
	val := 5
	
	fmt.Println(val)
	
	ptr := &val       //obj name for pointer 
	
	fmt.Println(ptr)  //pointer name
	
	fmt.Println(*ptr) //dereferencing
	
	*ptr =2           //reassigning
	
	fmt.Println(val)
	
}