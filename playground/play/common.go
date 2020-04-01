package play

import "fmt"

func println(a ...interface{}) {
	for _, item := range a {
		fmt.Print(item)
		fmt.Print(" ")
	}
	fmt.Println()
}

func printlnf(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a)
}

func Sum(a, b int) int {
	return a + b
}
