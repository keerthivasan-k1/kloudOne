package main

import "fmt"

func main() {

	var FName string = "kiki"
	LName := "vaz"
	const a float64 = 3.14159265
	b := 0

	fmt.Println(len(FName ))    //length of the string
	fmt.Println(FName + LName ) //concantaion
	fmt.Printf("%g \n", a)      //displays float as it is
	fmt.Printf("%f \n", a)      //round offs
	fmt.Printf("%.2f \n", a)    //shows upto n no of decimal points
	fmt.Printf("%T \n", FName)  //shows type of data
	fmt.Printf("%t \n", b==0 )  //returns bool

}