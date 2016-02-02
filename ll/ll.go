package ll

import (
	"bytes"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func (n *Node) String() string {
	var buf bytes.Buffer
	for n != nil {
		buf.WriteString(string(n.value))
		buf.WriteString(" ")
		n = n.next
	}
	return buf.String()
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(value int) {
	if l.head == nil {
		l.head = &Node{value: value}
	} else {
		temp := &Node{value: value}
		temp.next = l.head
		l.head = temp
	}
}

func (l *LinkedList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.value)
		if temp.next != nil {
			fmt.Print("->")
		}
		temp = temp.next
	}
}

func (l *LinkedList) sort() {
}

/* Merge two sorted linked lists.
This functions takes head as arguments. */
func MergeLL(n1, n2 *Node) *Node {
	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	} else {
		var temp *Node
		if n1.value < n2.value {
			temp = n1
			n1 = n1.next
		} else {
			temp = n2
			n2 = n2.next
		}
		for n1 != nil && n2 != nil {
			if n1.value < n2.value {
				temp.next = n1
				n1 = n1.next
			} else {
				temp.next = n2
				n2 = n2.next
			}
		}
		if n1 == nil {
			temp.next = n2
		} else {
			temp.next = n1
		}
		return temp
	}
}
