+++
title = "Nostr に旗を立てろ！"
date =  "2023-02-05T23:54:41+09:00"
description = "今回はアカウントを作っただけ"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter を追い出された（言い方！）なんとかドーシーさんが出資したとかいう Nostr (Notes and Other Stuff Transmitted by Relays) が流行りらしい。

- [GitHub - nostr-protocol/nostr: a truly censorship-resistant alternative to Twitter that has a chance of working](https://github.com/nostr-protocol/nostr)

Nostr 自身は，分散 SNS においてメッセージのリレーを行うためのプロトコルだそうで，このプロトコルを実装するためのリレー・サービスとクライアント・アプリケーションが必要らしい。

んで，そのリレー・サービスとクライアント・アプリケーションの一覧が以下にある。

- [Nostr-Clients-Features-List/Readme.md at main · vishalxl/Nostr-Clients-Features-List · GitHub](https://github.com/vishalxl/Nostr-Clients-Features-List/blob/main/Readme.md)

最近，私の周辺で Nostr が話題になったのは [Damus](https://damus.io/) が iOS アプリとしてクライアントを提供したことがきっかけのようだ。
「いや iOS とか関係ないし」とか思ってたが， Android 用のクライアントとして Amethyst ってのがあるそうな。

- [GitHub - vitorpamplona/amethyst: Nostr client for Android](https://github.com/vitorpamplona/amethyst)
- [Amethyst - Apps on Google Play](https://play.google.com/store/apps/details?id=com.vitorpamplona.amethyst)

ほんじゃあ，まぁ，試してみるべ。

## サインアップ？

{{< fig-img src="./amethyst.jpg" title="Amethyst" link="./amethyst.jpg" lang="en" width="720" >}}

んー。
分からん（笑） どうやってサインアップすんだ？

とりあえず「鍵」をつくるらしい。
“I sccept the term of use” にチェックを入れて “Generate a new key” のリンクをタップする。

{{< fig-img src="./signup-amethyst.jpg" title="sign up ?" link="./signup-amethyst.jpg" lang="en" width="720" >}}

いやいやいや。
何だこれ？ “empty” ってなんやねん？

ここでパニクった私はうっかりサインアウトしてしまったのよ。あとのまつり... `orz`

“Generate a new key” で公開鍵と秘密鍵の鍵ペアを自動で作るのだが， Amethyst は作成した鍵でそのままサインインして上の画面になってしまう。
んで，サインインした状態で（少なくとも）秘密鍵をメモっとかないと二度とサインインできない。
まじすか `orz`

ここでようやく “[Nostr Social Network](https://github.com/vishalxl/nostr_console/discussions/31 "What is Nostr, and how to start using Nostr · Discussion #31 · vishalxl/nostr_console")” を読み始める（先に読め！）。
ふむむ。
[astral.ninja](https://astral.ninja/) というのがあるのか。
こちらは Web ブラウザ・ベースのクライアントらしい。
早速アクセスしてみる。

{{< fig-img src="./astral-1.png" title="astral" link="https://astral.ninja/" lang="en" width="1442" >}}

ここで “GENERATE KEYS” ボタンを押す。
今度は間違えないぞ！

{{< fig-img src="./astral-2.png" title="astral" link="https://astral.ninja/" lang="en" width="1442" >}}

{{< div-gen class="center" >}}
（一応，鍵情報に目隠ししてます。あしからず）
{{< /div-gen >}}

んー。
これはこのまま “PROCEED” ボタンを押せばいいのか？ ポチッとな。

{{< fig-img src="./astral-3.png" title="astral - settings" link="https://astral.ninja/" lang="en" width="1442" >}}

{{< div-gen class="center" >}}
（一応，鍵情報に目隠ししてます。あしからず）
{{< /div-gen >}}

よしよし。
鍵情報が出てきたな。
`nsec` で始まる文字列が秘密鍵， `npub` で始まる鍵が秘密鍵である。
ちゃんとメモっておこう。

再度サインインする場合は秘密鍵を使う。
あとは設定画面の profile で Name と Picture URL をセットすれば表示名とアバターが表示される。
他はとりあえず後回しでよい。
なお `nsec` の秘密鍵を使えば他のクライアントでもサインインできるっぽい。
もちろん Amethyst でも使える。

## 秘密鍵は秘密のままに

当然ではあるが秘密鍵は絶対に公開してはいけない。
秘密鍵を公開するのは，その辺の Web サービスでユーザ ID とパスワードを公開するのと同じである。

逆に公開鍵の方は他ユーザにフォローしてもらえるよう公開しても OK。
ちなみに私の公開鍵は

```text
npub1v0yd6fxc0qczjj35rxu0m9gtkhhf4mkzl99tgs49f6e98y2ea6asl5rsxw
```

である。

ユーザ名には任意の名前を付けられる。
他のユーザとかぶっても構わない。
どうやら Nostr は鍵情報だけでユーザを識別しているようだ。
なので，いくらでもユーザを作ることができるし，なりすましだって可能である。
ちなみに某マスク氏の名前とアバターアイコンを使った（多分）偽ユーザとかもあるらしいっスよ（笑）

というわけで，作ったアカウントが自身のものである証明をする必要がある。
Nostr では Web サイトとの連携でこれを行うようだ。

## Web サイト連携によるユーザの証明

Nostr と Web サイトを連携するには，まず Web サイト側のルート直下に `/.well-known/nostr.json` というファイルを置く。
中身はこんな感じ。

```json
{
  "names": {
    "spiegel": "63c8dd24d87830294a3419b8fd950bb5ee9aeec2f94ab442a54eb2539159eebb"
  }
}
```

JSON 形式で `"names"` オブジェクト配下に `"<username>": "<pubkey>"` で記述されたオブジェクトを配置する。
`username` には任意の名前を記述する（多分 Nostr クライアントのユーザ名と揃えたほうがいいと思われ）。
`pubkey` には `npub` で始まる文字列ではなく公開鍵の生値を16進数文字列で記述する。
公開鍵データの変換は [damus key converter](https://damus.io/key/) で可能である。

あとはこれを Nostr 側に読み取らせればいいのだが，ここで [CORS (Cross-Origin Resource Sharing)](https://developer.mozilla.org/ja/docs/Web/HTTP/CORS) の設定が必要になる（つまり CORS の設定ができないサイトでは連携ができない）。

たとえば[さくらのレンタルサーバ](https://rs.sakura.ad.jp/ "さくらのレンタルサーバ | 高速・安定WordPressなら！無料2週間お試し")であれば ルート直下の `/.htaccess` ファイルに

```.htaccess
Header set Access-Control-Allow-Origin: "*"
```

の記述が必要（こんなドンブリ指定はちょっとヤバいかもなので，あとで修正するかも）。
設定ができたら，ちゃんと取得できるか確認してみよう。

```text
$ curl -sD - "https://baldanders.info/.well-known/nostr.json"
HTTP/2 200 
server: nginx
date: Sun, 05 Feb 2023 13:54:56 GMT
content-type: application/json
content-length: 103
last-modified: Sun, 05 Feb 2023 13:54:43 GMT
etag: "67-5f3f441b5a739"
accept-ranges: bytes
access-control-allow-origin: *
cache-control: no-store

{
  "names": {
    "spiegel": "63c8dd24d87830294a3419b8fd950bb5ee9aeec2f94ab442a54eb2539159eebb"
  }
}
```

よーし，うむうむ，よーし。
`curl` コマンドなんて知らねって方は [test-cors.org](https://www.test-cors.org/) でも確認できる。

Nostr クライアント側は profile 設定の “NIP-05 Identifier” 項目に `/.well-known/nostr.json` ファイルで指定した `username` とサイト名をセットする。
こんな感じ。

{{< fig-img src="./astral-4.png" title="astral - settings" link="https://astral.ninja/" lang="en" width="609" >}}

これで

{{< fig-img src="./astral-5.png" title="astral - profile" link="https://astral.ninja/" lang="en" width="609" >}}

という感じに連携できるようになった。

## Nostr に旗を立てろ！

これで最低限の設定ができた。
とはいえ，今のところあんまり使う気にならないんだよなぁ。
微妙に分かりにくいし使いにくい。
それに，やっぱ SNS は相手がいてはじめて成立しうるものだし？

まぁ，今回は “[Plant Your Flag, Mark Your Territory](https://krebsonsecurity.com/2018/06/plant-your-flag-mark-your-territory/ "Plant Your Flag, Mark Your Territory – Krebs on Security")” ってことで，旗を立てるだけにしておこう。

## ブックマーク

- [ツイッターライクなSNS「Nostr」を Damus アプリで使う。初期の設定と、認証バッジを付ける手順 | Lifehacking.jp](https://lifehacking.jp/2023/02/nostr-damus/)
- [Nostr. Nostrという新しい分散型のソーシャルプロトコルを試してみています。 | by Fumi | Feb, 2023 | Medium](https://fumi.medium.com/nostr-f8e6636b5724)
