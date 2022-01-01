// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
// 启动网页，然后直接执行
// go run main.go http://localhost:8000
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 获取 网页 URL 
	for _, url := range os.Args[1:] {
		// GET 
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 读取 GET 的内容
		b, err := ioutil.ReadAll(resp.Body)
		// 关闭
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// 打印夺取到的 字符串
		fmt.Printf("%s", b)
	}
}

//!-
