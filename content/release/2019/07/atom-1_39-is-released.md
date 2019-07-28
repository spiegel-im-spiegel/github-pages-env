+++
title = "ATOM エディタ v.1.39 がリリースされたのだが..."
date =  "2019-07-25T22:39:18+09:00"
description = "実験的に組み込まれている ripgrep がめっさ速い。 こりゃあええわ。"
image = "/images/attention/tools.png"
tags = ["atom", "editor", "ubuntu", "grep", "terminal"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨夜寝る前にメールをチェックしてたら [ATOM] エディタ v.1.39.0 がリリースされてた。

- [Release 1.39.0 · atom/atom · GitHub](https://github.com/atom/atom/releases/tag/v1.39.0)

ちなみに [ATOM] は公式の APT リポジトリを公開しているので [Ubuntu] ユーザはそのリポジトリを登録してインストールすればよい。

```text
$ wget -qO - https://packagecloud.io/AtomEditor/atom/gpgkey | sudo apt-key add -
$ sudo sh -c 'echo "deb [arch=amd64] https://packagecloud.io/AtomEditor/atom/any/ any main" > /etc/apt/sources.list.d/atom.list'
$ sudo apt update
$ sudo apt install atom
```

変更や修正は色々あるが，今回の目玉はベースになっている Electron や Chrome のバージョンを上げたことだろう。
うちの環境だとこんな感じ。

```text
$ atom -v
Atom    : 1.39.0
Electron: 3.1.10
Chrome  : 66.0.3359.181
Node    : 10.2.0
```

Electron が3系になったおかげで [Ubuntu] 環境でメニューバーの文字がちゃんと見えるようになった。
まぁ見えなくても別に困ってはいなかったのだが（笑）

悪いお知らせ。

[platformio-ide-terminal] が動かなくなった。

> node-pty-prebuilt@0.7.6 – The module '/home/username/.atom/packages/platformio-ide-terminal/node_modules/node-pty-prebuilt/build/Release/pty.node' was compiled against a different Node.js version using NODE_MODULE_VERSION 57. This version of Node.js requires NODE_MODULE_VERSION 64. Please try re-compiling or re-installing the module (for instance, using `npm rebuild` or `npm install`).

どうもコンパイル済みパッケージの内部バージョンが合わなくてコンフリクトを起こしているらしい。
これは最新バージョンが上がってくるのを待つしかないか。

エディタ上でターミナルが使えないのはめっさ不便なんだが。

朗報。

実験的に組み込まれている [ripgrep] がめっさ速い。
こりゃあええわ。

余談だが， [Ubuntu] では CLI の [ripgrep] は APT でインストールするのがいいようだ。

```text
$ sudo apt install ripgrep
$ rg -V
ripgrep 0.10.0
```

対応しているバージョンを見ると，一見 snap 版の [ripgrep] のほうが良さげだが

{{< fig-quote type="md" lang="en" title="BurntSushi/ripgrep: ripgrep recursively searches directories for a regex pattern" link="https://github.com/BurntSushi/ripgrep" >}}
{{< quote >}}N.B. Various snaps for ripgrep on Ubuntu are also available, but none of them seem to work right and generate a number of very strange bug reports that I don't know how to fix and don't have the time to fix. Therefore, it is no longer a recommended installation option.{{< /quote >}}
{{< /fig-quote >}}

なんだとさ。
残念。

しかし，まぁ，やはり時代は Rust かねぇ。
[Microsoft も興味を持ってるみたい](https://japan.zdnet.com/article/35140212/ "マイクロソフト、セキュアなコード実現に向けプログラミング言語「Rust」評価 - ZDNet Japan")だし。

## ブックマーク

- [Release 1.39.1 · atom/atom · GitHub](https://github.com/atom/atom/releases/tag/v1.39.1)

- [GitHub、フリーのコードエディター「Atom 1.39」をリリース ～あいまい検索が大幅に高速化 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1197983.html)

[ATOM]: https://atom.io/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[platformio-ide-terminal]: https://atom.io/packages/platformio-ide-terminal
[ripgrep]: https://github.com/BurntSushi/ripgrep "BurntSushi/ripgrep: ripgrep recursively searches directories for a regex pattern"
