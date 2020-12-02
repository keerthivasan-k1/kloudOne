package main

import (
	"fmt"
	"math/rand"
)

func ran_num(num chan<- int) {
	num <- rand.Intn(100)
}

func main() {
	num := make(chan int)
	go ran_num(num)
	n := <-num
	fmt.Println("Random number is", n)
}
