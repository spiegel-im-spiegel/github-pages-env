---
title: Unicode 部首の文字化けの話
closed: false
---



**✍️ Posted by Spiegel(@spiegel) on 2021/02/28**

  昨日（2021-02-27）だったかのとあるイベントで公開されたらしい資料：

  - [PDFのコピペが文字化けするのはなぜか？～CID/GIDと原ノ味フォント～](https://www.slideshare.net/trueroad_jp/pdfcidgid)

  これまでの復習のつもりで読んでたら，文字化けするのは康熙部首だけじゃなくて CJK 部首補助（U+2E80..U+2EFF）も含まれるらしい。なんだよそれ orz

  - [CJK Radicals Supplement - Wikipedia](https://en.wikipedia.org/wiki/CJK_Radicals_Supplement)
  - [CJK部首補助 - Wikipedia](https://ja.wikipedia.org/wiki/CJK%E9%83%A8%E9%A6%96%E8%A3%9C%E5%8A%A9)


-----------------------------------------

**✍️ Posted by Spiegel(@spiegel) on 2021/02/28**

  というわけで，早速以下のコードで確認

  ```go:sample3.go
  package main

  import (
      "fmt"

      "golang.org/x/text/unicode/norm"
  )

  func main() {
      for r := rune(0x2e80); r <= 0x2eff; r++ {
          rr := []rune(norm.NFKC.String(string([]rune{r})))
          if r != rr[0] {
              fmt.Printf("%#U -(NFKC)-> %#U\n", r, rr[0])
          }
      }
  }
  ```

  これを実行すると

  ```
  $ go run sample3.go 
  U+2E9F '⺟' -(NFKC)-> U+6BCD '母'
  U+2EF3 '⻳' -(NFKC)-> U+9F9F '龟'
  ```

  となった。まずい！ Unicode 正規化ではカバーしてないのか。

-----------------------------------------

**✍️ Posted by Spiegel(@spiegel) on 2021/02/28**

  なにか手がかりはないかなぁ，とググってみたら以下の記事を見つけた。

  - [やっかいな漢字 – CJK部首補助／康煕部首 – ものかの](https://tama-san.com/resolve-kanji/)

  この記事によると Unicode が公式に出している変換表があるそうな。

  - [EquivalentUnifiedIdeograph.txt](https://www.unicode.org/Public/UCD/latest/ucd/EquivalentUnifiedIdeograph.txt)

  でも，この変換表だと文字によってはグリフが大きく変わる場合があるらしい

  >これは「⺌」「⺍」を「小」にマッピングしていたりするので、見た目を変えたくないDTPでは使いにくいかもしれません。

  そこで上の記事では独自の[変換表を公開](https://app.box.com/s/3n450erc4jaibkvcrmnsfhhekbdcks58)されている。

  どうすっかなぁ。汎用で考えるなら Unicode 公式の変換表を使うのが筋だよね。ふむむむむー

-----------------------------------------
