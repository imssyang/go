package main

import "syscall/js"

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("Hi!")

	done := make(chan int, 0)
	syncMain()
	asyncMain()
	domMain()
	<-done
}

/*
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
$ go install github.com/shurcooL/goexec
$ goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'
*/
