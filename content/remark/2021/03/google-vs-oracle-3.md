+++
title = "Google vs Oracle の訴訟の行方 3"
date =  "2021-03-29T19:20:47+09:00"
description = "そういや，どうなったっけ？"
image = "/images/attention/kitten.jpg"
tags = ["code", "intellectual-property", "copyright", "java", "google", "android"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先週の mimemagic 騒ぎに関する戯れ言を [Zenn のほうに書いた](https://zenn.dev/spiegel/articles/20210327-copyright "「著作権」は何を制限しているのか")が，これで思い出したのが「そういや Google vs Oracle の訴訟ってどうなったっけ？」だった。

たしか Google による最高裁への上告が2019年に受理され（「コロナ」のせいで遅れに遅れて）翌2020年10月から審理が始まったんだよね。

- [グーグルとオラクルとの「Android」関連訴訟、米最高裁での審理始まる  - CNET Japan](https://japan.cnet.com/article/35160621/)

んで，まだ判決は出てないってことでOK？

この訴訟は，私が一介のプログラマとして Java というか Oracle を見限るキッカケになったものなので（もちろん仕事なら Java だろうが何だろうが喜んで引き受けますよw），当初から注目はしていた。
まぁ，ここ2,3年は追いかける余裕がなかったのだが。

せっかくなので[今まで書いた](https://baldanders.info/blog/000861/ "Google vs Oracle の訴訟の行方")分も含めて最初から書き直すか。

## Java API 特許権への侵害はなかったので...

元々この訴訟は Andorid を巡って Google が Java API の「特許権」を侵害している，というのが Oracle の言い分だった。
特許権侵害については2012年の連邦地裁の判決で決着がついている。

- [AndroidをめぐるOracle対Google裁判を振り返る（前編）～ Oracleが主張した特許侵害は認められず － Publickey](https://www.publickey1.jp/blog/12/androidoraclegoogle_oracle.html)
- [AndroidをめぐるOracle対Google裁判を振り返る（後編）～ 残る課題はAPI著作権と9行のコード － Publickey](https://www.publickey1.jp/blog/12/androidoraclegoogle_api9.html)

ここまでの経緯をかいつまんで紹介するとこんな感じ。

1. 最初は著作権ではなく特許権の侵害（7件）の有無が訴訟の中心だった
2. Oracle の申し立てた7件のうち5件は特許自体が無効とされ，残り2件については侵害は認められないとされた
3. この判決を受け，争点が特許権の侵害から著作権の侵害へと移る。論点は以下の通り
   1. 37の Java API パッケージの互換コードについて， Oracle は Google が著作物全体の「構造、順序、構成」（Structure, Sequence, Organization; SSO）を侵害したと証明したか
   2. 「`TimSort.java` および互換コードの `rangeCheck` メソッド」の著作権侵害があったか
4. 争点 3-1 については著作権侵害は認められるが「公正な利用（fair use）」については結論が出ず。争点 3-2 については侵害が認められた（ただし軽微）

上の争点 3-1 については，さらに2つの論点がある。

1. API の構成要素である SSO について著作権が適用されるのかどうか
2. 仮に著作権が適用されるとして公正な利用の範囲内かどうか

で，判決では，前者については yes，後者については不明（訴訟の範囲から外れるので）となった。
これに対し Google は「API に著作権なんかねーよ！」（←超意訳）と最高裁に上告したわけだ。

その結果が以下。

## 「API のコピー」は著作権に抵触するか

- [Oracle Am., Inc. v. Google Inc., No. 13-1021 (Fed. Cir. 2014) :: Justia](https://law.justia.com/cases/federal/appellate-courts/cafc/13-1021/13-1021-2014-05-09.html) : 判決文
- [グーグル対オラクルのJava訴訟、米最高裁がグーグルの上告を棄却--Reuters - CNET Japan](https://japan.cnet.com/article/35066650/)
- [OracleとGoogleの判決文を斜め読む - Qiita](https://qiita.com/shibukawa/items/9c99321b8edb6fc09cce)
- [Google v. Oracle API著作権裁判 · GitHub](https://gist.github.com/yudai/6f8f44ac878c41eaf7dc)

こんな感じで最高裁への上告は棄却されたので，連邦地裁での判決および議論はそのまま持ち越されたことになった。
ここは（プログラマとして）大事なところなので，もう少し詳しく見ていこう。

### 宣言コード

この「API のコピー」の論点には大きく2つあるようだ。
ひとつは「宣言コード（declaring code）」の扱いで，もうひとつは「SSO の非逐語的コピー（non-literal copy）」である。 

宣言コードというのは，たとえば

```java {hl_lines=[1]}
public static int max(int x, int y) {
    if (x > y) return x;
    else return y;
}
```

というコード（これは判決文の中に出てくる）の最初の行の部分を指す。
ちなみに，宣言コードに対する `{ ... }` 内の記述を「実装コード（implementing code）」と呼ぶ。

Google は，宣言コードは “method of operation” であり著作権は適用されないと主張したが[^moo1]，最終的には認められなかった。 

[^moo1]: 何故 “Method of operation” が著作権の適用外になるのかということについては「[コンピュータ関係の創作保護についての最近の米国での話題（via Internet Archive）](https://web.archive.org/web/20160915163851/http://homepage3.nifty.com/nmat/LOTUS.HTM)」あたりが参考になる。 なお「最近」と書かれているが1996年の記事である。

{{< fig-quote title="Oracle Am., Inc. v. Google Inc., No. 13-1021" link="https://law.justia.com/cases/federal/appellate-courts/cafc/13-1021/13-1021-2014-05-09.html" lang="en" >}}
{{% quote %}}Rather, the uncopyrightable "method of operation" or "system" or "process" is the underlying computer function triggered by the written code -- for example, an algorithm that the computer executes to sort a data set. The code itself, however, is eligible for copyright protection{{% /quote %}}.
{{< /fig-quote >}}

宣言コードは単なる「名前」ではなく，著作権の適用という点で宣言コードと実装コードを区別する意味はないということらしい。
その上で宣言コードの逐語的コピー（literal, verbatim copying of declaring code）があった，と認定されたわけだ。

### 構造、順序、構成

もうひとつの SSO だが， Google は “merger doctrine” や “scènes à faire doctrine” を盾に JDK の SSO を著作権の適用外とするよう求めていたが，この点についても認められなかった。

ちなみに “merger doctrine” や “scènes à faire doctrine” てのは，アイデアなどに対する「表現」が限られたバリエーションしかない場合や，古典的あるいは標準的な表現の場合には著作権の適用外となるというものだそうだが，控訴審が「命名や構造化には無数の方法がある」と述べたことに対し Google は反論しなかったようだ。

{{< fig-quote title="Oracle Am., Inc. v. Google Inc., No. 13-1021" link="https://law.justia.com/cases/federal/appellate-courts/cafc/13-1021/13-1021-2014-05-09.html" lang="en" >}}
{{% quote %}}The Court thus decided the case based on what was effectively a merger analysis. Here, by contrast, petitioner does not dispute the court of appeals' statement that there were "unlimited" ways that respondent could have named and structured its methods, Pet. App. 33, and nothing logically required petitioner to copy respondent's declaring code when it created the Android platform{{% /quote %}}.
{{< /fig-quote >}}

SSO のコピー自体は非逐語的コピーかもしれないが，もともと Android のプラットフォームや開発ツールが Java エンジニアへの利便性のために JDK に似せて作られている点と，宣言コードが明らかに逐語的コピーである点も合わせて，「独立した表現」であるとは言えない（SSO のコピーが逐語的であるかどうかの要件を問わない）ということのようだ。

ぶっちゃけて言うと「それ，ただの丸パクリだろ」ってことらしい。
身も蓋もないな（笑）

## 「公正な利用」は有効か（そして再び最高裁へ）

上の判断をふまえた上で API コードの利用[^use1] に “fair use doctrine” が認められるかどうかについては結論を先送りとし「上告を棄却」となった。
ただし，この訴訟については Section 102(b) （著作権の適用範囲）で争うのではなく， API コードに著作権があるとした上で， Section 107 “fair use doctrine” の可否で争うべきだとの意見が添えられた。

[^use1]: 著作権では「使用」と「利用」を分けて考える。著作権は「利用」を除く「使用」について関知しない。著作権の「利用」について厳密な話は大変だが，概ね「複製」「配布」「改変」の3つだと覚えておけばいいだろう。この記事では Java API の「複製」が問題となっているわけだが，逐語的コピーでないなら「改変」も含むし，それを広く公開しているのだから「配布」も絡んでくる。

{{< fig-quote title="Oracle Am., Inc. v. Google Inc., No. 13-1021" link="https://law.justia.com/cases/federal/appellate-courts/cafc/13-1021/13-1021-2014-05-09.html" lang="en" >}}
Rather, petitioner argues that programmers have become fluent in respondent's Java platform; that they will be deterred from writing programs for Android if they are required to learn all new commands; and that verbatim copying of respondent's declaring code was necessary for the familiar commands to work on the Android platform.

The general concerns that petitioner raises are substantial and important, but Section 102(b) is not the appropriate statutory provision to address them. Rather, legitimate concerns with interoperability and lock-in effects are far better addressed through the fair-use doctrine codified at Section 107.
{{< /fig-quote >}}

そこで争点は「API に著作権があるとした上で公正な利用を満たしているか」に移っていった。

米国の著作権における「公正な利用」は以下の4要件について社会的に「公正」であるかを（訴訟を通じて）議論する。

1. 利用の目的や本質
1. 原作品の本質
1. 抜粋の量や実質性
1. 原作品の価値への影響

つまり「公正な利用」に争点が移った時点で個別案件となったのだ。
故に，たとえこの件で Google が勝ったとしても，あらゆる局面で「API の利用は『公正な利用』である」とは言えない。

「自由でないソフトウェアによる API」の利用は常に著作権リスクを含んでいる。
このことは頭に入れておいて欲しい。

実は，2016年の時点では陪審員団に Java API の利用が公正な利用であると認められていた。

- [Google beats Oracle—Android makes “fair use” of Java APIs | Ars Technica](https://arstechnica.com:443/tech-policy/2016/05/google-wins-trial-against-oracle-as-jury-finds-android-is-fair-use/)
- [Google’s fair use victory is good for open source | Ars Technica](https://arstechnica.com:443/tech-policy/2016/06/googles-fair-use-victory-is-good-for-open-source/)
- [グーグル、Java API使用が「フェアユース」と認められる--対オラクル訴訟 - CNET Japan](https://japan.cnet.com/article/35083291/)
- [Googleの「公正使用」勝訴後も残る著作権に関する疑問  |  TechCrunch Japan](https://jp.techcrunch.com/2016/05/31/20160527copyright-questions-remain-after-googles-fair-use-victory/)
- [Oracle v. Google API Fair Use訴訟の話（連邦地裁編）](https://gist.github.com/yudai/c5906ca61d4fe367180a6e079c8fc309) : 日本語解説ならこちらがオススメ

しかし2018年の判決では一転して Oracle 側の勝訴となっている。

- [Oracle America, Inc. v. Google, Inc., No. 17-1118 (Fed. Cir. 2018) :: Justia](https://law.justia.com/cases/federal/appellate-courts/cafc/17-1118/17-1118-2018-03-27.html) : 判決文
- [OracleがJavaの著作権侵犯裁判でGoogleに勝利  |  TechCrunch Japan](https://jp.techcrunch.com/2018/03/28/2018-03-27-oracle-wins-appeal-against-google-in-copyright-case/)

そしてこの記事の最初の話に戻るわけだ。

## 【2021-04-11 追記】

判決出ちゃいました。

- [Google vs Oracle の訴訟の行方（最終章）]({{< ref "/remark/2021/04/google-vs-oracle-final.md" >}})

## ブックマーク

- [Google Books の Library Book Scan すら Fair Use と言われたのに...]({{< ref "/remark/2015/google-books-library-project.md" >}})
- [Google vs Oracle の訴訟の行方 2]({{< ref "/remark/2016/05/28-stories.md#api" >}} "週末スペシャル： Barack Obama 米国大統領の来広")

## 参考図書

{{% review-paapi "4641243336" %}} <!-- 著作権法 第3版 -->
{{% review-paapi "4757102852" %}} <!-- 著作権２．０ ウェブ時代の文化発展をめざして -->
{{% review-paapi "4757122349" %}} <!-- 〈反〉知的独占 -->
{{% review-paapi "B01CYDGUV8" %}} <!-- CODE VERSION 2.0 -->
{{% review-paapi "4622073455" %}} <!-- 〈海賊版〉の思想‐18世紀英国の永久コピーライト闘争 -->
{{% review-paapi "4797334673" %}} <!-- インターネットの法と慣習 -->
