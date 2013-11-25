package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// 设定 --help
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Progame ver...\n")
		fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
		flag.PrintDefaults()
	}

	// 常用的设定 flag
	file := flag.String("file", "default.txt", "to input a file name!")
	count := flag.Int("count", 0, "process count.")

	// 存储
	var zip bool
	flag.BoolVar(&zip, "zip", false, "is compress zip file.")

	flag.Parse() // 解析

	// 输出
	fmt.Printf("flags: %d", flag.NFlag())
	// 给出指针
	fmt.Println(*file, *count, zip)

	// Args
	fmt.Printf("Args: %d,%v\n", flag.NArg(), flag.Args())

	// set

	fmt.Printf("count now is : %d \n", *count)
	// 无论是什么,给出都是字符串,flag 会读应设置来取值
	flag.Set("count", "10")
	fmt.Printf("set count after is : %d \n", *count)

	// 直接取出name
	zip2 := flag.Lookup("zip")
	fmt.Printf("file is %s? %s \n", zip2.Name, zip2.Value.String())

}
