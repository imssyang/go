package main

import (
	"fmt"

	"imssyang.com/cgo/pkg/dlopen"
)

func main() {
	r, err := dlopen.AddMod([]string{"num/libnum.so"}, 10, 5, 12)
	fmt.Println(r, err)
}
