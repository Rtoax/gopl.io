// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	// 使用 join 汉书 将字符串连接
	// 这里不需要 for 循环
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//!-
