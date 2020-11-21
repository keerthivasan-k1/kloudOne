package main

import (
    "fmt"
    "time"
)

func main() {

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
	
	fmt.Println(time.Now().Weekday())
	
	
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:

	
        fmt.Println("Yaayyyy")
    default:
        fmt.Println("Booooo ")
    }

    t := time.Now().Hour()
	fmt.Println(t)
    switch {
    case t < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
}