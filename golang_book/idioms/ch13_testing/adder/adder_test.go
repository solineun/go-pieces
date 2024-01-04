package adder

import "testing"

func Test_add(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Error("incorrect result")
	}
}

