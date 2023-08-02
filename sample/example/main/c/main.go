package main

import "fmt"

func main() {
	// Slice
	ws := []string{"a", "b", "c"}
	for i := range ws { // 只迭代index（与for循环性能一致）
		if i == 1 {
			// 循环中修改切片，不在当前循环生效
			ws = append(ws, "d")
		}
		fmt.Println(i, ws[i]) // 0 a  1 b  2 c
	}
	for i, w := range ws { // 迭代index和value
		fmt.Println(i, w) // 0 a  1 b  2 c  3 d
	}

	// Map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		delete(m, "b")               // 迭代过程中，删除还未迭代到的键值对（b和c），则该键值对不会被迭代
		m["d"] = 4                   // 迭代过程中，新创建的键值对，不一定会被迭代
		fmt.Printf("%v: %v\n", k, v) // a 1  d 4  c 3
	}

	// Chan
	ch := make(chan string)
	go func() {
		ch <- "a"
		ch <- "b"
		ch <- "c"
		close(ch)
	}()
	for s := range ch {
		fmt.Println(s) // a  b  c
	}
}
