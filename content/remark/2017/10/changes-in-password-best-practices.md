+++
title = "「パスワードのベストプラクティス」が変わる"
date =  "2017-10-16T19:21:02+09:00"
update =  "2017-10-18T15:02:52+09:00"
description = "Bruce Schneier 氏の「Changes in Password Best Practices」の内容が簡潔だったので「そのうち紹介しなくちゃ」と思っていたが，先を越されたっぽい感じなので，便乗記事として上げておく（笑）"
tags        = [ "security", "authentication", "password", "risk", "management" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

Bruce Schneier 氏の “[Changes in Password Best Practices]” の内容が簡潔だったので「そのうち紹介しなくちゃ」と思っていたが，先を越されたっぽい感じなので，便乗記事として上げておく（笑）

- [やばいパスワード - 複雑なパスワードを強制、でも破られやすいという現実：ITpro](http://itpro.nikkeibp.co.jp/atcl/column/17/092800400/101200002/)

まずは “[Changes in Password Best Practices]” で挙げられている3つの要件を以下に紹介する。

1. Stop it with the annoying password complexity rules. They make passwords harder to remember. They increase errors because artificially complex passwords are harder to type in. And they don't help that much. It's better to allow people to **use pass phrases**. 
2. top it with password expiration. That was [an old idea for an old way](https://securingthehuman.sans.org/blog/2017/03/23/time-for-password-expiration-to-die "Security Awareness Blog | Time for Password Expiration to Die") we used computers. Today, **don't make people change their passwords unless there's indication of compromise**.
3. Let people **use password managers**. This is how we deal with all the passwords we need.

強調部分は私によるものである。

## 生成規則が複雑なだけのパスワードではダメ

最初の要件は，いたずらに複雑なパスワード生成規則を強要するな，というものだ。
ここで勘違いしてもらっては困るのだが，これは「複雑なパスワードはダメ」と言っているのではない。

パスワードの要件は

- できるだけ文字数が多いこと
- 出来るだけ文字種が多いこと
- 出来るだけランダムに文字を選ぶこと

の3つである。
これを人力で作って覚えるのは難しい。

パスワードの強度に関して IPA が2013年に発表した資料がある。

- [情報漏えいを防ぐためのモバイルデバイス等設定マニュアル：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ipg/documents/dev_setting_crypt.html)
    - {{< pdf-file title="情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編" link="https://www.ipa.go.jp/files/000026760.pdf" >}}

このうちの「{{< pdf-file title="解説編" link="https://www.ipa.go.jp/files/000026760.pdf" >}}」にパスワードの解読されやすさの一覧表がある。
以下に引用しよう。

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

見たらわかると思うが，文字種の多さより文字数の多いほうがインパクトがあることが分かる。
数英大小文字記号すべて使って8文字のパスワードを作っても**1日で解読完了させるコスト**は21万円ほどだが，これは数字だけで16文字からなるパスワードよりもコストが低い。

ただしこれは最もコストの高い「総当たり攻撃」の場合である（しかも4年も前の話だ）。

{{< div-gen >}}
<figure>
<blockquote>
<q>利用できる文字種類すべてを完全にランダムに選択して作ったパスワードを一つ一つ調べる全数探索により1日で解読しようとした際にかかるおおまかな想定攻撃コストを示しています。ここでは、全数探索(暗号鍵の総数256)でDES10を1日で解読するためのコストを約250万円と仮定します。また、パスワードを1つ検査するのとDESの暗号鍵を1つ検査するコストは同じであるとし、パスワードを求めるのに必要な計算量(検査する個数)が半分になればコストも半分、2倍になればコストも2倍になるものとしています。</q>
</blockquote>
<figcaption><q><a href='https://www.ipa.go.jp/files/000026760.pdf'>情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編 <sup><i class='far fa-file-pdf'></i></sup></a></q> 2.4.2.2項より</figcaption>
</figure>
{{< /div-gen >}}

したがって辞書にあるような単語を組み合わせたパスフレーズの場合は，余程の単語数が必要になる。
たとえば「[やばいパスワード]」で紹介されている方法では

```text
テレビは1日1時間 → Terebi ha 1 niti 1 Jikan  → Terebiha1niti1Jikan
```

といった感じで19文字のパスワードを生成していて，先程の表で照らし合わせると，総当たり攻撃なら，ほとんど天文学的なコストになるが，実際にはそれぞれの単語が攻撃側の辞書にある場合は，たったの6単語しかないわけで，やり方によっては解読コストをかなり引き下げられる可能性もある（犯罪者側の最近の事情をよく知らないので杞憂かもしれないが）。

パスフレーズというのは，もともと「辞書攻撃」などない長閑な時代に PGP などのセキュリティ製品に採用されていたもので（今でも GnuPG ではパスフレーズが使用できる）

```text
I could tell you my pass phrase, but then I would have to kill you.
```

みたいな比較的長い文（phrase）を使うことを想定している[^pgp1]。
したがってパスワード解読技術が向上している今の時代に単語を組み合わせただけのパスフレーズがどこまで効果的かは正直に言って分からない。
しかし “`Password!1`” みたいな「法令遵守の観点から社内規則に則ってはいるけど機械で容易に推測可能なパスワード」よりは遥かにマシということなのである。

[^pgp1]: この物騒なパスフレーズは Simson Garfinkel 氏の『[PGP]』に載っていたものである。パスフレーズは文字数や文字種の制限がないのが特徴である。もし入力に日本語（UTF-8）が使えるのなら，海外の犯罪者に対しては，かなり強力なパスフレーズができると思うんだけどねぇ。

## パスワードの有効期限など無意味

2番目の要件は既に散々言われていることなので今更であろう。

- [週末スペシャル： 「パスワードの定期変更はすべきでない」，他]({{< ref "/remark/2016/07/02-stories.md" >}})

### 【追記】 パスワード定期変更の起源？

「[JNSAメールマガジン　臨時号　2015.4.3.](http://www.jnsa.org/aboutus/jnsaml/ml-57special.html)」によると

{{< fig-quote title="JNSAメールマガジン　臨時号　2015.4.3." link="http://www.jnsa.org/aboutus/jnsaml/ml-57special.html" >}}
<q>パスワードの定期的な変更は、パスワード文字列が4文字だった時代にパスワードの総当たり攻撃(ブルートフォース攻撃)の対策として実施したことが起源と言われている。</q>
{{< /fig-quote >}}

なんだそうだ。
ホンマかいな（笑）

むしろ，ネットワーク管理者やセキュリティ管理者が自分たちの仕事（アカウント管理）を面倒臭がって「放っておいても期限切れになる」パスワード運用を強制した，のほうに1票いれるよ。

## パスワード管理ツールを使え

3番目の要件もおなじみのやつである。

最近はウイルス対策ソフトを提供しているセキュリティ企業がパスワード管理ツールも提供していたりするので，そちらを使う手もある。
私としては [KeePass] を是非オススメするが。

- [KeePass Password Safe](https://keepass.info/)

[KeePass] はオープンソースのパスワード管理ツールで Windows の .NET Framework 用だが Linux 用に [Mono] で動作するバージョンも存在する。
また Android や iOS で動作する互換アプリも存在する。

- [Keepass2Android Password Safe - Google Play](https://play.google.com/store/apps/details?id=keepass2android.keepass2android)

パスワード管理ツールにはデータ暗号化にマスタパスワードを要求するものが多いが， [KeePass] では暗号鍵ファイルで暗号化できるのでマスタパスワードも不要だ（マスタパスワードと組み合わせることも可能）。
暗号鍵ファイルさえ適切に管理すればパスワードを格納した DB ファイルをクラウドに置いて（他人に見えないところに置いてね）機器間で共有することもできる。

## パスワードを覚えようとか考えないこと

いつも言っていることだが「**パスワードを覚えるなんて脳みその無駄使い**」である。
最初に述べた3要件を満たすパスワードをツールで生成させてツールで管理すればよい。
どうしても不安ならば紙に書いて誰にも知られないように管理する手もある。
パスフレーズだって「覚えないといけない」という点では同じことなのだ。

「セキュリティと利便性のトレードオフ」なんてのは昔の話である。
適切な運用をすればセキュリティも利便性も両方確保できる。
それをしないのは単なる怠慢だ。

## ブックマーク

- [NIST SP 800-63 Digital Identity Guidelines](https://pages.nist.gov/800-63-3/)

- [崩れる「安全なパスワード」神話　否定される過去の基準、追従できない現場の課題 ｜ビジネス+IT](https://www.sbbit.jp/article/cont1/33969)
- [あの「面倒なパスワード作成ルール」、作った人も後悔していた - ZDNet Japan](https://japan.zdnet.com/article/35105725/)

[Changes in Password Best Practices]: https://www.schneier.com/blog/archives/2017/10/changes_in_pass.html "Changes in Password Best Practices - Schneier on Security"
[やばいパスワード]: http://itpro.nikkeibp.co.jp/atcl/column/17/092800400/101200002/ "やばいパスワード - 複雑なパスワードを強制、でも破られやすいという現実：ITpro"
[PGP]: http://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/ "Amazon | PGP―暗号メールと電子署名 | シムソン ガーフィンケル, Simson Garfinkel, ユニテック 通販"
[KeePass]: https://keepass.info/ "KeePass Password Safe"
[Mono]: http://www.mono-project.com/

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/5132396FFQL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/">PGP―暗号メールと電子署名</a></dt><dd>シムソン ガーフィンケル Simson Garfinkel </dd><dd>オライリー・ジャパン 1996-04</dd><dd>評価<abbr class="rating" title="3"><img src="http://g-images.amazon.com/images/G/01/detail/stars-3-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4756136494/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4756136494.09._SCTHUMBZZZ_.jpg"  alt="プログラミング作法"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320026926.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語C 第2版 ANSI規格準拠"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797350997/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797350997.09._SCTHUMBZZZ_.jpg"  alt="新版暗号技術入門 秘密の国のアリス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798132608/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798132608.09._SCTHUMBZZZ_.jpg"  alt="情報処理教科書 高度試験午後II論述 春期・秋期 (EXAMPRESS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798105538/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798105538.09._SCTHUMBZZZ_.jpg"  alt="エンタープライズ アプリケーションアーキテクチャパターン (Object Oriented Selection)"  /></a> </p>
<p class="description" >良書なのだが，残念ながら内容が古すぎた。 PGP の歴史資料として読むならいいかもしれない。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-10-16">2014/10/16</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
