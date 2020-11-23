package main

import "fmt"

type rect struct{
	length,
	width int
	 
}

func (m rect) area() int{
	return m.length* m.width 	
}
func (k rect) peri()int{
	return 2*k.length+2*k.width
}
func main(){
	p:= rect{5,6}
	fmt.Println(p.area())
	fmt.Println(p.peri())

}