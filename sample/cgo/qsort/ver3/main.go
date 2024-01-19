package main

import (
    "fmt"
    "imssyang.com/cgo/qsort"
)

func main() {
    values := []int32{42, 9, 101, 95, 27, 25}

    qsort.Slice(values, func(i, j int) bool {
        return values[i] < values[j]
    })

    fmt.Println(values)
}

