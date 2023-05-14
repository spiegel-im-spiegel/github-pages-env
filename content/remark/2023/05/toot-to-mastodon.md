+++
title = "コマンドラインで Mastodon に投稿する"
date =  "2023-05-14T13:20:56+09:00"
description = "これで出力の基本機能はできたので，今後は入力側の設計と実装だな。先は長い。"
image = "/images/attention/kitten.jpg"
tags = [ "golang", "tools", "mastodon" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今年の GW は（色々と大人の事情で）遠くに遊びに行くこともなく，ときどき自転車を乗り回す以外はおうちでコードを書くか，その辺に寝転がってラノベか Web 小説を読み耽るという，自堕落な生活を送っていた。

で，まぁ，その成果として以下の CLI ツールをリリースした。

- [goark/toolbox: A collection of miscellaneous commands][toolbox]

今のところ Mastodon と Bluesky への投稿機能のみ実装している。
最終的には自コンテンツ（主にブログ記事と Flickr 写真）の更新情報の投稿を自動化できればと思っているが，先は長い。

このうち今回は Mastodon への投稿機能について，覚え書きを兼ねて，記しておく。

今どきの流行りは Bluesky のほうぢゃねーのかとお思いでしょうが，利用している[公式の Go 用パッケージ](https://github.com/bluesky-social/indigo "bluesky-social/indigo: Go source code for Bluesky's atproto services. NOT STABLE (yet)")の作りが微妙で，ちょっとしたことでエラーを吐くためイマイチな出来。
みんな，あんな雑なエラーハンドリングでよく使えてるよなぁ。
とりあえず 400 を返せばいいみたいなのはどうにかしてほしい。
[大きな画像ファイルをアップロードしてエラーを返さない](https://zenn.dev/username/articles/20230506-downsizing-images "画像ファイルのサイズを縮小したい")のはもっと困るけど。

話がそれた。
では本題へ。

## アプリケーションの登録

Mastodon のハンドリングには以下のパッケージを使わせてもらっている。
ありがたや {{% emoji "ペコン" %}}

- [mattn/go-mastodon: mastodon client for golang][mattn/go-mastodon]

だがしかし，最初の「認証」部分でいきなり躓く。
あちこちの解説ページを覗いてみるに，どうやら Mastodon の認証は2段階あるらしい。

1. サーバに対してアプリケーション登録を行う
2. ユーザ認証を行い，アカウントに対してアプリケーションを認証する

1 を行うとクライアントIDとそのシークレットがもらえる。
1 でもらった情報を使って 2 を行うとアクセストークンがもらえる。
さらに 2 の認証にはパスワード認証と OAuth 認証の2つが用意されていて，後者についてはブラウザ操作が必要（？）
アクセストークンは永続的に有効らしい。
1 と 2 で取得した情報を使って Mastodon サーバに対して各種操作を行える。
 ...という感じ。

ここまで辿り着くのにかなり試行錯誤してしまった。

拙作 [toolbox] では `mastodon register` コマンドで 1, 2 の処理をまとめて行う。

```text
$ toolbox mastodon register -h
Register Mastodon application.

Usage:
  toolbox mastodon register [flags]

Aliases:
  register, reg

Flags:
  -h, --help   help for register

Global Flags:
      --bluesky-config string    Config file for Bluesky (default "/home/username/.config/toolbox/bluesky.json")
      --cache-dir string         Directory for cache files (default "/home/username/.cache/toolbox")
      --config string            Config file (default "/home/username/.config/toolbox/config.yaml")
      --debug                    for debug
      --log-dir string           Directory for log files (default "/home/username/.cache/toolbox")
      --log-level string         Log level [nop|error|warn|info|debug|trace] (default "nop")
      --mastodon-config string   Config file for Mastodon (default "/home/username/.config/toolbox/mastodon.json")
```

取得したアクセストークンは `--mastodon-config` オプションで指定したファイルに JSON 形式で保存される。
具体的には以下のように，サーバ名，ユーザID，パスワードを尋ねるプロンプトが表示されるので入力していけば OK。

```text
$ toolbox mastodon register
Server (e.g. mastodon.social) > fedibird.com
         User (email address) > jphn.do@exsample.com
                     Password > your_password

          server: https://fedibird.com
application name: github.com/goark/toolbox
         website: https://github.com/goark/toolbox
          scopes: read write follow

output: /home/username/.config/toolbox/mastodon.json
```

成功すれば 設定＞アカウント＞認証済みアプリ に認証されたアプリケーションが表示される（表示の仕方はサーバによって違うかも）。

{{< fig-img src="./authorized-apps.png" title="認証済みアプリ" link="./authorized-apps.png" width="1091" >}}

動作確認のために自身の profile を表示してみる。

```text
$ toolbox mastodon profile
      Username: spiegel
User ID (full): @spiegel@fedibird.com
           URL: https://fedibird.com/@spiegel
  Display name: Spiegel@がんばらない
    Created at: 2022-11-25 00:00:00 +0000 UTC
         Posts: 2289
       Follows: 46
     Followers: 98

<p>mstdn.jp から移住。職業プログラマ。<a href="https://fedibird.com/tags/golang" class="mention hashtag" rel="tag">#<span>golang</span></a> と <a href="https://fedibird.com/tags/flickr" class="mention hashtag" rel="tag">#<span>flickr</span></a> で遊んでいる人。暖かくなったので自転車で「お散歩カメラ」再開。</p><p>情報収集がメインだが最近は胡乱な発言もチラホラあるので，そこは許して（フォローの付け外しはご自由に）。フォローは特に意味なく頻繁に入れ換えますのであしからず。</p>
```

よーし，うむうむ，よーし。

## Mastodon に投稿する

以上で対象のサーバに [toolbox] でアクセスできるようになったので，さっそく何か投稿してみる。
Mastodon への投稿は `mastodon post` でできる。

```text
$ toolbox mastodon post -h
Post message to Mastodon.

Usage:
  toolbox mastodon post [flags]

Aliases:
  post, pst, p, toot, tt, t

Flags:
      --edit                  Edit message
  -h, --help                  help for post
  -i, --image-file strings    Image file
  -m, --message string        Message
      --pipe                  Input from standard-input
  -s, --spoiler-text string   Spoiler text
  -v, --visibility string     Visibility [public|unlisted|private|direct] (default "public")

Global Flags:
      --bluesky-config string    Config file for Bluesky (default "/home/username/.config/toolbox/bluesky.json")
      --cache-dir string         Directory for cache files (default "/home/username/.cache/toolbox")
      --config string            Config file (default "/home/username/.config/toolbox/config.yaml")
      --debug                    for debug
      --log-dir string           Directory for log files (default "/home/username/.cache/toolbox")
      --log-level string         Log level [nop|error|warn|info|debug|trace] (default "nop")
      --mastodon-config string   Config file for Mastodon (default "/home/username/.config/toolbox/mastodon.json")
```

- `-m` はコマンドライン上で1行メッセージを投稿するのに使うオプション
- `--pipe` は標準入力からの入力を投稿するオプション
- `--edit` は CUI で複数行編集ができるオプション
- `-m`, `--pipe`, `--edit` は排他オプションで同時に指定できない
- `-i` は画像ファイルをアップローするのに使うオプション。複数指定可能
- `-v` は表示範囲を指定するオプション

たとえばこんな感じで投稿する。

```text
$ toolbox mastodon post --edit -i lake-shinjiko.jpg -v direct
Input 'Ctrl+J' or 'Ctrl+Enter' to submit message
Input 'Ctrl+D' with no chars to stop
 1>はろー，ふぇでぃばーす！
 2>https://flic.kr/p/2nSUmaa
 3>
https://fedibird.com/@spiegel/110364957384850439
```

するとこんな感じに表示される。

{{< fig-img src="./toot-to-mastodon.png" title="Mastodon へ投稿" link="./toot-to-mastodon.png" width="627" >}}

よしよし。
ちなみに複数行編集には以下のパッケージを利用している。

- [hymkor/go-multiline-ny: Readline package supporting multi-lines](https://github.com/hymkor/go-multiline-ny)

マジ便利。
ありがたや {{% emoji "ペコン" %}}

Mastodon は表示範囲を指定できるのがいいよね。
とりあえず試し撃ちなら DM で投げればいい。
Bluesky は（多分まだ）表示範囲を指定できないのでテスト用のゴミ投稿も全部 TL に表示されてしまう。
しょうがないから招待コード使ってデバッグ用のアカウントをひとつ確保する羽目になった。

さて，これで出力の基本機能はできたので，今後は入力側の設計と実装だな。
先は長い。

## ブックマーク

- [マストドンのタイムラインをgo-mastodon のWebSocketを使用し取得する - Qiita](https://qiita.com/S-YOU/items/cf677ae282bd6f38fbbb)
- [mastodonのaccess tokenをauthorization_codeで取得する例 - Qiita](https://qiita.com/civic/items/7358dc1c54ff8e71c326)
- [Mastodon API の叩き方 · GitHub](https://gist.github.com/okapies/eab5c6fc217e914ed0cac6c944384e4d)
- [mastodonのtootを原始人くんが喋ってくれるbot作った - Qiita](https://qiita.com/shinderuman@github/items/c96161caa65c6a9e8ffc)

- [Fediverse 関連のブックマーク]({{< ref "/bookmarks/fediverse.md" >}})

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->

[toolbox]: https://github.com/goark/toolbox "goark/toolbox: A collection of miscellaneous commands"
[mattn/go-mastodon]: https://github.com/mattn/go-mastodon "mattn/go-mastodon: mastodon client for golang"
<!-- eof -->
