package main

import "fmt"

type stack []string

func (s *stack) size() int {
	return (len(*s))
}

func (s *stack) push(name string) {
	*s = append(*s, name)
}

func (s *stack) pop() string {
	sz := s.size()
	data := (*s)[sz-1]
	*s = (*s)[:sz-1]
	return data
}

func main() {
	// s := make(stack, 0)
	var s stack
	//  s := new(stack)
	s.push("51")
	s.push("62")
	s.push("2")
	s.push("48")
	fmt.Println("Toples")
	fmt.Println(s)
	t := s.pop()
	fmt.Println("Pop", t)
	fmt.Println(s)
	t = s.pop()
	fmt.Println("Pop", t)
	fmt.Println(s)
}
