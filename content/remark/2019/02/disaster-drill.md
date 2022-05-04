+++
title = "「避難訓練」の事前と事後"
date = "2019-02-11T11:10:15+09:00"
description = "とりあえず，この機会にパスワード管理について改めて見直すのがいいだろう。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "management", "password", "politics" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

最初に言い訳しておくと，私がセキュリティ管理者だったのはずいぶん昔の話だし今は職業エンジニアですらないので，セキュリティ関連の話題を深く追いかけることはしないようにしている。
そのせいかどうかは分からないが，日本政府が一般家庭に「サイバー攻撃」を仕掛けるというニュースを最初に聞いたのは先月末，しかも皮肉なことに海外ブログ経由だったりする。

- [Japanese Government Will Hack Citizens' IoT Devices - Schneier on Security](https://www.schneier.com/blog/archives/2019/01/japanese_govern.html)
- [Japanese government plans to hack into citizens' IoT devices | ZDNet](https://www.zdnet.com/article/japanese-government-plans-to-hack-into-citizens-iot-devices/)

元ネタは NHK ニュースらしいので多分「知ってる人は（とっくに）知ってる」状態なのだろう。
日本のいわゆる「セキュリティ・クラスタ」がどういう議論を行ってきたか（行わなかったか）は知らないし知る気もないが，海外からは

{{< fig-quote title="Japanese Government Will Hack Citizens' IoT Devices" link="https://www.schneier.com/blog/archives/2019/01/japanese_govern.html" lang="en" >}}
<q>I am interested in the results of this survey. Japan isn't very different from other industrialized nations in this regard, so their findings will be general.</q>
{{< /fig-quote >}}

という感じで注目されているらしい。
そして

{{< fig-quote title="Japanese Government Will Hack Citizens' IoT Devices" link="https://www.schneier.com/blog/archives/2019/01/japanese_govern.html" lang="en" >}}
<q>I am less optimistic about the country's ability to secure all of this stuff -- especially before the 2020 Summer Olympics.</q>
{{< /fig-quote >}}

と続く。

今回の日本政府による「サイバー攻撃」がどのようなものであれ，その結果を犯罪者側は知ろうとするだろうし確実に利用されるだろう[^s1]。
私も “less optimistic” という印象だ。

[^s1]: なにせ犯罪者側が普段は他の通信トラフィックに紛れてチマチマやっていることを白昼堂々と大規模にやってくれるんだから，彼らは大歓迎だろう。「日本政府からのアクセス」という Phishing もできそうだ。防衛という点に関しては[警察はあてにならない](https://piyolog.hatenadiary.jp/entry/2019/02/08/043308 "福岡県警本部で発生したマルウェア感染についてまとめてみた - piyolog")だろうし。

たとえば大企業などでは従業員に対して偽の Phishing メールを送ったりネットワークに接続する機器に「不正アクセス」を仕掛けたり，といったことが行われる。
これは一種の「避難訓練」で，接続機器の安全性を確かめたり従業員のセキュリティ意識を高めたりするのが目的である。

しかしこういった「避難訓練」は事前の教育と事後のフォローがあって初めて成り立つものである。
事前の教育は今更なのでしょうがないとして（日本ってホントに教育費をケチるよねぇ），日本政府は「リスト攻撃」で危険な機器のリストを作り上げて（おそらく膨大な数になるだろう），それからどうするつもりなのだろう。

サービスプロバイダや企業に対するものとは違い一般家庭への「あくてぃぶさいばーでぃふぇんす（笑）」はあまり現実的ではない気がするが... そいういった「事後」についてどうするのか，といったことがメディアからは聞こえてこない。
聞きそびれているのか（政府やメディアが）意図的に言わないのか，それとも何も考えてないのか（「何も考えてない」に100カノッサ）。

とりあえず，この機会にパスワード管理について改めて見直すのがいいだろう。
ルータやネットワーク家電といった所謂 IoT 機器に関しては

- 購入時の初期パスワードから変更しましょう（頻繁に変える必要はない）
- 機器ごとに異なるパスワードを設定しましょう
- パスワードの作成・管理で人間の脳みそはあてになりません。パスワード管理ツールを使いましょう。不安なら紙に書き出しておくのも可

といったところだろうか。
昔はルータの WAN 側で防衛できていればよかったが（引き続きそれは重要だが），今はスマホなどの携帯端末も含めていくらでも侵入経路があるからね[^lan1]。

[^lan1]: なので家庭内 LAN は用途ごとにセグメントを分けるのが吉。これからは家庭内にも検疫ネットワークが必要になるかも知れないねぇ。

ちなみに IPA ではパスワードに使う文字種と文字数ごとに[パスワードの解読コストを調べ](https://www.ipa.go.jp/security/ipg/documents/dev_setting_crypt.html "情報漏えいを防ぐためのモバイルデバイス等設定マニュアル：IPA 独立行政法人 情報処理推進機構")ていて[^ipa1]，ちょっとデータが古いが

[^ipa1]: 資料によると「利用できる文字種類すべてを完全にランダムに選択して作ったパスワードを一つ一つ調べる全数探索により1日で解読しようとした際にかかるおおまかな想定攻撃コスト」とある。「全数探索(暗号鍵の総数256)でDES10を1日で解読するためのコストを約250万円と仮定します」とあるが2013年頃の資料なので，今はクラウドの利用コストも下がってるし，もっと安上がりにできるかも知れない。

{{< security-strengths-for-password >}} <!-- 要 MathJax -->

という感じ。

文字種を増やしても8文字程度じゃ話にならないのがわかると思う。
英数字を混ぜて12文字以上ならとりあえず安全かな。

## ブックマーク

- [なぜ日本で「国による不正アクセス」が適法に？ アクティブサイバーディフェンスとは ｜ビジネス+IT](https://www.sbbit.jp/article/cont1/35989)
- [総務省のセキュリティ調査に「国が不正ログイン」と騒ぐ頭の悪い人たち／ひろゆき | ニコニコニュース](https://news.nicovideo.jp/watch/nw4807944)
- [安物のIoT機器は、たとえゴミ箱に叩き込んだあとでも持ち主を裏切り続ける  |  TechCrunch Japan](https://techcrunch.com/2019/01/30/cheap-internet-of-things-gadgets-betray-you-even-after-you-toss-them-in-the-trash/)
- [政府“IoT機器2億台にサイバー攻撃”が疑問視される4つの理由 - 「NOTICE」20日から実施 | Beyond（ビヨンド）](https://boxil.jp/beyond/a6051/)

- [リスク認知とトレードオフ]({{< ref "/remark/2016/02/risk-trade-off.md" >}})
- [「パスワードのベストプラクティス」が変わる]({{< ref "/remark/2017/10/changes-in-password-best-practices.md" >}})

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
