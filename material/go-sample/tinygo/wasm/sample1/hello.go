// +build js,wasm

package main

import "syscall/js"

func main() {
	ch := make(chan struct{})
	js.Global().Get("document").Call("getElementById", "hello").Set("innerHTML", "Hello, World!")
	<-ch // Code must not finish
}
