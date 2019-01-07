+++
date = "2015-11-07T22:01:10+09:00"
update = "2015-11-25T10:58:29+09:00"
description = "Git の repository を作る際に .gitignore をどうするかは悩みどころだが、 gibo というツールを使えば .gitignore の生成をかなり自動化できるらしい。"
draft = false
tags = ["git", "tools", "windows"]
title = "gibo による .gitignore 生成（再掲載）"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

以前 [Qiita で書いた](http://qiita.com/spiegel-im-spiegel/items/45a4619aafcacc161521)が，復習を兼ねて再掲載する。

Git の repository を作る際に `.gitignore` をどうするかは悩みどころだが（つか、大概は他の repository からコピってくるのだが）、 `gibo` というツールを使えば `.gitignore` の生成をかなり自動化できるらしい。

- [simonwhitaker/gibo](https://github.com/simonwhitaker/gibo)

導入方法は `README.md` に書かれているが，実体はスクリプトのみなので，面倒くさいならプラットフォームごとにスクリプト・ファイル（Windows なら `gibo.bat`）を取ってきてパスの通っているフォルダに放り込めばよい。
処理自体は簡単で， [github/gitignore] で公開されているテンプレートを取ってきて単純に連結しているだけである[^a]。

[^a]: とはいえ自分で処理を書くのは微妙に面倒くさいからね。このツールを作った方には感謝である。

以下は Windows での操作手順。
まずは動作確認と初期化から。

```bash
C:>gibo -v
gibo 1.0.4 by Simon Whitaker <sw@netcetera.org>
https://github.com/simonwhitaker/gitignore-boilerplates

C:>gibo -u
Cloning https://github.com/github/gitignore.git to C:\Users\username\AppData\Roaming\.gitignore-boilerplates
```

これで `C:\Users\username\AppData\Roaming\.gitignore-boilerplates` に [github/gitignore] の内容がセットされる。

使用するにはコマンドラインに言語名やフレームワーク名を並べればよい。

```bash
C:>gibo java eclipse
### java

*.class

# Mobile Tools for Java (J2ME)
.mtj.tmp/

# Package Files #
*.jar
*.war
*.ear

# virtual machine crash logs, see http://www.java.com/en/download/help/error_hotspot.xml
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

# STS (Spring Tool Suite)
.springBeans
```

出力は標準出力なので `.gitignore` にリダイレクトしてあげればよい。

```bash
C:>gibo java eclipse > .gitignore
```

既に `.gitignore` があるのなら追記で OK。

```bash
C:>gibo windows >> .gitignore
```

対応している言語等が知りたければ `-l` オプションを付けて起動すると一覧が表示される。

```bash
C:>gibo -l
=== Languages ===

Actionscript
Ada
Agda
Android
AppceleratorTitanium
AppEngine
ArchLinuxPackages
Autotools
C
C++
CakePHP
CFWheels
ChefCookbook
Clojure
CMake
CodeIgniter
CommonLisp
Composer
Concrete5
Coq
CraftCMS
CUDA
D
Dart
Delphi
DM
Drupal
Eagle
Elisp
Elixir
Elm
EPiServer
Erlang
ExpressionEngine
ExtJs
Fancy
Finale
ForceDotCom
Fortran
FuelPHP
Gcov
GitBook
Go
Gradle
Grails
GWT
Haskell
Idris
IGORPro
Java
Jboss
Jekyll
Joomla
KiCAD
Kohana
LabVIEW
Laravel
Leiningen
LemonStand
Lilypond
Lithium
Lua
Magento
Maven
Mercury
MetaProgrammingSystem
Nanoc
Nim
Node
Objective-C
OCaml
Opa
OpenCart
OracleForms
Packer
Perl
Phalcon
PlayFramework
Plone
Prestashop
Processing
Python
Qooxdoo
Qt
R
Rails
RhodesRhomobile
ROS
Ruby
Rust
Sass
Scala
SCons
Scrivener
Sdcc
SeamGen
SketchUp
Stella
SugarCRM
Swift
Symfony
SymphonyCMS
TeX
Textpattern
TurboGears2
Typo3
Umbraco
Unity
UnrealEngine
VisualStudio
VVVV
Waf
WordPress
Xojo
Yeoman
Yii
ZendFramework
Zephir

=== Global ===

Anjuta
Archives
BricxCC
Cloud9
CodeKit
CVS
DartEditor
Dreamweaver
Eclipse
EiffelStudio
Emacs
Ensime
Espresso
FlexBuilder
GPG
IPythonNotebook
JDeveloper
JetBrains
Kate
KDevelop4
Lazarus
LibreOffice
Linux
LyX
Matlab
Mercurial
MicrosoftOffice
ModelSim
Momentics
MonoDevelop
NetBeans
Ninja
NotepadPP
OSX
Otto
Redcar
Redis
SBT
SlickEdit
SublimeText
SVN
SynopsysVCS
Tags
TextMate
TortoiseGit
Vagrant
Vim
VirtualEnv
VisualStudioCode
WebMethods
Windows
Xcode
XilinxISE
```

## 追記

Ruby/Gem にも同様のツールがあるようです。

- [gemignore](https://rubygems.org/gems/gemignore "gemignore | RubyGems.org | your community gem host") 
    - [gitignoreの雛形を用意する - Qiita](http://qiita.com/nakaken0629/items/cd25b722d9eb15b4efcb)

[github/gitignore]: https://github.com/github/gitignore "github/gitignore"