+++
date = "2016-02-07T14:23:30+09:00"
update = "2016-02-14T14:03:24+09:00"
description = "Microsoft はもはや Windows を維持できない？ / LINE は「セキュリティ劇場」に気づいたか？ / 2017年の暦 / 鏡の国のアリス"
draft = false
tags = ["windows", "security", "risk", "line", "astronomy", "book", "aozora"]
title = "週末スペシャル： Microsoft はもはや Windows を維持できない？"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1月は行きました。
もう2月ですよ。
そうそう，タイトルは釣りです。

1. [Microsoft はもはや Windows を維持できない？]({{< relref "#win" >}})
1. [LINE は「セキュリティ劇場」に気づいたか？]({{< relref "#line" >}})
1. [2017年の暦]({{< relref "#naoj" >}})
1. [鏡の国のアリス]({{< relref "#alice" >}})

## Microsoft はもはや Windows を維持できない？{#win}

ついに Microsoft はユーザとの約束を反故にするようだ。

- [News ＆ Trend - 「Windows 7、最新CPUでは来夏でサポート終了」、波紋呼ぶMSの方針変更：ITpro](http://itpro.nikkeibp.co.jp/atcl/column/14/346926/020500437/?rt=nocnt)

{{< fig-quote title="「Windows 7、最新CPUでは来夏でサポート終了」、波紋呼ぶMSの方針変更" link="http://itpro.nikkeibp.co.jp/atcl/column/14/346926/020500437/?rt=nocnt" >}}
<q>Windows 7の本来のサポート期限は2020年1月。Windows 8.1は2023年1月までサポートが残っており、2017年7月時点ではメインストリームサポートすら終了していない。Skylakeより前の世代のCPUではWindows 7/8.1をそれぞれの期限まで使い続けられるが、これから新たにパソコンやタブレットを買おうという場合、Windows 7/8.1で使いたくても本来のサポート期限まで使えず、大きな制約を受けることになる。どうしても2020年まで使い続けたければ、旧型CPU搭載のパソコンやタブレットを探すという選択肢もあるが、既に大手メーカーの主力機種はSkylake搭載機に移っており、自由に機種を選べる状況からはほど遠い。</q>
{{< /fig-quote >}}

たしかにエンジニアにとって「保守」は（基本的に）つまらない仕事である。
なぜなら保守はプロダクトを「終わらせる」仕事だからだ。
Windwos 10 のように継続的に改良を続けていくならともかく， Widows 7 / 8.1 は既に終了時期が決まっている「終わるプロダクト」である。

しかしだからといってユーザとの約束を反故にしていいわけがない。
この影響を大きく受けるのは恐らく企業ユーザだろう。

2017年7月でサポートが終わるのであれば，現時点で最新のパソコンを導入できない。
おそらく多くの企業はまだ Windows 10 の評価が終わってない筈で（大きな企業だと評価が完了するのに通常1,2年かかる），評価結果によっては何らかの改修作業が発生する。
設備投資の面から考えても Windows 7 の期限である2020年1月は割とギリギリのラインである。

それでも「いつでも Windows 7 / 8.1 に戻すことができる」のなら段階的に新しい機械と OS を導入する手もあるが，今回の Microsoft の決定はそれも許さないということになる。

昨年から続く Windows 10 への強引なアップグレード要請も鼻についていたが（ちなみにボリュームライセンスではアップグレード要請は出ない）

- [また Windows 10 にヤラレタ（KB3112343 の恐怖）]({{< relref "remark/2015/windows-10-upgrade-problem.md" >}})
- [MS、「Windows 10」を「推奨される」更新プログラムとして提供開始 - CNET Japan](http://japan.cnet.com/news/service/35077208/)
- [Windows 7／8.1→Windows 10が“推奨される更新”に - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1602/02/news081.html)

どうもこれは営業的な問題ではなく， Microsoft はもはや Windows 7 / 8.1 を維持できないと見るべきかもしれない。
Windows 10 の他にも Microsoft は .NET Core や Visual Studio Code といったプロダクトのオープン化を進めていて，保守に人を割けない状態なのではないか。

- [Microsoft、「Visual Studio Code」（β）をオープンソースで公開　Googleの「Go」もサポート - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1511/19/news058.html)
- [［速報］オープンソース版.NETがリリース候補版に到達。Windows、MacOS X、Linuxで同一の.NETアプリが実行可能に。Microsoft Connect(); 2015 － Publickey](http://www.publickey1.jp/blog/15/netwindowsmacos_xlinux.html)

まぁ個人的には，次にパソコンを買い換えるときは Windows を捨てる予定なので全く困らないのであるが，「いよいよもってしぬがよいそしてさようなら」という感じである。
約束を守らない企業に明日はない。

## LINE は「セキュリティ劇場」に気づいたか？{#line}

- [LINE、アカウント乗っ取り対策の「PINコード」認証を廃止、機種変時などの引き継ぎ手順が変更 -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160204_742296.html)

だから私は最初から言ってるじゃない。
PIN コードは認証には使えないって。

- [セキュリティ劇場の喜劇王： LINE — Baldanders.info](http://www.baldanders.info/spiegel/log2/000718.shtml)

芸能人のゴシップには全く関心がないのだが（他人のプライベートがそんなに面白いのか），たまには役に立つということなのであろう。

{{< fig-quote title="LINE、アカウント乗っ取り対策の「PINコード」認証を廃止、機種変時などの引き継ぎ手順が変更" link="http://internet.watch.impress.co.jp/docs/news/20160204_742296.html" >}}
<q>具体的には、LINEの設定画面に「2段階認証」の「アカウントを引き継ぐ」というオプションを追加。これをオンにした後、24時間以内に限り、新端末側からの引き継ぎ操作を進められるようになっている。</q>
{{< /fig-quote >}}

まぁ遅まきながらではあるが， LINE が自身の道化っぷりに気づいたのであればよいことである。
[LINE を捨て](http://www.baldanders.info/spiegel/log2/000718.shtml)て1年半ほどになるが，ようやくまともな方に揺り戻したと考えていいだろうか。

もっとも1年半ほどの間に LINE がなくて困った事態には全く遭遇しなかったので，今後も導入しないと思うが。
LINE ってもうメッセージングというより広告アプリだしな（笑）

## 2017年の暦{#naoj}

- [平成29年（2017）暦要項を発表 | 国立天文台(NAOJ)](http://www.nao.ac.jp/news/topics/2016/20160201-rekiyoko.html)

毎年2月のはじめに国立天文台から翌年の暦が発表される。
祝日だけでなく24節気や朔望月等の暦象も発表される。

## 鏡の国のアリス{#alice}

- [aozorablog » 『するりと鏡を――ぬけてみて、アリスに見えたもの』巻頭詩](http://www.aozora.gr.jp/aozorablog/?p=3614)

大久保ゆうさんの今年の翻訳記事は「鏡の国のアリス」のようだ。
月1での更新らしい。
楽しみ！
