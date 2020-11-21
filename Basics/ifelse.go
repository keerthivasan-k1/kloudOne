package main

import "fmt"

func main() {
	
	
    if n :=9;n%2 == 0 {
        fmt.Println(n," is even")
    } else {
        fmt.Println(n," is odd")
    }
	
		
    if a,b := 8,4; a%b == 0 {
        fmt.Println(a," is divisible by ",b)
    }

    if num := -5; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}