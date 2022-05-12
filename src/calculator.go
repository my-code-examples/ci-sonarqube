package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Sum:", sum(1, 2))
	fmt.Println("Sub:", sub(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
	fmt.Println("Times:", times(1, 2))
}

func sum(i1, i2 int) int {
	return i1 + i2
}

func sub(i1, i2 int) int {
	return i1 - i2
}

func times(i1, i2 int) int {
	return i1 * i2
}

func div(i1, i2 int) int {
	if i2 == 0 {
		return 0
	} else if i1 < 0 {
		return -(-i1 / i2)
	} else {
		return i1 / i2
	}
}
