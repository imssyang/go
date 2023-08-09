package main

import (
	"strconv"
	"syscall/js"
	"time"
)

func asyncFib(this js.Value, args []js.Value) interface{} {
	var num = args[0].Int()
	var callback = args[len(args)-1]
	var ansEle = js.Global().Get("ans_async")
	ansEle.Set("innerHTML", "Waiting 3s @ "+strconv.Itoa(num))

	go func() {
		time.Sleep(3 * time.Second)
		ans := fib(num)
		ansEle.Set("innerHTML", "ans: "+strconv.Itoa(ans))
		time.Sleep(time.Second)
		callback.Invoke("Callback: " + strconv.Itoa(ans))
	}()
	return nil
}

func asyncMain() {
	js.Global().Set("asyncFib", js.FuncOf(asyncFib))
}
