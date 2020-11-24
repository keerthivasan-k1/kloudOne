package main

import (  
    "fmt"
    "time"
)

func num() {  
    for i:=0;i<=5;i++{
    time.Sleep(150 * time.Millisecond)
    fmt.Println(i)
    }
}
func alpha() {
   for i := 'a';i<='e';i++{
   time.Sleep(300 * time.Millisecond)
   fmt.Println(i)
   }
}
func main() {  
   go num()
   go alpha()
   time.Sleep(2000 * time.Millisecond)
   fmt.Println("The end")
   
}