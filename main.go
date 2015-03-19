// 這是一個基於 GoLang 所撰寫的網路爬蟲。
// ref: http://blog.enoir.tw/2014/09/15/%E6%B7%BA%E8%AB%87GO%E8%AA%9E%E8%A8%80%E4%B8%AD%E7%9A%84defer/
// ref: http://negaihoshi.logdown.com/posts/70147-go-way-of-language-learning-c
// ref: https://golang.org/doc/effective_go.html
// ref: http://zh.wikipedia.org/wiki/Go
// ref: http://golang.org/doc/install#windows
// ref: http://golang.org/pkg/net/http/
// ref: http://golang.org/pkg/bytes/
// ref: http://golang.org/pkg/regexp/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.hinet.net")
	if err != nil {
		fmt.Println("There are some error in this request.")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("There are same error in read data")
		return
	}

	fmt.Println()
	fmt.Printf("%s", body)

	fmt.Println()
	reg := regexp.MustCompile("<.*?>")

	fmt.Println()
	matchedList := reg.FindAllString(string(body), -1)

	fmt.Println()
	for idx, elm := range matchedList {
		fmt.Println(idx)
		fmt.Println(elm)
	}

	a := map[string]int{"a": 1}
	fmt.Println(a)
	const (
		c = iota
		e = iota
		d = iota
	)
	fmt.Println(c)
	fmt.Println(d)

	// test make T, len, capacity
	cc := make([]int, 4, 8)
	cc[0] = 1
	cc[1] = 2
	cc[2] = 2
	cc[3] = 2
	// redesing array in limit of his capacity.
	cc = cc[:6]
	// redesing array in limit of his capacity again.
	cc = cc[:8]

	// will throw error due to out of his capacity
	// cc = cc[:9]

	fmt.Println(cc)
	fmt.Println(cap(cc))
	fmt.Println(len(cc))

	// test new
	dd := new(string)
	*dd = "2"
	fmt.Println(*dd)
	// mean dd = is address point to where is a *dd value
	fmt.Println(dd)
	fmt.Println(&dd) // get address of this point dd, it's difference with dd.

	// Golang Pointer
	// http://blog.golang.org/gos-declaration-syntax#TOC_3.
	var ee int
	fmt.Println(ee)

	var ff *int
	fmt.Println(ff)
	ff = &ee
	*ff = 70
	fmt.Println(ee)
}
