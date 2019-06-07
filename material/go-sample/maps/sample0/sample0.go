package main

import "fmt"

type KeyValues map[string]string

func main() {
	kv := KeyValues{"foo": "bar"}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}
