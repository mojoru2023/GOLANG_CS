package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "https://www.baidu.com/s?ie=UTF-8&wd=github#test"
	// func Parse(rawurl string) (url *URL, err error)
	// Parse函数解析rawurl为一个URL结构体，rawurl可以是绝对地址，也可以是相对地址。
	u, err := url.Parse(urlString)

	if err != nil {
		return
	}
	fmt.Printf("%T, %v\n", u, u)            //*url.URL, https://www.baidu.com/s?ie=UTF-8&wd=github
	fmt.Printf("Scheme=%v\n", u.Scheme)     //Scheme=https
	fmt.Printf("Opaque=%v\n", u.Opaque)     //Opaque=
	fmt.Printf("User=%v\n", u.User)         //User=
	fmt.Printf("Host=%v\n", u.Host)         //Host=www.baidu.com
	fmt.Printf("Path=%v\n", u.Path)         //Path=/s
	fmt.Printf("RawPath=%v\n", u.RawPath)   //RawPath=
	fmt.Printf("RawQuery=%v\n", u.RawQuery) //RawQuery=ie=UTF-8&wd=github
	fmt.Printf("Fragment=%v\n", u.Fragment) //Fragment=test

	// func ParseRequestURI(rawurl string) (url *URL, err error)
	// ParseRequestURI函数解析rawurl为一个URL结构体，本函数会假设rawurl是在一个HTTP请求里，
	// 因此会假设该参数是一个绝对URL或者绝对路径，并会假设该URL没有#fragment后缀。
	// （网页浏览器会在去掉该后缀后才将网址发送到网页服务器）
	u, err = url.ParseRequestURI(urlString)
	fmt.Printf("Fragment=%v\n", u.Fragment) //Fragment=

	// IsAbs reports whether the URL is absolute.Absolute means that it has a non-empty scheme.
	// func (u *URL) IsAbs() bool
	fmt.Println(u.IsAbs()) //true

	// func (u *URL) Query() Values
	// Query方法解析RawQuery字段并返回其表示的Values类型键值对。
	fmt.Println(u.Query()) //map[ie:[UTF-8] wd:[github#test]]

	// func (u *URL) RequestURI() string
	// RequestURI方法返回编码好的path?query或opaque?query字符串，用在HTTP请求里。
	fmt.Println(u.RequestURI()) ///s?ie=UTF-8&wd=github#test

	// func (u *URL) String() string
	// String将URL重构为一个合法URL字符串。
	fmt.Println(u.String()) //https://www.baidu.com/s?ie=UTF-8&wd=github#test
}