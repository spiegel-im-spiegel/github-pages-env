+++
title = "国勢調査の思ひ出"
date =  "2020-10-01T12:09:11+09:00"
description = "まだ回答してない人でネット回答可能であれば，ネットのほうが断然楽。"
image = "/images/attention/kitten.jpg"
tags = [ "politics", "web", "security", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

知り合いのブログ記事で

- [国勢調査オンライン - 電気ウナギ的○○](https://blog.netandfield.com/shar/2020/09/post-3809.html)

というのがあって「そういや前回はネットで回答したな」と思い出して[作業用リポジトリ](https://github.com/spiegel-im-spiegel/github-pages-env "spiegel-im-spiegel/github-pages-env: Document Environment for spiegel-im-spiegel.github.io")に grep かけて探してみたらやっぱり書いてた。

- [国勢調査を片付けた]({{< ref "/remark/2015/0920-diary.md#census" >}} "今日の戯れ言：週末は掃除三昧")

なんでも記事にしておくものである。
5年前の私 good work!

## あの会社は今

当時は「国勢調査オンライン」のサイト証明書は Symantec 社が発行していて，ページにデカデカとロゴが貼り付けてあって苦笑したものだが，当の Symantec 社は CA 事業の不祥事で Web ブラウザのベンダ企業から信用されなくなり，2017年に事業を手放している。

- [Symantec→DigiCertでSSL/TLS証明書はどうなる？　日本国内にも認証局構築へ、IoT機器市場も見据え - INTERNET Watch](https://internet.watch.impress.co.jp/docs/news/1103161.html)

今回はどうしてるんだろうと見てみたら DigiCert 社だったよ。
何も変わってないな（笑）

そういや DigiCert 社のルート CA 証明書は RSA/2048 鍵で2031年まで有効なんだよな。
ちなみに [RSA/2048 鍵が acceptable なのは2030年まで]({{< ref "/remark/2017/10/key-parameters.md" >}} "暗号鍵関連の各種変数について")だ。
まぁ鼻の先は問題ないが，政府調達品でこの程度の認識しかない企業を使うのはどうなんだろう。

{{< div-box type="markdown" >}}
**【追記】**
Facebook で教えてもらったが DigiCert 社の CA 証明書はよりサイクルの短いものに順次切り替えていくそうだ。
Web 用のサーバ証明書は長くても1年程度のサイクルになる筈だし，数年かけて切り替えていくイメージだろうか。
{{< /div-box >}}

## なりすましと本人確認

そういや5年前はうっかり Phishing サイトを作って怒られた輩がいて話題になった。

- [国勢調査の“偽サイト”作った意図は？　総務省から削除依頼……「騒ぎになり深く反省」と制作者 (1/3) - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1509/15/news083.html)

最近の「ドコモロ系事案」を引くまでもなく，なりすましや本人確認は古くて新しい問題であり続ける。
このご時世に「オンライン推奨」とか現政権は勇気あるな（皮肉）

国勢調査が面白いのは，基本的に「戸」単位であり，その枠組みの中で識別できていれば「本人確認」は必ずしも必要条件ではないということだ。
そして，その「戸」単位の確認を行うのが「国勢調査員」である。
これは紙でもネットでも同じこと。
ネットで回答するにしても，そのアカウント情報は国勢調査員から貰わなければならない。

そもそも路上生活者とかにも国勢調査員が出向いて聞き取り調査するんだよ。
日本語が通じない人達だって相当数いるだろうに，ホンマ国勢調査員のご苦労は察するにあまりある。

故に「『コロナ』だからオンライン推奨」というのは全く以って詭弁である。
ちなみに，今回うちで国勢調査の回答をしたのは親父殿だが，パソコンもスマホも持ってないので紙で回答したらしい。
私？ 私はただの居候ですから（笑）

## 「国勢調査オンライン」はダサいか

どうも「国勢調査オンライン」のサイトをダサいとか文句を言ってる人がいるらしい。
まぁ今風ではないな（笑）

今回は私は回答してないのでどうだったかは知らないが，[前回ネットで回答]({{< ref "/remark/2015/0920-diary.md#census" >}} "今日の戯れ言：週末は掃除三昧")してめっさ楽だったのは覚えている。
ぶっちゃけて言うが

{{< div-gen type="markdown" class="center" >}}
**政府系サイトに {{% ruby "屁のつっぱり" %}}SEO{{% /ruby %}} は要らんですよ**
{{< /div-gen >}}

そもそも Web アクセシビリティってそんな簡単じゃない。
うちのサイトもそうだけど，本当に不特定に見易い使い易いサイトなんか無理である。
このブログサイトをダークモードにしてるのは，主に「私」のためだ。

できるだけ多くの人に対応しようとして結果として今風じゃないデザインになったとしても，利用者が文句を言う筋合いじゃないのだ。
機能要件やセキュリティ要件を満たしていない部分があるのなら，文句を言って然るべきだけど（笑）

というわけで，まだ回答してない人でネット回答可能であれば，ネットのほうが断然楽なので，お試しあれ。

## いまさら思い出したが...

このブログサイト，[正式オープン]({{< ref "/remark/2015/open-this-site.md" >}} "text.Baldanders.info 正式オープン")から5周年だわ。
何も考えてなかった。

まぁ，[この前買ったタブレット]({{< ref "/remark/2020/09/assemble-a-living-pc.md" >}} "整いました！")がご褒美ということで（笑）

## ブックマーク

- [NIST SP 800-207: “Zero Trust Architecture”]({{< ref "/remark/2020/09/nist-sp-800-207-zero-trust-architecture.md" >}})
- [Authenticator と AAL]({{< ref "/remark/2020/09/authenticator-and-aal.md" >}})

<!-- eof -->
