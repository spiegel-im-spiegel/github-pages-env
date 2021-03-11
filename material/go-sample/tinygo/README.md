# [TinyGo] で遊ぶ

- [tinygo-org/tinygo: Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM.](https://github.com/tinygo-org/tinygo)

2021-0310 時点で [0.17 がリリース](https://github.com/tinygo-org/tinygo/releases/tag/v0.17.0 "Release 0.17.0 · tinygo-org/tinygo")されている。

- tinygo0.17.0.darwin-amd64.tar.gz
- tinygo0.17.0.linux-amd64.tar.gz
- tinygo0.17.0.linux-arm.tar.gz
- tinygo0.17.0.windows-amd64.zip
- tinygo_0.17.0_amd64.deb
- tinygo_0.17.0_arm.deb

## インストール

[TinyGo] のインストールの前に，あらかじめ [Go] コンパイラがインストールされていること。
[Go] 1.16 以降がお勧めらしい。

### Ubuntu の場合

APT や Snap のリポジトリにはないので deb ファイルを取ってきて直接インストールする。

```
$ wget https://github.com/tinygo-org/tinygo/releases/download/v0.17.0/tinygo_0.17.0_amd64.deb
$ sudo dpkg -i tinygo_0.17.0_amd64.deb
```

組み込み用途だと他にもいくつかツールが必要なようだが今回はなし。

### Windows の場合

なんと Scoop でインストール可能。

```
$ scoop install tinygo
Installing 'tinygo' (0.17.0) [64bit]
...
Linking ~\scoop\apps\tinygo\current => ~\scoop\apps\tinygo\0.17.0
Creating shim for 'tinygo'.
'tinygo' (0.17.0) was installed successfully!
```

実行モジュールは tinygo.exe のみのようだ。




```
$ go build -o out1.exe -trimpath
$ GOOS=js GOARCH=wasm go build -o hello1.wasm -trimpath
$ tinygo build -o hello2.wasm -target wasm ./hello.go
```

```
$ ll
-rw-a--   80 Mar 10 13:21:50 hello.go
-rwxa-- 2.1M Mar 10 13:26:01 out1.exe*
-rw-a-- 2.2M Mar 10 13:26:25 out2.wasm
-rw-a-- 218K Mar 10 13:44:03 out3.wasm
```

凄ぇ，ホンマに10分の1だよ。



## ブックマーク

- [TinyGo - マイコンやWebAssemblyが作れる軽量なGo言語 MOONGIFT](https://www.moongift.jp/2019/09/tinygo-%E3%83%9E%E3%82%A4%E3%82%B3%E3%83%B3%E3%82%84webassembly%E3%81%8C%E4%BD%9C%E3%82%8C%E3%82%8B%E8%BB%BD%E9%87%8F%E3%81%AAgo%E8%A8%80%E8%AA%9E/)
- [TinyGoで始める組み込みプログラミング - 144Labグループ開発者ブログ](https://tech.144lab.com/entry/tinygo)
- [Go+WebAssemblyで時刻表示を実装してみた](https://zenn.dev/komisan19/articles/aec74dd8223095)

[TinyGo]: https://tinygo.org/ "TinyGo - Go on Microcontrollers and WASM"
[Go]: https://golang.org/ "The Go Programming Language"
