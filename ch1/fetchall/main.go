// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
// 并法获取多个 URL  
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// 开始
	start := time.Now()
	// 管道,当然也可以叫做通道
	ch := make(chan string)
	// 遍历输入的所有URL
	for _, url := range os.Args[1:] {
		// go 的并法机制 
		go fetch(url, ch) // start a goroutine
	}
	// 从通道中打印
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	// 计数
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	// 开始时间
	start := time.Now()
	// 从URL  GET 
	resp, err := http.Get(url)
	if err != nil {
		// 写入 通道
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	// 拷贝
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	// 计时
	secs := time.Since(start).Seconds()
	// 写入通道
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
