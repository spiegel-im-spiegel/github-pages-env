// +build run

package main

type ペンギン struct{}

func (p *ペンギン) 鳴く() string {
	return "ふるるー！"
}

func main() {
	println((&ペンギン{}).鳴く())
}
