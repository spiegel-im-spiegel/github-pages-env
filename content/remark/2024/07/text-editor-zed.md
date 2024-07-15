+++
title = "テキストエディタ Zed を試してみる"
date =  "2024-07-15T14:20:40+09:00"
description = "Atom エディタの元ユーザとして Zed には期待している。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "editor", "zed" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

かつての Atom エディタの開発者が作ったという Rust 製のテキストエディタ [Zed] の Linux 版（安定版）が登場したらしい。

- [Zed on Linux is here!](https://zed.dev/linux)
- [Linux when? Linux now.](https://zed.dev/blog/zed-on-linux)
- [Rust製のオープンソースエディタ「Zed」のLinux安定版が公開 － Publickey](https://www.publickey1.jp/blog/24/rustzedlinux.html)

既に macOS 版は存在する。
Windows 版はまだないようだ。

インストール方法は

- [Linux - Zed](https://zed.dev/docs/linux)

を参照のこと。
大抵の場合はインストール用の shell スクリプトをダウンロード，実行すればいいようだ。

```text
$ curl https://zed.dev/install.sh | sh
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  4220  100  4220    0     0   9353      0 --:--:-- --:--:-- --:--:--  9336
Downloading Zed
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   565    0   565    0     0    382      0 --:--:--  0:00:01 --:--:--   383
100 53.0M  100 53.0M    0     0  9866k      0  0:00:05  0:00:05 --:--:-- 15.9M
To run Zed from your terminal, you must add ~/.local/bin to your PATH
Run:
   echo 'export PATH=$HOME/.local/bin:$PATH' >> ~/.bashrc
   source ~/.bashrc
To run Zed now, '~/.local/bin/zed'
```

`~/.bashrc` ファイルにパスを追加しろって書いてあるが Ubuntu の場合は `~/.profile` ファイルに

```bash
# set PATH so it includes user's private bin if it exists
home_local_bin_path="$HOME/.local/bin"
if [ -d "$home_local_bin_path" -a -n "${PATH##*${home_local_bin_path}}" -a -n "${PATH##*${home_local_bin_path}:*}" ] ; then
  export PATH="$home_local_bin_path:$PATH"
fi
unset home_local_bin_path
```

みたいな記述があれば特に追加で何かする必要はないだろう。
ちなみに `~/.local/bin` ディレクトリに直接インストールされるわけではなく，実体は `~/.local/zed.app` ディレクトリ以下に展開され `~/.local/bin` ディレクトリにはシンボリックリンクが配置されている。

Ubuntu の場合，上述のスクリプトを使ってインストールする場合は制限があって Ubuntu 18 以前（2.29 より古い glibc を使用している）のバージョンでは動かないらしい。
また32ビットアーキテクチャでも動かない。
まぁ，ソースからビルドすれば行けるかもしれないが（Rust って古いディストリビューションって対応してたっけ？）。

あとアプリケーションの速度を確保するために [Vulkan] を通じて GPU と通信しているらしい。
このため GPU がないと動作が遅くなるとのこと。
GPU を確認する場合は

```text
$ sudo apt install vulkan-tools
```

で `vulkan-tools` パッケージをインストールし，ターミナルで `vkcube` コマンドを叩いてみる。
すると以下のウィンドウが起動して画面内の立方体が物凄い勢いで回転するのだが

{{< fig-img src="./vkcube.png" title="vkcube" link="./vkcube.png" width="528" >}}

ターミナル側には

```text
$ vkcube
Selected GPU 0: AMD Radeon Graphics (RADV RENOIR), type: IntegratedGpu
```

のように使用する GPU の ID が表示される。
これが `llvmpipe` だと GPU がないか認識してないってことらしい。
また AMD Radeon の場合はドライバによって “Broken Pipe” のエラーが出て上手く動かないことがあるそうな。
この場合はドライバを RADV にするといいとのこと（参照： [#13880](https://github.com/zed-industries/zed/issues/13880 "Arch Linux can't launch after build: Io error: Broken pipe (os error 32) · Issue #13880 · zed-industries/zed")）

インストールが完了してターミナルから `zed` コマンドが無事起動できたら以下の画面が表示される筈。

{{< fig-img src="./zed-1.png" title="vkcube" link="./zed-1.png" width="1044" >}}

最初から Vim モードを選択可能なのか。
Vimmer はどこにでもいるんだな（笑） その下の2つのチェックボックスはテレメータの設定なので，テレメータを送信したくないならチェックを外しておけばよい。

細かい設定は `~/.config/zed/settings.json` ファイルで可能。
設定可能な項目は “[Configuring Zed](https://zed.dev/docs/configuring-zed "Configuring Zed - Zed")” を参照のこと。
とりあえずフォントを好みに変えたかったのだが，フォント名（`buffer_font_family`）はひとつしか指定できないのかな，これ。
しょうがない。
[Bizin Gothic](https://github.com/yuru7/bizin-gothic "yuru7/bizin-gothic: Bizin Gothic は、ユニバーサルデザインフォントの BIZ UDゴシック と英文フォント Inconsolata を合成したプログラミング向けフォントです。") を入れるか（本当は日本語は明朝体で表示してほしいんだけどねぇ）。

Git のコミットやプッシュに相当するコマンドはなさそう（ターミナルで自力で起動しろってことかな）。
履歴をグラフィカルに見せる機能もなし（git blame の表示はできる）。
VSCode の拡張機能 “[Git Graph](https://marketplace.visualstudio.com/items?itemName=mhutchie.git-graph "Git Graph - Visual Studio Marketplace")” のような機能があればいいのだが [Zed] の拡張機能でもそれらしいのは見つからなかった。
サードパーティの拡張機能がもっとたくさん生えてくるといいんだろうねぇ。

Atom エディタの元ユーザとして [Zed] には期待している。
とはいえ，テキストエディタは機能的な性能以前に使ってて手に馴染むかどうかが決定的に重要なので，もうしばらく [Zed] は様子見かなぁ。

## ブックマーク

- [日本人プログラマー向けコーディングフォント「Bizin Gothic」が無償公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1596755.html)
- [Atom の落日]({{< ref "/remark/2022/06/sunsetting-atom.md" >}})

[Zed]: https://zed.dev/ "Zed - Code at the speed of thought"
[Vulkan]: https://www.vulkan.org/ "Vulkan | Cross platform 3D Graphics"
