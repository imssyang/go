package main

import "fmt"

type Person struct {
	name string
}

func newPerson(name string) *Person {
	p := new(Person)
	p.name = name
	return p // pointer作为返回值，导致局部变量 p 逃逸到堆上
}

func printPerson(p *Person) {
	fmt.Println(p.name) // 函数参数为 interface{} 时，编译期间很难确定其具体类型，会逃逸到堆上
}

func main() {
	john := newPerson("John")
	printPerson(john)
}

/*
$ go build -gcflags=-m ./main/c
# example.com/main/c
...
main/c/main.go:9:16: leaking param: name
main/c/main.go:10:10: new(Person) escapes to heap
main/c/main.go:15:18: leaking param content: p
main/c/main.go:16:13: ... argument does not escape
main/c/main.go:16:15: p.name escapes to heap
main/c/main.go:20:19: new(Person) does not escape
main/c/main.go:21:13: ... argument does not escape
main/c/main.go:21:13: p.name escapes to heap
*/
