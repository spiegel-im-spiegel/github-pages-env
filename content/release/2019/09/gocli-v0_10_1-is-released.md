+++
title = "spiegel-im-spiegel/gocli v0.10.1 のリリース"
date = "2019-09-07T11:55:58+09:00"
description = "書き直した結果 os.UserConfigDir() 関数で取得したパスにアプリケーション名と設定ファイル名をくっ付けただけの簡単なお仕事になった（笑）"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "cli" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Go 言語用 CLI プログラミング支援パッケージ [`gocli`] v0.10.1 をリリースした。

- [Release v0.10.1 · spiegel-im-spiegel/gocli · GitHub](https://github.com/spiegel-im-spiegel/gocli/releases/tag/v0.10.1)

実は『[改訂2版 みんなのGo言語](https://www.amazon.co.jp/exec/obidos/ASIN/B07VPSXF6N/baldandersinf-22/)』を読んで「そろそろ設定ファイルの置き場所を何とかしなくちゃ」と思ってこっそり [v0.10.0](https://github.com/spiegel-im-spiegel/gocli/releases/tag/v0.10.0 "Release v0.10.0 · spiegel-im-spiegel/gocli") をリリースしたのだが， [Go] 1.13 で [`os`]`.UserConfigDir()` 関数が用意されているのを知って「もうこれでいいぢゃん」てな感じになり，書き直した。こんな感じに使える。

```go
package main

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gocli/config"
)

func main() {
	path := config.Path("app", "config.json")
	fmt.Println(path)
	// Output:
	// /home/username/.config/app/config.json
}
```

書き直した結果 [`os`]`.UserConfigDir()` 関数で取得したパスにアプリケーション名と設定ファイル名をくっ付けただけの簡単なお仕事になった（笑）
私は設定ファイルと CLI との連携に [spf13/viper] を使っているので，これで必要十分。

[`gocli`] の使い方について詳しくは以下を参照のこと。

- [Go 言語用 CLI プログラミング支援パッケージ]({{< ref "/release/gocli-package-for-golang.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`gocli`]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[spf13/viper]: https://github.com/spf13/viper "spf13/viper: Go configuration with fangs"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%822%E7%89%88-%E3%81%BF%E3%82%93%E3%81%AA%E3%81%AEGo%E8%A8%80%E8%AA%9E-%E6%9D%BE%E6%9C%A8-%E9%9B%85%E5%B9%B8-ebook/dp/B07VPSXF6N?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07VPSXF6N"><img src="https://images-fe.ssl-images-amazon.com/images/I/51jif840ScL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%94%B9%E8%A8%822%E7%89%88-%E3%81%BF%E3%82%93%E3%81%AA%E3%81%AEGo%E8%A8%80%E8%AA%9E-%E6%9D%BE%E6%9C%A8-%E9%9B%85%E5%B9%B8-ebook/dp/B07VPSXF6N?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07VPSXF6N">改訂2版 みんなのGo言語</a></dt>
    <dd>松木 雅幸, mattn, 藤原 俊一郎, 中島 大一, 上田 拓也, 牧 大輔, 鈴木 健太</dd>
    <dd>技術評論社 2019-08-01 (Release 2019-08-01)</dd>
    <dd>Kindle版</dd>
    <dd>B07VPSXF6N (ASIN), 9784297107284 (EISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">改訂2版の目玉は7章の「データベースの扱い方」が追加されたことだろう。他の章では，大まかな構成は1版と同じだが細かい部分が変わっていて Go 1.12 への言及まであるのには驚いた。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-12">2019-08-12</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
