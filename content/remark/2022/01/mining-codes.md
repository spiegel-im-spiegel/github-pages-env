+++
title = "「俺が正義だ！」"
date =  "2022-01-22T13:24:25+09:00"
description = "今回の判決文はなかなか面白いので，特に職業エンジニアの方は読んでおくことをおすすめする。"
image = "/images/attention/kitten.jpg"
tags = [ "code", "law", "security", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[4年前]({{< ref "/remark/2017/12/28-stories.md#gray" >}} "サイトオーナーがページの広告掲載の代わりにマイニング・コードを仕込むのはヤクザの「みかじめ料」と同じ")にもちょろんと言及しているので，今回も一応言及しておこう。

例の裁判は最高裁による最終判断が出たようだ。

- [令和2(あ)457 不正指令電磁的記録保管被告事件](https://www.courts.go.jp/app/hanrei_jp/detail2?id=90869)
- {{< pdf-file title="判決文" link="https://www.courts.go.jp/app/files/hanrei_jp/869/090869_hanrei.pdf" >}}

Coinhive が提供する mining code による計算資源搾取は，当時は大騒ぎになったが，今ではブラウザ自身が拒否できるようになった（どこまで効くかは分からないけどw）。

{{< fig-img src="./firefox-settings.png" link="./firefox-settings.png" title="Firefox のセキュリティ設定例" width="676" >}}

今は「[インフラ処理でCPUを使ったら負け](https://ascii.jp/elem/000/004/070/4070140/ "ASCII.jp：ネットワークに特化したIPUのMount Evansでシェア拡大を狙うインテル　インテル CPUロードマップ (1/3)")」などと言われる時代である。
それだけ CPU/GPU 計算資源は価値が高いのだ。
もちろん，それを駆動する[電力](https://nextmoney.jp/?p=46494 "ロシアとウクライナが違法なマイニング施設を閉鎖| NEXTMONEY｜仮想通貨メディア")もね。

一方で，こう言っちゃあなんだが，日本は法治国家ではなく，刑罰に於ける法執行は基本的に「見せしめ主義」だ。
何故なら全国民や全法人に対して法を厳密に執行したら罪人や違反者で溢れかえってしまうから。
故に「[バレなきゃ犯罪じゃないんですよ](https://dic.nicovideo.jp/a/%E3%83%90%E3%83%AC%E3%81%AA%E3%81%8D%E3%82%83%E7%8A%AF%E7%BD%AA%E3%81%98%E3%82%83%E3%81%AA%E3%81%84%E3%82%93%E3%81%A7%E3%81%99%E3%82%88)」みたいな名言も生み出すことになる（笑）

しかも日本人は何故か「{{< ruby "ざまぁ" >}}私刑{{< /ruby >}}」に寛容で，積極的に行う輩すらいる野蛮国家である。
日本の法システムは壊れていると言ってもいいかもしれない。

今回の裁判は最高裁判決文の以下の文言に集約されていると言ってもいいだろう。

{{< fig-quote type="markdown" title="判決文" link="https://www.courts.go.jp/app/files/hanrei_jp/869/090869_hanrei.pdf" >}}
{{% quote %}}原判決は，不正指令電磁的記録の解釈を誤り，その該当性を判断する際に考慮すべき事情を適切に考慮しなかったため，重大な事実誤認をしたものというべきであり，これらが判決に影響を及ぼすことは明らかであって，原判決を破棄しなければ著しく正義に反すると認められる{{% /quote %}}
{{< /fig-quote >}}

ここで「原判決」というのは二審の有罪判決を指すが，遠回しに警察の所業も指していると解釈している。
どこぞのガンダムよろしく「俺が正義だ！」とばかりに法解釈を歪めて相手を陥れるというのは，勧善懲悪な時代劇ならスカッとするかもしれないが，リアルでやったら恐怖政治そのものだ。

個人的には今回の最高裁判決文は妥当なものと納得している。
その上で改めて「[サイトオーナーがページの広告掲載の代わりにマイニング・コードを仕込むのはヤクザの「みかじめ料」と同じ]({{< ref "/remark/2017/12/28-stories.md#gray" >}})」であると主張しておこう。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

今回の{{< pdf-file title="判決文" link="https://www.courts.go.jp/app/files/hanrei_jp/869/090869_hanrei.pdf" >}} はなかなか面白いので，特に職業エンジニアの方は読んでおくことをおすすめする。

今回の最高裁判決文のポイント（判示事項）は2つあるそうで

{{< fig-quote type="markdown" title="令和2(あ)457 不正指令電磁的記録保管被告事件" link="https://www.courts.go.jp/app/hanrei_jp/detail2?id=90869" >}}
1. 刑法１６８条の２第１項にいう「その意図に沿うべき動作をさせず，又はその意図に反する動作をさせるべき不正な指令を与える電磁的記録」に当たるか否かの判断方法
2. ウェブサイトの閲覧者の同意を得ることなくその電子計算機を使用して仮想通貨のマイニングを行わせるプログラムコードが不正指令電磁的記録に当たらないとされた事例
{{< /fig-quote >}}

この中の「意図に反する動作」と「不正な指令」がキーワードになっているのが分かるだろう。
この2つについて判決文には

{{< fig-quote type="markdown" title="判決文" link="https://www.courts.go.jp/app/files/hanrei_jp/869/090869_hanrei.pdf" >}}
すなわち，反意図性は，当該プログラムについて一般の使用者が認識すべき動作と実際の動作が異なる場合に肯定されるものと解するのが相当であり，一般の使用者が認識すべき動作の認定に当たっては，当該プログラムの動作の内容に加え，プログラムに付された名称，動作に関する説明の内容，想定される当該プログラムの利用方法等を考慮する必要がある。

また，不正性は，電子計算機による情報処理に対する社会一般の信頼を保護し，**電子計算機の社会的機能を保護するという観点から，社会的に許容し得ないプログラムについて肯定される**ものと解するのが相当であり，その判断に当たっては，当該プログラムの動作の内容に加え，その動作が電子計算機の機能や電子計算機による情報処理に与える影響の有無・程度，当該プログラムの利用方法等を考慮する必要がある。
{{< /fig-quote >}}

（強調はわたしがやりました）

とアンダーライン付きで記されている（大事なことらしい）。
これを踏まえて

{{< fig-quote type="markdown" title="判決文" link="https://www.courts.go.jp/app/files/hanrei_jp/869/090869_hanrei.pdf" >}}
{{% quote %}}本件プログラムコードは，反意図性は認められるが，不正性は認められないため，不正指令電磁的記録とは認められない{{% /quote %}}
{{< /fig-quote >}}

となったわけだ。
「反意図性」だけでは「不正指令電磁的記録」の要件を満たさないということだろう。
例えば，プログラムのバグや脆弱性を指して，いちいち「不正指令電磁的記録」などと訴訟を起こされては堪らないからな。

実際，4年前に Coinhive の mining code が登場した当初から，これを malware (malicious software) と見なすかどうかには戸惑いがあった。

{{< fig-quote title="FinTech？マルウエア？無断でスマホCPU使う謎のサービス" link="http://itpro.nikkeibp.co.jp/atcl/column/14/346926/110801194/" >}}
<q>しかし、現実としてはトレンドマイクロのウイルスバスターはCoinhiveをブロックしている。CoinhiveのJavaScriptが埋め込まれたWebサイトにアクセスすると、警告メッセージを表示してスクリプトの実行を止める。シマンテックなどほかのベンダーのセキュリティソフトも同様だ。<br>
マルウエアには当たらないが、好ましくない動作を行う可能性がある「グレーウエア」に分類されているためだ。</q>
{{< /fig-quote >}}

もっとも今ではウイルス対策ツールにすら mining code が仕込まれる有様だけどね（笑）

- [Norton 360 Now Comes With a Cryptominer – Krebs on Security](https://krebsonsecurity.com/2022/01/norton-360-now-comes-with-a-cryptominer/)
- [Once Opted Into Norton Crypto, You Can't Easily Uninstall | Digital Trends](https://www.digitaltrends.com/computing/no-easy-way-to-uninstall-norton-crypto/)
- [Yes, Norton 360 has a built in cryptominer. Deletion is easy • The Register](https://www.theregister.com/2022/01/05/norton_360_cryptominer_deletion/)
- [Norton’s Antivirus Product Now Includes an Ethereum Miner - Schneier on Security](https://www.schneier.com/blog/archives/2022/01/nortons-antivirus-product-now-includes-an-ethereum-miner.html)

今回の最高裁判決文を読む限り「不正性」の証明はなかなかに難しい印象を受ける。
今後もこの境界を巡って法的な議論があるだろうが，{{< ruby "孤独な正義" >}}独善{{< /ruby >}}を振りかざして真っ当なエンジニアを陥れるのは勘弁していただきたいところである。

一方で私は，エンジニアは「善を実装する者」であると認識している。
なにを以って「善」とするかは難しいところだが，この一点に於いて「[エンジニアこそ『狂狷の徒』たれ]({{< ref "/remark/2017/12/hacker-ethic.md" >}})」と思うのである。

{{< fig-quote title="はやぶさ―不死身の探査機と宇宙研の物語" link="https://www.amazon.co.jp/dp/4344980158?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
<q>理学は、真理の探究であり、工学は善の実現である。そして、藝術は美の表現である－－これで所謂「真美善」が揃う</q>
{{< /fig-quote >}}

おあとがよろしいようで {{< emoji "ペコン" >}}

## ブックマーク

- [コインハイブ事件の有罪判決、破棄自判で「無罪」に　最高裁 - 弁護士ドットコム](https://www.bengo4.com/c_1009/n_14033/)
- [Are We Really Engineers? • Hillel Wayne](https://www.hillelwayne.com/post/are-we-really-engineers/)

## 参考図書

{{% review-paapi "B01ESA9R5Q" %}} <!-- 機動戦士ガンダム00 -->
{{% review-paapi "B07K356C43" %}} <!-- 転スラ Another Colony -->
