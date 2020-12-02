package main

import (
	"fmt"
	"regexp"
)

	
func main() {
	valid := regexp.MustCompile("[A-Z]{1}[a-z]{4}[0-9]{1}")
	
	p1:= "Kikiv11"
	p2:= "Kloud1"
	p3:= "pswd1111"
	fmt.Println(valid.FindString(p2))
	fmt.Println(valid.FindString(p1))
	fmt.Println(valid.MatchString(p3))
}