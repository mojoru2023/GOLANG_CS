package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawurl := "ie=UTF-8&wd=github#test"
	v, _ := url.ParseQuery(rawurl)
	fmt.Println(v) //map[ie:[UTF-8] wd:[github#test]]
	rawurl = "https://www.baidu.com/s?ie=UTF-8&wd=github#test"

	// func ParseQuery(query string) (m Values, err error)
	// ParseQuery函数解析一个URL编码的查询字符串，并返回可以表示该查询的Values类型的字典。
	u, _ := url.Parse(rawurl)
	v = u.Query()
	fmt.Println(v) //map[ie:[UTF-8] wd:[github]]

	// func (v Values) Set(key, value string)
	// Set方法将key对应的值集设为只有value，它会替换掉已有的值集。
	v.Set("name", "Tom")

	// func (v Values) Add(key, value string)
	// Add将value添加到key关联的值集里原有的值的后面。
	v.Add("name", "Jack")
	fmt.Println(v) //map[ie:[UTF-8] name:[Tom Jack] wd:[github]]

	// func (v Values) Del(key string)
	// Del删除key关联的值集。
	v.Del("ie")
	v.Del("ig")
	fmt.Println(v) //map[name:[Tom Jack] wd:[github]]

	// func (v Values) Get(key string) string
	// Get会获取key对应的值集的第一个值。如果没有对应key的值集会返回空字符串。获取值集请直接用map。
	fmt.Println(v.Get("ig"))   //
	fmt.Println(v.Get("name")) //Tom

	// func (v Values) Encode() string
	// Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序
	fmt.Println(v.Encode()) //name=Tom&name=Jack&wd=github
}