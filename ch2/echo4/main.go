// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	// flag 包
	// 指针对于flag 包是很有用的
	// flag 包使用程序的命令行参数来设置整个程序内某个变量的值
	"flag"
	"fmt"
	"strings"
)

// -n
var n = flag.Bool("n", false, "omit trailing newline")
// -s
var sep = flag.String("s", " ", "separator")

func main() {
	// 解析 -n 和 -s 参数
	// 输入 -h 将打印所有信息
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

//!-
