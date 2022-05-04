+++
draft = false
description = "個人的に TypeScript や CoffeeScript の構文は好みじゃないし（もちろん仕事ならやりますよ）， Dart でそういったものの代わりになるのなら悪くないと思ったのだ。"
tags = [
  "programming",
  "language",
  "dart",
]
date = "2016-10-30T01:39:04+09:00"
title = "Dart 言語に関する覚え書き"

[author]
  github = "spiegel-im-spiegel"
  tumblr = ""
  facebook = "spiegel.im.spiegel"
  instagram = "spiegel_2007"
  url = "https://baldanders.info/profile/"
  license = "by-sa"
  flattr = ""
  flickr = "spiegel"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  avatar = "/images/avatar.jpg"
  twitter = "spiegel_2007"
+++

最初の頃は Dash とか呼ばれていたこの言語だが，正直に言ってあまり関心は高くなかった。
昔 Microsoft が JScript/VBScript や ActiveX で似たようなことをやろうとして大失敗したのを見ていたので「大失敗の2番煎じとか（笑）」という感じだったのだ。

でも

- [GoogleのDartプログラミング言語に再びスポットライトが…その高い生産性にまず社内で人気が盛り上がる | TechCrunch Japan](https://jp.techcrunch.com/2016/10/26/20161026googles-dart-programming-language-returns-to-the-spotlight/)

を見てちょっと調べてみることにした。
個人的に TypeScript や CoffeeScript の構文は好みじゃないし（もちろん仕事ならやりますよ）， [Dart] でそういったものの代わりになるのなら悪くないと思ったのだ。

[Dart] 環境は以下から取得できる。

- [Install Dart | Dart](https://www.dartlang.org/install)

Windows の場合は Chocolatey 経由で導入するかサード・パーティのインストーラが用意されている。
今回はサード・パーティのインストーラを使ってみた。
インストール後にコマンドプロンプト等で

```text
$ dart --version
Dart VM version: 1.20.1 (Wed Oct 12 15:07:45 2016) on "windows_x64"
```

てな感じになれば成功である。
さっそく，みんな大好き “Hello World” から

```dart
void main() {
    print('Hello World!');
}
```

これを DartVM 上で実行する。

```text
$ dart hello.dart
Hello World!
```

よーし，うむうむ，よーし。

じゃあ，これを JavaScript コードに変換してみようか。

```text
$ dart2js -ohello.js hello.dart
Dart file (hello.dart) compiled to JavaScript: hello.js
```

うひゃ！ なんか凄いコード吐いたな。
元の3行のコードに対して300行くらいあるぞ。
でも node.js[^n] に食わせると一応ちゃんと出力される。

[^n]: [node.js](https://nodejs.org/) は最近 v7.0.0 が出ている。

```text
$ node hello.js
Hello World!
```

出力された JavaScript コードの一部を抜粋するとこんな感じのコードを吐いている。

```javascript
var dart = [["dart2js._js_primitives", "dart:_js_primitives",, H, {
  "^": "",
  printString: function(string) {
    if (typeof dartPrint == "function") {
      dartPrint(string);
      return;
    }
    if (typeof console == "object" && typeof console.log != "undefined") {
      console.log(string);
      return;
    }
    if (typeof window == "object")
      return;
    if (typeof print == "function") {
      print(string);
      return;
    }
    throw "Unable to print message: " + String(string);
  }
}], ["", "hello.dart",, G, {
  "^": "",
  main: function() {
    H.printString("Hello World!");
  }
}, 1]];
setupProgram(dart, 0);
```

いや，うーん。
いいのか，これ。
“Hello World” ごときでこれって，もう少しスリムなコードを吐けないのだろうか。

たとえば仕事で [Dart] を使うようなプロジェクトでもあれば面白そうだが，それ以外で積極的に使おうという気にはならないかなぁ，これは。
制御用の言語を統一するのなら JavaScript (ES6) で十分だし，バックエンド側のみということであれば [Go 言語]のほうがよさ気だし。
うーん。

というわけで，もう少し様子見。

## ブックマーク

- [Dart programming language | Dart](https://www.dartlang.org/)
- [Dart · GitHub](https://github.com/dart-lang)
- [Dartプログラミング言語仕様書邦訳版](http://www.cresc.co.jp/tech/java/Google_Dart/DartLanguageSpecification_about.html)
- [dartrefjp](https://sites.google.com/site/dartrefjp/)
- [小山博史のJavaを楽しむ（16）：JavaとJavaScriptの良いとこ取り？ 「Dart」超入門 (1/3) - ＠IT](http://www.atmarkit.co.jp/ait/articles/1208/29/news120.html)
- [Dart入門してみる。インストール～ブラウザでHelloWorldまで - Qiita](http://qiita.com/alucky0707/items/76aaf819a86eda7d6c4d)
- [2015年にDart言語はどう変わってどこに向かっていってるのか - Qiita](http://qiita.com/takyam/items/3dd2c1948f1fa7968a01)
- [Dart 1.19リリースノート - Qiita](http://qiita.com/sh4869/items/55d1ad5cd011113ed543)

[Dart]: https://www.dartlang.org/ "Dart programming language | Dart"
[Go 言語]: https://golang.org/ "The Go Programming Language"
