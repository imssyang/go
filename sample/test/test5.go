package main

var c = make(chan int) // (unbuffered channel)
var a string

func f() {
	a = "Hi"
	<-c // (synchronized: receiver is executed after that sender is complete)
}

func main() {
	go f()
	c <- 0   // (sender)
	print(a) // Hi
}
