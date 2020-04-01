package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/typd/go-playground/playground/play"
)

var (
	g int
)

func println(a interface{}) {
	fmt.Println(a)
}

func fprintln(format string, a interface{}) {
	fmt.Printf(format+"\n", a)
}

func tryIf() {
	if 1 < 2 {
		fmt.Printf("aaaa %d    ", 3, 4)
	}
}

func trySwitch() {
	var a = 3
	switch {
	case 30 > a:
		fmt.Println("caaaa", "aa")
	case 1 < 2:
		fmt.Println("aaaa", "aa")
	}
	fmt.Println("xcc")
}

func init() {
	println("--------------")
	println("Now init() is called ...")
	fprintln("os.Args: %s", os.Args)
	fileInfo, err := os.Stat(".")
	if err != nil {
		println("error of calling os.Stat('.')")
		return
	}
	fprintln("fileInfo of '.': %s", fileInfo)
	println("--------------")
}

func tryInt() {
	var i int
	fmt.Printf("initial value of i: %d\n", i)
	i += 1
	for i > 0 {
		i += i
		fmt.Printf("now i is: %d\n", i)
	}
	/*
	output as below:
		...
		now i is: 4611686018427387904
		now i is: -9223372036854775808
	*/
	var i64 int64
	i64 += 1
	for i64 > 0 {
		i64 += i64
		fmt.Printf("now i64 is: %d\n", i64)
	}
}

func multiParamReturnValueFunc(a, b int) (sum, diff int) {
	return a + b, a - b
}

func array() {
	var a = make([]string, 100)
	fmt.Printf("a is %v, len of a is: %d\n", a, len(a))

	b := [...]float32{3.4, 5.6}
	fmt.Printf("b is %v, len of b is: %d\n", b, len(b))
}

type AType struct {
	i int
	f float32
	s string
}

func aboutBool() {
	var a bool
	fmt.Println("default bool value is:", a) // false
}

func tryStd() {
	fmt.Fprintln(os.Stdout, "something to Stdout")
	fmt.Fprintln(os.Stderr, "something to Stderr")
}

func (t AType) String() string {
	// below is wrong because of dead loop
	// return fmt.Sprintf("This is String() of AType: %v", t)
	return fmt.Sprintf("This is String() of AType: %v/%v/%q", t.i, t.f, t.s)
}

func tryType() {
	a := &AType{3, 4, "ss"}
	println(a)
}

func tryAppend() {
	x := []int{1, 2, 3}
	b := []int{}
	b = append(x, 4, 5, 9)
	fprintln("b is : %v", b)
	fprintln("x is : %v", x)
}

type Seq []int

func (s Seq) Len() int {
	return len(s)
}
func (s Seq) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Seq) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func trySort() {
	seq := Seq{1, 2, 9, 3, 5}
	fprintln("before sort: %v", seq)
	sort.Sort(seq)
	// the other way
	sort.IntSlice(seq).Sort()
	fprintln("after sort : %v", seq)
}

func tryCastImpl(value interface{}) {
	str, ok := value.(string)
	if ok {
		fmt.Printf("string value is: %q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}

func tryCast() {
	tryCastImpl(2)
	tryCastImpl("strss")
}

func main() {
	play.TryShowType()
	//play.TryDeferAndReturnTime()
	//play.TryDeferValueAssignment()
	//play.TryTimeout()
	// play.TryTime()
	// play.Sum(1, 3)
	// play.struct.Try()
	// tryCast()
	// trySort()
	// tryAppend()
	// tryType()
	// tryInt()
	// tryStd()
	// aboutBool()
	// tryIf()
	// trySwitch()
	// fmt.Println(multiParamReturnValueFunc(2, 4))
	// array()
}
