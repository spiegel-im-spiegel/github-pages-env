とりあえず，メモ。

- [Go 1.5 is released - The Go Blog](https://blog.golang.org/go1.5)
- [Go 1.5 Release Notes - The Go Programming Language](https://golang.org/doc/go1.5)

[ブックマーク](http://qiita.com/spiegel-im-spiegel/items/98d49ac456485b007a15)からこちらも挙げてくか。

- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)
- [Heroku、Go言語の正式サポートを発表 － Publickey](http://www.publickey1.jp/blog/15/herokugo.html)
- [Google App Engineも「Go言語」の正式サポートを発表 － Publickey](http://www.publickey1.jp/blog/15/google_app_enginego_1.html)
- [golang - GVM で go1.5rc1 のインストール - Qiita](http://qiita.com/msaito3/items/3aef86e9864990b16b4c)

## Windows では上書きインストールできない？

Windows 版（64bit）のインストールパッケージでインストールしようとしたらエラーになった。

[![Install Error](https://farm6.staticflickr.com/5759/20692976166_a38bee50d8_o.png "Install Error")](https://www.flickr.com/photos/spiegel/20692976166/)

この場合は，コントロールパネルの「プログラムと機能」で既存のバージョンをアンインストールしてから最新バージョンをインストールし直せば OK。ただし，バージョンアップによる副作用のほどは不明。（きっと識者の方がそのうち記事を出すだろう）

```shell
C:>go version
go version go1.5 windows/amd64
```
