+++
title = "「パスワードのベストプラクティス」が変わる"
date =  "2017-10-16T19:21:02+09:00"
description = "Bruce Schneier 氏の「Changes in Password Best Practices」の内容が簡潔だったので「そのうち紹介しなくちゃ」と思っていたが，先を越されたっぽい感じなので，便乗記事として上げておく（笑）"
tags = [ "security", "authentication", "password", "risk", "management", "nist" ]

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

{{< security-strengths-for-password >}} <!-- 要 MathJax -->

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

- [Authenticator と AAL]({{< ref "/remark/2020/09/authenticator-and-aal.md" >}})

[Changes in Password Best Practices]: https://www.schneier.com/blog/archives/2017/10/changes_in_pass.html "Changes in Password Best Practices - Schneier on Security"
[やばいパスワード]: http://itpro.nikkeibp.co.jp/atcl/column/17/092800400/101200002/ "やばいパスワード - 複雑なパスワードを強制、でも破られやすいという現実：ITpro"
[PGP]: https://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/ "Amazon | PGP―暗号メールと電子署名 | シムソン ガーフィンケル, Simson Garfinkel, ユニテック 通販"
[KeePass]: https://keepass.info/ "KeePass Password Safe"
[Mono]: http://www.mono-project.com/

## 参考図書

{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
