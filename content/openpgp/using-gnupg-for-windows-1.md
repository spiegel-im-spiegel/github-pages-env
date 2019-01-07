+++
date = "2017-12-01T17:48:32+09:00"
update = "2018-06-24T13:55:41+09:00"
title = "GnuPG for Windows インストール編"
description = "Windows 版 GnuPG のインストールについて。"
draft = false
tags = ["security", "cryptography", "openpgp", "gnupg", "tools"]
image = "/images/attention/openpgp.png"

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
  url = "https://baldanders.info/spiegel/profile/"

[scripts]
  mathjax = false
  mermaidjs = false
+++

この記事は以下の記事を最新版 [GnuPG] 用に再構成したものです。

- [GnuPG 2.1.0 (modern) for Windows のインストール — Baldanders.info](https://baldanders.info/spiegel/log2/000770.shtml)

ECC への対応については以下の Gist ページを参照のこと。
（これもそのうち再構成してここで公開する予定）

- [Windows 版 GnuPG 2.1.x を使ってみる](https://gist.github.com/spiegel-im-spiegel/f177c02af04d3b34ade0)

## OpenPGP 実装としての GnuPG

[OpenPGP] の起源は [Phil Zimmermann] さんによる PGP (Pretty Good Privacy) と呼ばれる暗号ツールである。
PGP の最初のバージョンは1991年に公開された[^1991]。
当時の [Phil Zimmermann] さんは反核運動家で，政府等の組織からデータやメッセージ（特に電子メール）を保護するための手段として PGP を開発し，最終的にそれをフリーで公開した[^pgp]。

[^1991]: 当時の PGP の仕様は（公開年にちなんでか） [RFC 1991](http://tools.ietf.org/html/rfc1991 "RFC 1991 - PGP Message Exchange Formats") として公開されている。
[^pgp]: もともと [Phil Zimmermann] さんは PGP をシェアウェアとして売り出すつもりだったらしい。しかし米国内で事実上暗号を禁止する法案が提出され，法案の可決を阻止する目的もあり PGP をフリーで公開した。ところが [Phil Zimmermann] さんが暗号に関する特許について迂闊だったことや PGP が ftp サーバを通じて海外に漏洩してしまった（当時は暗号製品には輸出規制があり強い暗号製品は米国外に持ち出せなかった）ことなどもあって，しばらくの間 [Phil Zimmermann] さんと PGP は不遇の身の上となる。当時の輸入規制に「書籍」は含まれていなかったため，最新版の PGP コードを書籍として出版し海外でコンパイルする国際化プロジェクトがあった。何もかも懐かしい（笑） ちなみに現在の PGP は無料ではない。

その後 PGP はいくつか改良を重ね，1998年に [RFC 2440] つまり OpenPGP として標準化された[^op]。
また特許上の制限や国際政治上の問題も2000年を機に大幅に緩和され PGP を含む多くの暗号製品が本格的に使われるようになった。

[^op]: 現在は [RFC 4880] にアップデートされ，更に改良が進められている。

[GnuPG] は [OpenPGP] をベースにドイツで生まれた製品である。
特定の個人・組織が独占することのないよう [GNU] プロジェクトの一環として現在も開発が行われている[^gpl]。

[^gpl]: [GnuPG] の著作権は [FSF] に帰属し [GNU GPL] でライセンスされている。

[GnuPG] の最新バージョンは 2.2 系である。
2.0 系（旧 stable version）および 2.1 系（旧 modern version）は 2.2 系に統合された。
また 2.0 系は2017年末でサポートが終了した。

なお classic version である 1.4 系はレガシー・システムとの互換性維持のためにメンテナンスが継続されるが， Windows で新たに導入するのであれば 2.2 系を強くお勧めする。

## 【事前準備】インストーラのダウンロード

[ダウンロードページ](https://gnupg.org/download/ "GnuPG - Download") の “GnuPG binary releases” にある Windows 用のバイナリへのリンクから “current GnuPG” をダウンロードする（2017年11月20日時点で v2.2.3 が最新）。
必ずインストーラ本体と署名ファイルをセットでダウンロードすること。

前バージョンの [GnuPG] を持っている場合はインストーラの署名検証を行い，正しいファイルであることを確認すること。

```text
$ gpg --verify  gnupg-w32-2.2.3_20171120.exe.sig
gpg: 署名されたデータが'gnupg-w32-2.2.3_20171120.exe'にあると想定します
gpg: 11/20/17 21:25:34 東京 (標準時)に施された署名
gpg:                RSA鍵D8692123C4065DEA5E0F3AB5249B39D24F25E3B6を使用
gpg: "Werner Koch (dist sig)"からの正しい署名 [充分]
```

署名検証用の公開鍵は以下にある。

- [Signature Key](https://gnupg.org/signature_key.html)

公開鍵は鍵サーバから取得することもできる。

```text
$ gpg --keyserver keys.gnupg.net --recv-keys 0x4F25E3B6
```

### Classic Version 削除のススメ

現行バージョンのファイル構成は classic version と互換性がない。
Windows で現行バージョンを利用するのなら classic version は削除するのがお勧めである。

1. Classic version の鍵束（keyring; `pubring.gpg`, `secring.gpg`, `trustdb.gpg`）は別の場所に退避させておき，現行バージョンのインストール後にインポートする[^imp]。インポートの方法は後述する
1. Classic version  アンインストール後に環境変数 `PATH` に `gpg.exe` へのパスが残っている場合は念のためこれも削除しておく。環境変数の変更方法がわからない方は無理に削除しなくてもいい
1. Classic version  アンインストール後にレジストリ `HKEY_CURRENT_USER\Software\GNU\GnuPG` が残っている場合は，これも削除してしまうのがよいだろう。ただしレジストリ操作に自信のない人はこれも無理に触らなくてよい

[^imp]: 実は classic version の鍵束をそのまま使っても自動的にファイルが移行されるため大抵は問題ないのだが，旧鍵束にはバグが混入しているそうで，安全のため明示的にインポート作業を行うほうがいいらしい。なお現行バージョンの [GnuPG] は，移行時以外は classic version の `secring.gpg` を参照しないため，Classic version と混在させるのであれば取り扱いに注意が必要である。（`gpg-v21-migrated` ファイルを削除すると再度移行処理が走るらしい）

なお `trustdb.gpg` は以下のコマンドでテキストファイルにエクスポートしておくとよい[^t]。

[^t]: `trustdb.gpg` ファイルはそのまま使うのではなく， `--export-ownertrust` オプションでテキストファイルにエクスポートしたものを使うのが安全なようだ。

```text
$ gpg --export-ownertrust > trust.txt
```

## インストーラの実行

準備ができたところでインストールを始めよう。
ダウンロードしたインストーラを起動する（スクリーンショットが古いがご容赦）。

{{< fig-img src="https://farm2.staticflickr.com/1502/24974542243_4e83a1d7b1.jpg" title="Installing GnuPG for Windows (1)" link="https://www.flickr.com/photos/spiegel/24974542243/" >}}

英語だけど無問題。
ほとんど選択肢はないので `[Next]` ボタンで先に進めていけばいい。

{{< fig-img src="https://farm2.staticflickr.com/1545/25482633892_d9dc023e1a.jpg" title="Installing GnuPG for Windows (2)" link="https://www.flickr.com/photos/spiegel/25482633892/" >}}

{{< fig-img src="https://farm2.staticflickr.com/1695/24974542073_20408e1079.jpg" title="Installing GnuPG for Windows (3)" link="https://www.flickr.com/photos/spiegel/24974542073/" >}}

<!--
{{< fig-img src="https://farm2.staticflickr.com/1472/25305629970_6f5dcb4ef0.jpg" title="Installing GnuPG for Windows (4)" link="https://www.flickr.com/photos/spiegel/25305629970/" >}}

インストール先のフォルダを変えたい場合はここで変更する。

{{< fig-img src="https://farm2.staticflickr.com/1449/25601226555_b07b73e7fa.jpg" title="Installing GnuPG for Windows (5)" link="https://www.flickr.com/photos/spiegel/25601226555/" >}}

-->

{{% div-box %}}
**【追記 2018-06-24】**
最近の GnuPG ではインストール先フォルダが `C:\Program Files (x86)\gnupg` 固定になっている（64bit版 GnuPG for Windows のバイナリ提供はない）。
古いバージョンで左記のフォルダ以外にインストールしている場合はそのフォルダに上書きインストールされる。
では続きをどうぞ。
{{% /div-box %}}

{{< fig-img src="https://farm2.staticflickr.com/1633/25575126816_f090b537bf.jpg" title="Installing GnuPG for Windows (6)" link="https://www.flickr.com/photos/spiegel/25575126816/" >}}

{{< fig-img src="https://farm2.staticflickr.com/1587/24970753344_5da4faf427.jpg" title="Installing GnuPG for Windows (7)" link="https://www.flickr.com/photos/spiegel/24970753344/" >}}

`[Finish]` ボタンを押してインストール完了。

この時点で `PATH` も通っているため，コマンドプロンプトから

```text
$ gpg --version
gpg (GnuPG) 2.2.3
libgcrypt 1.8.1
Copyright (C) 2017 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: C:/Users/username/AppData/Roaming/gnupg
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

と入力すればバージョン情報が表示される。

### ホームディレクトリの変更（必要に応じて）

インストール直後は `%APPDATA%\gnupg` が [GnuPG] のホームディレクトリになっている[^gh]。
通常はこれで問題ないが，他のフォルダに変更したい場合は環境変数 `GNUPGHOME` でフォルダを指定する。
また `gpg.exe` 起動時に `--homedir` オプションでホームディレクトリを直接指定することもできる（`--homedir` オプションが優先）。

[^gh]: [GnuPG] のホームディレクトリにはいわゆる鍵束（key ring）の情報が格納される。 [GnuPG] はこの鍵束から暗号鍵を取得して暗号化および復号を行うわけだ。ちなみに，環境変数 `APPDATA` には通常 `C:\Users\username\AppData\Roaming` （`username` はログインユーザの名前）がセットされている。ちなみに UNIX 系のプラットフォームでは `~/.gnupg` が [GnuPG] 既定のホームディレクトリだが Windows は構成が異なるためこのようになっている。

```text
$ gpg --version --homedir C:\usr\home
gpg (GnuPG) 2.2.3
libgcrypt 1.8.1
Copyright (C) 2017 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: C:/usr/home
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

インストール直後のホームディレクトリはまだ空である。

### Classic Version の鍵束のインポート（移行時のみ）

Classic version からアップグレードした人は旧鍵束（`pubring.gpg`, `secring.gpg`, `trustdb.gpg` → `trust.txt`）をあらかじめ退避していると思うが，これを現行バージョンへインポートする。
手順は以下のとおり。

```text
$ gpg --import-options import-local-sigs --import pubring.gpg
$ gpg --import secring.gpg
$ gpg --import-ownertrust trust.txt
```

秘密鍵（`secring.gpg`）のインポートでは鍵の数だけパスフレーズの入力をを要求される。

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1507/25316582890_9ff8c3d2ea_o.png" title="GnuPG pinentry" link="https://www.flickr.com/photos/spiegel/25316582890/" >}}

このプロンプト画面（Pinentry）については[次回]に `gpg-agent` の話と絡めて説明する。

上手くインポートできていれば以下のように鍵を表示することができる。

```text
$ gpg --list-keys 0x4F25E3B6
pub   rsa2048 2011-01-12 [SC] [有効期限: 2019-12-31]
      D8692123C4065DEA5E0F3AB5249B39D24F25E3B6
uid           [  充分  ] Werner Koch (dist sig)
sub   rsa2048 2011-01-12 [A] [有効期限: 2019-12-31]
```

インポートにより [GnuPG] のホームディレクトリには以下のフォルダ・ファイルができているはずである。

- `pubring.kbx` ファイル[^kbx]
- `trustdb.gpg` ファイル
- `gpg-v21-migrated` ファイル
- `private-keys-v1.d` フォルダ

[^kbx]: kbx は keybox の略らしい。 バージョン 2 以降では OpenPGP の鍵束だけでなく S/MIME （X.509）や OpenSSH の鍵も格納できる。

`private-keys-v1.d` フォルダにはインポートした秘密鍵の数だけファイルが作成されている。
`gpg-v21-migrated` ファイルは鍵束が現行バージョンへ移行したことを示すフラグである。

## 参考になる（かもしれない） Web ページ

[ブックマークはこちら]({{< ref "/openpgp/bookmark.md" >}})に移動した。

[次回]: {{< ref "/openpgp/using-gnupg-for-windows-2.md" >}} "GnuPG for Windows : gpg-agent について"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 2440]: http://tools.ietf.org/html/rfc2440 "RFC 2440 - OpenPGP Message Format"
[Phil Zimmermann]: https://www.philzimmermann.com/ "Phil Zimmermann's Home Page"
[GNU]: https://www.gnu.org/ "The GNU Operating System and the Free Software Movement"
[FSF]: http://www.fsf.org/ "Front Page — Free Software Foundation — working together for free software"
[GNU GPL]: http://www.gnu.org/licenses/licenses.html#GPL "The GNU General Public License"
[Gpg4win]: https://www.gpg4win.org/ "Gpg4win - Secure email and file encryption with GnuPG for Windows"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/">暗号化 プライバシーを救った反乱者たち</a></dt><dd>スティーブン・レビー 斉藤 隆央 </dd><dd>紀伊國屋書店 2002-02-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487593100X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/487593100X.09._SCTHUMBZZZ_.jpg"  alt="ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484111160/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484111160.09._SCTHUMBZZZ_.jpg"  alt="グーグル ネット覇者の真実 追われる立場から追う立場へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/410215972X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/410215972X.09._SCTHUMBZZZ_.jpg"  alt="暗号解読〈上〉 (新潮文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159738/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159738.09._SCTHUMBZZZ_.jpg"  alt="暗号解読 下巻 (新潮文庫 シ 37-3)"  /></a> </p>
<p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-03-09">2015/03/09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/5132396FFQL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/">PGP―暗号メールと電子署名</a></dt><dd>シムソン ガーフィンケル Simson Garfinkel </dd><dd>オライリー・ジャパン 1996-04</dd><dd>評価<abbr class="rating" title="3"><img src="http://g-images.amazon.com/images/G/01/detail/stars-3-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4756136494/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4756136494.09._SCTHUMBZZZ_.jpg"  alt="プログラミング作法"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320026926.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語C 第2版 ANSI規格準拠"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797350997/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797350997.09._SCTHUMBZZZ_.jpg"  alt="新版暗号技術入門 秘密の国のアリス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798132608/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798132608.09._SCTHUMBZZZ_.jpg"  alt="情報処理教科書 高度試験午後II論述 春期・秋期 (EXAMPRESS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798105538/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798105538.09._SCTHUMBZZZ_.jpg"  alt="エンタープライズ アプリケーションアーキテクチャパターン (Object Oriented Selection)"  /></a> </p>
<p class="description" >良書なのだが，残念ながら内容が古すぎた。 PGP の歴史資料として読むならいいかもしれない。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-10-16">2014/10/16</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
