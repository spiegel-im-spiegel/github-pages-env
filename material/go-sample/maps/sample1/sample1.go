package main

import "fmt"

type KeyValues map[string]string

func (kv KeyValues) Set(k, v string) {
	kv[k] = v
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}

func main() {
	kv := KeyValues{}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
	kv.Set("foo", "bar")
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}
