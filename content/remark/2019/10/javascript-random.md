+++
title = "「今あなた以外に○○人が見ています」デモ"
date =  "2019-10-21T22:58:57+09:00"
description = "これ最初に考えた奴は天才だろ（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "programming", "javascript", "random" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ぶちウケた！

- [旅行予約サイトの「今あなた以外に○○人が見ています」はウソだったことが判明 - GIGAZINE](https://gigazine.net/news/20191021-one-travel-number-rundom/)

この記事によると

{{< fig-quote type="markdown" title="旅行予約サイトの「今あなた以外に○○人が見ています」はウソだったことが判明" link="https://gigazine.net/news/20191021-one-travel-number-rundom/" >}}
{{< quote >}}JavaScriptで以下のような部分を発見。完全に28から44までの数字がランダムに生成され、それが表示されているだけだったことが判明しました。つまり、One Travelの「○○人がこの搭乗券をチェックしています」という部分は全くのウソであり、ユーザーにフライトの予約を急がせるためのものだったというわけです。{{< /quote >}}
{{< /fig-quote >}}

なんだって。

ちうわけで，戯れにデモ・コードを書いてみた。
こんな感じでどうだろう。

```html
<div id='demo' class='box'></div>
<script>
  let rn = Math.floor( Math.random() * 17 ) + 28;
  let p = document.createElement('p');
  p.appendChild(document.createTextNode('今あなた以外に'+rn+'人が見ています'));
  p.setAttribute('class', 'center');
  document.getElementById('demo').appendChild(p);
</script>
```

実際に試してみよう。

{{< fig-gen title="「今あなた以外に○○人が見ています」デモ" >}}
<div id='demo' class='box'></div>
<script>
  let rn = Math.floor( Math.random() * 17 ) + 28;
  let p = document.createElement('p');
  p.appendChild(document.createTextNode('今あなた以外に'+rn+'人が見ています'));
  p.setAttribute('class', 'center');
  document.getElementById('demo').appendChild(p);
</script>
{{< /fig-gen >}}

おー。
できたできた。

最小の労力で最大の効果。
これ最初に考えた奴は天才だろ（笑）

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4873115531?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51HQWyI1aUL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4873115531?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">JavaScriptリファレンス 第6版</a></dt>
    <dd>David Flanagan (著), 木下 哲也 (翻訳)</dd>
    <dd>オライリージャパン 2012-08-10</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4873115531 (ASIN), 9784873115535 (EAN), 4873115531 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="3">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">最初に私が JavaScript を勉強したのって，これよりも更に古い版だったんだよなぁ。この版でも今となっては古すぎて使い物にならないけど（笑）</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-10-21">2019-10-21</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
