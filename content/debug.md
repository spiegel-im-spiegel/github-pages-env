+++
title = "画面確認用デバッグ・ページ"
date =  "2020-08-11T10:05:48+09:00"
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

{{< div-box type="markdown" >}}
- &#x845B; (U+845B)
- &#x845B;&#xE0100; (U+845B U+E0100)
{ .cloud .center .huge }
{{< /div-box >}}

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

{{< fig-quote type="markdown" title="著作権法" link="https://elaws.e-gov.go.jp/document?lawid=345AC0000000048" >}}
{{% quote %}}著作者は、その著作物の原作品に、又はその著作物の公衆への提供若しくは提示に際し、その実名若しくは変名を著作者名として表示し、又は著作者名を表示しないこととする権利を有する。その著作物を原著作物とする二次的著作物の公衆への提供又は提示に際しての原著作物の著作者名の表示についても、同様とする{{% /quote %}}。
{{< /fig-quote >}}

{{< fig-img-quote src="/images/attention/summer.jpg" title="In the beach | Flickr" link="https://www.flickr.com/photos/spiegel/32118886/" width="640" lang="en" >}}

### 図表

{{< fig-img src="/images/attention/summer.jpg" title="In the beach" link="/images/attention/summer.jpg" width="640" >}}

{{< fig-youtube id="Yh1_wHdUx3Y" title="Comet NEOWISE from ISS - YouTube" >}}

{{< comparable-security-strengths >}}

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->

### 数式（MathJax）

エネルギーと質量には $E=mc^2$ の関係がある。

エネルギーと質量には $$E=mc^2$$ の関係がある。

\[
C_p[\ce{H2O(l)}] = \pu{75.3 J // mol K}
\]

`<figure>` 要素で囲ってみる。

{{< fig-math >}}
\begin{aligned}
KL(\hat{y} || y) &= \sum_{c=1}^{M}\hat{y}_c \log{\frac{\hat{y}_c}{y_c}} \\
JS(\hat{y} || y) &= \frac{1}{2}(KL(y||\frac{y+\hat{y}}{2}) + KL(\hat{y}||\frac{y+\hat{y}}{2}))
\end{aligned}
{{< /fig-math >}}

さらに四角で囲ってみる。

{{< fig-math class="box" >}}
\begin{aligned}
KL(\hat{y} || y) &= \sum_{c=1}^{M}\hat{y}_c \log{\frac{\hat{y}_c}{y_c}} \\
JS(\hat{y} || y) &= \frac{1}{2}(KL(y||\frac{y+\hat{y}}{2}) + KL(\hat{y}||\frac{y+\hat{y}}{2}))
\end{aligned}
{{< /fig-math >}}

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

### フローチャート（Mermaid）

{{< fig-mermaid title="Web of trust" >}}
graph LR
  Alice["Alice's Public Key"]
  Bob["Bob's Public Key"]
  Chris["Chris's Public Key"]

  Alice-- Digital Sign -->Bob
  Alice-. validate! .->Chris
  Alice-. trust .->Bob
  Bob-- Digital Sign -->Alice
  Bob-- Digital Sign -->Chris
  Chris-- Digital Sign -->Bob
{{< /fig-mermaid >}}

### クラス図（Mermaid）

{{< fig-mermaid title="フレンズ" >}}
classDiagram
    friends<|--serval
    friends<|--raccoon
    friends<|--fennec
    serval : +Waai()
    raccoon : +Omakase-nanoda()
    fennec : +Haiyo()
{{< /fig-mermaid >}}

### Git コミット図（Mermaid）

```mermaid
%%{init: { 'gitGraph': {'showCommitLabel': false}} }%%
gitGraph
    commit tag:"v1.0.0" type:REVERSE
    branch featureA
    commit
    checkout main
    branch hotfix
    commit
    commit type: HIGHLIGHT
    checkout main
    merge hotfix
    commit tag:"v1.0.1"
    checkout featureA
    merge hotfix
    commit type: HIGHLIGHT
    checkout main
    merge featureA
    commit tag:"v1.1.0"
```

## アイコン

### CC Licenses

| コードと表示                                                               |
| -------------------------------------------------------------------------- |
| `{{</* cc-syms "cc" "pd" "zero" "by" "sa" "nc" "nc-jp" "nc-eu" "nd" */>}}` |
| {{< cc-syms "cc" "pd" "zero" "by" "sa" "nc" "nc-jp" "nc-eu" "nd" >}}       |

### 各種サービス

| コードと表示                                                                                                                            |
| --------------------------------------------------------------------------------------------------------------------------------------- |
| `{{</* icons "github" "twitter" "twitter-sq" "mastodon" "bluesky" "medium" "tumblr" "instagram" "flickr" "facebook" "facebook-sq" "linkedin" "pocket" */>}}` |
| {{< icons "github" "twitter" "twitter-sq" "mastodon"  "bluesky" "medium" "tumblr" "instagram" "flickr" "facebook" "facebook-sq" "linkedin" "pocket" >}}      |

### 汎用

| コードと表示                                                                                                                     |
| -------------------------------------------------------------------------------------------------------------------------------- |
| `{{</* icons "square" "square-check" "check" "rss" "share" "share-alt" "share-nodes" "email" "pdf" "css3" "erlang" "html5" "java" "nodejs" "php" "python" "rust" "swift" */>}}`    |
| {{< icons "square" "square-check" "check" "rss" "share" "share-alt" "share-nodes" "email" "pdf" "css3" "erlang" "golang" "html5" "java" "nodejs" "php" "python" "rust" "swift" >}} |


### 拡大・色付け

|                                  コードと表示                                  |
| :----------------------------------------------------------------------------: |
| `{{</* span class="fa-4x" */>}}{{</* cc-syms "cc" "by" */>}}{{</* /span */>}}` |
|        {{< span class="fa-4x" >}}{{< cc-syms "cc" "by" >}}{{< /span >}}        |

|                                       コードと表示                                        |
| :---------------------------------------------------------------------------------------: |
| `{{</* span class="fa-4x  golang-color" */>}}{{</* icons "golang" */>}}{{</* /span */>}}` |
|        {{< span class="fa-4x golang-color" >}}{{< icons "golang" >}}{{< /span >}}         |

### 絵文字

|            字形            | Short Code                       |
| :------------------------: | -------------------------------- |
|   {{< emoji "ゴメン" >}}   | `{{</* emoji "ゴメン" */>}}`     |
|   {{< emoji "ふむむ" >}}   | `{{</* emoji "ふむむ" */>}}`     |
|  {{< emoji "おやすみ" >}}  | `{{</* emoji "おやすみ" */>}}`   |
|    {{< emoji "おこ" >}}    | `{{</* emoji "おこ" */>}}`       |
|  {{< emoji "まじすか" >}}  | `{{</* emoji "まじすか" */>}}`   |
|   {{< emoji "はぁと" >}}   | `{{</* emoji "はぁと" */>}}`     |
|     {{< emoji "星" >}}     | `{{</* emoji "星" */>}}`         |
|    {{< emoji "電球" >}}    | `{{</* emoji "電球" */>}}`       |
|    {{< emoji "音符" >}}    | `{{</* emoji "音符" */>}}`       |
| {{< emoji "キーボード" >}} | `{{</* emoji "キーボード" */>}}` |
|  {{< emoji "はなまる" >}}  | `{{</* emoji "はなまる" */>}}`   |
|  {{< emoji "錠前と鍵" >}}  | `{{</* emoji "錠前と鍵" */>}}`   |
|  {{< emoji "蝶" >}}  | `{{</* emoji "蝶" */>}}`   |

|                            サイズ                             | コード                                                                      |
| :-----------------------------------------------------------: | --------------------------------------------------------------------------- |
|  {{< span class="mini" >}}{{< emoji "電球" >}}{{< /span >}}   | `{{</* span class="mini" */>}}{{</* emoji "電球" */>}}{{</* /span */>}}`    |
| {{< span class="smaller" >}}{{< emoji "電球" >}}{{< /span >}} | `{{</* span class="smaller" */>}}{{</* emoji "電球" */>}}{{</* /span */>}}` |
|                     {{< emoji "電球" >}}                      | `{{</* emoji "電球" */>}}`                                                  |
| {{< span class="larger" >}}{{< emoji "電球" >}}{{< /span >}}  | `{{</* span class="larger" */>}}{{</* emoji "電球" */>}}{{</* /span */>}}`  |
|  {{< span class="huge" >}}{{< emoji "電球" >}}{{< /span >}}   | `{{</* span class="huge" */>}}{{</* emoji "電球" */>}}{{</* /span */>}}`    |

<!-- end of file -->
