package main

import (
	"fmt"
)
type student struct{
	Name string
	Loc string
	id int
}
func main() {
	fmt.Println("kiki","chennai",54)
	stud := student{"kumar","singh",22}
	fmt.Println(stud)
	stud_ptr := &stud
	fmt.Println(stud_ptr)
	stud_ptr.Name = "kabir"
	fmt.Println(stud)
}