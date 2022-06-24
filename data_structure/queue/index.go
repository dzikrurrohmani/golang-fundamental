package main

import "fmt"

type queue []string

func (q *queue) enqueue(name string) {
	*q = append(*q, name)
}

func (q *queue) dequeue() string {
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func main() {
	// q := make(queue, 0)
	var q queue
	//  q := new(queue)
	q.enqueue("51")
	q.enqueue("62")
	q.enqueue("2")
	q.enqueue("48")
	fmt.Println("List")
	fmt.Println(q)
	q.dequeue()
	q.dequeue()
	fmt.Println(q)
}
