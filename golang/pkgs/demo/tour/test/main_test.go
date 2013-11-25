package main

import (
	"math/rand"
	"os"
	"testing"
)

// 使用 *testing.T 用于输出信息或中断测试
func TestSum(t *testing.T) {
	if rand.Intn(50) != 33 {
		t.Fatal("输出日志。。。")
		os.Exit(200)
	}
	t.Log("go...")
}
