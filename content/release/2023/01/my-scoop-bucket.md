+++
title = "オレオレ Scoop Bucket を作ってみた"
date =  "2023-01-08T15:10:56+09:00"
description = "これならボクにもできそう"
image = "/images/attention/tools.png"
tags  = [ "tools", "windows", "install", "scoop" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[hymkor](https://zenn.dev/zetamatta) さんの

- [おれおれ scoop バケットを作ろう｜scoop / nyagos で始めるコマンドライン生活](https://zenn.dev/zetamatta/books/5ac80a9ddb35fef9a146/viewer/cccccc)

を見て「これならボクにもできそう」と思ったので作ってみた。

まずは，[テンプレート・リポジトリ](https://github.com/ScoopInstaller/BucketTemplate "ScoopInstaller/BucketTemplate: Template Bucket for Scoop Installer")の GitHub ページで “Use this template” → “Create a new repository” で一式を fork する。
既定ブランチが `master` なのでご注意を。

あとは `README.md` を適当に修正して `bucket/` フォルダにマニフェスト・ファイル（`*.json`）を入れるだけ。

マニフェスト・ファイルはパッケージごとに作成する。
たとえば，拙作 [gpgpdump] ならファイル名を `bucket/gpgpdump.json` として，こんな感じに記述する。

```json
{
    "version": "0.15.0",
    "description": "OpenPGP packet visualizer",
    "homepage": "https://github.com/goark/gpgpdump",
    "license": "Apache-2.0",
    "architecture": {
        "64bit": {
            "url": "https://github.com/goark/gpgpdump/releases/download/v0.15.0/gpgpdump_0.15.0_Windows_64bit.zip",
            "hash": "be818119dc650f245aa8665f1af155b9d14c17c70e617517e817d81acb244151"
        },
        "arm64": {
            "url": "https://github.com/goark/gpgpdump/releases/download/v0.15.0/gpgpdump_0.15.0_Windows_ARM64.zip",
            "hash": "73d999250dc4a03b2298aa88271a34db2ff1cd3013428243b2e28afaed95aa5e"
        }
    },
    "bin": "gpgpdump.exe",
    "checkver": {
        "url": "https://github.com/goark/gpgpdump",
        "regex": "tag/([\\d\\._]+)"
    },
    "autoupdate": {
        "architecture": {
            "64bit": {
                "url": "https://github.com/goark/gpgpdump/releases/download/v$version/gpgpdump_$version_Windows_64bit.zip"
            },
            "arm64": {
                "url": "https://github.com/goark/gpgpdump/releases/download/v$version/gpgpdump_$version_Windows_ARM64.zip"
            }
        }
    }
}
```

シングルバイナリなので簡単！

必要な数だけマニフェスト・ファイルができたら commit & push で GitHub に送る。
すると GitHub Actions が走り出すので，終わるまで待ってエラーが出なければ無問題である。

早速 Windows 環境で試してみよう。

まずは，既存の bucket に [gpgpdump] がないことを確認する。

```text
$ scoop search gpgpdump
WARN  No matches found.
```

次に，作成したオレオレ buket を登録する。
そのまま GitHub リポジトリのパスを指示すれば OK。

```text
$ scoop bucket add goark https://github.com/goark/scoop-bucket
Checking repo... OK
The goark bucket was added successfully.
```

うんうん。
ちなみに bucket 名は（名前が被らなければ）任意に指定できる[^scoop1]。

[^scoop1]: `scoop bucket known` で表示されるリストは公式（？）の bucket 名なので，それ以外の名前を使うのがいいだろう。 Bucket を削除する場合は `scoop bucket rm <bucket name>` でいける。

この状態でもう一度 [gpgpdump] を探してみる。

```text
$ scoop search gpgpdump
Results from local buckets...

Name     Version Source Binaries
----     ------- ------ --------
gpgpdump 0.15.0  goark
```

おー。
あるやないかい。
では実際にインストールして動作確認してみよう。

```text
$ scoop install gpgpdump
Installing 'gpgpdump' (0.15.0) [64bit] from goark bucket
gpgpdump_0.15.0_Windows_64bit.zip (3.1 MB) [==================================================================] 100%
Checking hash of gpgpdump_0.15.0_Windows_64bit.zip ... ok.
Extracting gpgpdump_0.15.0_Windows_64bit.zip ... done.
Linking ~\scoop\apps\gpgpdump\current => ~\scoop\apps\gpgpdump\0.15.0
Creating shim for 'gpgpdump'.
'gpgpdump' (0.15.0) was installed successfully!

$ gpgpdump.exe version
gpgpdump v0.15.0
repository: https://github.com/goark/gpgpdump
```

よーし，うむうむ，よーし。

今のところ以下のツールを登録している。

- [gpgpdump](https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer") - OpenPGP packet visualizer
- [depm](https://github.com/goark/depm "goark/depm: Visualize depndency packages and modules") - Visualize depndency packages and modules
- [gnkf](https://github.com/goark/gnkf "goark/gnkf: Network Kanji Filter by Golang") - Network Kanji Filter by Golang
- [ml](https://github.com/goark/ml "goark/ml: Make Link with Markdown Format") - Make Link with Markdown Format

問題は登録しているツールのバージョンを上げるたびに対応するマニフェスト・ファイルも更新しないといけないところかな。
手で直すのは手間なので，何らかのバッチ処理を考えないとな。

## 【2022-01-09 追記】make-scoop-manifest でマニフェスト・ファイルを生成する

- [hymkor/make-scoop-manifest: Make scoop manifest （そのまんまや！）](https://github.com/hymkor/make-scoop-manifest)

これを使えばマニフェスト・ファイルが簡単に生成できるようだ。
ありがたや {{< emoji "ペコン" >}}

私の場合は [Usage-3](https://github.com/hymkor/make-scoop-manifest#usage-3) でいけそうかな[^msm1]。

[^msm1]: 私は [GoReleaser](https://goreleaser.com/) + [GitHub Actions](https://github.com/goreleaser/goreleaser-action "goreleaser/goreleaser-action: GitHub Action for GoReleaser") でバイナリを自動生成してしまうので，手元にバイナリはないのだ。

Windows 用のバイナリは[提供されている](https://github.com/hymkor/make-scoop-manifest/releases "Releases · hymkor/make-scoop-manifest")が，他のプラットフォーム用バイナリはないので，とりあえず以下のコマンドラインで試してみた。

```text
$ go run github.com/hymkor/make-scoop-manifest@latest -D -g goark/gpgpdump
go: downloading github.com/hymkor/make-scoop-manifest v0.2.0
/tmp/go-build3806744147/b001/exe/make-scoop-manifest  for linux/amd64 by go1.19.4
Get: https://api.github.com/repos/goark/gpgpdump/releases
Download: https://github.com/goark/gpgpdump/releases/download/v0.15.0/gpgpdump_0.15.0_Windows_64bit.zip
Download: https://github.com/goark/gpgpdump/releases/download/v0.15.0/gpgpdump_0.15.0_Windows_ARM64.zip
Get: https://api.github.com/repos/goark/gpgpdump
{
    "version": "0.15.0",
    "description": "OpenPGP packet visualizer",
    "homepage": "https://github.com/goark/gpgpdump",
    "license": "Apache License 2.0",
    "architecture": {
        "64bit": {
            "url": "https://github.com/goark/gpgpdump/releases/download/v0.15.0/gpgpdump_0.15.0_Windows_64bit.zip",
            "hash": "be818119dc650f245aa8665f1af155b9d14c17c70e617517e817d81acb244151"
        },
        "arm64": {
            "url": "https://github.com/goark/gpgpdump/releases/download/v0.15.0/gpgpdump_0.15.0_Windows_ARM64.zip",
            "hash": "73d999250dc4a03b2298aa88271a34db2ff1cd3013428243b2e28afaed95aa5e"
        }
    },
    "bin": [
        "gpgpdump.exe"
    ],
    "checkver": {
        "github": "https://github.com/goark/gpgpdump",
        "regex": "tag/([\\d\\._]+)"
    },
    "autoupdate": {
        "architecture": {
            "64bit": {
                "url": "https://github.com/goark/gpgpdump/releases/download/v$version/gpgpdump_$version_Windows_64bit.zip"
            },
            "arm64": {
                "url": "https://github.com/goark/gpgpdump/releases/download/v$version/gpgpdump_$version_Windows_ARM64.zip"
            }
        }
    }
}
```

おー。
いけるいける。

というわけで，早速 [Task] 用の設定ファイルを書いて一括処理できるようにする。
こんな感じ。

```yaml
version: '3'

tasks:
  default:
    deps: [depm, gnkf, gpgpdump, ml]

  depm:
    dir: bucket/
    deps: [build-make-scoop-manifest]
    cmds:
      - make-scoop-manifest -D -g goark/depm > depm.json
    generates:
      - depm.json

  gnkf:
    dir: bucket/
    deps: [build-make-scoop-manifest]
    cmds:
      - make-scoop-manifest -D -g goark/gnkf > gnkf.json
    generates:
      - gnkf.json

  gpgpdump:
    dir: bucket/
    deps: [build-make-scoop-manifest]
    cmds:
      - make-scoop-manifest -D -g goark/gpgpdump > gpgpdump.json
    generates:
      - gpgpdump.json

  ml:
    dir: bucket/
    deps: [build-make-scoop-manifest]
    cmds:
      - make-scoop-manifest -D -g goark/depm > ml.json
    generates:
      - ml.json

  build-make-scoop-manifest:
    desc: Build make-scoop-manifest command.
    run: once
    cmds:
      - go install github.com/hymkor/make-scoop-manifest@latest
```

もうちょっとスマートに書きたいが ...まぁいいか。

これでバージョンアップの度に手書きで更新しなくてよくなった。
めでたし！

## ブックマーク

- [scoop / nyagos で始めるコマンドライン生活](https://zenn.dev/zetamatta/books/5ac80a9ddb35fef9a146)
- [GitHub - hymkor/scoop-bucket](https://github.com/hymkor/scoop-bucket)
- [Scoop で利用できる Bucket の解説（`scoop bucket known` 限定） - Qiita](https://qiita.com/nimzo6689/items/5ead753169dbad72e4eb)
- [GitHubで自分以外の人がmainブランチに直接PUSHするのを禁止する](https://zenn.dev/ttani/articles/github-approval-self)

[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[Task]: https://taskfile.dev/
