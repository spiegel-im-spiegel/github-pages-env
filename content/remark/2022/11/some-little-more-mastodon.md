+++
title = "もうちょこっと Mastodon"
date =  "2022-11-20T14:45:05+09:00"
description = "アイコン / サイト連携 / 公式アプリ / シェア / Go パッケージ"
image = "/images/attention/kitten.jpg"
tags = [ "mastodon", "tools", "font", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./the-return-of-mastodon.md" >}} "Mastodon の復活")に引き続き，もうちょこっとだけ Mastodon の話。

## Font Awesome

Font Awesome に Mastodon のアイコンがないかなぁと思ってたら，ありました。

- [Mastodon Icon | Font Awesome](http://fontawesome.com/icons/mastodon)

たとえば，こんな感じに使える。

| <span style="font-variant:normal;font-size:smaller;">`<span class="fa-4x" style="color:#3088d4;"><i class="fa-brands fa-mastodon"></i></span>`</span> |
| :-------------------------------------------------------------------------------------: |
|  <span class="fa-4x" style="color:#3088d4;"><i class="fa-brands fa-mastodon"></i></span>  |

## サイト連携

プロフィールの補足情報として各種サイトの URL を載せることができるのだが

{{< fig-img src="./mastodon-profile.png" title="プロフィール補足情報" link="./mastodon-profile.png" width="828" >}}

上の図の右側にも書かれている通り，指定 URL の Web ページで

```html
<a rel="me" href="https://hostname/@username">Mastodon</a>
```

のように Mastodon プロフィール・ページを `<a>` 要素で指示するか，あるいは `<head>` 要素内に

```html
<link rel="me" href="https://hostname/@username">
```

などと `<link>` 要素を設定することで Mastodon と（緩い）連携ができる。
ポイントは `rel="me"` の属性を必ず付けること。
上手くいけばプロフィール・ページで

{{< fig-img src="./mastodon-profile-2.png" title="プロフィール補足情報" link="./mastodon-profile-2.png" >}}

のように認証された URL にチェックマークが付くようだ。
先に連携先のサイトで `rel="me"` 属性付きのリンクを設定し，そのあとでプロフィール補足情報を設定すると上手くいくっぽい。

## 公式アプリ

他の方の toot を見て初めて知ったのだが，スマホ用の公式アプリがあるらしい。

- [Get an app for Mastodon - Mastodon](https://joinmastodon.org/apps)

試しにインストールして使ってみたが意図的に機能を簡略化してるように見える。
まぁ，ローカルとか連合（federation）とか follow してもないユーザの toots とか要らんよね。

というわけで，しばらくこれで遊んでみることにしよう。

## Share ボタンを付けてみたかったのだが...

Web ページに Mastodon への share ボタンがあれば面白いと一瞬思ったのだが...

たとえばブラウザで

```text
https://mstdn.jp/share?text=hello+world
```

という URL を叩けば

{{< fig-img src="./mastodon-share.png" link="./mastodon-share.png" >}}

という画面が開く。
ただしこれは [mstdn.jp](https://mstdn.jp/) にアカウントを持っていてサインインしている状態であれば，の話。
分散型の Mastodon ではユーザごとに所属しているサーバは異なるんだから，上のような固定の URL をボタンに紐付けても意味がない。
まぁ，ブラウザに自分専用のブックマークレットでも設置すればいけるのかもしれないが。

あと `web+mastodon` scheme ってのがあるそうで，この scheme による連携が Mastodon アプリとブラウザとの間でできていれば

```text
web+mastodon://share?text=hello+world
```

などと指定することでいけるらしい。
でも，これも一般的じゃないよなぁ...

Android 機のブラウザであればシェア {{% icons "share-nodes" %}} 機能で任意のアプリと連携できるし， Web ページ側でボタンを設置するのは諦めるかな。

## Go による Mastodon クライアント・パッケージ

そうそう。
Go のパッケージに

- [mattn/go-mastodon: mastodon client for golang](https://github.com/mattn/go-mastodon)

というのがあるのだが，今もこまめにメンテナンスされてるっぽい。
Go で Mastodon クライアント機能を組み込みたいのであれば参考になるかもしれない。

## ブックマーク

- [Mastodon向け簡易シェアボタン - Qiita](https://qiita.com/mod_poppo/items/d80ff225b4cc93318ee8)
