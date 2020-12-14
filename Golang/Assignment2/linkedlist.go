package main

import "fmt"

type node struct {
	data int
	next *node
	// prev *node
}

type linkedlist struct {
	length int
	head   *node
	//	tail   *node
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedlist) display() {
	dis := l.head
	for l.length != 0 {
		fmt.Printf("%d ", dis.data)
		dis = dis.next
		l.length--
	}
	fmt.Println("")
}

func (l *linkedlist) delete(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	beDelete := l.head
	for beDelete.next.data != value {
		if beDelete.next.next == nil {
			return
		}
		beDelete = beDelete.next

	}
	beDelete.next = beDelete.next.next
	l.length--
}

func main() {
	list := linkedlist{}
	node1 := &node{data: 42}
	node2 := &node{data: 23}
	node3 := &node{data: 15}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.display()
	list.delete(23)
	list.display()

}
