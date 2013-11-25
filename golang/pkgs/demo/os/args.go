package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%T , %v\n", os.Args, os.Args)
}
