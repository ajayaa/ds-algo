package stack

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
