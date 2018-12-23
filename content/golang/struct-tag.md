+++
date = "2016-02-05T00:11:53+09:00"
update = "2017-09-25T13:08:33+09:00"
description = "Struct で正規化できる情報であれば，タグを使うことでアプリケーション外部とのやり取りがずっと楽になる。"
draft = false
tags = ["golang", "struct", "tags", "marshal"]
title = "Struct タグについて"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

たとえば [struct] で構造化されている情報を特定のファイルやデータベースに出力したり，逆にファイルやデータベースの情報を [struct] に流し込みたい場合に [struct] の各フィールドに目印になる情報があると便利である。
この目印として機能するのが [struct] タグである[^an]。

[^an]: 「アノテーション（annotation）」と呼ぶ人もいる。たぶん Java の annotation 機能を意識しているんだろう。

 [struct] タグは以下のように記述する。

{{< fig-quote title="reflect - The Go Programming Language" link="https://golang.org/pkg/reflect/#example_StructTag" lang="en" >}}
<q>By convention, tag strings are a concatenation of optionally space-separated key:"value" pairs. Each key is a non-empty string consisting of non-control characters other than space (U+0020 ' '), quote (U+0022 '"'), and colon (U+003A ':'). Each value is quoted using U+0022 '"' characters and Go string literal syntax.</q>
{{< /fig-quote >}}

```go
type Server struct {
    Host      string `elem:"host"`
    IPAddress string `elem:"ip_address"`
    Port      int    `elem:"port"`
    Note      string `elem:"note"`
}
```

このタグ情報を取得するには [`reflect`] パッケージを使う。
たとえばこんな感じ。

```go
package main

import (
    "fmt"
    "reflect"
)

type Server struct {
    Host      string `elem:"host"`
    IPAddress string `elem:"ip_address"`
    Port      int    `elem:"port"`
    Note      string `elem:"note"`
}

func main() {
    s := Server{}
    t := reflect.TypeOf(s)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Name=%s , tag(elem)=%s\n", field.Name, field.Tag.Get("elem"))
    }
}
```

これを実行するとこうなる。

```
Name=Host , tag(elem)=host
Name=IPAddress , tag(elem)=ip_address
Name=Port , tag(elem)=port
Name=Note , tag(elem)=note
```

実際には [`reflect`] を直接使う局面は少なく，既にあるパッケージを利用することが多い。
たとえば [struct] による構造化データを [JSON] 形式に出力する [`encoding/json`] パッケージがある。

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    Host      string `json:"host"`
    IPAddress string `json:"ip_address"`
    Port      int    `json:"port"`
    Note      string `json:"note"`
}

func main() {
    s := Server{Host: "localhost", IPAddress: "127.0.0.1", Port: 8080, Note: "Web Application"}
    j, err := json.MarshalIndent(s, "", "  ")
    if err != nil {
        return
    }
    fmt.Println(string(j))
}
```

これを実行するとこうなる。

```
{
  "host": "localhost",
  "ip_address": "127.0.0.1",
  "port": 8080,
  "note": "Web Application"
}
```

`Server` の内容が [JSON] 形式で出力されているのが分かるだろう。
[JSON] の要素名がタグで指定した名前になっていることを確認してほしい。

反対もやってみよう。

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    Host      string `json:"host"`
    IPAddress string `json:"ip_address"`
    Port      int    `json:"port"`
    Note      string `json:"note"`
}

func main() {
    svr := []byte(`{
  "host": "localhost",
  "ip_address": "127.0.0.1",
  "port": 8080,
  "note": "Web Application"
}`)
    var s Server
    if err := json.Unmarshal(svr, &s); err != nil {
        return
    }
    fmt.Println(s)
}
```

実行結果はこうなる。

```
{localhost 127.0.0.1 8080 Web Application}
```

きれいに [struct] に値が入っているのが分かると思う。

ちなみにタグの書式は `key:"value"` だが，間違って記述しても単に無視されるだけでコンパイル時も実行時もエラーにならないので注意が必要である。
なおタグ書式の文法ミスについては，静的検査ツールの [vet] でチェックできる。

タグは複数列挙することができる。
たとえばサンプルの構造体を [TOML] にも対応させたいなら

```go
type Server struct {
    Host      string `json:"host" toml:"host"`
    IPAddress string `json:"ip_address" toml:"ip_address"`
    Port      int    `json:"port" toml:"port"`
    Note      string `json:"note" toml:"note"`
}
```

などとする（デリミタは空白文字）。
じゃあ，先ほどと同じようにして [TOML] で出力してみる。
[TOML] を扱うには [`github.com/BurntSushi/toml`](https://github.com/BurntSushi/toml) パッケージを使うとよい。

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/BurntSushi/toml"
)

type Server struct {
    Host      string `json:"host" toml:"host"`
    IPAddress string `json:"ip_address" toml:"ip_address"`
    Port      int    `json:"port" toml:"port"`
    Note      string `json:"note" toml:"note,omitempty"`
}

func main() {
    s := Server{Host: "localhost", IPAddress: "127.0.0.1", Port: 8080, Note: ""}
    t := new(bytes.Buffer)
    if err := toml.NewEncoder(t).Encode(s); err != nil {
        return
    }
    fmt.Println(t.String())
}
```

実行結果は以下の通り。

```
host = "localhost"
ip_address = "127.0.0.1"
port = 8080
```

`omitempty` オプションはフィールドが空（`nil` または空文字列）の場合に出力を省略できる[^oz]。
このオプションは [`encoding/json`] パッケージでも使える。

[^oz]: 数値の場合は `omitzero` オプションを付けると 0 のときに出力を省略できる。ただし [`BurntSushi/toml`](https://github.com/BurntSushi/toml) パッケージでは [`Decode()` がうまく動かない](http://qiita.com/reiki4040/items/6556d4eba797329e9f51)らしい。実は `omitempty` オプションも `Decode()` 時の挙動が怪しいんだよなぁ。 [TOML] パーサの別実装としては [naoina/toml](https://github.com/naoina/toml) というのもある。これは最新の [TOML] 仕様に追随しているようだが `omitzero` オプションには対応していない。

ついでに反対もやってみよう。

```go
package main

import (
    "fmt"

    "github.com/BurntSushi/toml"
)

type Server struct {
    Host      string `json:"host" toml:"host"`
    IPAddress string `json:"ip_address" toml:"ip_address"`
    Port      int    `json:"port" toml:"port"`
    Note      string `json:"note" toml:"note,omitempty"`
}

func main() {
    svr := `
host = "localhost"
ip_address = "127.0.0.1"
port = 8080
note = "Web Application"
`
    var s Server
    if _, err := toml.Decode(svr, &s); err != nil {
        return
    }
    fmt.Println(s)
}
```

結果は以下の通り。

```
{localhost 127.0.0.1 8080 Web Application}
```

このように [struct] で正規化できる情報であれば，タグ機能を使うことでアプリケーション外部とのやり取りがだいぶ楽になる。

## ブックマーク

- [Go で struct のタグ情報を取得する - hiyosi's blog](http://hiyosi.tumblr.com/post/100922038678/go-%E3%81%A7-struct-%E3%81%AE%E3%82%BF%E3%82%B0%E6%83%85%E5%A0%B1%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
- [struct にアノテーションつけてたら go vet . すべき - Qiita](http://qiita.com/amanoiverse/items/fcd25db64f341ad2471f)
- [Goのjson.Marshal/Unmarshalの仕様を整理してみる - I Will Survive](http://blog.restartr.com/2014/08/13/golang-json-marshal-unmarshal/)
- [BurntSushi/tomlを使ってハマったこと - Qiita](http://qiita.com/reiki4040/items/6556d4eba797329e9f51)
- [GoでJSONの一部分を利用者が定義した構造体に読み込める便利な手法を見つけた - Qiita](http://qiita.com/hnakamur/items/ba363e82332d4dbdf34a)
- [Go 言語 1つの構造体に複数の validation を適応する - Qiita](http://qiita.com/iktakahiro/items/2e240147ca3188948a17) : [struct] タグに validation 情報を埋め込んで利用する
- [Goのencoding/xmlを使いこなす - Qiita](http://qiita.com/ono_matope/items/70080cc33b75152c5c2a)
- [Go言語でJSON内の整数は10進数6桁しか表現できない - Qiita](http://qiita.com/toast-uz/items/52f0c86716493ad3ca12)
- [GolangでEnumをフィールドに持つstructをいい感じにjsonエンコード / デコードする - 一から勉強させてください(￣ω￣;)](http://dangerous-animal141.hatenablog.com/entry/2017/01/19/004650)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[struct]: https://golang.org/ref/spec#Struct_types "Struct types"
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"
[`encoding/json`]: https://golang.org/pkg/encoding/json/ "json - The Go Programming Language"
[JSON]: https://tools.ietf.org/html/rfc7159 "RFC 7159 - The JavaScript Object Notation (JSON) Data Interchange Format"
[TOML]: https://github.com/toml-lang/toml "toml-lang/toml: Tom's Obvious, Minimal Language"
[vet]: https://golang.org/cmd/vet/ "vet - The Go Programming Language"
