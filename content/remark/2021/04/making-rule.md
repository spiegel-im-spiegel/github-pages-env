+++
title = "MISRA-C の思ひ出（または「守られないルールはルール自体に問題がある」）"
date =  "2021-04-25T11:35:57+09:00"
description = "ルール・メイキングは難しい"
image = "/images/attention/kitten.jpg"
tags = [ "code", "engineering", "security", "risk", "management", "golang", "lint" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## オンライン開催お疲れさまでした

2021-04-21 に開催された [Go Conference 2021 Spring](https://gocon.connpass.com/event/208896/) は初のオンラインだったそうで，週末の土曜日ということもあって，出不精で人見知りな私でも気軽に参加できたのがありがたかった。
（リアルタイム視聴も含めて）動画や音声のコンテンツはまとまった時間で拘束されてしまうのであまり好きではないのだが，たまにはこういうものに参加するのもよろしかろう。

そのなかのひとつである

- [Go をセキュアに書き進めるための「ガードレール」を整備しよう / Let's Build Security Guardrails For Your Go Programs! - Speaker Deck][ガードレール]

は興味深く視聴させてもらった。
特に [go-ruleguard] は面白そうだ。
個人でも使う機会があるかもしれない。

ちなみに，[スライド][ガードレール]にも書かれているが， [gosec] は [golangci-lint] にも組み込まれているので

```text
$ golangci-lint run --enable gosec ./...
```

といった感じに使うことができる。

## MISRA-C の思ひ出

この発表を視聴しながら思い出していたのは大昔に車載系のプロジェクトに参加したときのことだ。
あのときはアセンブラでベクタテーブルからゴリゴリ書いてたよなぁ（遠い目）

今は違うと思うが，当時は MISRA-C というガイドラインがあって，これが車載系における事実上の C 言語コーディング基準（criteria）になっていた。

詳細は割愛するが， MISRA-C には127個のルールがあって「必要（Required）」と「推奨（Advisory）」のいずれかに分類されている。
このうち「必要」ルールは強制的に課せられるルールで，このルールからの逸脱（deviation）を許容する場合には，手続きを踏んで文書化と承認を行わなければならない。
各ルールにはルールの詳細（何故そのルールが必要か）と逸脱可否の判断基準とサンプルコードが載っているので，それを参考に逸脱の可否を判定する。

MISRA-C が定めるルール自体は合理的な内容なのだが，127個ものルールを机上でチェックするのは不毛なので MISRA-C 対応の lint ツールを使うことになる（当時はこれがバカ高くてねぇ）。
で，当然ながら lint ツールは製品の差別化のために MISRA-C 以外のルールもチェックできるようになっていて，最終的に数百ものルールをチェックすることになる。

## 守られないルールはルール自体に問題がある

問題は lint を実施した結果，大量の警告が出た場合である。
あるルールについて大量の逸脱が発生する理由は大きく2つある。

1. プログラム設計が根本的に間違っている
2. ルールが間違っている

プログラマのスキルによるかもしれないが，経験上こういうときに前者が理由であることはほとんどない。
なので，まずは「ルールが間違っている」のではないかと疑ってみるのが定石である。

基本的に「ルールは守られるべきもの」であるが，ルールもまた人間が考えたものであり，間違っている可能性を常に考慮すべきである。
何故なら人間は間違いを犯す生き物なのだから。

私は「悪法も法」という考え方には与しない。
「悪法は悪法」であり正すべきだ。
そして「悪法」の判断基準のひとつが「守られないルールはルール自体に問題がある」である。
ルールもリファクタリングの対象となる「コード」なのだ。

最初に挙げた[発表][ガードレール]では [reviewdog] を使って変更部分のみチェックする方法が紹介されていたが， [reviewdog] を常用するのは個人的にお勧めできない。
これが常態化すると「動いてるコードは触るな」という方向に行きがちで，それによってリファクタリングの機会を失うこととなり，最終的に[技術的負債]({{< ref "/remark/2020/12/technical-debt-and-hacker.md" >}} "技術的負債とハッカー")の返済が遅れることになる。
リファクタリングに厚いのが Go の持ち味なので，これを抑圧するような運用は避けるべきだろう。

ことほど左様にルール・メイキングというのは難しいのである。
できるなら，煩わしいことは機械に任せて楽しくコードを書きたいものである。

## ブックマーク

- [golangci-lint を GitHub Actions で使う]({{< ref "/golang/using-golangci-lint-action.md" >}})

[ガードレール]: https://speakerdeck.com/lmt_swallow/lets-build-security-guardrails-for-your-go-programs "Go をセキュアに書き進めるための「ガードレール」を整備しよう / Let's Build Security Guardrails For Your Go Programs! - Speaker Deck"
[go-ruleguard]: https://github.com/quasilyte/go-ruleguard "quasilyte/go-ruleguard: Define and run pattern-based custom linting rules."
[gosec]: https://github.com/securego/gosec "securego/gosec: Golang security checker"
[golangci-lint]: https://golangci-lint.run/
[reviewdog]: https://github.com/reviewdog/reviewdog "reviewdog/reviewdog: 🐶 Automated code review tool integrated with any code analysis tools regardless of programming language"

## 参考図書

{{% review-paapi "4542503461" %}} <!-- 組込み開発者におくるMISRA‐C:2004 -->
{{% review-paapi "B01CYDGUV8" %}} <!-- CODE VERSION 2.0 -->
