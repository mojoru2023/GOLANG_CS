package main


//使用http客户端发送请求
//使用http.Client控制请求头部等
//使用httputil简化工作
//其他标准库
//bufio
//log
//encoding/json
//regexp
//time
//string/math/rand
//标准库网站
//godoc -http:8888 (需要安装godoc)
//http://www.studygolang/pkgdoc

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	_"net/http/pprof"
)

func main()  {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	// resp, err := http.Get("http://www.imooc.com")
	if err!=nil{
	panic(err)}
}
	respClient := http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	fmt.Println("Redirect:", req)
	return nil
	},
	}
	resp , errClientDo := respClient.Do(request)
	if errClientDo != nil{
		panic(err)
}
	defer resp.Body.Close()
	s,errDumpResponse := httputil.DumpResponse(resp,true)
	if errDumpResponse := nil{
		panic(errDumpResponse)
}
	fmt.Printf("%s",s)
}
