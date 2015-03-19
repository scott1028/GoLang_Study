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

	// Golang GoTo Statement
	count := 0
Savepoint:
	count += 1
	if count <= 10 {
		fmt.Println("Execute goto statement.")
		goto Savepoint
	}
	fmt.Println(count)

	// test go func return multi value
	kk := add1(1, 3)
	tt, yy := add2(10, 20)
	fmt.Println(kk)
	fmt.Println(yy)
	fmt.Println(tt)
	fmt.Println(add3(99, 2, 3))

	// 在 defer 后指定的函数会在函数退出前调用。(可用於 open-close file ...etc)
	defer fmt.Println("This is a defer logic #1")
	defer fmt.Println("This is a defer logic #2 (invoked first)")
	fmt.Println("Before defer logic")

	// pass a lambda func into func
	fmt.Println(add4(100, add1))

	// work for panic()
	defer func() {
		fmt.Println(recover())
	}()

	// not work for panic()
	//defer fmt.Println(recover())
	test()
}

// function name(args) returnType { ... }
func add1(x int, y int) int {
	return x + y
}

func add2(x int, y int) (int, int) {
	return x + y, x * y
}

// 不定參數
func add3(xs ...int) int {
	fmt.Println(len(xs))
	return xs[0]
}

// define a customize func type for using later
type lambda func(int, int) int

func add4(x int, fn lambda) int {
	return x + fn(x, x)
}

// test panic
func test() {
	panic("Throw some error.")
}
