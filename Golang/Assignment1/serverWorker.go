package main

import (
	"fmt"
	"time"
)
type T = struct{}

func worker(id int, ready <-chan T, done chan<- T){
	<-ready
	fmt.Println("worker number", id , "job started")
	time.Sleep(time.Second * time.Duration(id +1))
	fmt.Println("worker number", id , "job finished")
	done<- T{}
	

}

func main(){
	ready := make(chan T)
	done := make(chan T)
	
	go worker(1,ready,done)
	go worker(2,ready,done)
	go worker(3,ready,done)
	
	time.Sleep(time.Second)
	ready<- T{};ready<- T{};ready<- T{}
	<-done;<-done;<-done
}