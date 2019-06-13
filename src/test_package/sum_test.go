package test_package

import "testing"

func TestSum3(t *testing.T) {
	v := Sum(1, 2)
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}
