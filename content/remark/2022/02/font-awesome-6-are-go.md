+++
title = "Font Awesome 6 を導入した"
date =  "2022-02-23T14:44:22+09:00"
description = "Go 言語のアイコンもあった。"
image = "/images/attention/kitten.jpg"
tags = [ "site", "font", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 [Font Awesome] 6 がリリースされた。

- [Font Awesome 6 Is Live! See the New Icons, Styles and Upgrade Without the Hassle](https://blog.fontawesome.com/font-awesome-6-2/)

というわけで，このブログにも導入した。
まぁ，見た目はほぼ変わらないんだけどね。

[Font Awesome] 6 の目玉は [Thin Style](https://fontawesome.com/search?s=thin) が登場したことだろうか。
これはなかなか面白いと思う。

なんとなく思いついて検索してみたら [Go] 言語のアイコンもあった。

{{< fig-img-quote src="./font-awesome-golang.png" title="Go - brands | Font Awesome" link="https://fontawesome.com/icons/golang?s=brands" width="1002" lang="en" >}}

たとえば，こんな感じに使える。

| <span style="font-variant:normal;font-size:smaller;">`<span class="fa-4x" style="color:#00aed9;"><i class="fa-brands fa-golang"></i></span>`</span> |
| :-------------------------------------------------------------------------------------: |
|  <span class="fa-4x" style="color:#00aed9;"><i class="fa-brands fa-golang"></i></span>  |

HTML + Font Awesome でスライドを書いてる人は，はかどるかもしれない（笑） ちなみに Rust とかのアイコンは以前からあったんだよねぇ。

> {{< span class="fa-2x" >}}{{< icons "css3" "erlang" "html5" "java" "nodejs" "php" "python" "rust" "swift" >}}{{< /span >}}

## ブックマーク

- [Go's New Brand - The Go Programming Language](https://go.dev/blog/go-brand) : [Go] ロゴの取り扱いについてはこちらを参照のこと

[Font Awesome]: https://fontawesome.com/
[Go]: https://go.dev/
