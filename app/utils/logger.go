package utils

import (
	"fmt"
)

// Info 显示信息
func Info(args ...interface{}) {
	print(args...)
}

// Debug 调试信息
func Debug(args ...interface{}) {
	print(args...)
}

// print
func print(args ...interface{}) {
	fmt.Println(args...)
}
