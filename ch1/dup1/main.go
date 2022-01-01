// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
// 找出重复行
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建 映射
	counts := make(map[string]int)
	// 创建 扫描器
	input := bufio.NewScanner(os.Stdin)
	// 遍历扫描器
	for input.Scan() {
		// 计数+1
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		// >1 的才打印
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
