package main

import (
    "fmt"
    "time"
)

func pinger (p1,p2 chan int){
	//p1=make(chan int)
	//p2=make(chan int)
	for{
	   <-p1
	fmt.Println("ping")
	time.Sleep(time.Second)
	p2<-1
	}
}
func ponger (p1,p2 chan int){
	//p1=make(chan int)
	//p2=make(chan int)
	for{
	   <-p2
	fmt.Println("pong")
	time.Sleep(time.Second)
	p1<-1
	}
}

func main() {
	ping:=make(chan int)
	pong:=make(chan int)
	
	go pinger(ping,pong)
	go ponger(ping,pong)
	
	ping<-1
	
	for{
	time.Sleep(time.Second)
	}
}