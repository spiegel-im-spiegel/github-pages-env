+++
title = "Gocli Package v0.7.0 Released"
date = "2018-03-03T21:12:41+09:00"
update = "2018-03-04T10:50:47+09:00"
description = "v0.7.0 では SIGNAL 制御を行う gocli/signal サブパッケージを追加した。 具体的には context パッケージと組み合わせてキャンセル・イベントとして実装している。"
image = "/images/attention/tools.png"
tags  = [ "golang", "package", "cli" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[spiegel-im-spiegel/gocli] パッケージ v0.7.0 をリリースした。

- [spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface](https://github.com/spiegel-im-spiegel/gocli)

v0.7.0 では [SIGNAL] 制御を行う [`gocli`]`/signal` サブパッケージを追加した。
具体的には [context] パッケージと組み合わせてキャンセル・イベントとして実装している。

例えば，こんな感じで使う。

{{< highlight go "hl_lines=20-22 33" >}}
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spiegel-im-spiegel/gocli/signal"
)

func ticker(ctx context.Context) error {
	t := time.NewTicker(1 * time.Second) // 1 second cycle
	defer t.Stop()

	for {
		select {
		case now := <-t.C: // ticker event
			fmt.Println(now.Format(time.RFC3339))
		case <-ctx.Done(): // cancel event from context
			fmt.Println("Stop ticker")
			return ctx.Err()
		}
	}
}

func Run() error {
	errCh := make(chan error, 1)
	defer close(errCh)

	go func() {
		child, cancelChild := context.WithTimeout(
			signal.Context(context.Background(), os.Interrupt), // cancel event by SIGNAL
			10*time.Second,                                     // timeout after 10 seconds
		)
		defer cancelChild()
		errCh <- ticker(child)
	}()

	err := <-errCh
	fmt.Println("Done")
	return err
}

func main() {
	if err := Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
{{< /highlight >}}

Windows 環境では [SIGNAL] 周りのテストが出来ないので結構困ってたり。

[spiegel-im-spiegel/gocli] パッケージは CLI (Command-Line Interface) を組む際に（主に自分が）便利な細々とした機能を収録している。
他の人には使いにくいかもしれないし大した内容でもないため [CC0](https://creativecommons.org/publicdomain/zero/1.0/ "Creative Commons — CC0 1.0 Universal") で公開している。
一切の権利を放棄しているので自由に持っていって弄っていただいて構わない。

## ブックマーク

- [time.Ticker で遊ぶ]({{< ref "/golang/ticker.md" >}})
- [コマンドライン・インタフェースとファサード・パターン]({{< ref "/golang/cli-and-facade-pattern.md" >}})
- [Cobra の使い方とテスト]({{< ref "/golang/using-and-testing-cobra.md" >}})

[spiegel-im-spiegel/gocli]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[`gocli`]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[SIGNAL]: https://linuxjm.osdn.jp/html/LDP_man-pages/man7/signal.7.html "Man page of SIGNAL"
[context]: https://golang.org/pkg/context/ "context - The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
