+++
title = "ChatGPT という壁打ち"
date =  "2023-03-19T21:43:57+09:00"
description = "改めて OpenAI にアカウントを作った"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日から『[はじめて学ぶ ビデオゲームの心理学]』という本をちょっとずつ読み始めている。
この本の感想はいずれ書くとして，この本に

{{< fig-quote type="markdown" title="『はじめて学ぶ ビデオゲームの心理学』p.39" link="https://www.amazon.co.jp/dp/4571210450?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
ビデオゲームをはじめ各種のゲームは、ツールでないという点で、ほかの製品と少し違います。ビデオゲームは外部の目標に向かってプレイするものではなく、プレイ自体が目的になる、自己目的的な活動です。私たちはゲームを通じてシステムに働きかけ、その反応を楽しみます。
{{< /fig-quote  >}}

と書かれているのを見て，その前に見た

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">おもちゃが手に入ったときの草の根活動は日本最強だと思うんですよね。事業化するのがとにかく下手 <a href="https://t.co/d43QS8F4h7">https://t.co/d43QS8F4h7</a></p>&mdash; odashi (@odashi_t) <a href="https://twitter.com/odashi_t/status/1634252815882391554?ref_src=twsrc%5Etfw">March 10, 2023</a></blockquote>
{{< /fig-gen >}}

という tweet を思い出した。

もの凄い個人的な意見を言わせていただければ AI でプログラマが失業するかもしれないなんて話は完全に想定内なのよ。
つい40年前には「キーパンチャー」とか「コーダー」とかいう（今は絶滅した）職業が存在してたんだぜ。
プログラマみたいな（頭脳）労働集約的な職業が近い将来なくなったとしても全く不思議ではないだろう。

たとえば「Linux のメンテナを AI に一任することになりました」とか「Eマスク氏が AI を使って Twitter を自分好みに全面改修しました」みたいな話が起きれば本格的に（職業をかわるという意味での）転職を考えないといけなくなるだろうが，たかだか学部生の宿題レベルが解ける程度なら，鼻の先は驚異でも脅威でもない（日進月歩という意味では凄いと思うけど）。

話が逸れた。
逸れたついでに...

恥ずかしながら中学時代はテニス部だったのだが[^s1]，運動神経が壊滅してる私は「壁打ち」で練習することが多かった。
「壁打ち」って奥が深いよね。
どの方向にどういう強さ打って，球にどう回転を加えれば，思い通りの球が返ってくるか。
まさに物理学だよね。

[^s1]: 親に無理やり入れさせられた。高校に入ったら一切運動部には関わらないという条件で渋々入った。テニス部を選んだのは一番楽そうに見えたからだ（←シロート考えw）。まぁ，それで身に着けた体力の余剰分でこの歳まで生き延びたんだから，そこだけは当時の親に感謝しておこう。今だに運動も体育会系のノリもまっぴらゴメンって感じだが。

私の中で ChatGPT は何となく「壁打ち」なイメージなんだよなぁ。
インプットに対するアウトプットを見ながら次のインプットを調整していく。
フィードバックを行うのは機械（＝壁）じゃなくて人間のほうだよね。

というような思いつきを Mastodon に投稿したら

{{< fig-quote  class="nobox center" >}}
<iframe src="https://social.hyuki.net/@hyuki/110046769146123924/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe>
{{< /fig-quote >}}

というリアクションをいただいた。
ありがたや。

結城浩さんといえば最近

- [ChatGPT と結城浩の対話（矛盾や反復を含んだ対話によってAIと人間の識別は行えるか） · GitHub](https://gist.github.com/hyuki/f7218870ae47847eab066545b8b51d05)

というのを公開されていて，これは凄いと思った。
対話というのはお互いの間で「文脈（context）」を生成・共有する必要があるが，パッと見た限りでは両者の間に「文脈」があるように見える。
もっとも ChatGPT については

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">ChatGPTと会話していると、記憶力があるように錯覚するけれど、APIで使う人はわかるように、記憶はなく（ステートレスで）、今までの会話の全記録（ただし長さ制限あり）と新しい質問を入力して、新しい回答を生成するだけ。密かに悪巧みを考えることもできない（SFに出てくるAGIではない）</p>&mdash; Haruhiko Okumura (@h_okumura) <a href="https://twitter.com/h_okumura/status/1637280594534211584?ref_src=twsrc%5Etfw">March 19, 2023</a></blockquote>
{{< /fig-gen >}}

なんだそうで，そうなると少なくとも ChatGPT 側は「文脈」と言えるものを持ってないことになる？

「ゲーム」の面白さのひとつは，「ゲーム」というブラックボックスに対して何某かの入力を行いそれに対する反応をフィードバックして次の入力を行う，という繰り返しで対象への「理解」を深めることにあり，それ自体が目的化している点にあると思う。

何だかよく分からないものを理解していくというプロセスは，自身の「『有能さ』への欲求」を満たす行為であり，まさに「ゲーム」に対する内発的動機になり得るものだ。

であるなら ChatGPT も「仕事に使えるか？」とか要らないことは考えずに，純粋に「ゲーム」だと思えば楽しく遊べるのではないか，と考えを切り替えることにした。
競馬の予想屋と大して変わらないなんちゃらアナリストの発言に右往左往するのはもう止めよう。

というわけで，改めて（Google アカウントではない）アカウントを作ってみた。
Subscription で決済が発生するのなら Google アカウントは使いたくなかったので。
まずは無料枠分で「壁打ち」するところから始めようか（笑）

{{< fig-quote type="markdown" title="『はじめて学ぶ ビデオゲームの心理学』p.71" link="https://www.amazon.co.jp/dp/4571210450?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
年齢に関わらず、遊びは私たちの精神を鋭敏に保つために重要です。
[...]
遊ぶことは学ぶことです。
{{< /fig-quote  >}}

## ブックマーク

- [ChatGPTのAPIをGolangで実装する - Qiita](https://qiita.com/yukiaprogramming/items/538a18bb3581245857e5)
- [VisualStudio CodeとGoogle Apps ScriptでChatGPT(gpt-3.5-turbo)をより安全快適に使う](https://zenn.dev/o_ob/articles/gas-chatgpt-api)

[はじめて学ぶ ビデオゲームの心理学]: https://www.amazon.co.jp/dp/4571210450?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "はじめて学ぶ ビデオゲームの心理学 脳のはたらきとユーザー体験（UX） | セリア ホデント, 山根 信二, 成田 啓行 |本 | 通販 | Amazon"

## 参考図書

{{% review-paapi "4571210450" %}} <!-- はじめて学ぶ ビデオゲームの心理学 -->
{{% review-paapi "B0BXQ2HMQ5" %}} <!-- 日経サイエンス2023年5月号（特集：対話するAI ChatGPT） -->
{{% review-paapi "B00LF90DZW" %}} <!-- 賢者の弟子を名乗る賢者 -->
