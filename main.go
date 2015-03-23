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
	stdout "fmt"
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

	// use alia module name
	stdout.Println("use alia for 'fmt' modules")

	// use stuct
	obj := new(person)
	obj.age = 1
	obj.name = "scott"
	fmt.Println(obj.age)
	fmt.Println(obj.name)
	fmt.Println(obj)

	var obj2 person
	obj2.age = 2
	obj2.name = "alan"
	fmt.Println(obj2)

	obj3 := person{age: 30, name: "Donald"}
	fmt.Println(obj3)

	// test golang OOP
	obj4 := new(Student)
	obj4.age = 30
	obj4.name = "Lulu"
	obj4.skill = "study"
	obj4.echo = echo_action
	fmt.Println(obj4.age)
	fmt.Println(obj4.echo())
	fmt.Println(obj4)
	obj4.test_self_action()
	obj4.test_run()
	obj4.test_run_try()
	//	obj4.test_echo = test_self_action
	//	obj4.test_echo()

	obj5 := new(Rect)
	obj5.height = 1.5
	obj5.width = 1.5
	fmt.Print("Rect area->")
	fmt.Println(obj5.area())

	// Interface Study Start
	_a := new(A)
	_a.name = "Scott"
	_a.echo()

	_b := new(B)
	_b.name = "Alan"
	_b.echo()

	var _c A_and_B
	_c = _b
	_c.echo()
	_c = _a
	_c.echo()

	// 作為通用容器用途, 但是所有特性[Method, Property]都會不能存取
	var _d interface{}
	_d = _a
	fmt.Println(_d)
	// _c = _d
	// _c.echo()			// 將出錯, 因為先前已將特性去除

	// Interface Study End

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
	panic("Throw some error")
}

// define my struct
type person struct {
	name string
	age  int
}

// golang OOP by struct
type Human struct { // parent schema
	name string
	age  int
}

func (r *Human) test_run() {
	fmt.Println("Human test_run by " + r.name)
}

func (r *Human) test_run_try() {
	fmt.Println("Human test_run_try by " + r.name)
}

type echo func() bool // method mixin
type test_echo func() // method mixin

type Student struct { // final schema
	Human
	echo
	test_echo // 可以不定義, 就像 test_run Method 一樣
	skill     string
}

func echo_action() bool {
	fmt.Println("student echo!")
	return true
}

func (r *Student) test_self_action() {
	fmt.Print("this is instance method get age->")
	fmt.Println(r.age)
}

// override - func (r *Human) test_run()
func (r *Student) test_run() {
	fmt.Println("test_run by " + r.name)
}

// OOP Style
type Rect struct {
	width, height float64
}

func (self *Rect) area() float64 {
	return self.height * self.width
}

// OOP Style End

// Interface Style

type A struct {
	name string
}

type B struct {
	name string
}

func (self A) echo() { // 這樣才是正確的, 如果用 *A 會變成 Reference 但是於 Golang 內會自動轉換依然可以正常 Work！
	fmt.Println(self)
	fmt.Println(self.name)
	fmt.Println("I'm A echoing!")
}

func (self B) echo() { // 這樣才是正確的, 如果用 *A 會變成 Reference 但是於 Golang 內會自動轉換依然可以正常 Work！
	fmt.Println(self)
	fmt.Println(self.name)
	fmt.Println("I'm B echoing!")
}

type A_and_B interface {
	echo()
}

// Interface Style End
