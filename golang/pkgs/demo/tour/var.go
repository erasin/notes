package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 64
	// 所有类型都是数字
	fmt.Println(strconv.Itoa(a))
	fmt.Println(string(a))
}
