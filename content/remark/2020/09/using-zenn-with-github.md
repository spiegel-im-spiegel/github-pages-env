+++
title = "近ごろ流行りらしい “Zenn” のアカウントを作ってみた"
date =  "2020-09-19T17:09:47+09:00"
description = "とりあえず私も Qiita に置いてる記事の一部を移行してみるか。"
image = "/images/attention/kitten.jpg"
tags = [ "web", "engineering", "github" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

きっかけは [@tsuyoshi_cho](https://qiita.com/tsuyoshi_cho "tsuyoshi_cho - Qiita") さんの

- [Gitのaliasを晒す - Qiita](https://qiita.com/tsuyoshi_cho/items/f615dbd4631957334ef3)

で，最近の更新に

{{< fig-quote type="markdown" title="Gitのaliasを晒す - Qiita" link="https://qiita.com/tsuyoshi_cho/items/f615dbd4631957334ef3" >}}
Zennへ移植改訂しました。<br>
https://zenn.dev/tsuyoshicho/articles/git-aliases-revised
{{< /fig-quote >}}

とあって「[Zenn] ってなんじゃら？」と調べてみた。

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">プログラマーのための新しい情報共有サービス「Zenn」をリリースしました🎉<br><br>コンセプトは、有益な情報を共有する人がもっと対価を得られること。<br>noteのようにお互いを金銭的にサポートしたり、知見を本にまとめて販売したりできるプラットフォームです。<a href="https://t.co/l0lxlW24Ug">https://t.co/l0lxlW24Ug</a> <a href="https://t.co/qrddHoCWsr">pic.twitter.com/qrddHoCWsr</a></p>&mdash; catnose (@catnose99) <a href="https://twitter.com/catnose99/status/1306160371468627968?ref_src=twsrc%5Etfw">September 16, 2020</a></blockquote>
{{< /fig-gen >}}

おおっ！ 最近 launch したサービスだったのか。

パッと見の印象は「Qiita ＋ note」という感じ。
今や [note が出版社御用達のプラットフォームになっている](https://yamdas.hatenablog.com/entry/20200914/publishers-note "出版社のnote活用事例まとめ完全版（2020年9月時点） - YAMDAS現更新履歴")ように [Zenn] は（決済可能な）エンジニア御用達のプラットフォームになれればいいねぇ。

というわけで，とりあえずアカウントを作ってみた。

- [Spiegel | Zenn](https://zenn.dev/spiegel)

日本のサービスは `spiegel` 名義でアカウントが取れるのが素敵♡

ただ，決済情報は未入力のままにしてある。
できればクレカや銀行口座より PayPal 決済に対応して欲しい。
[達人出版会](https://tatsu-zine.com/)も PayPal 決済だし，技術系のサービスなんだからその方がいいと思うのだが...

## “Tech” と “Idea” という色分け

[Zenn] では “Tech” と “Idea” の2つの固定カテゴリが用意されていて，投稿する記事は必ずどちらかのカテゴリに含まれる。
[説明](https://zenn.dev/tech-or-idea "投稿カテゴリー「Tech」「Idea」の選び方 | Zenn")によると

{{< fig-quote type="markdown" title="投稿カテゴリー「Tech」「Idea」の選び方 | Zenn" link="https://zenn.dev/tech-or-idea" class="nobox" >}}
Tech
: {{% quote %}}プログラミングやソフトフェア開発、インフラなどに関する技術記事ならTechを選びます{{% /quote %}}

Idea
: {{% quote %}}個人的な意見やポエム、キャリアについての記事、キュレーション記事ならIdeaを選びます{{% /quote %}}
{{< /fig-quote >}}

なんだそうだ。
Qiita で技術記事と所謂「ポエム」が入り混じって出てくる状況を見れば妥当な措置だろう。
まぁ，悩んだら “Idea” にするのがいいんだろうねぇ。

## [GitHub] との連携

[Zenn] では [GitHub] リポジトリから記事を deploy することができる。
具体的な手順は以下の記事を参照のこと。

- [GitHubリポジトリでZennのコンテンツを管理する | Zenn](https://zenn.dev/zenn/articles/connect-to-github)

ただし，いくつか制限があって

- リポジトリ上の記事を削除しても [Zenn] に反映されない
- 一度 [Zenn] に deploy された記事の slug は変更できない（別の記事として扱われる）
- 既に [Zenn] でオン書きしたコンテンツは [GitHub] に反映されない

ようだ。

[GitHub] でリポジトリを作る際は，リポジトリ名は任意だが， `.gitignore` や `README.md` は作成しなくてよい。
これらは後述する `zenn-cli` ツールで作成される。

## zenn-cli ツールの導入

まずは [Ubuntu] 環境に node.js をインストールしてしまおう（まだ導入していない場合）。
こんな感じでいいだろう。

```text
curl -sL https://deb.nodesource.com/setup_current.x | sudo -E bash -
sudo apt install -y nodejs
```

次に，作成した [GitHub] リポジトリを適当な場所に `git clone` し，リポジトリのあるディレクトリに移動する。
まずは npm の初期化から。

```text
$ cd ~/workspace
$ gh repo clone spiegel-im-spiegel/zenn-docs
$ cd zenn-docs
$ npm init --yes
Wrote to /home/username/workspace/zenn-docs/package.json:

{
  "name": "zenn-docs",
  "version": "1.0.0",
  "description": "## Links",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "repository": {
    "type": "git",
    "url": "git+https://github.com/spiegel-im-spiegel/zenn-docs.git"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "bugs": {
    "url": "https://github.com/spiegel-im-spiegel/zenn-docs/issues"
  },
  "homepage": "https://github.com/spiegel-im-spiegel/zenn-docs#readme"
}
```

`package.json` は弄らなくて大丈夫。
続けて `zenn-cli` ツールをインストールする。

```text
$ npm install zenn-cli
...
+ zenn-cli@0.1.29
added 5 packages from 3 contributors, removed 3 packages, updated 3 packages, moved 1 package and audited 905 packages in 42.516s
found 5 low severity vulnerabilities
  run `npm audit fix` to fix them, or `npm audit` for details
```

なんか不穏なメッセージが見えるが大丈夫か，これ。 ...まぁいいや[^issue1]。
次いってみよう。

[^issue1]: [Zenn] 側に問い合わせてみたが，セキュリティ上の問題はないらしい。「あまり精神的に良いものではないのでそのうちなんとかしようと思っています」とのこと（笑）

`zenn-cli` ツールがインストールできたらリポジトリ内を初期化する。

```text
$ npx zenn init

  🎉Done!
  早速コンテンツを作成しましょう

  👇新しい記事を作成する
  $ zenn new:article

  👇新しい本を作成する
  $ zenn new:book

  👇表示をプレビューする
  $ zenn preview
```

これでリポジトリ内に `articles/` および `books/` ディレクトリが作成され，さらに `.gitignore` および `README.md` ファイルも作成される。
ちなみに `.gitignore` の中身はこんな感じ。

```text
node_modules
.DS_Store
```

何ともシンプルだが，これで `zenn-cli` インストール時に作成される `node_modules/` ディレクトリはリポジトリから除外される。

ここまで出来たら一度コミットしておいたほうがいいだろう。

## 記事の作成

入力ファイルの作成には以下のコマンドを打つ。

```text
$ npx zenn new:article
📄d309af5057a827deda35.md created.
```

このファイル名がそのまま slug として URL のパスになる。
Slug は `zenn-cli` ツールが適当に生成するのでユーザは考えなくともよい。

もし slug を指定したいのであれば `--slug` オプションを使う。

```text
$ npx zenn new:article --slug hello-zenn-world
📄hello-zenn-world.md created.
```

ただし slug 文字列には以下の制限がある。

- 半角英数字（a-z, 0-9）とハイフン（-）の 12〜50 字の組み合わせのみ有効
- `articles` 以下のファイルはディレクトリ階層に出来ない（フラットな構成）
- `books` の場合は「本」ごとに slug を指定できる。本の slug 以下はフラットな構成

Slug 文字列が短いとエラーになるのでご注意を。

```text
$ npx zenn new:article --slug hello
エラー：slugの値（hello）が不正です。半角英数字（a-z0-9）とハイフン（-）の12〜50字の組み合わせにしてください
```

作成したファイルの中身は，以下のように front matter のひな型のみが書き込まれている。

```yaml
---
title: ""
emoji: "🎉"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: []
published: true
---
```

`emoji` 項目は，記事のアテンションに使われるのだが，毎回ランダムな文字で初期化されるようだ。
絵文字を直接入力することはないのだが [GitHub markdown](https://www.webfx.com/tools/emoji-cheat-sheet/ "🎁 Emoji cheat sheet for GitHub, Basecamp, Slack & more") のように文字列で指定できないものかねぇ。

## プレビューが素敵！

{{< div-box type="markdown" >}}
**【2020-09-20 更新】**
以前プレビュー機能が動かないと書いたが，他でも [Issue が上がっていた](https://github.com/zenn-dev/zenn-roadmap/issues/37 "npx zenn preview でエラーが大量に吐き出される · Issue #37 · zenn-dev/zenn-roadmap")らしく，対応されていた。
感謝！

この節は，以前の内容から差し替えている。
なお，障害報告は [GitHub] の [zenn-dev/zenn-roadmap/issues](https://github.com/zenn-dev/zenn-roadmap/issues) に上げてほしいとのこと。

[GitHub]: https://github.com/
{{< /div-box >}}

以下のコマンドでプレビュー用のローカルサーバが起動する。

```text
$ npx zenn preview
👀Preview on http://localhost:8000
```

画面はこんな感じ。

{{< fig-img src="./preview.png" link="./preview.png" width="1181" >}}

おおっ！ 間違いまで指摘してくれるのか。
こりゃあ，ええ。

エディタの markdown プレビュー機能でもある程度は見れるけど，やっぱサービス立ち上げてブラウザで確認できるのがいいよね。

うんうん。

<!--
プレビュー用のローカルサーバを起動しようとしたら

```text
$ npx zenn preview
(node:126485) UnhandledPromiseRejectionWarning: Error: Could not find a valid build in the '/home/username/workspace/zenn-docs/.next' directory! Try building your app with 'next build' before starting the server.
    at Server.readBuildId (/home/username/workspace/zenn-docs/node_modules/next/dist/next-server/server/next-server.js:113:355)
    at new Server (/home/username/workspace/zenn-docs/node_modules/next/dist/next-server/server/next-server.js:3:120)
    at Object.createServer [as default] (/home/username/workspace/zenn-docs/node_modules/next/dist/server/next.js:2:638)
    at /home/username/workspace/zenn-docs/node_modules/zenn-cli/dist/cli/preview/build.js:53:41
    at step (/home/username/workspace/zenn-docs/node_modules/zenn-cli/dist/cli/preview/build.js:33:23)
    at Object.next (/home/username/workspace/zenn-docs/node_modules/zenn-cli/dist/cli/preview/build.js:14:53)
    at /home/username/workspace/zenn-docs/node_modules/zenn-cli/dist/cli/preview/build.js:8:71
    at new Promise (<anonymous>)
    at __awaiter (/home/username/workspace/zenn-docs/node_modules/zenn-cli/dist/cli/preview/build.js:4:12)
    at Object.exports.build (/home/username/workspace/zenn-docs/node_modules/zenn-cli/dist/cli/preview/build.js:48:12)
(Use `node --trace-warnings ...` to show where the warning was created)
(node:126485) UnhandledPromiseRejectionWarning: Unhandled promise rejection. This error originated either by throwing inside of an async function without a catch block, or by rejecting a promise which was not handled with .catch(). To terminate the node process on unhandled promise rejection, use the CLI flag `--unhandled-rejections=strict` (see https://nodejs.org/api/cli.html#cli_unhandled_rejections_mode). (rejection id: 3)
(node:126485) [DEP0018] DeprecationWarning: Unhandled promise rejections are deprecated. In the future, promise rejections that are not handled will terminate the Node.js process with a non-zero exit code.
```

てな感じに例外を吐いて落っこちた。

まぁ，プレビューなくても問題はないのだが...
-->

## とりあえず私も...

まずは Qiita に置いてる記事の一部を移行してみるか。
古すぎて使えない記事はダメだけど（笑）

## ブックマーク

- [GitHubリポジトリでZennのコンテンツを管理する | Zenn](https://zenn.dev/zenn/articles/connect-to-github)
- [Zenn CLIをインストールする | Zenn](https://zenn.dev/zenn/articles/install-zenn-cli)
- [Zenn CLIを使ってコンテンツを作成する | Zenn](https://zenn.dev/zenn/articles/zenn-cli-guide)
- [ZennのMarkdown記法 | Zenn](https://zenn.dev/zenn/articles/markdown-guide)

- [GitHub CLI 1.0 is now available - The GitHub Blog](https://github.blog/2020-09-17-github-cli-1-0-is-now-available/)
    - [GitHub CLI 1.0 がリリースされたので使ってみた | Zenn](https://zenn.dev/morinokami/articles/hello-github-cli)
- [Ubuntu/Debianに最新のNode.jsをインストールする一番良い方法 | LFI](https://linuxfan.info/install_nodejs_on_ubuntu_debian)
- [🎁 Emoji cheat sheet for GitHub, Basecamp, Slack & more](https://www.webfx.com/tools/emoji-cheat-sheet/)

[Zenn]: https://zenn.dev/ "Zenn｜プログラマーのための情報共有コミュニティ"
[GitHub]: https://github.com/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
