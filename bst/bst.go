package bst

import (
	"fmt"
	"github.com/ajayaa/ds-algo/stack"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{root: nil}
}

func (tree *BST) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{value: value}
	} else {
		//temp := tree.root
		var prev *Node
		for temp := tree.root; temp != nil; {
			prev = temp
			if temp.value < value {
				temp = temp.right
			} else if temp.value > value {
				temp = temp.left
			} else {
				panic("I don't support duplicates!")
			}
		}
		if prev.value < value {
			prev.right = &Node{value: value}
		} else {
			prev.left = &Node{value: value}
		}
	}
}

func (tree *BST) InOrder() {
	tree.root.InOrder()
}

func (tree *BST) PreOrder() {
	tree.root.PreOrder()
}

func (tree *BST) PostOrder() {
	tree.root.PostOrder()
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.left.InOrder()
	fmt.Println(n.value)
	n.right.InOrder()
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	n.left.PreOrder()
	n.right.PreOrder()
}

func (n *Node) PostOrder() {
	if n == nil {
		return
	}
	n.left.PostOrder()
	n.right.PostOrder()
	fmt.Println(n.value)
}

func (tree *BST) IterativeInOrder() {
	if tree.root == nil {
		return
	}
	stack := stack.Stack{}

	push_all_left := func(n *Node) {
		for {
			if n == nil {
				break
			}
			stack.Push(n)
			n = n.left
		}
	}
	temp := tree.root
	push_all_left(temp)
	for !stack.IsEmpty() {
		popped := stack.Pop().(*Node)
		fmt.Println(popped.value)
		if popped.right != nil {
			temp = popped.right
			push_all_left(temp)
		}
	}
}
