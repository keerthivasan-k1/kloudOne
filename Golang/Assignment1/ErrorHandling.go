package main

import (
	"errors"
	"fmt"
)

func add(start , end int) (int,error){
	if start > end {
		return 0, errors.New("Start is greater than End")
	}
	tot:= 0
	for i:= start;i<=end;i++ {
	tot += i
	}
	return tot, nil  
}

func main(){
	total,err:=add(6,5)
	if err != nil {
	fmt.Println("There is an error...REASON:",err)
	}else {
		fmt.Println(total)
	}
	
}