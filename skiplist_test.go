package skiplist

import (
	"testing"
)

func TestSkipList_Basic(t *testing.T) {
	sl := NewSkipList()
	// insert/get one element
	sl.Insert("1", 1)
	if sl.Size() != 1 {
		t.Errorf("fail to insert element")
	}
	val, exit := sl.Get("1")
	if !exit || val != 1 {
		t.Errorf("fail to get element")
	}
	// update an exist element
	sl.Insert("1", 11)
	if sl.Size() != 1 {
		t.Errorf("fail to insert element")
	}
	val, exit = sl.Get("1")
	if !exit || val != 11 {
		t.Errorf("fail to get element")
	}
	// add another element
	sl.Insert("2", 2)
	if sl.Size() != 2 {
		t.Errorf("fail to insert element")
	}
	val, exit = sl.Get("2")
	if !exit || val != 2 {
		t.Errorf("fail to get element")
	}
	// and another one...
	sl.Insert("3", 3)
	if sl.Size() != 3 {
		t.Errorf("fail to insert element")
	}
	val, exit = sl.Get("3")
	if !exit || val != 3 {
		t.Errorf("fail to get element")
	}
}
