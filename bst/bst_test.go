package bst

import (
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	tree := BST{}
	nums := []int{5, 1, 2, 6}
	for _, num := range nums {
		tree.Insert(num)
	}
	assert := func(a, b int) {
		t.Log(a, b)
		if reflect.DeepEqual(a, b) != true {
			t.Fatal("error")
		}
	}
	assert(tree.root.value, 5)
	assert(tree.root.left.value, 1)
	assert(tree.root.left.right.value, 2)
	assert(tree.root.right.value, 6)
}
