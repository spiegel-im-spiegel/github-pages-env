+++
title = "#shimanego より： 配列とスライスとソート"
date =  "2020-02-13T19:51:39+09:00"
description = "description"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "communication", "shimane", "engineering", "presentation" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

今月も [Shimane.go#04] に参加して図々しくもまた喋ってきた。
以下にそのとき使ったスライドを公開しておく。

- [配列とスライスとソート | slide.Baldanders.info](https://slide.baldanders.info/shimane-go-2020-02-13/)

ターゲットとしては [A Tour of Go] をひととおり終わらせて「なんとなく」 [Go] が分かってきたかなぁ，という感じの人。
さすが松江は「お膝元」なので Ruby 経験者は多いが [Go] には馴染みのない人が多いようなので。

実はソートの速度とかベンチマークを取ってやろうかとも考えたのだが，いい感じのデータが作れず（ただのランダムなデータ列ならいくらでも作れるけど，多分そうじゃない），諦めた。
いい方法を考えたらそのうちブログ記事にするかも。

まぁでも，やっぱ座学は退屈だよねぇ。
プログラムは書いてナンボだし。

なんか面白い遊びを提示できれば面白いんだけど。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Shimane.go#04]: https://shimane-go.connpass.com/event/165192/ "Shimane.go#04 - connpass"
[A Tour of Go]: https://go-tour-jp.appspot.com/

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
