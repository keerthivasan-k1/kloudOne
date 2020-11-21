package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7  
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

   // delete(m, "k2")
   //fmt.Println("map:", m)

   //to check whether the value is there in the map in boolean

    i, prs := m["k2"]
    fmt.Println("prs:", prs,i)
    
    //Arranged in alphabetical order based on key
 
    n := map[string]int{"foo": 1, "bar": 2, "app": 3}
    fmt.Println("map:", n)
}