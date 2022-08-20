+++
title = "名誉毀損とリンク"
date =  "2022-08-20T11:43:10+09:00"
description = "リンクは単なる道具・手段"
image = "/images/attention/kitten.jpg"
tags = [ "code", "law", "web", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

オーストラリアの話であるが Google の検索結果およびそのリンクによって名誉毀損を受けたとして裁判になっていたそうで，先日その判決結果が出たらしい。

発端はこんな感じ。

{{< fig-quote type="markdown" title="Linking to news doesn’t make Google liable for defamation, Australia court rules | Ars Technica" link="https://arstechnica.com/tech-policy/2022/08/linking-to-news-doesnt-make-google-liable-for-defamation-australia-court-rules/" lang="en" >}}
The [case](https://www.hcourt.gov.au/cases/case_m86-2021) relates to a Google search result that linked to a 2004 article published by The Age with the title, "Underworld loses valued friend at court." The article described Melbourne-based lawyer George Defteros, who was charged with conspiracy to murder and incitement to murder the day before it was published. The charge was withdrawn in 2005.

Defteros sued Google after becoming aware that a Google search of his name produced a link to the article and a snippet. Google refused to remove the article from search results despite a request from Defteros in 2016.
{{< /fig-quote >}}

これに対して下級裁判所は

{{< fig-quote type="markdown" title="Linking to news doesn’t make Google liable for defamation, Australia court rules | Ars Technica" link="https://arstechnica.com/tech-policy/2022/08/linking-to-news-doesnt-make-google-liable-for-defamation-australia-court-rules/" lang="en" >}}
Lower courts decided that Google "published the defamatory matter because the provision of the Search Result was instrumental to the communication of the content of the Underworld article to the user, in that it lent assistance to its publication," [according to a summary of today's ruling](https://cdn.hcourt.gov.au/assets/publications/judgment-summaries/2022/hca-27-2022-08-17.pdf) provided by the High Court of Australia.
{{< /fig-quote >}}

として Google 側に賠償金の支払いを命じた。
その後 (当然ながら) Google は上告したようで，最終的に

{{< fig-quote type="markdown" title="Linking to news doesn’t make Google liable for defamation, Australia court rules | Ars Technica" link="https://arstechnica.com/tech-policy/2022/08/linking-to-news-doesnt-make-google-liable-for-defamation-australia-court-rules/" lang="en" >}}
Google cannot be held liable for defamation simply for providing hyperlinks to other webpages, Australia's highest court [ruled today](https://eresources.hcourt.gov.au/downloadPdf/2022/HCA/27). By itself, providing a URL is not "participation in the communication of defamatory matter which happens to be at that address... In reality, **a hyperlink is merely a tool which enables a person to navigate to another webpage**," the High Court of Australia ruling said.
{{< /fig-quote >}}

という結果になったそうな。
ちなみに強調は私がやりました。

まぁ，当たり前っちゃあ当たり前なんだけど

{{< fig-quote type="markdown" title="Linking to news doesn’t make Google liable for defamation, Australia court rules | Ars Technica" link="https://arstechnica.com/tech-policy/2022/08/linking-to-news-doesnt-make-google-liable-for-defamation-australia-court-rules/" lang="en" >}}
Today's ruling could have been different if Google had been paid to promote The Age article. The appeal "does not present the occasion to consider whether the conclusion would be different in respect of those hyperlinks that, by agreement with a third party, are promoted by the appellant following a search request," the ruling said.
{{< /fig-quote >}}

とあるようにリンクそのものではなく，リンクを含むコンテンツの文脈とかその背後にある利害関係が重要かな，と思ったりする。
たとえば6年前に「[欧州司法裁、侵害コンテンツへのリンクを著作権侵害と判断](https://p2ptk.org/copyright/553 "欧州司法裁、侵害コンテンツへのリンクを著作権侵害と判断 | p2ptk[.]org")」という話があったが，この前提のもとに改めて記事を見直せば「なるほど」と思うかもしれない。

「リンク」という言葉だけを抜き出してシロかクロか論じることに意味はない，ということだろう。
願わくば日本のマスメディアもこうした視点で報道していただきたいものである。
