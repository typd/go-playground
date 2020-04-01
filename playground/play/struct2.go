package play

import "time"

// can add method to a struct in the same package
func (a A) x() string {
	return "x method of A"
}

// cant' add method to a struct to a different package
// func (a A) x() string{
//	return "x method of A"
//}

func TryStruct2() {
	var t time.Duration
	_ = t
	a := A{}
	println(a.x())
}
