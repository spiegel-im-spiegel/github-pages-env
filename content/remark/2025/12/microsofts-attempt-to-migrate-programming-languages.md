+++
title = "Microsoft によるプログラミング言語移行の試み"
date =  "2025-12-25T17:06:11+09:00"
description = "Microsoft のエンジニアである Galen Hunt 氏のチームによる研究プロジェクトらしい"
image = "/images/attention/kitten.jpg"
tags = [ "engineering", "programming", "language", "rust", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

昨日 Bluesky の私の TL に面白い話が流れてきて

- [Microsoft to Replace All C/C++ Code With Rust by 2030 - Thurrott.com](https://www.thurrott.com/dev/330980/microsoft-to-replace-all-c-c-code-with-rust-by-2030)

Microsoft は2030年末を目標にして，社内の C/C++ コードをすべて Rust に置き換える計画を立てていて，そのためのエンジニアを募集しているらしい。
その後も

- [Microsoft wants to replace its entire C and C++ codebase • The Register](https://www.theregister.com/2025/12/24/microsoft_rust_codebase_migration/)

といった記事も上がっている。

一方で

- [Microsoft is not rewriting Windows in Rust | InfoWorld](https://www.infoworld.com/article/4111553/microsoft-is-not-rewriting-windows-in-rust.html)

みたいな記事も上がっている。
これによると，全社的なプロジェクトではなく Microsoft のエンジニアである Galen Hunt 氏のチームによる研究プロジェクトらしい。

{{< fig-quote type="markdown" title="Microsoft is not rewriting Windows in Rust" link="https://www.infoworld.com/article/4111553/microsoft-is-not-rewriting-windows-in-rust.html" >}}
“My team’s project is a research project. We are building tech to make migration from language to language possible,” he wrote in an **[update to his LinkedIn](https://www.linkedin.com/posts/galenh_principal-software-engineer-coreai-microsoft-activity-7407863239289729024-WTzf/)** post. His intent, he said, was to find like-minded engineers, “not to set a new strategy for Windows 11+ or to imply that Rust is an endpoint.”
{{< /fig-quote >}}

Rust のようなメモリ安全なプログラミング言語を採用するという圧力はあるようで

{{< fig-quote type="markdown" title="Microsoft is not rewriting Windows in Rust" link="https://www.infoworld.com/article/4111553/microsoft-is-not-rewriting-windows-in-rust.html" >}}
Pressure to **[ditch C and C++ in favor of memory-safe languages](https://www.infoworld.com/article/2336216/white-house-urges-developers-to-dump-c-and-c.html)** such as Rust comes right from the top, with research by Google and Microsoft showing that around 70 percent of all security vulnerabilities in software are caused by memory safety issues.
{{< /fig-quote >}}

最初の方の記事はその辺の忖度が働いたのかねぇ，と邪推したりする。

Rust は随分前からメモリ安全なコンパイル言語として注目されているが，最近になって注目されているのはやはり生成 AI との「協働」が期待されているからだろう。

- [Rustは生成AI時代の覇権言語になる説](https://zenn.dev/ryugotoo/articles/691c08f844a5e1)

しかし

{{< fig-quote type="markdown" title="Microsoft is not rewriting Windows in Rust" link="https://www.infoworld.com/article/4111553/microsoft-is-not-rewriting-windows-in-rust.html" >}}
However, **[using AI to rewrite code](https://www.infoworld.com/article/3844351/how-to-simplify-app-migration-with-generative-ai-tools.html)**, even in a memory-safe language, may not make things more secure: AI-generated code typically contains more issues than code written by humans, according to **[research by CodeRabbit](https://www.infoworld.com/article/4109129/ai-assisted-coding-creates-more-problems-report.html)**.

That’s not stopping some of the biggest software developers **[pushing ahead with AI-powered software development](https://www.computerworld.com/article/3975705/from-prompts-to-production-ai-will-soon-write-most-code-reshape-developer-roles.html)**, though. Already, **[AI writes 30% of Microsoft’s new code](https://www.pcworld.com/article/2768784/microsoft-ceo-claims-30-of-new-code-is-written-by-ai.html)**, Microsoft CEO Satya Nadella said in April.

{{< /fig-quote >}}

などとも言われている。
「Microsoft の新規コードの30%を AI が書いている」と言うが，裏を返せば70%はまだ使い物にならないわけで，簡単でないのは確かであろう。
