+++
date = "2015-11-28T19:39:24+09:00"
description = "DELL よ，お前もか / たしかに「プライバシーマーク」はクソだけど / Google は Android 端末を解除できる"
tags = ["security", "cryptography", "risk", "pki", "x509", "privacy", "grigori"]
title = "週末スペシャル： DELL よ，お前もか"

[scripts]
  mathjax = false
  mermaidjs = false
+++

1. [DELL よ，お前もか]({{< relref "#pki" >}})
1. [たしかに「プライバシーマーク」はクソだけど]({{< relref "#privacy" >}})
1. [Google は Android 端末を解除できる]({{< relref "#android" >}})

## DELL よ，お前もか{#pki}

あーあ。

- [Vulnerability Note VU#870761 - Dell Foundation Services installs root certificate and private key (eDellRoot)](http://www.kb.cert.org/vuls/id/870761)
- [Vulnerability Note VU#925497 - Dell System Detect installs root certificate and private key (DSDTestProvider)](http://www.kb.cert.org/vuls/id/925497)
- [JVNVU#91791008: Dell Foundation Services (DFS) がルート証明書と秘密鍵 (eDellRoot) をインストールする問題](http://jvn.jp/vu/JVNVU91791008/)
- [JVNVU#99824449: Dell System Detect (DSD) がルート証明書と秘密鍵 (DSDTestProvider) をインストールする問題](http://jvn.jp/vu/JVNVU99824449/)
- [Dell製PCで確認されたeDellRoot証明書の関連情報をまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20151124/1448366156)
- [デル製PCに「意図せぬ脆弱性」--プリインストールされたルート証明書で - ZDNet Japan](http://japan.zdnet.com/article/35073924/)
- [Dell、ルート証明書の脆弱性で対応表明　別の問題発覚 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1511/25/news055.html)
- [DellのPCに不審なルート証明書、LenovoのSuperfishと同じ問題か - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1511/24/news048.html)
- [「Windows Defender」、デルのルート証明書問題に対応 - ZDNet Japan](http://japan.zdnet.com/article/35074095/)
- [デル、PC証明書脆弱性の影響範囲などのより詳細な情報を開示 ～32bit版も対応開始 - PC Watch](http://pc.watch.impress.co.jp/docs/news/20151201_733070.html)
- [不注意で公開されたデジタル証明書により、なりすましが行われる - マイクロソフト セキュリティ アドバイザリ 3119884](https://technet.microsoft.com/ja-jp/library/security/3119884)
    - [MSが証明書信頼リストを更新--デルのルート証明書問題に対応 - ZDNet Japan](http://japan.zdnet.com/article/35074332/)

[Lenovo のとき](https://baldanders.info/blog/000809/)も書いたけど， PC ベンダ側がどう言い訳しようとも，これは純然たる「悪意」なの！

{{< fig-quote title="Malware Spoofing HTTPS — Baldanders.info" link="https://baldanders.info/blog/000809/" >}}
<q>セキュリティ企業や組織は「脆弱性（vulnerability）」などと比較的穏当な表現をしているが，オレオレ・ルート証明書をインストールする行為自体が明確な「悪意」である。 なぜなら（上述した通り）これは PKI の信用モデルに対する攻撃（破壊）だからだ。</q>
{{< /fig-quote >}}

ちなみに EV SSL も前提が崩れ始めていることをお忘れなく。

- [踊る PKI — Baldanders.info](https://baldanders.info/blog/000828/)

Lenovo のときは「中国企業だから」なんてなこともちょっと思ったが， DELL もやっているとなると X.509 がいよいよ「バベルの塔」になる日も近いということかねぇ。

## たしかに「プライバシーマーク」はクソだけど{#privacy}

- [News ＆ Trend - なぜCCCはプライバシーマークを返上し、T会員規約を改訂したのか（前編）：ITpro](http://itpro.nikkeibp.co.jp/atcl/column/14/346926/112000384/?rt=nocnt)
- [News ＆ Trend - なぜCCCはプライバシーマークを返上し、T会員規約を改訂したのか（後編）：ITpro](http://itpro.nikkeibp.co.jp/atcl/column/14/346926/112600387/?cx)
- [高木浩光＠自宅の日記 - CCCはお気の毒と言わざるをえない](http://takagi-hiromitsu.jp/diary/20151121.html#p01)

この問題が複雑そうに見えるのは「プライバシーマーク」が本当にただの「マーク」で，ユーザから見てなんの役にも立っていないということだ。
本来この手の認証サービスは相手が信頼できるかどうか不明な場合に第3者がそれを担保する仕組みなのだが，「プライバシーマーク」はそのようには機能していない。
なぜなら，認証を受けた企業・組織が実際に条件を逸脱する行為をしたとしても「プライバシーマーク」はその逸脱に対して事実上なにもしないからだ。

企業が利益を追求すること自体は evil ではない。
だがそれは他の企業・組織やユーザとの間に社会的な信頼関係があってはじめて成立し得る。
その企業・組織の信頼性を客観的に示す方法はない。
地道に「過去の実績」を積み上げていくしかないのである。
これを「ブランド」という。

CCC のブランド・イメージはとっくの昔に壊れている。
それは公共図書館運営に対する「嫌悪」とも言える反発を見ても分かることだ。
「信用がない」どころかマイナスに振りきれている。
それで「うちは独自のプライバシー・ポリシーで運営します」と言ったところで「あー，また CCC がやらかしてら」くらいにしか思われなくて当然だろう。

まぁ，私個人は今後も CCC は（プライバシー情報の運用に対して）信用しないし， [CCC に加担する企業・組織](http://nukalumix.hateblo.jp/entry/tcardoptoutlist)も信用しない。

## Google は Android 端末を解除できる{#android}

- [Googleはユーザーが所有するスマホの端末ロックをリモート解除可能と発覚 - GIGAZINE](http://gigazine.net/news/20151125-google-remote-unlock/)

なお，これには条件があって

{{< fig-quote title="Googleはユーザーが所有するスマホの端末ロックをリモート解除可能と発覚" link="http://gigazine.net/news/20151125-google-remote-unlock/" >}}
<q>レポートで判明した事実について、IT関連メディアの<a href="http://thenextweb.com/google/2015/11/22/google-can-remotely-bypass-the-passcode-of-at-least-74-of-android-devices-if-ordered/">The Next Web</a>がGoogleに確認したところ、Googleが端末のロック解除をリモートで行えるのは「パターンで画面ロックされているAndroid 4.4搭載端末のみ」であり、PINおよびパスワードでロックしている端末についてはリモートで解除を行えないとのことです。</q>
{{< /fig-quote >}}

ということらしい。
Andorid 5.x は暗号化されている建前だが，パフォーマンスがどうこう言うお馬鹿ユーザが多いせいか，暗号化されてない端末も多いようだ。
Andorid 5.x でも暗号化されてない端末には同様のリスクがある。

Google はユーザに対して「[デバイスマネージャ](https://www.google.com/android/devicemanager)」のサービスを提供しており， Google が画面ロックをリモートで解除可能だとしてもおかしくはない。
また PIN コードやパスワードでリモート解除できないのであれば，それほどのインパクトではない。

対策としては

1. Android 端末は必ず暗号化を行うこと
1. 盗難・紛失・没収の際は「[デバイスマネージャ](https://www.google.com/android/devicemanager)」を使い強いパスワードでロックすること

というところだろう。
あと，端末の紛失を想定して「避難訓練」しておくのは有効である。

もちろんデータをクラウドに置いている場合はクラウドごとに個別に対処する必要があるし，そもそも携帯端末は tracking が容易であり警察・諜報組織は端末の中身が見れなくてもさほど困らないということをお忘れなく。

- [スマートフォンのセキュリティ管理 — Baldanders.info](https://baldanders.info/blog/000516/)
- [「パスワードを覚える」なんて脳みその無駄遣い — Baldanders.info](https://baldanders.info/blog/000739/)
- [「オーウェルが描いた悪夢のような監視社会をさまざまな点で超えてしまっているこの世界」で私たちはいかにして生き残るか — Baldanders.info](https://baldanders.info/blog/000768/)
- [警察がテロリストの捨てた携帯からわかること : ギズモード・ジャパン](http://www.gizmodo.jp/2015/11/what-police-can-learn-from-a-terrorist.html)

## 参考図書

{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
