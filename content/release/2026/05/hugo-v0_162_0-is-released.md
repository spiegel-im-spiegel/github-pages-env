+++
title = "Hugo v0.162.0 がリリースされた【Security fix あり】"
date =  "2026-05-27T16:01:18+09:00"
description = "このバージョンから既定で HTML フォーマットの入力ファイルを拒否するようになった。"
isCJKLanguage = true
image = "/images/attention/tools.png"
tags  = [ "tools", "site", "hugo", "security", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Hugo] v0.162.0 がリリースされた。

- [Release v0.162.0 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.162.0)

[Go] コンパイラおよび関連パッケージのセキュリティ・アップデートがあるのはいつものことなのだが，今回は注意すべき修正がある。

{{< fig-quote type="markdown" title="Release v0.162.0 · gohugoio/hugo" link="https://github.com/gohugoio/hugo/releases/tag/v0.162.0" lang="en" >}}
**Security fixes and hardening in Hugo**

The following changes either fix a concrete issue or reduce the default attack surface of `hugo` builds.

- Disallow `text/html` content files by default ([e41a064](https://github.com/gohugoio/hugo/commit/e41a06447d)). A new `security.allowContent` policy gates which content media types may be used for pages under `/content`. `text/html` is denied by default; sites that rely on hand-authored or adapter-emitted HTML content can opt back in with `security.allowContent = ['.*']`.
- Re-check `security.http.urls` on every redirect hop in `resources.GetRemote` ([86fbb0f](https://github.com/gohugoio/hugo/commit/86fbb0f7a8)).
- Reject symlinked entries in `resources.Get` ([f8b5fa0](https://github.com/gohugoio/hugo/commit/f8b5fa09a6)).
{{< /fig-quote >}}

このうち，最初の項目にまんまと引っかかった。
こんな感じ：

```text {hl_lines=[7]}
ERROR error building site: assemble: failed to create page from pageMetaSource 
/remark/2015/gnupg-2-1-8: "/path/to/content/remark/2015/gnupg-2-1-8.html:1:1": 
access denied: "text/html" is not whitelisted in policy "security.allowContent"; 
the current security configuration is:

[security]
  allowContent = ['! ^text/html$']
  enableInlineShortcodes = false
  ...

```

どうも既定では HTML ファイルを拒否する設定になったらしい。
それならば HTML を許可すればいいだろうと， “`!`” を取って

```toml
[security]
  allowContent = ["^text/html$"]
```

と `hugo.toml` に書いたら markdown ファイルが全拒否になってしまった（笑）。HTML と markdown のみを許可するには

```toml
[security]
  allowContent = ["^text/markdown$", "^text/html$"]
```

などとする必要がある。
ちなみに [Hugo] では

{{< fig-quote type="markdown" title="Content formats" link="https://gohugo.io/content-management/formats/" lang="en" >}}
Create your content using Markdown, HTML, Emacs Org Mode, AsciiDoc, Pandoc, or reStructuredText.
{{< /fig-quote >}}

と多様なフォーマットに対応している。
この中で HTML 形式を受け入れるのは（malware を埋め込まれたりとか）リスクが高いと判断して安全側に倒したものと思われる。
実運用では `security.allowContent` に，実際に使用するドキュメントフォーマットのみを列挙するのが正しいのだろう。

エラーの出た `gnupg-2-1-8.html` はこのブログにおける[最初の記事]({{< ref "/remark/2015/gnupg-2-1-8.html" >}} "GnuPG 2.1.8 released")で，元々は Movable Type 用に書いた記事で「そのまま読み込ませてコンパイルできるかなぁ」と試してみた記事だった。
記事本文の HTML 記述はともかく front matter を整えるのが面倒くさそうなので旧ブログからの移行は諦めたが。
今なら変換ツールを AI に書かせたんだろうなぁ（笑）

最近は GitHub Copilot に [Hugo] のアップデートと redeploy をさせているのだが，今回の障害の分析と対処はすべてお任せで作業を完遂できた。
「Hugo のマイナーバージョンアップは問題が出やすいから気をつけて」と指示に添えたのが功を奏したようだ。
自力でやったら悩みまくって，[ドキュメント](https://gohugo.io/documentation/ "Hugo Documentation")を調べまくって30分くらい時間を無駄にしてただろう。
いい時代になったものである。

[Hugo]: https://gohugo.io/ "The world's fastest framework for building websites"
[Go]: https://go.dev/
