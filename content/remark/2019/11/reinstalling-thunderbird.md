+++
title = "結局 Thunderbird もインストールし直すことにした"
date =  "2019-11-10T16:43:28+09:00"
description = "どうやら Ubuntu は Thunderbird をまともにメンテナンスする気がないらしいとの結論に至った。そこで APT による管理はすっぱり諦めて自前で導入・管理することにした。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "thunderbird", "install", "mua" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] が [19.10 にアップグレード]({{< ref "/release/2019/10/upgrade-ubuntu-19_10.md" >}} "Ubuntu 19.10 にアップグレードする")されて [Thunderbird] が 68 ベースになったまではよかったが，そこから全くメンテされなくなってしまった。

[APT] 管理下の [Thunderbird] は2019年11月時点で

```text
$ apt show thunderbird
Package: thunderbird
Version: 1:68.1.2+build1-0ubuntu1
Priority: optional
Section: mail
Origin: Ubuntu
...
```

となっているが，実際には既に 68.2.2 がリリースされている。

- [Thunderbird — Release Notes (68.2.2) — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.2.2/releasenotes/)

そういえば 60.9.0 がリリースされたときも [APT] で配信可能になるまで1ヶ月かかったんだよな。

ちうわけで，どうやら [Ubuntu] は [Thunderbird] をまともにメンテナンスする気がないらしいとの結論に至った。
以前に「[デフォルトのメールクライアントをGearyに変更する提案](https://kledgeb.blogspot.com/2019/04/ubuntu-1910-3-geary.html "Ubuntu 19.10 その3 - デフォルトのメールクライアントをGearyに変更する提案 - kledgeb")」とかあったようなので仕方がないのかもしれないが，（鍵管理を含めて）まともに OpenPGP 暗号化が使える MUA は殆どない[^mua1] ので [Thunderbird] が「使えない」のは非常に困るわけですよ。

[^mua1]: ブラウザ上で動作する Web メール用の暗号化ツールとしては [Mailvelope](https://www.mailvelope.com/) が有名で，私も一応インストールしているが，鍵管理が独特で気持ち悪いので実際には使ってない。

そこで [APT] による管理はすっぱり諦めて自前で導入・管理することにした。
といっても[ダウンロードページ](https://www.thunderbird.net/en-US/thunderbird/all/ "Download Thunderbird in your language — Mozilla")でプラットフォームと言語に応じたバイナリをダウンロードしてローカルの適当な場所で展開するだけだけどね[^inst1]。

[^inst1]: [Thunderbird] はインストール先ディレクトリを変えると別のインスタンスとみなしてそれぞれ別のプロファイルを作成するので注意が必要だ。この辺をコントロースするには `~/thunderbird/` ディレクトリにある `installs.ini` および `profiles.ini` 各ファイルを弄る必要があるらしい。

設定およびメールデータは `~/thunderbird/` ディレクトリ以下に格納されているが，バックアップをとった上で，バッサリ削除して最初からやり直した。
実際には，パソコンと携帯端末とでメールデータを共有するために，IMAP で管理しているので（だからこそ E2E 暗号化ができることが重要なんだけどね）特に問題はなかった。
[Enigmail] も問題なく動作することを確認済み。

よーし，うむうむ，よーし。

## ブックマーク

- [Thunderbird — 68.0 System Requirements — Mozilla](https://www.thunderbird.net/en-US/thunderbird/68.0/system-requirements/)
- [How to Install and Setup Thunderbird Email Client in Ubuntu](https://vitux.com/how-to-install-and-setup-thunderbird-email-client-in-ubuntu/)
- [GNOME3でアプリケーションメニューにランチャーを追加する (alacarteを使用しない方法) – Crow's eye](https://kzmmtmt.pgw.jp/?p=842)

- [Thunderbird, Enigmail and OpenPGP | The Mozilla Thunderbird Blog](https://blog.mozilla.org/thunderbird/2019/10/thunderbird-enigmail-and-openpgp/) : 将来バージョンで [Enigmail] は [Thunderbird] に組み込まれるそうだ

- [Thunderbird でメール暗号化]({{< ref "/openpgp/encrypted-mail-with-thunderbird.md" >}}) : Windows 用に書いたが Linux でも同じように使えるので大丈夫

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[APT]: https://debian-handbook.info/browse/ja-JP/stable/apt.html "第 6 章 メンテナンスと更新、APT ツール"
[Thunderbird]: https://www.thunderbird.net/ "Thunderbird — Software made to make email easier. — Mozilla"
[Enigmail]: https://addons.thunderbird.net/thunderbird/addon/enigmail/ "Enigmail :: Add-ons for Thunderbird"
