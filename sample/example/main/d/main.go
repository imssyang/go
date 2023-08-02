package main

import "fmt"

func main() {
	// ForValue
	ts := []struct{ id int }{{id: 1}, {id: 2}, {id: 3}}
	for _, s := range ts {
		s.id += 10 // range对每个迭代值都创建了一个拷贝
	}
	for i := 0; i < len(ts); i++ {
		ts[i].id += 100
	}
	fmt.Println(ts) // [{101} {102} {103}]

	// ForPointer
	ps := []*struct{ id int }{{id: 1}, {id: 2}, {id: 3}}
	for _, s := range ps {
		s.id += 10
	}
	for i := 0; i < len(ps); i++ {
		ps[i].id += 100
	}
	fmt.Println(ps[0], ps[1], ps[2]) // &{111} &{112} &{113}
}
