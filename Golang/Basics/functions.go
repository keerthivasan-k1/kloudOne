package main

import "fmt"

func add(a,b int) int{
    return a + b 
}

func sub(a,b int) int{
    return a - b 
}
func mul(a,b int) int{
    return a * b 
}
func div(a,b int) int{
    return a / b 
}

func main() {
    x,y := 3,4
    sum := add(x,y)
    sub := sub(x,y)
    mul := mul(x,y)
    div := div(x,y)
    
    fmt.Println("sum of ",x ,"and",y, "is",sum)
    fmt.Println("sub of ",x ,"and",y, "is",sub)
    fmt.Println("mul of ",x ,"and",y, "is",mul)
    fmt.Println("div of ",x ,"and",y, "is",div)
}