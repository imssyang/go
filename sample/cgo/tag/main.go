// +build linux,debug

package main

import (
        "fmt"
)

var buildMode = "debug"

func main() {
        fmt.Print(buildMode)
}
