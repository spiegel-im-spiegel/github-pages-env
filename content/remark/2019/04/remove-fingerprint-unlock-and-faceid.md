+++
title = "指紋認証と FaceID は解除せよ"
date = "2019-04-08T22:54:17+09:00"
description = " セキュリティのために「個人の自由」を売り渡してはいけないのだ。"
image = "/images/attention/kitten.jpg"
tags = [ "code", "security", "privacy", "risk", "politics", "grigori" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

タイトルは釣りです，もちろん。

[EFF] から面白い記事が出てた。

- [Attending a Protest | Surveillance Self-Defense](https://ssd.eff.org/en/module/attending-protest)

この記事は抗議行動を含む市民活動においてセキュリティないしはプライバシーのリスクとなりうるものを軽減するためのアドバイスである。
日付が4月1日になってたのでエイプリルフールネタかと思ったのだが，どうもマジ話らしい。

英語不得手なので記事の全部を紹介することは出来ないが “[Before the Protest](https://ssd.eff.org/en/module/attending-protest#1)” の節だけちょろんと紹介しよう（だれか翻訳しないかな。ちなみに原文は CC BY で公開されている）。

## Enable full-disk encryption on your device

まぁ，これは当然だよね。

抗議活動以前に携帯端末は紛失・盗難のリスクが高い。
（犯罪者だろうが国家だろうが）データを盗もうとする側は「貴方」の携帯端末に気を遣ったりしない。
最悪の場合，端末をぶっ壊してでもデータを吸い上げようとするだろう。

注意しないといけないのは外部ストレージ（SD カードなど）の暗号化は別途行わなければならないということ。
あるいはそういったものの暗号化が出来ないかもしれない。

もっと言うとストレージの暗号化が出来るデジカメはあまりないらしい。

{{< fig-quote tittle="Attending a Protest" link="https://ssd.eff.org/en/module/attending-protest" lang="en" >}}
<q>In addition, many digital cameras lack the ability to encrypt. It is safe to assume that photos and videos taken with digital cameras will be stored unencrypted, unless explicitly stated otherwise.</q>
{{< /fig-quote >}}

本気で携帯端末を暗号化するのはけっこう大変かもしれない。

## Remove fingerprint unlock and FaceID

つい先日も Twitter か Facebook の TL で書いたような気がするが，いわゆる生体情報（biometric）は秘密情報ではない。
故に生体情報を認証に使うべきではない。

生体情報を認証に使うのなら相手（犯罪者かもしれないし国家かもしれない）はそれを取得するために躊躇なく物理的手段を講じるだろう[^i1]。
グミで指紋のコピーを作ったりどっかから顔写真を入手する必要はない。
本人を連れてきて無理やり認証してしまえばいいのだ。
本人が抵抗するなら首か手首を切り取ってしまえばいいのだ[^bm1]。

[^i1]: もちろん，そうするだけのインセンティブがあればの話だよ。鵜呑みにしないように（笑）
[^bm1]: 海外では実際にそういう事例があった。指紋認証が使える高級車を盗むためにオーナーの指を切り取ったそうだ。

{{< fig-quote tittle="Attending a Protest" link="https://ssd.eff.org/en/module/attending-protest" lang="en" >}}
<q>In the U.S., using a biometric—like your face scan or fingerprint—to unlock your phone compromises protections for the contents of your phone afforded to you under the Fifth Amendment privilege against compelled incrimination. A police officer may try to intimidate you into “consenting” to unlock your phone, whether you use a biometric or a memorized passcode. But if you exercise your right to refuse and biometric unlocking functionality is turned on, an officer may physically force you to biometrically unlock your device.</q>
{{< /fig-quote >}}

アメリカめっさ怖いな！ いや，日本でも（昨今の[ケーサツの暴走っぷり]({{< ref "/remark/2019/03/the-crawling-chaos.md" >}} "暗い水曜日，他")から考えると）ありうる話かもしれないが。

## Install Signal

「信号」じゃなくてセキュリティ・ツールの [Signal] ね。

- [Signal · GitHub](https://github.com/signalapp)

最近のバージョンではグループチャットも出来るようになったらしい。

[Signal] の利点は会話履歴をネット上のどこにも残さない点である。
存在しないデータは取得しようがない。

{{< fig-quote tittle="Attending a Protest" link="https://ssd.eff.org/en/module/attending-protest" lang="en" >}}
<q>In 2016, a grand jury in the Eastern District of Virginia issued a subpoena to Open Whisper Systems, the developers of Signal. Because of the architecture of Signal, which limits the user metadata stored on the company’s servers, the only data they were able to provide was "the date and time a user registered with Signal and the last date of a user's connectivity to the Signal service."</q>
{{< /fig-quote >}}

つまり，ユーザが「[Signal] を使った」という履歴は残るけど，会話の内容は一切記録されないというわけだ。

私も [Signal] をメインのメッセージング・アプリとして使っている。
[Facebook の Messenger は削除した]({{< ref "/remark/2019/03/withdrawal-from-facebook.md" >}})。
スマホ標準の SMS アプリと置き換えることもできるので，まずはそこから始めてみてもいいだろう。

## Back up your data

まぁこれも当たり前。
バックアップ先が商用のクラウドストレージでは意味がないからね（国家はサービス・プロバイダに命令できる）。

## Buy a prepaid, disposable phone

アメリカではプリペイド SIM を買うのに ID を提示しなくていいらしい。
なんと羨ましい。

{{< fig-quote tittle="Attending a Protest" link="https://ssd.eff.org/en/module/attending-protest" lang="en" >}}
<q>In the United States, at the time this guide was written, current federal regulation does not require you to show your ID to purchase a prepaid SIM card (but your state might). Most countries require you to provide a form of ID to purchase a prepaid SIM card, thus linking the card to your identity and removing the possibility of anonymity.</q>
{{< /fig-quote >}}

しかし，そうするといわゆる「プリペイド携帯」を買えって話になるけど，日本ではプリペイド携帯も身分の提示が必要になるんじゃなかったっけ。
つか，そもそもプリペイド携帯って今も売ってるのか？ 最近の状況を知らないのでよく分からない。

“[Attending a Protest]” には「機内モードを有効にしろ（Enable airplane mode）」とも書かれていて，普段はケータイを機内モードにしておけば少なくとも行動追跡のリスクは減るかも知れない。

## 国家は国家のためにしか駆動しない

「自分の身は自分で守れ」というのはいかにもアメリカらしいなぁ，と思ったりする。

しかし結局のところ，国家は国家のためにしか駆動しない。
個人である私達を守れるのは最終的に私達自身しかいないのだ。

そういう意味で「プライバシー」は「個人の自由」を守るための大切な権利である，という点は忘れてはいけないと思う。
セキュリティや利便性のために「個人の自由」を売り渡してはいけないのだ。
「[それは、ぜったいに、ぜったいです](https://dic.nicovideo.jp/a/%E3%82%B7%E3%83%B3%E3%83%95%E3%82%A9%E3%82%AE%E3%82%A2%E3%82%BB%E3%83%AB%E3%83%95%E3%83%91%E3%83%AD%E3%83%87%E3%82%A3%E9%9B%86)」。

## ブックマーク

- [クーリエ連載；エコノミスト紹介、自由のためなら人が死んでもいい](https://cruel.org/economist/courier200712.html)

[EFF]: https://www.eff.org/ "Electronic Frontier Foundation | Defending your rights in the digital world"
[Signal]: https://www.signal.org/
[Attending a Protest]: https://ssd.eff.org/en/module/attending-protest "Attending a Protest | Surveillance Self-Defense"

## 参考図書

<div class="hreview" >
	<div class="photo"><a class="item url" href="https://tatsu-zine.com/books/infoshare2"><img src="https://tatsu-zine.com/images/books/877/cover_s.jpg" alt="photo"></a></div>
    <dl class="fn">
      <dt><a href="https://tatsu-zine.com/books/infoshare2">もうすぐ絶滅するという開かれたウェブについて 続・情報共有の未来</a></dt>
      <dd>yomoyomo</dd>
      <dd>達人出版会 2017-12-25</dd>
      <dd>評価&nbsp;<abbr class="rating fa-sm" title="4">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="far fa-star"></i>
      </abbr></dd>
    </dl>
    <p class="description"><a href="https://wirelesswire.jp/author/yomoyomo/">WirelessWire News 連載</a>の書籍化。感想は<a href="/remark/2019/01/infoshare2/">こちら</a></p>
	<p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018-12-31</abbr></p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.aozora.gr.jp/cards/000021/card4307.html"><img src="https://text.baldanders.info/images/aozora/card4307.svg" width="110" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.aozora.gr.jp/cards/000021/card4307.html">グリゴリの捕縛</a></dt>
    <dd>白田 秀彰</dd>
    <dd> 2001-11-26 (Release 2014-09-17)</dd>
    <dd>青空文庫</dd>
    <dd>4307 (図書カードNo.)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">白田秀彰さんの「<a href="http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm">グリゴリの捕縛</a>」が青空文庫に収録されていた。
内容は <delete>怪獣大決戦</delete> おっと憲法（基本法）についてのお話。
古代社会 → 中世社会 → 近代社会 → 現代社会 と順を追って法と慣習そして力（power）との関係について解説し，その中で憲法（基本法）がどのように望まれ実装されていったか指摘してる。
その後，現代社会の次のパラダイムに顕現する「情報力」と社会との関係に言及していくわけだ。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-30">2019-03-30</abbr> (powered by <a href="https://aozorahack.org/">aozorahack</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51-pZ52JsUL._SL160_.jpg" width="107" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">セキュリティはなぜやぶられたのか</a></dt>
    <dd>ブルース・シュナイアー (著), 井口 耕二 (翻訳)</dd>
    <dd>日経BP 2007-02-15</dd>
    <dd>単行本</dd>
    <dd>4822283100 (ASIN), 9784822283100 (EAN), 4822283100 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">原書のタイトルが “<a href="https://www.amazon.co.jp/dp/B000PY3NB4?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">Beyond Fear: Thinking Sensibly About Security in an Uncertain World</a>” なのに対して日本語タイトルがどうしようもなくヘボいが中身は名著。とりあえず読んどきなはれ。ゼロ年代当時 9.11 およびその後の米国のセキュリティ政策と深く関連している内容なので，そのへんを加味して読むとよい。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-02-11">2019-02-11</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
