package play

func TryDeferValueAssignment() {
	i := 0
	defer printlnf("original i==0, defer i: %v", i)
	i = 1
	printlnf("i: %v", i)
}

func TryDeferAndReturnTime() {
	println(c()) // will be 2
}

func c() (i int) {
	defer func(){i++}() // this will be executed after the return statement
	return 1 // will return 2 actually
}
