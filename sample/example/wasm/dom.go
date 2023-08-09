package main

import (
	"strconv"
	"syscall/js"
)

var (
	document = js.Global().Get("document")
	numEle   = document.Call("getElementById", "num_dom")
	ansEle   = document.Call("getElementById", "ans_dom")
	btnEle   = js.Global().Get("btn_dom")
)

func domFib(this js.Value, args []js.Value) interface{} {
	v := numEle.Get("value")
	if num, err := strconv.Atoi(v.String()); err == nil {
		ansEle.Set("innerHTML", js.ValueOf(fib(num)))
	}
	return nil
}

func domMain() {
	btnEle.Call("addEventListener", "click", js.FuncOf(domFib))
}
