+++
title = "Mastodon と GitHub との連携"
date =  "2023-02-02T09:52:56+09:00"
description = "“Social accounts” の項目に URL をセットすればいいのか"
image = "/images/attention/kitten.jpg"
tags = [ "mastodon", "github" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[以前にも紹介した]({{< ref "/remark/2022/11/some-little-more-mastodon.md" >}} "もうちょこっと Mastodon")が， Mastodon と他サイトを連携させるには， Mastodon 側のプロフィール編集の「プロフィール補足情報」で

{{< fig-img src="./fedibird-profile.png" title="プロフィールを編集 - Fedibird" link="./fedibird-profile.png" width="812" >}}

こんな風に URL を列挙し[^p1]，対応するサイトで

[^p1]: Mastodon の標準では補足情報の URL は4つまで登録できるが Fedibird は8つまで拡張されている。

```html
<a rel="me" href="https://hostname/@username">Mastodon</a>
```

とか

```html
<link rel="me" href="https://hostname/@username">
```

とかいった感じのリンクを記述すればいい。
のだが， GitHub のような SaaS ではサービス側が対応してくれないと難しかったりする。

GitHub の場合 `github.com/username/username` リポジトリを作ってその中の `README.md` ファイルに任意のリンクを載せられるのだが

- [GitHub プロファイルを（ちょっとだけ）カッコよくしてみる]({{< ref "/remark/2020/09/using-github-profile-readme.md" >}})

外部サイトへのリンクには強制的に `rel="nofollow"` が上書き設定されてしまうみたいなのね。
まぁ，気持ちは分かるので「しょうがないか」と諦めていたのだが， Mastodon の TL で

{{< fig-gen >}}
<iframe src="https://hachyderm.io/@nova/109790530971147702/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe>
{{< /fig-gen >}}

という投稿を見かけたので真似してみることにした。

具体的には GitHub の自ユーザページの “Edit profile” ボタンを押して編集モードにし， “Social accounts” の項目で

{{< fig-img src="./github-profile.png" title="Edit profile - GitHub" link="./github-profile.png" >}}

という感じに Mastodon のプロフィール・ページの URL をセットすればよい。
`mstdn.jp` はアイコンが {{< icons "mastodon" >}} に変わるのに `fedibird.com` は変わらないんだな。
...まぁいいか。

この設定により GitHub の自ユーザページに

```html
<a rel="nofollow me" class="Link--primary" href="https://fedibird.com/@spiegel">https://fedibird.com/@spiegel</a>
```

あるいは

```html
<a rel="nofollow me" class="Link--primary" href="https://mstdn.jp/@spiegel">@spiegel@mstdn.jp</a>
```

といった感じのリンクが張られる。
一方で Mastodon の自プロフィールページには

{{< fig-img src="./my-profile.png" title="プロフィール - Fedibird" link="./my-profile.png" >}}

という感じに GitHub の URL にチェックマークが付く。

めでたし！


<!-- eof -->
