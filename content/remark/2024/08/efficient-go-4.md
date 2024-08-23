+++
title = "『効率的な Go』読書会 4回目"
date =  "2024-08-23T16:12:17+09:00"
description = "description"
image = "/images/attention/go-logo_blue.png"
tags = [ "book", "golang", "engineering" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日『[効率的な Go]』の4回目の読書会が行われた。

- [第61回横浜Go読書会（オンライン） - connpass](https://yokohama-go-reading.connpass.com/event/325866/)

実は3回目は私用で欠席した。
あとで復習しないとなー。

前回の3回目と今回の4回目は3章の「効率化の攻略」という大事な章を含んでいるのだが，この辺の話を書き出すと（量的に）大変なことになるので，この記事では4章の「Go での CPU リソースの（ちょっとした）解説」の最初の方に出てくる Go コードをアセンブラコードに変換して眺める話を実際に手を動かして試すに留める。

なお，今回の [Go] コンパイラのバージョンは先日[リリース]({{< ref "/release/2024/08/go-1_23_0-is-released.md" >}} "Go 1.23 のリリース")されたばかりの 1.23.0 を使用している。
このため『[効率的な Go]』で例示されている内容とはだいぶ異なっている点に注意してほしい。

## サンプルコード

まずはサンプルコードを書いてみる。

```go
package main

import (
    "bytes"
    "fmt"
    "os"
    "strconv"
)

func Sum(fileName string) (ret int64, _ error) {
    b, err := os.ReadFile(fileName)
    if err != nil {
        return 0, err
    }
    for _, line := range bytes.Split(b, []byte("¥n")) {
        num, err := strconv.ParseInt(string(line), 10, 64)
        if err != nil {
            return 0, err
        }
        ret += num
    }
    return ret, nil
}

func main() {
    n, err := Sum("numbers.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println(n)
}
```

これは『[効率的な Go]』4.2節に出てくるコードに `main()` 関数を付加したものである。
なんで `main()` 関数を付加したかというと，そうしないと `go build` が通らないから。

## アセンブラコードを出力しながらビルドする

コンパイルオプションに `-S` を付加すると標準エラー出力にアセンブラコードを吐き出しながらビルドを行う。

```text
$ go build -gcflags="-S" sum.go
```

実際に吐き出した内容はこんな感じ。

```text
# command-line-arguments
main.Sum STEXT size=237 args=0x10 locals=0x88 funcid=0x0 align=0x0
    0x0000 00000 (/path/to/sum.go:10)    TEXT    main.Sum(SB), ABIInternal, $136-16
    0x0000 00000 (/path/to/sum.go:10)    LEAQ    -8(SP), R12
    0x0005 00005 (/path/to/sum.go:10)    CMPQ    R12, 16(R14)
    0x0009 00009 (/path/to/sum.go:10)    PCDATA    $0, $-2
    0x0009 00009 (/path/to/sum.go:10)    JLS    207
    0x000f 00015 (/path/to/sum.go:10)    PCDATA    $0, $-1
...
```

各インストラクションの意味については以下の記事が参考になる。

- [解説&翻訳 - A Quick Guide to Go's Assembler](https://zenn.dev/hsaki/articles/godoc-asm-ja)

ちなみにこのインストラクションは AMD (or Intel) 系のもので，たとえば Mx macOS 用にビルドすると異なるインストラクションが吐き出されるようだ。

## コンパイラによる自動最適化の決定を表示する

コンパイルオプションに `-m=<number>` を付加するとコンパイラによる自動最適化の決定を標準エラー出力に表示しながらビルドしてくれる。
こんな感じ。

```text
$ go build -gcflags="-m=1" sum.go
# command-line-arguments
/usr/local/go1.23.0/src/sync/atomic/type.go:63:6: can inline atomic.(*Pointer[go.shape.struct { os.mu sync.Mutex; os.buf *[]uint8; os.nbuf int; os.bufp int }]).CompareAndSwap
/usr/local/go1.23.0/src/sync/atomic/type.go:60:6: can inline atomic.(*Pointer[go.shape.struct { os.mu sync.Mutex; os.buf *[]uint8; os.nbuf int; os.bufp int }]).Swap
/usr/local/go1.23.0/src/sync/atomic/type.go:57:6: can inline atomic.(*Pointer[go.shape.struct { os.mu sync.Mutex; os.buf *[]uint8; os.nbuf int; os.bufp int }]).Store
/usr/local/go1.23.0/src/sync/atomic/type.go:54:6: can inline atomic.(*Pointer[go.shape.struct { os.mu sync.Mutex; os.buf *[]uint8; os.nbuf int; os.bufp int }]).Load
/usr/local/go1.23.0/src/sync/atomic/type.go:63:6: can inline atomic.(*Pointer[os.dirInfo]).CompareAndSwap
/usr/local/go1.23.0/src/sync/atomic/type.go:60:6: can inline atomic.(*Pointer[os.dirInfo]).Swap
/usr/local/go1.23.0/src/sync/atomic/type.go:57:6: can inline atomic.(*Pointer[os.dirInfo]).Store
/usr/local/go1.23.0/src/sync/atomic/type.go:54:6: can inline atomic.(*Pointer[os.dirInfo]).Load
./sum.go:15:34: inlining call to bytes.Split
./sum.go:31:13: inlining call to fmt.Println
/usr/local/go1.23.0/src/sync/atomic/type.go:63:6: inlining call to atomic.(*Pointer[go.shape.struct { os.mu sync.Mutex; os.buf *[]uint8; os.nbuf int; os.bufp int }]).CompareAndSwap
/usr/local/go1.23.0/src/sync/atomic/type.go:60:6: inlining call to atomic.(*Pointer[go.shape.struct { os.mu sync.Mutex; os.buf *[]uint8; os.nbuf int; os.bufp int }]).Swap
/usr/local/go1.23.0/src/sync/atomic/type.go:57:6: inlining call to atomic.(*Pointer[go.shape.struct { os.mu sync.Mutex; os.buf *[]uint8; os.nbuf int; os.bufp int }]).Store
/usr/local/go1.23.0/src/sync/atomic/type.go:54:6: inlining call to atomic.(*Pointer[go.shape.struct { os.mu sync.Mutex; os.buf *[]uint8; os.nbuf int; os.bufp int }]).Load
./sum.go:10:10: leaking param: fileName
./sum.go:15:45: ([]byte)("¥n") does not escape
./sum.go:16:39: string(line) does not escape
./sum.go:28:15: ... argument does not escape
./sum.go:31:13: ... argument does not escape
./sum.go:31:14: n escapes to heap
```

いやぁ，この機能は知らんかった。
めっさお役立ちやん。
この情報を考慮してアセンブラコードを見れば実際にどういう風にコンパイルされているか分かりやすくなる。

## バイナリコードを逆アセンブルする

『[効率的な Go]』ではバイナリの逆アセンブルに `objdump` コマンドを使用しているが，これって多分 Windows に（標準では）ないコマンドだよね。
あと，今どきの Linux には GCC 等の開発環境が入ってなかったりする。
たとえば，以下のように開発環境を入れれば `objdump` コマンドもインストールされる。

```text
$ sudo apt install build-essential
```

逆アセンブルする場合は `objdump` コマンドを生で使うのではなく `go tool` コマンドから呼び出すようにするといいようだ。

```text
$ go tool objdump -S sum
```

これで逆アセンブルした結果が標準出力に表示される。
こんな感じ。

```text
TEXT main.Sum(SB) /path/to/sum.go
func Sum(fileName string) (ret int64, _ error) {
  0x49e000        4c8d6424f8        LEAQ -0x8(SP), R12    
  0x49e005        4d3b6610        CMPQ R12, 0x10(R14)    
  0x49e009        0f86c0000000        JBE 0x49e0cf        
  0x49e00f        55            PUSHQ BP        
  0x49e010        4889e5            MOVQ SP, BP        
  0x49e013        4883c480        ADDQ $-0x80, SP        
  0x49e017        4889842490000000    MOVQ AX, 0x90(SP)    
    b, err := os.ReadFile(fileName)
  0x49e01f        90            NOPL            
  0x49e020        e87b74ffff        CALL os.ReadFile(SB)    
    if err != nil {
  0x49e025        4885ff            TESTQ DI, DI        
  0x49e028        752c            JNE 0x49e056        
    for _, line := range bytes.Split(b, []byte("¥n")) {
  0x49e02a        66c7442465c2a5        MOVW $-0x5a3e, 0x65(SP)    
  0x49e031        c64424676e        MOVB $0x6e, 0x67(SP)    
func Split(s, sep []byte) [][]byte { return genSplit(s, sep, 0, -1) }
  0x49e036        488d7c2465        LEAQ 0x65(SP), DI    
  0x49e03b        be03000000        MOVL $0x3, SI        
  0x49e040        4989f0            MOVQ SI, R8        
  0x49e043        4531c9            XORL R9, R9        
  0x49e046        49c7c2ffffffff        MOVQ $-0x1, R10        
  0x49e04d        e84e54fdff        CALL bytes.genSplit(SB)    
    for _, line := range bytes.Split(b, []byte("¥n")) {
  0x49e052        31c9            XORL CX, CX        
  0x49e054        eb2b            JMP 0x49e081        
        return 0, err
  0x49e056        31c0            XORL AX, AX        
  0x49e058        4889fb            MOVQ DI, BX        
  0x49e05b        4889f1            MOVQ SI, CX        
  0x49e05e        4883ec80        SUBQ $-0x80, SP        
  0x49e062        5d            POPQ BP            
  0x49e063        c3            RET            
    for _, line := range bytes.Split(b, []byte("¥n")) {
  0x49e064        488b542478        MOVQ 0x78(SP), DX    
  0x49e069        4883c218        ADDQ $0x18, DX        
  0x49e06d        488b5c2470        MOVQ 0x70(SP), BX    
  0x49e072        48ffcb            DECQ BX            
        ret += num
  0x49e075        488b742468        MOVQ 0x68(SP), SI    
  0x49e07a        488d0c06        LEAQ 0(SI)(AX*1), CX    
    for _, line := range bytes.Split(b, []byte("¥n")) {
  0x49e07e        4889d0            MOVQ DX, AX        
  0x49e081        4885db            TESTQ BX, BX        
  0x49e084        7e3c            JLE 0x49e0c2        
  0x49e086        48895c2470        MOVQ BX, 0x70(SP)    
        ret += num
  0x49e08b        48894c2468        MOVQ CX, 0x68(SP)    
    for _, line := range bytes.Split(b, []byte("¥n")) {
  0x49e090        4889442478        MOVQ AX, 0x78(SP)    
  0x49e095        488b18            MOVQ 0(AX), BX        
  0x49e098        488b4808        MOVQ 0x8(AX), CX    
        num, err := strconv.ParseInt(string(line), 10, 64)
  0x49e09c        488d442445        LEAQ 0x45(SP), AX            
  0x49e0a1        e8fa93fcff        CALL runtime.slicebytetostring(SB)    
  0x49e0a6        b90a000000        MOVL $0xa, CX                
  0x49e0ab        bf40000000        MOVL $0x40, DI                
  0x49e0b0        e8ab6dfdff        CALL strconv.ParseInt(SB)        
        if err != nil {
  0x49e0b5        4885db            TESTQ BX, BX        
  0x49e0b8        74aa            JE 0x49e064        
            return 0, err
  0x49e0ba        31c0            XORL AX, AX        
  0x49e0bc        4883ec80        SUBQ $-0x80, SP        
  0x49e0c0        5d            POPQ BP            
  0x49e0c1        c3            RET            
    return ret, nil
  0x49e0c2        4889c8            MOVQ CX, AX        
  0x49e0c5        31db            XORL BX, BX        
  0x49e0c7        31c9            XORL CX, CX        
  0x49e0c9        4883ec80        SUBQ $-0x80, SP        
  0x49e0cd        5d            POPQ BP            
  0x49e0ce        c3            RET            
func Sum(fileName string) (ret int64, _ error) {
  0x49e0cf        4889442408        MOVQ AX, 0x8(SP)            
  0x49e0d4        48895c2410        MOVQ BX, 0x10(SP)            
  0x49e0d9        e882c3fcff        CALL runtime.morestack_noctxt.abi0(SB)    
  0x49e0de        488b442408        MOVQ 0x8(SP), AX            
  0x49e0e3        488b5c2410        MOVQ 0x10(SP), BX            
  0x49e0e8        e913ffffff        JMP main.Sum(SB)            
```

ここでは `Sum()` 関数部分のみを抜粋している。
たとえば，前節で `bytes.Split()` 関数がインライ化されていることが示されているが，実際にどのようなコードになっているかよく分かる。

それにしても『[効率的な Go]』とはだいぶ違うなぁ。

[Go]: https://go.dev/
[効率的な Go]: https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "効率的なGo ―データ指向によるGoアプリケーションの性能最適化 | Bartłomiej Płotka, 山口 能迪 |本 | 通販 | Amazon"

## 参考

{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4863544006" %}} <!-- 基礎から学ぶ TinyGo の組込み開発 -->
