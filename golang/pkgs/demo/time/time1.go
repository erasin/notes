package main

import (
	"fmt"
	// "strconv"
	"time"
)

func main() {
	fmt.Println("time Unix: ", time.Now().Unix())
	//时间戳int转Time实例
	t := time.Unix(1362984425, 0000000000)
	fmt.Println(t)
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().String())
	// 格式化采用时间为 go doc 上的示例时间
	fmt.Println(time.Now().Format("2006 01 02 15:04:05"))
}
