+++
title = "Android 版 Firefox の検索窓に DuckDuckGo を指定する"
date =  "2020-08-19T19:32:12+09:00"
description = "Mozilla がユーザのプライバシーを重視しているというのは嘘っぱちである。"
image = "/images/attention/kitten.jpg"
tags = [ "privacy", "risk", "firefox", "search", "web" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

他の方はどうか知らないが，何故か私のスマホの Firefox は古い ESR 版（68 相当）のまま一向にバージョンアップする気配がなかったのだが，やっとバージョンが上がったようだ。
これでよーやくダーク・テーマが使えるぜ！

で，例によって設定が初期化されたので（怒），検索窓を [DuckDuckGo] にし直そうと思ったら選択できないぢゃん。
しょうがないから手入力で登録したよ。
前はもっと簡単にできたような気がするのだが...

設定の「検索エンジンの追加」で「その他」を選び「検索クエリ―文字列」に以下を入力する。

```text
https://duckduckgo.com/?q=%s
```

検索エンジン名は適当に（笑）

パソコン用に比べてセキュリティ・プライバシー関連の設定項目が少ないのは意図的なのかねぇ。
ホンマ何とかしてくれよ。
これで “[Changing Mozilla]” とか笑かす。

まぁ，私のスマホ既定ブラウザは [Firefox Focus](https://play.google.com/store/apps/details?id=org.mozilla.focus) なので致命的ではないのが救いだが， “[Changing Mozilla]” でコレまでリストラされないことを祈るよ。

## ブックマーク

- [Laying the foundation for Rust's future | Rust Blog](https://blog.rust-lang.org/2020/08/18/laying-the-foundation-for-rusts-future.html) : Rust は更に Mozilla への依存度を下げるのかな？
    - [「Rust Foundation」が年内にも設立。Rust言語のコアチームとMozillaが発表 － Publickey](https://www.publickey1.jp/blog/20/rust_foundationrustmozilla.html)
- [Android版「Firefox」が新設計のニューバージョンにスイッチ中！　旧版と何が違う？ - やじうまの杜 - 窓の杜](https://forest.watch.impress.co.jp/docs/serial/yajiuma/1271727.html) : 私の環境だけ古いバージョンというわけじゃなかったらしい（笑）
- [Firefox の DoH は無効にすべきか（もしくは水売りと水道局）]({{< ref "/remark/2019/09/should-disable-doh-in-firefox.md" >}}) : Mozilla がユーザのプライバシーを重視しているというのは嘘っぱちである

[DuckDuckGo]: https://duckduckgo.com/
[Changing Mozilla]: https://blog.mozilla.org/blog/2020/08/11/changing-world-changing-mozilla/ "Changing World, Changing Mozilla - The Mozilla Blog"
