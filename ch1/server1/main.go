// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
// 一个简单的 web http服务
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 处理 /, 也就是 localhost:8000/ 中的 "/"
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q, 这是一个测试Web服务器\n", r.URL.Path)
}

//!-
