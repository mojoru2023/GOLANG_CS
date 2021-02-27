package main

import "fmt"

func main() {
    fmt.Println("Hello, world")
}


import "net/http"
...
resp, err := http.Get("http://wwww.baidu.com")
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    fmt.Println("Hello, world")
    resp, err := http.Get("http://www.baidu.com/")
    if err != nil {
        fmt.Println("http get error", err)
        return
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read error", err)
        return
    }
    fmt.Println(string(body))
}

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    fmt.Println("Hello, world")
    url := "http://www.baidu.com/"
    download(url)
}

func download(url string) {
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    // 自定义Header
    req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("http get error", err)
        return
    }
    //函数结束后关闭相关链接
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read error", err)
        return
    }
    fmt.Println(string(body))
}

package main

import (
    "fmt"
    "github.com/jackdanger/collectlinks"
    "net/http"
)

func main() {
    fmt.Println("Hello, world")
    url := "http://www.baidu.com/"
    download(url)
}

func download(url string) {
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    // 自定义Header
    req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("http get error", err)
        return
    }
    //函数结束后关闭相关链接
    defer resp.Body.Close()

    links := collectlinks.All(resp.Body)
    for _, link := range links {
        fmt.Println("parse url", link)
    }
}

package main

import (
    "fmt"
    "github.com/jackdanger/collectlinks"
    "net/http"
)

func main() {
    fmt.Println("Hello, world")
    url := "http://www.baidu.com/"

    queue := make(chan string)
    go func() {
        queue <- url
    }()
    for uri := range queue {
        download(uri, queue)
    }
}

func download(url string, queue chan string) {
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    // 自定义Header
    req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("http get error", err)
        return
    }
    //函数结束后关闭相关链接
    defer resp.Body.Close()

    links := collectlinks.All(resp.Body)
    for _, link := range links {
        fmt.Println("parse url", link)
        go func() {
            queue <- link
        }()
    }
}

package main

import (
    "fmt"
    "github.com/jackdanger/collectlinks"
    "net/http"
    "net/url"
)

func main() {
    fmt.Println("Hello, world")
    url := "http://www.baidu.com/"

    queue := make(chan string)
    go func() {
        queue <- url
    }()
    for uri := range queue {
        download(uri, queue)
    }
}

func download(url string, queue chan string) {
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    // 自定义Header
    req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("http get error", err)
        return
    }
    //函数结束后关闭相关链接
    defer resp.Body.Close()

    links := collectlinks.All(resp.Body)
    for _, link := range links {
        absolute := urlJoin(link, url)
        if url != " " {
            fmt.Println("parse url", absolute)
            go func() {
                queue <- absolute
            }()
        }
    }
}

func urlJoin(href, base string) string {
    uri, err := url.Parse(href)
    if err != nil {
        return " "
    }
    baseUrl, err := url.Parse(base)
    if err != nil {
        return " "
    }
    return baseUrl.ResolveReference(uri).String()
}


package main

import (
    "fmt"
    "github.com/jackdanger/collectlinks"
    "net/http"
    "net/url"
)

var visited = make(map[string]bool)

func main() {
    fmt.Println("Hello, world")
    url := "http://www.baidu.com/"

    queue := make(chan string)
    go func() {
        queue <- url
    }()
    for uri := range queue {
        download(uri, queue)
    }
}

func download(url string, queue chan string) {
    visited[url] = true
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    // 自定义Header
    req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("http get error", err)
        return
    }
    //函数结束后关闭相关链接
    defer resp.Body.Close()

    links := collectlinks.All(resp.Body)
    for _, link := range links {
        absolute := urlJoin(link, url)
        if url != " " {
            if !visited[absolute] {
                fmt.Println("parse url", absolute)
                go func() {
                    queue <- absolute
                }()
            }
        }
    }
}

func urlJoin(href, base string) string {
    uri, err := url.Parse(href)
    if err != nil {
        return " "
    }
    baseUrl, err := url.Parse(base)
    if err != nil {
        return " "
    }
    return baseUrl.ResolveReference(uri).String()
}

