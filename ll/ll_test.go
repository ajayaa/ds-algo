package ll

import (
	"reflect"
	"testing"
)

func TestMergenode(t *testing.T) {
	n1 := &node{value: 1}
	n2 := &node{value: 2}
	exp := &node{value: 1}
	exp.next = &node{value: 2}
	res := mergeLL(n1, n2)
	t.Log(res)
	if reflect.DeepEqual(res, exp) != true {
		t.Fatal("error")
	}
	//(&linkList{head: res}).print()
}
