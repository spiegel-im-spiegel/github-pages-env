+++
title = "「避難訓練」の事前と事後"
date = "2019-02-11T11:10:15+09:00"
description = "とりあえず，この機会にパスワード管理について改めて見直すのがいいだろう。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "management", "password", "politics" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

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

{{< div-gen >}}
<figure>
<table>
<thead>
<tr>
<th colspan='4'>利用する文字種類数と内訳</th>
<th colspan='4'>パスワード長</th>
</tr>
<tr>
<th>種類数</th>
<th>数字</th>
<th>文字</th>
<th>シンボル</th>
<th>4文字</th>
<th>8文字</th>
<th>12文字</th>
<th>16文字</th>
</tr>
</thead>
<tbody>
<tr><td>10種</td><td>0-9</td><td>なし</td>      <td>なし</td><td>1円未満（$2^{13.3}$）</td><td>1円未満（$2^{26.6}$）</td>  <td>約35円（$2^{39.9}$）</td>     <td>約35万円（$2^{53.2}$）</td></tr>
<tr><td>36種</td><td>0-9</td><td>a-z</td>       <td>なし</td><td>1円未満（$2^{20.7}$）</td><td>約100円（$2^{41.4}$）</td>  <td>約1.65億円（$2^{62.0}$）</td> <td>約276兆円（$2^{82.7}$）</td></tr>
<tr><td>62種</td><td>0-9</td><td>a-z<br>A-Z</td><td>なし</td><td>1円未満（$2^{23.8}$）</td><td>約7,500円（$2^{47.6}$）</td><td>約1,120億円（$2^{71.5}$）</td><td>約165京円（$2^{95.3}$）</td></tr>
<tr><td>94種</td><td>0-9</td><td>a-z<br>A-Z</td><td><code style='font-size:smaller;'>! " # $ %<br>&amp; ' ( ) =<br>~ | - ^ `<br>¥ { @ [<br>+ * ] ; :<br>} &lt; &gt; ? _<br>, . /</code></td>
                                                             <td>1円未満（$2^{26.2}$）</td><td>約21万円（$2^{52.4}$）</td> <td>約16.5兆円（$2^{78.7}$）</td> <td>約129,000京円（$2^{104.9}$）</td></tr>
</tbody>
</table>
<figcaption>パスワード解読の想定コスト例（<q><a href='https://www.ipa.go.jp/files/000026760.pdf'>情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編 <sup><i class='far fa-file-pdf'></i></sup></a></q> 2.4.2.2項より）</figcaption>
</figure>
{{< /div-gen >}}

という感じ。

文字種を増やしても8文字程度じゃ話にならないのがわかると思う。
英数字を混ぜて12文字以上ならとりあえず安全かな。

## ブックマーク

- [なぜ日本で「国による不正アクセス」が適法に？ アクティブサイバーディフェンスとは ｜ビジネス+IT](https://www.sbbit.jp/article/cont1/35989)
- [総務省のセキュリティ調査に「国が不正ログイン」と騒ぐ頭の悪い人たち／ひろゆき | ニコニコニュース](https://news.nicovideo.jp/watch/nw4807944)
- [安物のIoT機器は、たとえゴミ箱に叩き込んだあとでも持ち主を裏切り続ける  |  TechCrunch Japan](https://jp.techcrunch.com/2019/02/01/2019-01-30-cheap-internet-of-things-gadgets-betray-you-even-after-you-toss-them-in-the-trash/)

- [リスク認知とトレードオフ]({{< ref "/remark/2016/02/risk-trade-off.md" >}})

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3%E3%81%AF%E3%81%AA%E3%81%9C%E3%82%84%E3%81%B6%E3%82%89%E3%82%8C%E3%81%9F%E3%81%AE%E3%81%8B-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC/dp/4822283100?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4822283100"><img src="https://images-fe.ssl-images-amazon.com/images/I/51-pZ52JsUL._SL160_.jpg" width="107" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3%E3%81%AF%E3%81%AA%E3%81%9C%E3%82%84%E3%81%B6%E3%82%89%E3%82%8C%E3%81%9F%E3%81%AE%E3%81%8B-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC/dp/4822283100?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4822283100">セキュリティはなぜやぶられたのか</a></dt>
	<dd>ブルース・シュナイアー</dd>
	<dd>井口 耕二 (翻訳)</dd>
    <dd>日経BP社 2007-02-15</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4822283100, EAN: 9784822283100</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">原書のタイトルが “<a href="https://www.amazon.co.jp/Beyond-Fear-Thinking-Sensibly-Uncertain-ebook/dp/B000PY3NB4?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B000PY3NB4">Beyond Fear: Thinking Sensibly About Security in an Uncertain World</a>” なのに対して日本語タイトルがどうしようもなくヘボいが中身は名著。とりあえず読んどきなはれ。ゼロ年代当時 9.11 およびその後の米国のセキュリティ政策と深く関連している内容なので，そのへんを加味して読むとよい。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-02-11">2019-02-11</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
