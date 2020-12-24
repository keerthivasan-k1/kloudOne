package main

import "fmt"

func main() {
	arr, even, odd := []int{}, []int{}, []int{}
	limit := 10

	for i := 0; i <= limit; i++ {
		arr = append(arr, i)

	}
	for i := range arr {
		if arr[i]%2 == 0 {

			even = append(even, arr[i])
		} else {

			odd = append(odd, arr[i])
		}

	}

	fmt.Println("Even numbers are ", even)
	fmt.Println(" odd numbers are ", odd)
}
