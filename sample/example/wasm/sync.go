package main

import "syscall/js"

func syncFib(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(fib(args[0].Int()))
}

func syncMain() {
	js.Global().Set("syncFib", js.FuncOf(syncFib))
}
