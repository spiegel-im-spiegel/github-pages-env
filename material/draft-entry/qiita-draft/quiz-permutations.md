# 日能研の問題をGo言語で力づくで解いてみた(2018/11)

[2018年11月の日能研の広告の問題](https://www.nichinoken.co.jp/shikakumaru/201811_sa)を力づくで解くページを見かけた

- [日能研の問題を解く(2018/11) - Qiita](https://qiita.com/sora410/items/bd349c48e7920e7f2573)
- [日能研の問題をCで解いてみた(2018/11) - Qiita](https://qiita.com/98hira/items/7a704fa0d1082c2e1190)

ので，私は [Go 言語]で解いてみる。

```go
package main

import "fmt"

func Permutations(cards []int) <-chan []int {
	ch := make(chan []int)
	go func() {
		defer close(ch)
		perm(ch, make([]int, 0, len(cards)), cards)
	}()
	return ch
}
func dup(list []int) []int {
	l := make([]int, len(list), cap(list))
	copy(l, list)
	return l
}
func perm(ch chan<- []int, list []int, rest []int) {
	if len(rest) == 0 {
		ch <- dup(list)
		return
	}
	for i, v := range rest {
		restx := dup(rest)
		restx = append(restx[:i], restx[i+1:]...)
		listx := append(list, v)
		perm(ch, listx, restx)
	}

}

func list2num(list []int) int {
	v := 0
	for i, t := len(list)-1, 1; i >= 0; i, t = i-1, t*10 {
		v += list[i] * t
	}
	return v
}
func main() {
	for p := range Permutations([]int{1, 2, 3, 4, 5, 6}) {
		res := true
		for i := 2; i <= 6; i++ {
			if list2num(p[:i])%i != 0 {
				res = false
				break
			}
		}
		if res {
			fmt.Println(list2num(p))
		}
	}
}
```

これで[問題なく動作する](https://play.golang.org/p/z0fgDhSWzE4)。

標準ライブラリには順列を生成する関数はないし reduce のような気の利いた高階関数もないので，ひたすら for 文で回します（笑）

特に [Go 言語]っぽいトピックはないけど，強いて挙げるなら「配列は値だけど slice は配列への参照（単なるポインタじゃない）を表す」というところだろうか。なので slice の複製を作るのに `dup()` 関数みたいなのを用意する必要があるわけですね。

あとは順列を生成するのに goroutine でジェネレータ・パターンを構成していることだろうか。

## 参考

- [Go言語で順列組み合わせ(Permutations and Combinations) - Qiita](https://qiita.com/kokardy/items/a18b574d548c212e58e6)
- [gitchander/permutation: Simple permutation package for golang](https://github.com/gitchander/permutation)

[Go 言語]: https://golang.org/ "The Go Programming Language"
