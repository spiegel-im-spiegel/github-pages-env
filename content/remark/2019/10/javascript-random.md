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

{{% review-paapi "4873115531" %}} <!-- JavaScriptリファレンス 第6版 -->
