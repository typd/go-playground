package play

import (
	"os"
	"reflect"
)

func TryDiffType() {
	var d1 [3]int
	var d2 [2]int

	//d1 = d2 // wrong
	_ = d1
	_ = d2
}

func TryShowType() {
	intArray := [...]int{1, 2}
	intArray2 := [...]int{1, 2, 3}

	// intArray = intArray2 // can't do this because they're different types
	_ = intArray2

	intSlice2 := []int{1, 2}
	intSlice3 := []int{1, 2, 3}
	intSlice2 = intSlice3 // this can be done as slices type doesn't care length
	_ = intSlice2

	printlnf("intArray is type: %v", reflect.TypeOf(intArray))
	printlnf("len(intArray) is: %v", len(intArray))
	intSlice := intArray[:]
	printlnf("intSlice is type: %v", reflect.TypeOf(intSlice))
	printlnf("len(intSlice) is: %v", len(intSlice))
}

func TryTypeCast() {
	var e error
	pe, ok := e.(*os.PathError)
	_ = pe
	if ok {
		println("cast successfully")
	} else {
		println("cast fail")
	}
}
