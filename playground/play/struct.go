package play

type A struct{
}

func (a A) a() string{
	return "a method of A"
}

func (a A) c() string{
	return "c method of A"
}

type B struct{
}

func (b B) b() string{
	return "b method of B"
}

func (b B) c() string{
	return "c method of B"
}

func (b B) d() string{
	return "d method of B"
}

type AB struct {
	A
	B
}

func TryStruct() {
	println("in play/struct/try")
	ab := AB{}
	ab.A = A{}
	ab.B = B{}
	println(ab.a())
	println(ab.b())
	println(ab.A.c())
	println(ab.B.c())
	println(ab.d())
	// ambiguous selector ab.c
	// can't call "ab.c()" because both ab.A and ab.B has c()
	// println(ab.c())
	printlnf("this is ab: %v", ab)
}

