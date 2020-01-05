+++
title = "Hash 値を計算するパッケージを作ってみた"
date =  "2017-10-31T14:31:04+09:00"
description = "Windows で hash 値を計算するいい感じのツールがないので，もう自分で作っちゃったよ。といっても自前の部分は殆どないけどね。"
tags = [ "golang", "programming", "cryptography", "hash" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

Windows で hash 値を計算するいい感じのツールがない[^std] ので，もう自分で作っちゃったよ。
といっても自前の部分は殆どないけどね。
    
[^std]: Windows 標準ツールとしては， PowerShell 4.0 以上が入っている PC なら， `Get-FileHash` コマンドレットが使える。 Windows 7 の場合は “[Windows Management Framework 4.0](https://www.microsoft.com/ja-jp/download/details.aspx?id=40855)” をインストールすることで PowerShell 4.0 にアップグレードできる。

- [spiegel-im-spiegel/hash: Calculating Hash Value](https://github.com/spiegel-im-spiegel/hash)

詳しくは [README.md](https://github.com/spiegel-im-spiegel/hash/blob/master/README.md "hash/README.md at master · spiegel-im-spiegel/hash") を見ていただくとして，実際に計算をするのはこの関数。

```go
package hash

import (
    "crypto"
    "io"
    "strings"

    "github.com/pkg/errors"
)

var (
    //ErrNoImplement is error "no implementation"
    ErrNoImplement = errors.New("no implementation")
)

//Value returns hash value string from io.Reader
func Value(r io.Reader, alg crypto.Hash) ([]byte, error) {
    if !alg.Available() {
        return nil, errors.Wrap(ErrNoImplement, "error "+algoString(alg))
    }
    h := alg.New()
    io.Copy(h, r)
    return h.Sum(nil), nil
}
```

呼び出し側はこんな感じ（空文字列の SHA1 値を取得する場合）。

```go
v, err := hash.Value(bytes.NewBuffer([]byte("")), crypto.SHA1)
if err != nil {
    return
}
fmt.Printf("%x\n", v)
// Output:
// da39a3ee5e6b4b0d3255bfef95601890afd80709
```

[`crypto`]`.Hash.New()` 関数で [`hash`]`.Hash` のインスタンスを生成している。
[`hash`]`.Hash` はこんな感じの [interface] 型である。

```go
// Hash is the common interface implemented by all hash functions.
type Hash interface {
    // Write (via the embedded io.Writer interface) adds more data to the running hash.
    // It never returns an error.
    io.Writer

    // Sum appends the current hash to b and returns the resulting slice.
    // It does not change the underlying hash state.
    Sum(b []byte) []byte

    // Reset resets the Hash to its initial state.
    Reset()

    // Size returns the number of bytes Sum will return.
    Size() int

    // BlockSize returns the hash's underlying block size.
    // The Write method must be able to accept any amount
    // of data, but it may operate more efficiently if all writes
    // are a multiple of the block size.
    BlockSize() int
}
```

つまり，このインタフェースを備えていれば自前の hash アルゴリズムを簡単に組み込むことができるわけだ。
言い方を変えると， `hash.Value()` 関数で実際に hash 値を計算するにはこのパッケージだけではダメで， hash アルゴリズムを実装するパッケージをインポートする必要がある。

以下にアルゴリズム毎に必要なパッケージを示す。

| hash algorithm | import package |
|:---------------|:---------------|
| `crypto.MD4`         | `golang.org/x/crypto/md4` |
| `crypto.MD5`         | `crypto/md5` |
| `crypto.SHA1`        | `crypto/sha1` |
| `crypto.SHA224`      | `crypto/sha256` |
| `crypto.SHA256`      | `crypto/sha256` |
| `crypto.SHA384`      | `crypto/sha512` |
| `crypto.SHA512`      | `crypto/sha512` |
| `crypto.SHA512_224`  | `crypto/sha512` |
| `crypto.SHA512_256`  | `crypto/sha512` |
| `crypto.RIPEMD160`   | `golang.org/x/crypto/ripemd160` |
| `crypto.SHA3_224`    | `golang.org/x/crypto/sha3` |
| `crypto.SHA3_256`    | `golang.org/x/crypto/sha3` |
| `crypto.SHA3_384`    | `golang.org/x/crypto/sha3` |
| `crypto.SHA3_512`    | `golang.org/x/crypto/sha3` |
| `crypto.BLAKE2s_256` | `golang.org/x/crypto/blake2s` |
| `crypto.BLAKE2b_256` | `golang.org/x/crypto/blake2b` |
| `crypto.BLAKE2b_384` | `golang.org/x/crypto/blake2b` |
| `crypto.BLAKE2b_512` | `golang.org/x/crypto/blake2b` |

この中から必要なパッケージを `main` パッケージでブランク・インポートする[^lint1]。
全部インポートするならこんな感じ。

[^lint1]: ブランク・インポートは `main` パッケージでしないと [golint] に怒られるのよ。まぁ言いたいことは分かるけど。

```go
package main

import (
    _ "crypto/md5"
    _ "crypto/sha1"
    _ "crypto/sha256"
    _ "crypto/sha512"

    _ "golang.org/x/crypto/blake2b"
    _ "golang.org/x/crypto/blake2s"
    _ "golang.org/x/crypto/md4"
    _ "golang.org/x/crypto/ripemd160"
    _ "golang.org/x/crypto/sha3"
)
```

これを CLI (Command-Line Interface) にしたのが以下に示す hash コマンドである。

```text
$ hash -h
Usage:
  hash [flags] [binary file]

Flags:
  -a, --algo string      hash algorithm (default "sha256")
  -c, --compare string   compare hash value
  -h, --help             help for hash
```

サイズがゼロの空ファイル `empty.txt` を作って試してみると

```text
$ hash -a sha1 empty.txt
da39a3ee5e6b4b0d3255bfef95601890afd80709
```

となる。
パイプにも対応してるので

```text
$ cat empty.txt | hash -a sha1
da39a3ee5e6b4b0d3255bfef95601890afd80709
```

とすることもできる。
さらに `-c` オプションで hash 値の計算結果をリファレンスの値と比較できる。

```text
$ hash -a sha1 empty.txt -c da39a3ee5e6b4b0d3255bfef95601890afd80709
matched
```

なお，アルゴリズムには `md4`, `md5`, `sha1`, `sha224`, `sha256`, `sha384`, `sha512`, `sha512/224`, `sha512/256`, `ripemd160`, `sha3-224`, `sha3-256`, `sha3-384`, `sha3-512`, `blake2s`, `blake2b/256`, `blake2b/384`, `blake2b/512` を指定できるようにした。
既定は `sha256`。

例えば，これで [dep] の実行モジュールの正当性確認が少し楽になる。
[dep] の[リリースページ](https://github.com/golang/dep/releases/latest)で Windows 用の実行モジュールと SHA256 値を記述したファイルをダウンロードする。

- `dep-windows-amd64`
- `dep-windows-amd64.sha256`

`dep-windows-amd64.sha256` の中身が

```text
034f8cf6c225fde51aa025376df12450832f111b39050a7ec451a9ec2ce2cb54  release/dep-windows-amd64
```

とするなら

```text
$ hash dep-windows-amd64 -c 034f8cf6c225fde51aa025376df12450832f111b39050a7ec451a9ec2ce2cb54
matched
```

で一発確認できる。
確認できたら `dep-windows-amd64` を `dep.exe` にリネームして使えばよい。

## ブックマーク

- [Glide から Dep への移行を検討する]({{< relref "consider-switching-from-glide-to-dep.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`crypto`]: https://golang.org/pkg/crypto/ "crypto - The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[`hash`]: https://golang.org/pkg/hash/ "hash - The Go Programming Language"

[interface]: https://golang.org/ref/spec#Interface_types "Interface types"
[golint]: https://github.com/golang/lint "golang/lint: This is a linter for Go source code."
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
