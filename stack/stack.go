package stack

import "fmt"

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Node{value: value, next: s.top}
	s.size++
}

func (s *Stack) Pop() (ret interface{}) {
	if s.top == nil {
		return nil
	} else {
		ret, s.top = s.top.value, s.top.next
		s.size--
		return
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Print() {
	temp := s.top
	for temp != nil {
		fmt.Printf("%s ", temp)
	}
	fmt.Printf("\n")
}

func (s *Stack) Top() interface{} {
	if s.size != 0 {
		return s.top
	}
	return nil
}
