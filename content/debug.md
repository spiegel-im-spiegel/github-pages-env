+++
title = "画面確認用デバッグ・ページ"
date =  "2020-07-24T09:49:31+09:00"
description = "description"
image = "/images/attention/site.jpg"
tags = ["site", "web"]
pageType = "text"
draft = true

[scripts]
  mathjax = true
  mermaidjs = true
+++

## 見出し2
### 見出し3
#### 見出し4
##### 見出し5
###### 見出し6

## 各部品の表示

通常の \<p\>aragraph[^tn1].

[^tn1]: 脚注。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

*強調1* **強調2** `Code表示012` ~~打ち消し~~

{{< div-gen class="caution" >}}
!!! caution !!!
{{< /div-gen >}}

1. 番号付きリスト1
    - リンク : [WEB色見本 原色大辞典 - HTMLカラーコード](https://www.colordic.org/)
    - PDF : {{< pdf-file title="CRYPTREC Report 2018: 暗号技術評価委員会報告" link="https://www.cryptrec.go.jp/report/cryptrec-rp-2000-2018.pdf" >}}
2. 番号付きリスト2

### 定義リスト

定義1
: 内容1

定義2
: 内容2 

### フェンス付きコード・ブロック

```go {hl_lines=[3]}
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
    //Output: Hello World!
}
```

### 引用

{{< fig-quote type="markdown" title="著作権法" link="https://elaws.e-gov.go.jp/search/elawsSearch/elaws_search/lsg0500/detail?lawId=345AC0000000048" >}}
{{% quote %}}著作者は、その著作物の原作品に、又はその著作物の公衆への提供若しくは提示に際し、その実名若しくは変名を著作者名として表示し、又は著作者名を表示しないこととする権利を有する。その著作物を原著作物とする二次的著作物の公衆への提供又は提示に際しての原著作物の著作者名の表示についても、同様とする{{% /quote %}}。
{{< /fig-quote >}}

### 図表

{{< fig-img src="/images/attention/summer.jpg" title="In the beach" link="/images/attention/summer.jpg" width="640" >}}

{{< fig-youtube id="Yh1_wHdUx3Y" title="Comet NEOWISE from ISS - YouTube" >}}

{{< comparable-security-strengths >}}

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->

### 数式（MathJax）

エネルギーと質量には $E=mc^2$ の関係がある。

エネルギーと質量には $$E=mc^2$$ の関係がある。

### シーケンス図（Mermaid）

{{< fig-mermaid title="カバとカバン" >}}
sequenceDiagram
    カバ->>+カバン: あなた，泳げまして？
    カバン-->>-カバ: いえ
    カバ->>+カバン: 空は飛べるんですの？
    カバン-->>-カバ: いえ
    カバ->>+カバン: じゃあ，足が速いとか？
    カバン-->>-カバ: いえ
    カバ->>カバン: あなた，何にもできないのねぇ
    loop ひとりヘコむ
        カバン->>カバン: ううっ
    end
{{< /fig-mermaid >}}

### アイコン

CC Licenses : {{< cc-syms "cc" "pd" "zero" "by" "sa" "nc" "nc-jp" "nc-eu" "nd" >}}

各種サービス : {{< icons "github" "twitter" "twitter-sq" "medium" "tumblr" "instagram" "flickr" "facebook" "facebook-sq" "linkedin" "pocket" >}}

汎用 : {{< icons "check" "rss" "share" "email" "pdf" "css3" "erlang" "html5" "java" "nodejs" "php" "python" "rust" "swift" >}}







<!-- end of file -->
