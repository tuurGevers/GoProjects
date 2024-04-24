package main

import (
	"testing"
)

func TestMainGenericTypes(t *testing.T) {
	list := &List[int]{}

	// Initially insert elements
	list.InsertBack(1)
	list.InsertBack(2)
	list.InsertBack(3)

	if list.String() != "1 ->2 ->3 ->" {
		t.Errorf("Expected list to be '1 ->2 ->3 ->', got '%s'", list)
	}

	list.InsertFront(0)

	if list.String() != "0 ->1 ->2 ->3 ->" {
		t.Errorf("Expected list to be '0 ->1 ->2 ->3 ->', got '%s'", list)
	}

	list.InsertAt(2, 1)

	if list.String() != "0 ->1 ->1 ->2 ->3 ->" {
		t.Errorf("Expected list to be '0 ->1 ->1 ->2 ->3 ->', got '%s'", list)
	}

	list.RemoveBack()

	if list.String() != "0 ->1 ->1 ->2 ->" {
		t.Errorf("Expected list to be '0 ->1 ->1 ->2 ->', got '%s'", list)
	}

	list.RemoveFront()

	if list.String() != "1 ->1 ->2 ->" {
		t.Errorf("Expected list to be '1 ->1 ->2 ->', got '%s'", list)
	}

	list.RemoveAt(1)

	if list.String() != "1 ->2 ->" {
		t.Errorf("Expected list to be '1 ->2 ->', got '%s'", list)
	}

	// Check finding the element
	node, err := list.Find(2)
	if err != nil {
		t.Errorf("Expected to find element 2, but got error: %v", err)
	} else if node.Value != 2 {
		t.Errorf("Expected node value to be 2, but got %v", node.Value)
	}

	if list.String() != "1 ->2 ->" {
		t.Errorf("Expected list to be '1 ->2 ->', got '%s'", list)
	}
}
