+++
title = "「マガジン航」サイト終了のお知らせ"
date =  "2026-10-07T12:23:27+09:00"
description = "あとは magazine-k.jp ドメインが spam や phishing に転用されないよう，あるいは Internet Archive が市場や時の為政者に踏み潰されないよう祈るばかりである。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "site", "media", "internet" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[yomoyomo さんの記事](https://yamdas.hatenablog.com/entry/20261007/magazine-ko "「マガジン航」終了のお知らせに寄せて - YAMDAS現更新履歴")経由で知ったのだが

- [「マガジン航」終了のお知らせ ｜ 仲俣暁生](https://note.com/solar1964/n/n18a48f8d6412)
- [「マガジン航」終了のお知らせ](https://web.archive.org/web/20261003032524/https://magazine-k.jp/2026/10/01/cease_publication/)

えっ「[マガジン航]」って続いてたのか？ と一瞬思ったが，そういえば止めたという話も聞いてなかったな。

{{< fig-quote type="markdown" title="「マガジン航」終了のお知らせ" link="https://note.com/solar1964/n/n18a48f8d6412" >}}
しかし2023年秋に最後の記事更新をして以来、今後このメディアをどのように活かして行くべきか、編集発行人としてその方向性を見出すことができずにいました。

この間にインターネットの環境そのものが、ウェブからSNSへ、さらに生成AIへと急速にシフトしており、そのようななかで過去記事のアーカイヴをどのように管理すべきかについても検討を重ねましたが、いまの状況のもとではウェブ上にいつまでも置き続けること自体がよいことだとは思えず、サイトの終了と完全閉鎖を選ぶことにしました。
{{< /fig-quote >}}

なるほど。
というわけで，今月いっぱいでドメインごとなくなるらしい。

{{< fig-quote type="markdown" title="「マガジン航」終了のお知らせ" link="https://note.com/solar1964/n/n18a48f8d6412" >}}
これまでの記事のアーカイヴは、[Internet ArchiveのWayback Machine](https://web.archive.org/web/20260415124052/https://magazine-k.jp/)に委ねることとし、ドメイン名も含めてすべて解消いたします。

[...]

これまでにご寄稿いただいたテキストについては、ローカルファイルでもバックアップをとっておりますので、万一、Internet Archiveからのサルベージができなかった場合は、責任をもってこちらからコピーをお送りします。
{{< /fig-quote >}}

私の[本家サイト](https://baldanders.info/ "Baldanders.info")で塩漬けになっている日記や旧ブログはいまさら弄る気も起きないのだが，ここのブログは直しておいたほうがいいか。
とりあえず [GitHub Copilot] に調査と作業計画の立案をしてもらおう。
結果はこんな感じ。

{{< div-ai type="markdown" >}}
**Plan: マガジン航リンクのWayback化**

`content` 内のマガジン航リンクを，指定日時を基準にしたWayback URLへ置き換えます。現状は66件・44ファイルで，記事リンクのほか，サイトトップ・作者・カテゴリへのリンクも含まれています。

1. `content` の対象URLを確定します。生成物の `public` と，過去のDelicious書き出し資料は直接変更しません。
2. リンク先を `https://web.archive.org/web/20261003095755/` と元URLをつないだ形式に置き換えます。リンク文字列や引用文，元URLのパス・クエリ・フラグメントは保ち，Markdownリンクと `fig-quote` などの属性も対象にします。
3. Waybackで記事ページとトップ・作者・カテゴリページの保存状態を確認します。保存がない，または適切でないURLは元URLを残して例外として報告します。
4. `content` を再検索し，例外以外に元ドメインが残っていないことと，差分がリンク先の変更に限られることを確認します。
5. `build.sh` を実行し，ビルド成功と生成HTMLのリンク先を確認します。コミット・公開は行いません。公開する場合はリポジトリの作業手順に従います。
{{< /div-ai >}}

`content` は [Hugo] 環境で記事ファイルを置くディレクトリのこと。
`fig-quote` は引用ブロックを書くために独自に作成したショートコードである。
`build.sh` は名前の通りビルド用のバッチファイル（シェルスクリプト）。

リンク先の変更自体は（大抵のエディタが備えている）単純な置換処理でできるのだが，対象範囲の制限と変更後の Internet Archive への疎通テストもお願いした。
作業計画は問題なさそうなので，そのまま進めてもらう。

Internet Archive 側のリンクでパスの修正が必要だったが（古い日付にリダイレクトされる），それも含めてほぼ手放しで修正作業を完了した。
いやぁ，いい時代になったねぇ。
手作業なら面倒くさくなって途中で投げ出していたところだ。

私がこのブログで最後に参照した記事は2016年の[工藤郁子]さんによる以下の記事だ。

- [メディアは（常に）スパムか？](https://web.archive.org/web/20260123182520/https://magazine-k.jp/2016/01/25/spam-and-media/)

これは書籍『[スパム[spam]:インターネットのダークサイド](https://www.amazon.co.jp/dp/430924744X?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: スパム[spam]:インターネットのダークサイド : フィン・ブラントン, 生貝直人, 成原慧, 松浦俊輔: 本")』の書評記事なのだが，件の本がめっちゃ読みにくいため，[工藤郁子さんの書評記事](https://web.archive.org/web/20260123182520/https://magazine-k.jp/2016/01/25/spam-and-media/ "メディアは（常に）スパムか？")のほうを参考にさせてもらっている。
これは今後も変わらないだろう。

「[マガジン航]」の終了自体は時代的なものもあり「しょうがないか」と思うこともあるが，サイト自体が消えるのは残念な気持ちではある。
あとは `magazine-k.jp` ドメインが spam や phishing に転用されないよう，あるいは Internet Archive が市場や時の為政者に踏み潰されないよう祈るばかりである。

ところで[仲俣暁生](https://note.com/solar1964 "仲俣暁生｜note")さんは Bluesky か Mastodon あたりに来てはいただけないのだろうか。
いや {{% emoji "X" %}} の TL は見ないので，個人的な願望なのだが...

[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[yomoyomo さんの記事]: https://yamdas.hatenablog.com/entry/20261007/magazine-ko "「マガジン航」終了のお知らせに寄せて - YAMDAS現更新履歴"
[マガジン航]: https://web.archive.org/web/20261003095755/https://magazine-k.jp/ "マガジン航[kɔː]"
[Hugo]: https://gohugo.io/ "The world's fastest framework for building websites"
[工藤郁子]: https://bsky.app/profile/fumikok.bsky.social "工藤郁子 Fumiko Kudo (@fumikok.bsky.social) — Bluesky"

## 参考

{{< review-paapi "430924744X" >}} <!-- スパム -->
