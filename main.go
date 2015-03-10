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
}
