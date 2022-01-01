// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// := 是声明字符串
	s, sep := "", ""
	//一般采用_ 放上没用的参数
	// 并且 _ 不能直接使用, 改称 line 使用
	// 1: 从第一个开始
	for line, arg := range os.Args[1:] {
		//fmt.Println(_)
		s += sep + arg
		sep = " "
		fmt.Printf("%d %s\n", line, arg)
	}
	fmt.Println(s)
}

//!-
