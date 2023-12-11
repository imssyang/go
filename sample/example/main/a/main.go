package main

import "fmt"

// 基类
type Base struct {
	Name string
}

// 基类的方法
func (b *Base) Foo() {
	fmt.Println("Base Foo")
}

// 派生类
type Derived struct {
	Base
	Age int
}

// 派生类的方法
func (d *Derived) Bar() {
	fmt.Println("Derived Bar")
}

func main() {
	// 创建派生类实例
	d := &Derived{
		Base: Base{Name: "Derived"},
		Age:  20,
	}

	// 调用继承的方法
	d.Foo()
	d.Bar()
}
