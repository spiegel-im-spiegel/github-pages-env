package main

import "fmt"

type KeyValues map[string]string

func (kv *KeyValues) Initialize(k, v string) {
	*kv = KeyValues{k: v}
	fmt.Printf("%p: %p: %v\n", kv, *kv, *kv)
}

func main() {
	kv := KeyValues{}
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
	kv.Initialize("foo", "bar")
	fmt.Printf("%p: %p: %v\n", &kv, kv, kv)
}
