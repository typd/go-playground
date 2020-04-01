package play

func aboutMap() {
	var m = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	printlnf("This is %v, len of it is: %d\n", m, len(m))

	for k, v := range m {
		println("map item: ", k, v)
	}

	key := "randomKey"
	if _, has := m[key]; has {
		println("m has key: ", key)
	} else {
		println("m doesn't have key: ", key)
	}
	key = "UTC"
	if _, has := m[key]; has {
		println("m has key: ", key)
	} else {
		println("m doesn't have key: ", key)
	}

	println()

	var m2 = map[string]string{
		"aa": "bb",
	}
	println("m2[aa]", m2["aa"])
	println("m2[b]", m2["b"])
}

func TryMap() {
	aboutMap()
}
