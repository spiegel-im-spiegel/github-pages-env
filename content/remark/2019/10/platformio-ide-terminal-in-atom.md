+++
title = "ATOM 1.41 にアップデートしたら platformio-ide-terminal が動かねー！ と思ったら"
date =  "2019-10-25T22:33:09+09:00"
description = "色々と試行錯誤してみたが，どうやら Ubuntu の APT でインストールしている node.js と ATOM 内部の node.js が衝突しているらしい。"
image = "/images/attention/kitten.jpg"
tags = ["atom", "editor", "ubuntu", "terminal", "nodejs"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

{{< div-box type="markdown" >}}
**【2020-11-04 追記】**
既に私は [platformio-ide-terminal] に見切りをつけて [x-terminal] に乗り換えました。

- [とりあえず ATOM エディタ内ターミナルを x-terminal に乗り換えた]({{< ref "/remark/2020/05/x-terminal-with-atom.md" >}})

メインで使ってるのは [Ubuntu] ですが， Windows 環境でも問題なく動きます。
[NYAGOS] もちゃんと動く。
ブラボー！

[platformio-ide-terminal]: https://atom.io/packages/platformio-ide-terminal
[x-terminal]: https://atom.io/packages/x-terminal
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[NYAGOS]: https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"
{{< /div-box >}}

最近は猫も杓子も Vim か VSCode で寂しい限りだが， [ATOM] 1.41 がリリースされ，ようやく [Electron] 4 ベースの構成になった。

- [Release 1.41.0 · atom/atom · GitHub](https://github.com/atom/atom/releases/tag/v1.41.0)

それはいいのだが，また [platformio-ide-terminal] が動かなくなった。
しかも [1.39 のとき]({{< ref "/release/2019/07/atom-1_39-is-released.md" >}} "TOM エディタ v.1.39 がリリースされたのだが...
")とは様子が異なるようだ。

色々と試行錯誤してみたが，どうやら [Ubuntu] の APT でインストールしている [node.js] と [ATOM] 内部の [node.js] が衝突しているらしい。

```text
$ sudo apt purge nodejs
```

で [Ubuntu] 側の [node.js] をいったん削除した上で [platformio-ide-terminal] を入れ直したら問題なく動いた。

ちなみに APT で管理されている [node.js] は素の設定で

```text
$ apt show nodejs
Package: nodejs
Version: 10.15.2~dfsg-2ubuntu1
Priority: extra
Section: universe/web
Origin: Ubuntu
...
```

だった。
一方 [ATOM] 1.41 の構成は

```text
$ atom -v
Atom    : 1.41.0
Electron: 4.2.7
Chrome  : 69.0.3497.128
Node    : 10.11.0

$ apm -v
apm  2.4.3
npm  6.2.0
node 10.2.1 x64
atom 1.41.0
...
```

と何とも微妙な感じである。
大丈夫か，これ。

更に更に余談だが [Ubuntu] の APT に最新の [node.js] を組み込むには，以下のようにスクリプトを取ってきて実行すればいいらしい（以下は LTS 版の [node.js] 12 を組み込む場合）。

```text
$ curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
```

ただし現時点（2019-10-25）で [Ubuntu] 19.10 には対応してなかった `orz`

## ブックマーク

- [nodesource/distributions: NodeSource Node.js Binary Distributions](https://github.com/nodesource/distributions)

[ATOM]: https://atom.io/
[Electron]: https://electronjs.org/ "Electron | Build cross platform desktop apps with JavaScript, HTML, and CSS."
[platformio-ide-terminal]: https://atom.io/packages/platformio-ide-terminal
[node.js]: https://nodejs.org/en/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考（にならない）図書

{{% review-paapi "B00B47FIDC" %}} <!-- PLUTO -->
{{% review-paapi "B00YRVO8EC" %}} <!-- アトム ザ・ビギニング -->
