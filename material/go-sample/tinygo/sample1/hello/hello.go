package main

import "syscall/js"

func main() {
	js.Global().Get("document").Call("getElementById", "hello").Set("innerHTML", "Hello, World!")
	select {}
}
