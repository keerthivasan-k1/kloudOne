package main

import "fmt"

type stack struct {
	items []int
}

func (s *stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *stack) pop() int {
	l := len(s.items) - 1
	remove := s.items[l]
	s.items = s.items[:l]
	return remove
}

func main() {
	myStack := stack{}
	fmt.Println(myStack)
	myStack.push(12)
	myStack.push(36)
	myStack.push(43)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)

}
