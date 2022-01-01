// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
// 打印输入中多次出现的行的个数和文本，
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 字符串:计数 映射
	counts := make(map[string]int)
	// 文件
	files := os.Args[1:]
	if len(files) == 0 {
		// 计数
		countLines(os.Stdin, counts)
	} else {
		// 遍历所有文件
		for _, arg := range files {
			// 打开文件
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// 计数
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
