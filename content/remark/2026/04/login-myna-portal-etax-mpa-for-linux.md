+++
title = "MPA for Linux でログイン検証（Linux で個人番号カードを読む 2）"
date =  "2026-04-01T15:25:15+09:00"
description = "これで来年は自宅 Linux 機で確定申告できるな。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "my-number", "tools", "install", "web", "pki", "authentication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[前回]の続き。

[myna のリポジトリ][myna]によると，サブプロジェクトの [MPA for Linux] を使って Linux の Web ブラウザでマイナポータルや e-Tax のサイトに個人番号カードを使ってログインできるらしい。
素晴らしい！

## Rust ツールチェーンのインストール

事前準備として [Rust] ツールチェーンのインストールを行う。

[Rust] の[基礎勉強]({{< rlnk "/rust-lang/" >}})をしてたのはもう6年も前で，仕事に結びつくこともなかったので完全に放置していた。
しかも，あれから自宅パソコンを[買い替え]({{< ref "/remark/2021/06/new-machine-here.md" >}} "自宅マシンを買うた（これで私も人並みに...）")たりして開発環境もなくなったので，インストールからやり直すことに。

[インストールページ](https://rust-lang.org/tools/install/ "Install Rust - Rust Programming Language")に従って，以下のスクリプトをダウンロード&実行する。

```text
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
info: downloading installer

Welcome to Rust!

This will download and install the official compiler for the Rust
programming language, and its package manager, Cargo.

...

   default host triple: x86_64-unknown-linux-gnu
     default toolchain: stable (default)
               profile: default
  modify PATH variable: yes

1) Proceed with standard installation (default - just press enter)
2) Customize installation
3) Cancel installation
>
```

今回はこのまま {{% keytop %}}Enter{{% /keytop %}} キーを押して続行。

```text
info: profile set to default
info: default host triple is x86_64-unknown-linux-gnu
info: syncing channel updates for stable-x86_64-unknown-linux-gnu
info: latest update on 2026-03-26 for version 1.94.1 (e408947bf 2026-03-25)

...

Rust is installed now. Great!

To get started you may need to restart your current shell.
This would reload your PATH environment variable to include
Cargo's bin directory ($HOME/.cargo/bin).

To configure your current shell, you need to source
the corresponding env file under $HOME/.cargo.

This is usually done by running one of the following (note the leading DOT):
. "$HOME/.cargo/env"            # For sh/bash/zsh/ash/dash/pdksh
source "$HOME/.cargo/env.fish"  # For fish
source "~/.cargo/env.nu"  # For nushell
source "$HOME/.cargo/env.tcsh"  # For tcsh
. "$HOME/.cargo/env.ps1"        # For pwsh
source "$HOME/.cargo/env.xsh"   # For xonsh
```

これで `~/.rustup/` および `~/.cargo/` ディレクトリ以下にツールチェーンがインストールされた。
`PATH` 設定が `~/.cargo/env` ファイルに記述されていて `~/.profile` と `~/.bashrc` が `~/.cargo/env` を読み込むよう書き換えられている。
必要に応じて内容を調整する。

とりあえず，今すぐ `PATH` を通したいなら

```text
$ . ~/.cargo/env
```

とすればよい[^b1]。

[^b1]: コマンドの `.` は `source` と同じ意味で，指定したファイルを現在の shell で実行する。

起動確認しておこう。

```text
$ rustc --version
rustc 1.94.1 (e408947bf 2026-03-25)
```

うんうん。
問題なさそうやね。
[Rust] の勉強もやり直すかなぁ。

## MPA for Linux のインストール

いよいよ [MPA for Linux] をインストールする。

適当なディレクトリに [`github.com/jpki/myna`][myna] リポジトリを clone する。

```text
$ git clone https://github.com/jpki/myna.git
```

リポジトリ内の `mpa` ディレクトリに移動して `cargo install` を実行する。

```text
$ cd myna/mpa
$ cargo install --path .
  Installing mpa v23.0.0 (/home/username/path/to/myna/mpa)
    Updating crates.io index
     Locking 107 packages to latest Rust 1.94.1 compatible versions
      Adding der v0.7.10 (available: v0.8.0)
      Adding generic-array v0.14.7 (available: v0.14.9)
      Adding sha1 v0.10.6 (available: v0.11.0)
      Adding sha2 v0.10.9 (available: v0.11.0)

         ...

   Compiling myna v0.6.4 (/home/username/path/to/myna)
   Compiling mpa v23.0.0 (/home/username/path/to/myna/mpa)
    Finished `release` profile [optimized] target(s) in 16.86s
  Installing /home/username/.cargo/bin/mpa
   Installed package `mpa v23.0.0 (/home/username/path/to/myna/mpa)` (executable `mpa`)
```

これで `~/.cargo/bin/` ディレクトリにホストアプリケーション `mpa` がインストールされた。

次に同ディレクトリにある `install.sh` を実行してホストアプリケーションをブラウザに登録する。
ブラウザのプロファイルが既定の場所にあるのであれば，引数なしで起動すればよい。

```text
$ ./install.sh
=== Installing Native Messaging Host manifests ===
Host path: /home/username/.cargo/bin/mpa

Installed: /home/username/.config/google-chrome/NativeMessagingHosts/com.github.jpki.mpa.json
Installed: /home/username/.config/chromium/NativeMessagingHosts/com.github.jpki.mpa.json
Installed: /home/username/.mozilla/native-messaging-hosts/com.github.jpki.mpa.json
```

ブラウザのプロファイルが既定のの場所にない場合は

```text
./install.sh --user-data-dir /path/to/datadir
```

のように指定できるらしい。
今回は既定のままでOK。

次にブラウザのほうほうにも拡張機能をインストールする必要があるのだが，正規ルートからはインストールできないようなので Developer mode で強制的に行う。
手順は以下の通り。

{{< fig-quote type="markdown" title="ブラウザ拡張のインストール" link="https://github.com/jpki/myna/tree/master/mpa#%E3%83%96%E3%83%A9%E3%82%A6%E3%82%B6%E6%8B%A1%E5%BC%B5%E3%81%AE%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB" >}}
### Chrome
- `chrome://extensions/`を開く
- 右上のディベロッパーモードをON
- `パッケージ化されていない拡張機能を読み込む`で`./mpa/extension`を読み込む
- 拡張機能のメニューから`MPA for Linux`を開く
- 動作確認ボタンを押してエラーが出なければOK

### Firefox(一時的)
- `about:debugging`の`このFirefox`を開く
- `一時的なアドオンを読み込む`で`./mpa/extension/manifest.json`を読み込む
- 拡張機能のメニューから`MPA for Linux`を開く
- 動作確認ボタンを押してエラーが出なければOK
{{< /fig-quote >}}

メインで使ってる Firefox に入れるのは怖いので，サブとして使ってる [ungoogled-chromium] で試した。
やり方はたぶん Chrome と同じでいいよね。
こんな感じ？

{{< fig-img-quote src="./chromium-extensions.png" link="./chromium-extensions.png" width="761" >}}

ちなみに Developer mode を OFF にするとこの拡張機能も無効になる。
これがあるからオススメしにくいんだよなぁ。

これで [MPA for Linux] の導入は完了。
上手くログインできるかなぁ。

## マイナポータルサイトにログインする

{{< fig-img-quote src="./mp-top.png" title="マイナポータル" link="https://myna.go.jp/" width="1030" >}}

左サイドにあるログインボタンを押してログイン画面に移動する。

{{< fig-img-quote src="./mp-login.png" title="ログイン | マイナポータル" link="./mp-login.png" width="1030" >}}

ここで暗証番号の入力を求められる。

{{< fig-img-quote src="./mpa-auth.png" title="暗証番号入力" link="./mpa-auth.png" >}}

利用者証明用パスワード（4文字の数字）を入力して {{% keytop %}} OK {{% /keytop %}} ボタンをクリックする。
ボタンをクリックせず {{% keytop %}}Enter{{% /keytop %}} キーを押すとなにも起こらず処理が止まってしまうので注意（困るなぁ）。

{{< fig-img-quote src="./mp-main.png" title="ホーム | マイナポータル" link="https://myna.go.jp/" width="1030" >}}

よし。
上手くいった！

## e-Tax サイトにログインする

[e-Tax] サイトのログインは個人用と法人用がある。
私は個人用からログインする。

{{< fig-img-quote src="./etax-indilogin.png" title="個人ログイン | e-Tax" link="https://login.e-tax.nta.go.jp/login/reception/loginIndividual" width="1030" >}}

下の方にスクロールすると

{{< fig-img-quote src="./etax-indilogin-2.png" title="個人ログイン | e-Tax" link="https://login.e-tax.nta.go.jp/login/reception/loginIndividual" width="1030" >}}

「ICカードリーダーで読み取り」ボタンがあるので，これをクリックする。
あとは前節と同じように暗証番号の入力を求められるので，利用者証明用パスワード（4文字の数字）を入力して {{% keytop %}} OK {{% /keytop %}} ボタンをクリックする。

{{< fig-img-quote src="./etax-main.png" title="TOP | e-Tax" link="https://mypage.e-tax.nta.go.jp/" width="1030" >}}

こちらも問題なく入れた！

## これで Linux 機で確定申告できる！

[マイナポータル]および [e-Tax] の両サイトへのログインを確認できたので，ブラウザ拡張機能の Developer mode を OFF に戻しておく。

これで来年は自宅 Linux 機で確定申告できるな。
もうスマホで確定申告するのは嫌なのよ。
スマホは入力端末としては向かないっスよ。

[ミニ PC]({{< ref "/remark/2025/01/win11pro-on-minipc.md" >}} "Mini PC を衝動買いした") の Windows 機はますますゲーム専用機になっていくな（笑） まぁ，それはそれで重宝しているからいいか。

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

[前回]: {{< ref "/remark/2026/03/read-individual-number-card-on-linux.md" >}} "Linux で個人番号カードを読む"
[myna]: https://github.com/jpki/myna "jpki/myna: マイナンバーカード・ユーティリティ・JPKI署名ツール · GitHub"
[MPA for Linux]: https://github.com/jpki/myna/tree/master/mpa "MPA for Linux"
[Rust]: https://rust-lang.org/ "Rust Programming Language"
[ungoogled-chromium]: https://github.com/ungoogled-software/ungoogled-chromium "ungoogled-software/ungoogled-chromium: Google Chromium, sans integration with Google"
[マイナポータル]: https://myna.go.jp/ "マイナポータル"
[e-Tax]: https://www.e-tax.nta.go.jp/ "【e-Tax】国税電子申告・納税システム(イータックス)"

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
