package main

// Example

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	a := Add(1, 10)
	b := Add(1, 20)
	fmt.Println(a)
	fmt.Println(b)

	// Output:
	// 11
	// 21

}
