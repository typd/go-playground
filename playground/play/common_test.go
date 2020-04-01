package play

import (
	"testing"
)

func TestSum(t *testing.T) {
	testSumImpl(3, 4, t)
	testSumImpl(2, 1.0, t)
	testSumImpl(-1, 1.0, t)
	testSumImpl(0, 0, t)
}

func testSumImpl(a, b int, t *testing.T){
	var r int
	// r = Sum(a, b) // TODO why can't compile here?
	r = a + b
	exp := a + b
	if r != exp{
		t.Errorf("Expected %v, but got %v", exp, r)
	}
}

