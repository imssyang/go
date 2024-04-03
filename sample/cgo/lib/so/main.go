package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -Wl,-rpath=./ -lnum

#include "num.h"
*/
import "C"
import "fmt"

func main() {
    fmt.Println(C.add_mod(10, 5, 12))
}

