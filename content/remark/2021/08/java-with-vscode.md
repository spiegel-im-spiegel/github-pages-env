+++
title = "Java と VS Code"
date =  "2021-08-18T22:20:30+09:00"
description = "どうせなら VS Code で環境を作るのがいいよねってことで"
image = "/images/attention/kitten.jpg"
tags = ["vscode", "editor", "java", "web"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

1. [パソコンに Visual Studio Code を導入する（再チャレンジ）]({{< ref "/remark/2021/02/installing-vscode-again.md" >}})
2. [Go と VS Code]({{< ref "/remark/2021/02/golang-with-vscode.md" >}})
3. [Markdown と VS Code]({{< ref "/remark/2021/02/markdown-with-vscode.md" >}})
4. [Java と VS Code]({{< ref "/remark/2021/08/java-with-vscode.md" >}}) ← イマココ

仕事で Spring Boot なコードを書くことになって，今更ながら基本から勉強し直している。
で，どうせなら [VS Code] で環境を作るのがいいよねってことで，覚え書きとして記しておく。

## [Extension Pack for Java]

これをインストールすると以下のパッケージも併せてインストールされる。

- [Language Support for Java(TM) by Red Hat - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=redhat.java)
- [Debugger for Java - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=vscjava.vscode-java-debug)
- [Test Runner for Java - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=vscjava.vscode-java-test)
- [Maven for Java - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=vscjava.vscode-maven)
- [Project Manager for Java - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=vscjava.vscode-java-dependency)
- [Visual Studio IntelliCode - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=VisualStudioExptTeam.vscodeintellicode)

これだけあれば基本的な道具は一通り揃う。

### みんな大好き “Hello, World!”

ここで簡単に動作確認しておこう。
以下の手順でプロジェクトの雛形を作成する。

1. コマンドパレットから “Java: Create Java Project” を選択する
2. プロジェクト・タイプの一覧が表示されるので “No build tools” を選択する
3. 作業ディレクトリを選択する
4. プロジェクト名を指定する（ここでは `hello` と入力）

これで

```text
$ tree hello
hello
├── README.md
├── lib
└── src
    └── App.java

2 directories, 2 files
```

という感じにディレクトリ・ファイルが作成される。
`App.java` の中身はこんな感じ。

```java
public class App {
    public static void main(String[] args) throws Exception {
        System.out.println("Hello, World!");
    }
}
```

エディタ上は

{{< fig-img src="./hello-java-code.png" link="/hello-java-code.png" width="528" >}}

という感じに表示されているので `main()` 関数直上の `Run` のリンクをクリックすればコンパイル＆実行してくれる。
Java Process Console に `Hello, World!` と表示されれば無問題。

## [Spring Boot Extension Pack]

これも以下のパッケージを含んでいるようだ。

- [Spring Boot Tools - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=Pivotal.vscode-spring-boot)
- [Spring Initializr Java Support - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=vscjava.vscode-spring-initializr)
- [Spring Boot Dashboard - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=vscjava.vscode-spring-boot-dashboard)

### 雛形プロジェクトを作る

これも動作確認しておこう。

1. コマンドパレットから “Spring Initializr: Generate a Maven Project” を選択する
2. Spring Boot のバージョンを選択（2021-08-18 時点の最新は 2.5.3）
3. 使用言語を選択。もちろん Java で
4. Group Id を入力。ここはデフォルトの `com.example` のままにしておく
5. Artifact Id を入力。ここもデフォルトの `demo` のままにしておく
6. パッケージタイプを選択。 `Jar` と `War` がある。とりあえす `Jar` にしておこうか
7. Java のバージョンを選択。無難に LTS 版の 11 を選択しておくか（JDK のインストールは別途行うこと）
8. 依存パッケージを選択。 Spring Web と Lombok は必須。あとは必要に応じて
   - Spring Web (必須)
   - Lombok (必須)
   - Spring Boot DevTools
   - Thymeleaf

{{< fig-img src="./choose-dependencies.png" link="/choose-dependencies.png" width="630" >}}

あとは作業ディレクトリを指定すれば完了。
作業ディレクトリ直下に Artifact Id で指定した名前でディレクトリが掘られ，ディレクトリ・ファイルが展開される。

```text
$ tree demo
demo
├── HELP.md
├── mvnw
├── mvnw.cmd
├── pom.xml
└── src
    ├── main
    │   ├── java
    │   │   └── com
    │   │       └── example
    │   │           └── demo
    │   │               └── DemoApplication.java
    │   └── resources
    │       ├── application.properties
    │       ├── static
    │       └── templates
    └── test
        └── java
            └── com
                └── example
                    └── demo
                        └── DemoApplicationTests.java
```

`DemoApplication.java` の中身はこんな感じ。

```java
package com.example.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class DemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}

}
```

このままだと何も表示できないので controller クラスと対応するテンプレートファイルを用意する。

まずは `demo/src/main/java/com/example/demo/controller/DemoController.java` ファイルを作る。
中身はこんな感じ。

```java
package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class DemoController {

	@GetMapping("/")
	public String demo() {
		return "demo";
	}

}
```

次に `demo/src/main/resources/templates/demo.html` ファイルを作る。
中身はこんな感じ。

```html
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Demo</title>
</head>
<body>
  <h1>Hello World!</h1>
</body>
</html>
```

全体の構成はこんな感じ。

```text
$ tree demo
demo
├── HELP.md
├── mvnw
├── mvnw.cmd
├── pom.xml
└── src
    ├── main
    │   ├── java
    │   │   └── com
    │   │       └── example
    │   │           └── demo
    │   │               ├── DemoApplication.java
    │   │               └── controller
    │   │                   └── DemoController.java
    │   └── resources
    │       ├── application.properties
    │       ├── static
    │       └── templates
    │           └── demo.html
    └── test
        └── java
            └── com
                └── example
                    └── demo
                        └── DemoApplicationTests.java

```

テストとか端折ってるけどご容赦ね。

あとは Spring Boot Dashboard から Start すれば OK。

{{< fig-img src="./dashboard.png" link="/dashboard.png" >}}

Web ブラウザから `http://localhost:8080/` を叩いて `Hello, World!` と表示されれば無問題。

## ブックマーク

- [VSCodeで作るJava開発環境＆Spring Bootアプリケーション入門 - Qiita](https://qiita.com/takumi_links/items/fe71cfeb4dfaa76fbe31)
- [Spring Boot アプリケーションのデプロイ - リファレンス](https://spring.pleiades.io/spring-boot/docs/current/reference/html/deployment.html)
- [SpringBootをVSCodeで使ってみる - Qiita](https://qiita.com/koukibuu3/items/77734596483ffd788931)
- [【Maven編】Spring Bootのビルドとデプロイ方法 | こへいブログ](https://kohei.life/spring-boot-build-deploy/)
- [VSCodeとDockerでSpring Boot + PostgreSQL開発環境を作る(2) | Sales8開発者日記](https://ameblo.jp/kazusa-g/entry-12536838291.html)
- [SpringBoot ログイン画面作成](https://learning-collection.com/springboot-%e3%83%ad%e3%82%b0%e3%82%a4%e3%83%b3%e7%94%bb%e9%9d%a2%e4%bd%9c%e6%88%90/)
- [Spring Boot + Spring Data JPA ～サンプルアプリ実装～ - Qiita](https://qiita.com/t-shin0hara/items/eaf44e4f48341616ab97)
- [JPA (Java Persistence API)のアノテーション](http://itref.fc2web.com/java/jpa/annotation.html)
- [Thymeleafチートシート - Qiita](https://qiita.com/NagaokaKenichi/items/c6d1b76090ef5ef39482)
- [SpringBootアプリにBootstrap4を追加（WebJars使用） – One IT Thing](https://one-it-thing.com/2074/)
- [SpringブートSLF4Jロギングの例](https://www.codeflow.site/ja/article/spring-boot__spring-boot-slf4j-logging-example)
- [SpringブートSLF4Jロギングの例 - 開発者ドキュメント](https://ja.getdocs.org/spring-boot-spring-boot-slf4j-logging-example/)
- [Spring bootでHttpSessionを使用する - m_shige1979のささやかな抵抗と欲望の日々](https://m-shige1979.hatenablog.com/entry/2016/11/30/080000)
- [Spring Boot で Ajax を実装する単純なサンプル - Qiita](https://qiita.com/t-yama-3/items/572fabc873b4b6a0fc7c)



[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[OpenJDK]: http://openjdk.java.net/
[Extension Pack for Java]: https://marketplace.visualstudio.com/items?itemName=vscjava.vscode-java-pack "Extension Pack for Java - Visual Studio Marketplace"
[Spring Boot Extension Pack]: https://marketplace.visualstudio.com/items?itemName=Pivotal.vscode-boot-dev-pack "Spring Boot Extension Pack - Visual Studio Marketplace"

## 参考図書

{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
