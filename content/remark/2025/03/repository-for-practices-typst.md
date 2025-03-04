+++
title = "Typst 練習用のリポジトリを作った，他"
date =  "2025-03-03T20:52:58+09:00"
description = "Typst 練習用のリポジトリを作った / GitHub Copilot Pro に加入した"
image = "/images/attention/kitten.jpg"
tags = [ "typst", "github", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## Typst 練習用のリポジトリを作った

[Typst] のお勉強用に書いたコードをまとめたリポジトリを作った。

- [GitHub - spiegel-im-spiegel/practices-typst: Typst による組版の練習](https://github.com/spiegel-im-spiegel/practices-typst)

[MIT-0](https://spdx.org/licenses/MIT-0.html "MIT No Attribution | Software Package Data Exchange (SPDX)") ライセンスで公開しているので，再利用等ご自由にどうぞ（再利用する価値があるかどうかは別として）。

今後も思いついたことをチマチマ残していく予定。
ブログのほうは「[Typst に関する雑多な話]({{< ref "/typst/x-miscellaneous.md" >}})」を随時更新で書き足していくつもり。
これで [Typst] については一区切り付いたかな。

## GitHub Copilot Pro に加入した

今回 [Typst] で遊んでたら [GitHub Copilot] Free 版の制限いっぱいになってしまいまして。
10秒くらい考えて Pro 版に入ることにした。
とりあえず様子見で月 10USD ずつ払う。

やっぱ本格的に使うならお金を払わんとダメか。

主に VS Code でコーディング支援をしてもらってるが [Typst] と TypeScript で混乱してるのか，ウソの提案が多い。
大抵はコンパイルエラーになるので致命的ではないが。
ちゃんと統計をとってないけど体感で3割くらいしか正しくない感じ。
プロ野球選手なら打率3割でエース級だが，提案が7割使えないのはちょっと困る。
Go のコードならそんなに外れはないんだけどねぇ。

[Copilot][GitHub Copilot] に関してはローンチ当初から色々言われているが，最近でも

- [Thousands of exposed GitHub repositories, now private, can still be accessed through Copilot | TechCrunch](https://techcrunch.com/2025/02/26/thousands-of-exposed-github-repositories-now-private-can-still-be-accessed-through-copilot/)
- [GitHubで非公開にされたはずのリポジトリがMicrosoftのAIアシスタント「Copilot」を通じて公開されていたという指摘 - GIGAZINE](https://gigazine.net/news/20250303-github-private-repositories-accessed-copilot/)

みたいな話があって正直いまでもビミョーな気分なんだけど，なんだかんだと便利に使っている自分がいる。
生成 AI 周りはホンマ（かつての FinTech 流行時のように）山師みたいなのが多いのであまり深入りしないようにしているが，若い人が会議のテキスト起こしと議事録を作成するのに生成 AI を便利に使ってる話とか聞くと「もうそういう時代なんだなぁ」としみじみする。

まぁ，年寄りの冷水なので聞き流してください（笑）

## ブックマーク

- [GitHub、あらゆるエディタやIDEとGitHub Copilotとの統合を可能にする「Copilot Language Server SDK」を一般公開|CodeZine（コードジン）](https://codezine.jp/article/detail/20981)
- [GitHub for Beginners: How to get started with GitHub Copilot - The GitHub Blog](https://github.blog/ai-and-ml/github-copilot/github-for-beginners-how-to-get-started-with-github-copilot/)

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
