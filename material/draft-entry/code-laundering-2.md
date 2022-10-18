+++
title = "GitHub Copilot とコード洗浄"
date =  "2022-10-18T09:13:23+09:00"
description = "description"
image = "/images/attention/site.jpg"
tags = [ "code", "intellectual-property", "github", "artificial-intelligence" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

今年に入ってなかなかにお騒がせな GitHub Copilot であるが，今回もまたひとつ（笑）

- [ソースコードの続きを自動補完するGitHub Copilotに「著作権で保護されたコードを出力している」という指摘が寄せられる - GIGAZINE](https://gigazine.net/news/20221018-github-copilot-emit-copyrighted-code/)

ちなみに私の問題意識はちょっとずれたところにあって，これについては既に[記事にしている][記事]ので，よろしかったらどうぞ。

では本題へ。

## 著作権は本当は何を「保護」しているのか

そもそもの話，既に公有 [{{% cc-syms "pd" %}}](https://creativecommons.org/publicdomain/mark/1.0/ "Creative Commons — Public Domain Mark 1.0") であるものや [CC0 {{% cc-syms "cc" "zero" %}}](https://creativecommons.org/publicdomain/zero/1.0/ "Creative Commons — CC0 1.0 Universal") 献呈（dedication）されたものを除き，著作権で「保護」されて**いない**著作物などないのである。
それは FOSS 製品のプログラム・コードでも同じ。

いや，最近つくづく思うんだけどさ。
「著作権で『保護』された」などという言い回しが誤解の元なんじゃないかねぇ。

著作権による「保護」は自然保護や動物愛護みたいなのとは根本的に異なる。
本来は公有である「表現」を一定の制限下で「独占」することを（法により）許可するのが著作権の本質であり，その「独占」という特権状態を実際にどう運用するかは著作（権）者に委ねられている。
そのための許諾契約が「ライセンス」なのだ。
このように著作権システムは二重三重に捻じれている。

だから「著作権で保護された」などという言い回しより「独占状態を著作権により保護（保障）された」あるいはもっと直截に「著作権で独占された」とはっきり言うべきなんじゃないの？ こんな持って回った言い方をするから某著作権管理組織の中の人みたいに「著作権は [...] 憲法で保障された基本的人権」などと言い出したりするんだよ。
著作権が基本的人権というなら，それを管理組織に委譲させて著作者本人からも銭を巻き上げるのは明確に人権の損害じゃんか（笑）

...いや，愚痴の続きは後日ということで。

## コード洗浄

GitHub Copilot が知識ソースとしているのは GitHub 上の公開された FOSS コードである。

本来 FOSS 製品のコードを「利用」するのであれば帰属とライセンスを明示する必要がある。
コードを利用するユーザから見ると，たとえば GNU GPL なコードを自分の製品に「利用」するのなら 自分のコードも同じく GNU GPL でライセンスしなければならない。

GitHub Copilot は，そういったコードにまつわる諸々をまるっと無視して「洗浄」し「ボクが考えたコード」みたいな顔をして提示するが，ユーザ側はそんなコードを提示されても利用の可否を判断できず，コードの品質以前に「怖くて使えない」ということになる。

以前に紹介した SFC の “[Give Up GitHub](https://sfconservancy.org/blog/2022/jun/30/give-up-github-launch/ "Give Up GitHub: The Time Has Come! - Conservancy Blog - Software Freedom Conservancy")” でも

{{< fig-quote type="markdown" title="Give Up GitHub: The Time Has Come!" link="https://sfconservancy.org/blog/2022/jun/30/give-up-github-launch/" lang="en" >}}
One recent preliminary finding was that [AI-assisted software development tools can be constructed in a way that by-default respects FOSS licenses](https://lists.copyleft.org/pipermail/ai-assist/2022-June/000015.html). We will continue to support the committee as they explore that idea further, and, with their help, we are actively monitoring this novel area of research. While Microsoft's GitHub was the first mover in this area, by way of comparison, early reports suggest that Amazon's new CodeWhisperer system ([also launched last week](https://www.theregister.com/2022/06/23/amazon_codewhisperer/)) seeks to provide proper attribution and licensing information for code suggestions.
{{< /fig-quote >}}

と述べていて，少なくとも「ライセンスのどんぶり勘定」部分については改善する意思を示さないと，本当に「GitHub を諦める」ことになりかねないと思うぞ。

## 創作者は特権者ではない

AI 関係で今年の話題と言えば「AI 絵画」だろう。
最近では Microsoft が Azure で Dall·E 2 の提供を開始した。
AI コード生成と同じく， AI 絵画も既にクラウドで PaaS 化されている。

- [［速報］マイクロソフト、Azure OpenAI Serviceで「Dall·E 2」の提供を開始、テキストから画像を生成可能。Ignite 2022 － Publickey](https://www.publickey1.jp/blog/22/azure_openai_servicedalle_2ignite_2022.html)

「著作物」という観点でプログラム・コードが小説・音楽・絵画等と違うのは「作家性」の強度だろうか。

たとえば「Linux は Linus Torvalds 作」みたいな感じで製品と作者がダイレクトに繋がっている場合もあるが，実際には一から十まで彼が全部プログラム・コードを書いているわけではなく，多数の参加者による貢献（contribution）の上に成り立っている。
この辺の話は，今や古典となった「[伽藍とバザール](http://cruel.org/freeware/cathedral.html)」あたりを読むのが手っ取り早いだろう。

ネットとはコピーと改変の連鎖でできている。
FOSS はこれを積極的に受け入れることで IT 産業構造をも変えたわけだ。

AI がコードを書くようになったとしても，それはコード貢献を行うプレイヤーに「機械」が加わっただけで（生成されるコードが使えるかどうかはともかく）さしたるインパクトではないのかもしれない。
そうなってくると前節までの懸念こそが杞憂であり，独占型著作権システムは文化・技術の発展を妨げる「老害」に成り果てる？

一方，小説・音楽・絵画等の作者は creator vs. consumer の対置で特権的な立場で扱われることが多いし，著作権システムもそれを後押ししている。
作者側も創作物を自身の identity の一部として特別視している人が多いように見える（個人の感想です）。

「AI 絵画」の台頭は creator vs. consumer という構造を解体する。
それは作者を特権者の立ち位置から引きずり下ろし，彼らの創作物を基礎とした identity を揺るがす。

これが何をもたらすのかは私には分からない。
「作者などいない」になるかもしれないし「製作委員会方式」に機械がクレジットされるかもしれない。
あるいはター〇ネーターよろしく機械に徹底抗戦する未来かもしれない（笑）

## 機械に出来ることを人間が「頑張る」のはリソースの無駄遣い

私は，機械に出来ることを人間が「頑張る」のは，人間という貴重なリソースの無駄遣いと思っている。
ましてや機械のために人間が「頑張る」のは以ってのほかだろう。

せっかく文明社会に生きてるんだから人と機械が上手く連携し，社会の中で**人間**が幸せになる未来を設計していきたいし，それが工学の目指す「善」だと思う。

{{< fig-quote title="はやぶさ―不死身の探査機と宇宙研の物語" link="https://www.amazon.co.jp/dp/4344980158?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
<q>理学は、真理の探究であり、工学は善の実現である。そして、藝術は美の表現である－－これで所謂「真美善」が揃う</q>
{{< /fig-quote >}}

[記事]: {{< ref "/remark/2022/07/code-launderring.md" >}} "GitHub Copilot は貢献者から貢献を奪うか？"

## 参考図書

{{% review-paapi "B099RTG3J7" %}} <!-- 著作権は文化を発展させるのか: 人権と文化コモンズ -->

[伽藍とバザール](https://www.aozora.gr.jp/cards/000029/card227.html)
