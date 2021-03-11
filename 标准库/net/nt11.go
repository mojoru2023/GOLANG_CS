package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawurl := "ie=UTF-8&wd=github#test"
	// func QueryEscape(s string) string
	// QueryEscape函数对s进行转码使之可以安全的用在URL查询里。
	stdurl := url.QueryEscape(rawurl)
	fmt.Println(stdurl) //ie%3DUTF-8%26wd%3Dgithub%23test
	// func QueryUnescape(s string) (string, error)
	// QueryUnescape函数用于将QueryEscape转码的字符串还原。它会把%AB改为字节0xAB，将'+'改为' '。如果有某个%后面未跟两个十六进制数字，本函数会返回错误。
	fmt.Println(url.QueryUnescape(stdurl)) //ie=UTF-8&wd=github#test <nil>
}