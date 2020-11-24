package main

import (  
    "fmt"
    "math"
)
     
type geometry interface{
	area() float64
	peri() float64
} 
type rect struct{
	length,
	breadth float64
}
type circle struct{
	radius float64
}

func (r rect) area() float64{
	return r.length * r.breadth
} 

func (c circle) area() float64{
	return math.Pi *c.radius *c.radius 
}

func (r rect) peri() float64{
	return 2* r.length + 2*r.breadth
}

func (c circle) peri() float64{
	return 2*math.Pi*c.radius
}

func measure(g geometry){
	//fmt.Println(g)
	fmt.Println("area :",g.area())
	fmt.Println("perimetre :",g.peri())
	
}

func main(){
	a:=rect{3,4}	
	b:=circle{5}
	
	measure(a)
	measure(b)
}