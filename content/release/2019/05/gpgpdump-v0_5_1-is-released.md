+++
title = "gpgpdump v0.5.1 をリリースした（v0.5.2 もリリースした）"
date =  "2019-05-01T13:37:53+09:00"
description = "リテラルパケットのファイル名に含まれる改行コードなどの制御コードを符号点表示に展開するようにした。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump", "security", "vulnerability"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を視覚化する [gpgpdump] の v0.5.1 をリリースした。

- [Release v0.5.1 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.5.1)

今回はバグ修正（？）版。

いや [OpenPGP のメーリングリスト](https://www.ietf.org/mailman/listinfo/openpgp)で以下の記事が流れてきたのだが

- [Spoofing OpenPGP and S/MIME Signatures in Emails](https://mailarchive.ietf.org/arch/msg/openpgp/_JZSrjvFaoPMJoV-z8Lz5hYLtBM)

[GnuPG] の[脆弱性の話なんて昨年のこと]({{< ref "/release/2018/06/gnupg-2_2_8-and-libgcrypt-1_8_3-are-released.md" >}} "GnuPG 2.2.8 および Libgcrypt 1.8.3 がリリース（セキュリティ・アップデート）")だし他の問題にしたって MUA の実装上の不具合ではあろうけど脆弱性とは言い難いし「なんだかなぁ」という感じで眺めていたのだが，「リテラルパケット（tag11）については [gpgpdump] も対処しておいたほうがいいかな」と思い立ち修正してみた。

具体的にはリテラルパケットのファイル名に含まれる改行コードなどの制御コードを `(U+000A)` のように符号点表示に展開するようにした。
制御コードの判定には [Go 言語]標準の [`unicode`]`.IsControl()` 関数を使っている（手抜き実装`w`）。

他にも UTF-8 以外の文字エンコーディングの場合は `"invalid text string"` と表記するようにした。
まぁ Shift-JIS とか EUC とか軒並み引っかかっちゃうんだけど，もう気にしないことにした（今までだって文字化けしてまともに表示できなかった筈）。

それにしても [Ubuntu] は快適だね。
今まで Windows で作業してたのが馬鹿みたいだよ。
Microsoft が今後 PWA (Progressive Web Apps) へのシフトを進めていくと Windows はどんどん「{{% ruby "Progammable Controller" %}}コントローラ{{% /ruby %}}」になっていくだろうし，そうなると「{{% ruby "Personal Computer" %}}パソコン{{% /ruby %}}」と言えるのは macOS や [Ubuntu] のような UNIX 系のデスクトップ OS だけになるかもしれないねぇ[^subs1]。

[^subs1]: と考えると Windows のサブシステムに Linux を入れたのは本当にお馬鹿な選択だったと言わざるを得ない。包含関係が逆だよ。 Linux 環境下で Windows がサブシステムとして動作できるようにしないと。

## 【追記】 [gpgpdump] v0.5.2 をリリースした

v0.5.1 をリリースしたばっかだが v0.5.2 を出した。

- [Release v0.5.2 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.5.2)

いやぁ，よく考えたらリテラルのテキストを扱うのはリテラルパケットだけじゃなかった。
ちうわけでコード内を浚って生データをそのまま string にキャストして出力してる部分を修正した。
アホだなぁ，私。

これで大丈夫なはず。

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[`gpgpdump`]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`unicode`]: https://golang.org/pkg/unicode/ "unicode - The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
