package main

import (
	"fmt"
	"time"
)

func count(sound string) {
	for i :=0;true;i++ {
	fmt.Println(sound)
	time.Sleep(time.Millisecond * 500)
	}
}
func main(){
	go count("Pong")
	//time.Sleep(time.Millisecond * 500)
	count("Ping")

}