package main

import (
	"fmt"
	"time"
)

func cook(food string){

fmt.Println("cooking" ,food)
time.Sleep(time.Millisecond * 2000)
fmt.Println("Finished" ,food)

}


func main() {
	foods:= []string {"idly","pongal","dosa","poori",}
	res:= make(chan bool)
	for _,f := range foods{
	//fmt.Println("cooking" ,f)
		go func(f string) {
			cook(f)
			res<-true
		}(f)
	}
	for i:=0;i<len(foods);i++{
		<-res
		
	}	
}