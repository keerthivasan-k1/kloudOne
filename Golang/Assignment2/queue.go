package main

import "fmt"

type queue struct {
	items []int
}

func (q *queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *queue) Dequeue() int {
	remove := q.items[0]
	q.items = q.items[1:]
	return remove

}

func main() {
	myQueue := queue{}
	myQueue.Enqueue(4)
	myQueue.Enqueue(9)
	myQueue.Enqueue(5)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
