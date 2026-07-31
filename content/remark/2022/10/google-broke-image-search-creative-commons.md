+++
title = "Google 画像検索の CC Licenses 絞り込み機能は壊れている？"
date =  "2022-10-02T12:26:46+09:00"
description = "ルークよ， Openverse を使え！"
image = "/images/attention/kitten.jpg"
tags = ["google", "creative-commons", "search", "image", "photography", "flickr"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

これも Twitter TL から拾ったネタ。

- [Google画像検索にある「クリエイティブ・コモンズ ライセンス」の絞り込み機能は壊れているとの指摘 - GIGAZINE](https://gigazine.net/news/20220930-google-broke-image-search-creative-commons/)

というわけで，実際に試してみる。

まず Google の検索窓で "lock" をキーに画像検索してみる。
このときツール・オプションで「クリエイティブ・コモンズライセンス」を選択する。

{{< fig-img src="./search-options.png" link="./search-options.png" width="865" >}}

結果はこんな感じ。

{{< fig-img src="./lock-images-by-google-1.png" link="./lock-images-by-google-1.png" width="1371" >}}

いやいやいや（笑）

このときの URL はこんな感じ（余計なオプションは省いている）

```text
https://www.google.com/search?q=lock&tbm=isch&tbs=il:cl&hl=ja
```

言語パラメータの `hl=ja` を省くと検索結果が変わるのがアレなのだがそれはともかく，どう見ても CC licenses とは関係ないし結果の件数が少なすぎるよね。
ここで件の記事に従って `tbs` オプションの値を `sur:fmc,il:cl` に書き直してみる。

{{< fig-img src="./lock-images-by-google-2.png" link="./lock-images-by-google-2.png" width="1368" >}}

おー，出るやないかい！ 画像左下の「ライセンス可能」のアイコンをクリックすると

{{< fig-img src="./lock-images-by-google-3.png" link="./lock-images-by-google-3.png" width="1368" >}}

Wikimedia Commons の画像であることが分かる（ちなみにこの錠前の画像は [CC0 {{< cc-syms "cc" "zero" >}}](https://creativecommons.org/publicdomain/zero/1.0/deed "Creative Commons — CC0 1.0 Universal") 献呈された画像である）。

なお，この状態で検索ワードを変えるとツール・オプションがなかったことにされる。

{{< div-gen type="markdown" class="center" >}}
**使えん！ `orz`**
{{< /div-gen >}}

悪意なのか？ それとも馬鹿なのか？

気を取り直して... 最初に挙げた記事には「[Openverse] を使え」とある。
早速試してみよう。

{{< fig-img src="./openverse.png" link="./openverse.png" width="833" >}}

結果はご覧の通り。

{{< fig-img src="./lock-images-by-openverse-1.png" link="./lock-images-by-openverse-1.png" width="1344" >}}

レスリングの画像ががが。
確かに「アームロック」の画像だけど（笑）

画像部分をクリックすると

{{< fig-img src="./lock-images-by-openverse-2.png" link="./lock-images-by-openverse-2.png" width="950" >}}

という感じに詳細情報が出る。
クレジット表記の雛形も表示してくれるんだねぇ。
ちなみに日本語の検索ワードもちゃんと通る。

{{< fig-img src="./lock-images-by-openverse-3.png" link="./lock-images-by-openverse-3.png" width="1361" >}}

ところで [Openverse] って何だろうと思ったが，どうも2017年に登場した [CC Search](https://ccsearch.creativecommons.org/) の後継もしくは派生サービスのようだ。
`https://ccsearch.creativecommons.org/` から [Openverse] にリダイレクトされるし。
近年は Creative Commons の話題を真面目に追わなくなったせいか，全く気が付かなかった。

ま，まぁ，ともかく， Google 画像検索は捨て！ ということで（笑）

## Virtual Photography

話は変わるが [Flickr] に画像をアップロードする際に “Photos” とか “Screenshots” とか “Illustration” とかいった種類を指定することができるのだが，これに新たに “Virtual Photography / Machinima” が加わった。

- [Discover Virtual Photography on Flickr | Flickr Blog](https://blog.flickr.net/en/2022/09/13/discover-virtual-photography-on-flickr/)

いわゆる「AI 絵画」の流行を受けてのものだろう。
まぁ，どの種類を指定するかはユーザの胸三寸なんだけどね。

もし AI 絵画を「自由なライセンス」で公開したいのであれば [Flickr] の利用も検討してください，ということで。

## ブックマーク

- [新しい CC Search が登場，他]({{< ref "/remark/2017/02/new-cc-search.md" >}})
- [Flickr は監視資本主義に向かわない]({{< ref "/remark/2019/03/flickr-has-not-turned-to-surveillance-capitalism.md" >}})

[Openverse]: https://wordpress.org/openverse/ "Openly Licensed Images, Audio and More | Openverse"
[Flickr]: https://www.flickr.com/

## 参考図書

{{% review-paapi "B099RTG3J7" %}} <!-- 著作権は文化を発展させるのか: 人権と文化コモンズ -->
