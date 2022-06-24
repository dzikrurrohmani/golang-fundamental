package main

import (
	"errors"
	"fmt"
)

/*
Jenis linked list:
1. Single
2. Double
3. Circular
*/

type node struct {
	data string
	next *node
}

type mySingleLinkedList struct {
	size int
	head *node
}

func (sl *mySingleLinkedList) addToHead(name string) {
	if isSame, err := sl.isSame(name); isSame {
		fmt.Println(err.Error())
		return
	}

	newNode := &node{
		data: name,
	}

	if sl.head == nil {
		sl.head = newNode
	} else {
		newNode.next = sl.head
		sl.head = newNode
	}
	sl.size++
}

func (sl *mySingleLinkedList) addToTail(name string) {
	if isSame, err := sl.isSame(name); isSame {
		fmt.Println(err)
		return
	}

	newNode := &node{
		data: name,
	}

	if sl.head == nil {
		sl.head = newNode
	} else {
		current := sl.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	sl.size++
}

func (sl *mySingleLinkedList) isSame(name string) (bool, error) {
	for node := sl.head; node != nil; node = node.next {
		if node.data == name {
			errorNya := fmt.Sprintf("Data %s sudah ada, masukkan data berbeda.. ;)", name)
			return true, errors.New(errorNya)
		}
	}
	return false, nil
}

func (sl *mySingleLinkedList) iterateList() {
	for node := sl.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

func NewLinkedList() *mySingleLinkedList {
	return &mySingleLinkedList{}
}

func main() {
	singleList := NewLinkedList()
	singleList.addToHead("Satu")
	singleList.addToHead("Dua")
	singleList.addToHead("Tiga")
	singleList.addToTail("Nol")
	singleList.addToHead("Nol")
	singleList.addToTail("Nol")
	singleList.iterateList()

	fmt.Println("size", singleList.size)
}

// return true, errors.New(fmt.Sprintf("Data %s sudah ada, masukkan data berbeda\n", name))
