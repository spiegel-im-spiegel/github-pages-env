# gibo (gitignore-boilerplates) による .gitignore 生成（Windows でもできるよ）

git の repository を作る際に `.gitignore` をどうするかは悩みどころだが（つか、他の repository からコピってくることが多いのだが）、 `gibo` というツールを使えば `.gitignore` の生成をかなり自動化できるらしい。

- [simonwhitaker/gibo](https://github.com/simonwhitaker/gibo)

導入は `README.md` に書かれているのでそちらを参考にすればいいのだが、実体はスクリプトなので、まぁどうにでもなるだろう。

Windows の場合は scoop を使わなくても repository を clone し、そのフォルダにパスを通せばいい。

導入したら動作確認。

```shell
C:>gibo -v
gibo 1.0.4 by Simon Whitaker <sw@netcetera.org>
https://github.com/simonwhitaker/gitignore-boilerplates

C:>gibo -u
Cloning https://github.com/github/gitignore.git to C:\Users\username\AppData\Roaming\.gitignore-boilerplates
```

これで `C:\Users\username\AppData\Roaming\.gitignore-boilerplates` に情報がセットされる。使用するにはコマンドラインに言語名やフレームワーク名を並べればよい。

```shell
c:>gibo java eclipse
### java

*.class

# Mobile Tools for Java (J2ME)
.mtj.tmp/

# Package Files #
*.jar
*.war
*.ear

# virtual machine crash logs, see http://www.java.com/en/download/help/error_hot
spot.xml
hs_err_pid*


### eclipse

*.pydevproject
.metadata
.gradle
bin/
tmp/
*.tmp
*.bak
*.swp
*~.nib
local.properties
.settings/
.loadpath

# Eclipse Core
.project

# External tool builders
.externalToolBuilders/

# Locally stored "Eclipse launch configurations"
*.launch

# CDT-specific
.cproject

# JDT-specific (Eclipse Java Development Tools)
.classpath

# Java annotation processor (APT)
.factorypath

# PDT-specific
.buildpath

# sbteclipse plugin
.target

# TeXlipse plugin
.texlipse
```

出力は標準出力なので適当にリダイレクトしてあげればよい。

```shell
C:>gibo java eclipse > .gitignore
```

対応している言語等が知りたければ `-l` オプションを付けて起動すると一覧が表示される（めちゃめちゃ多いので動作例はパス）。
