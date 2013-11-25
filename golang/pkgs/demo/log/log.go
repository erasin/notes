package main

import (
	"log"
	"os"
)

func init() {
	logfile, _ := os.Create("demo.log")
	// 输出到file
	log.SetOutput(logfile)
	log.SetPrefix("[test]")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func main() {
	log.Println(log.Flags())
	log.Println("第一行日志")

	log.SetPrefix("[dier]")
	log.Println("第二行日志")

	log.Panicln("Error ?")
	log.Fatalln("输出后退出")
	log.Panicln("不会被打印")
}
