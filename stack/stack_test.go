package stack

import (
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	s := Stack{}
	str := "aj"
	s.Push(str)
	val := s.Pop()
	t.Logf("%T\n", val)
	if reflect.DeepEqual(str, val) != true {
		t.Fatal("error")
	}
}
