+++
title = "Philosophers” を Go で解く（CLI）"
date =  "2025-09-03T18:34:53+09:00"
description = "実際に問題を解く前にコマンドライン・インタフェースを整えておく。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "concurrency", "games", "cli" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

## 目次

1. [“Philosophers” を Go で解く（問題編）]({{< ref "./philosophers-1.md" >}})
2. [“Philosophers” を Go で解く（CLI）]({{< ref "./philosophers-2.md" >}}) ← イマココ

## コマンドライン・インタフェース

実際に[問題]を解く前にコマンドライン・インタフェースを整えておく。
CLI には [`spf13/cobra`] パッケージを使用する。
詳細については，ちょっと古い内容だが「[Cobra でテストしやすい CLI を構成する](https://zenn.dev/spiegel/articles/20201018-cli-with-cobra-and-golang)」を参考にどうぞ。
この記事では結果だけ示す。

```text
$ go run github.com/spiegel-im-spiegel/philo -h
Simulation for philosophers' problem.

Usage:
  philo [flags]
  philo [command]

Available Commands:
  help        Help about any command
  version     print the version number

Flags:
      --debug   for debug
  -h, --help    help for philo

Use "philo [command] --help" for more information about a command.
```

なおリポジトリは以下にある。

- [Philosophers - I never thought philosophy would be so deadly](https://github.com/spiegel-im-spiegel/philo)

## パラメータの定義

[問題]で指示されているパラメータを今一度挙げる。

{{< fig-quote class="nobox" type="markdown" >}}
| シンボル | 型 | 説明 |
| :--- | :---: | :--- |
| `NumPhilos` | `int` | 哲学者およびフォークの数（$\ge 2$） |
| `TimeDie` | `int` | 餓死するまでの時間（$\gt 0\\,\mathrm{ms}$） |
| `TimeEat` | `int` | 「食べる」時間（$\gt 0\\,\mathrm{ms}$） |
| `TimeSleep` | `int` | 「眠る」時間（$\gt 0\\,\mathrm{ms}$） |
| `TimeThink` | `int` | 「考える」時間（$\ge 0\\,\mathrm{ms}$） |
| `NumEat` | `int` | 各哲学者が最低限食事を行う回数（$\ge 0$） |
{{< /fig-quote >}}

これをコマンドラインから設定できるようにする。
先に指定したパラメータを格納する型を定義しよう。

```go
// Params holds tunable configuration values for a Dining Philosophers simulation.
type Params struct {
	NumPhilos int // Total number of philosophers (and forks). Must be >= 2.
	TimeDie   int // Maximum time in milliseconds a philosopher can go without eating before being considered dead.
	TimeEat   int // Time in milliseconds a philosopher spends eating (holding both forks).
	TimeSleep int // Time in milliseconds a philosopher sleeps after eating.
	TimeThink int // Optional target time in milliseconds a philosopher thinks after sleeping.
	NumEat    int // Optional target number of meals each philosopher must consume; if == 0 the simulation runs until a death occurs.
}

// String returns a human-readable summary of the Params configuration,
// listing the number of philosophers, each lifecycle duration (die, eat,
// sleep, think in milliseconds), and the required number of meals.
// Example:
//
//	Philosophers: 5, Die: 800ms, Eat: 200ms, Sleep: 200ms, Think: 0ms, Eat Count: 3
func (p *Params) String() string {
	return fmt.Sprintf("Philosophers: %d, Die: %dms, Eat: %dms, Sleep: %dms, Think: %dms, Eat Count: %d",
		p.NumPhilos, p.TimeDie, p.TimeEat, p.TimeSleep, p.TimeThink, p.NumEat)
}
```

いやぁ，今どきは AI がコメントを書いてくれるから楽でいいね（笑）
これを前節の CLI と接続する。
こんな感じ。

```text {hl_lines=["15-20"]}
$ go run github.com/spiegel-im-spiegel/philo -h
Simulation for philosophers' problem.

Usage:
  philo [flags]
  philo [command]

Available Commands:
  help        Help about any command
  version     print the version number

Flags:
      --debug              for debug
  -h, --help               help for philo
  -m, --must-eat int       number of times each philosopher must eat (default 3)
  -p, --philosophers int   number of philosophers (default 5)
  -d, --to-die int         time to die (ms) (default 200)
  -e, --to-eat int         time to eat (ms) (default 20)
  -s, --to-sleep int       time to sleep (ms) (default 80)
  -t, --to-think int       time to think (ms) (default 80)

Use "philo [command] --help" for more information about a command.
```

デフォルト値は適当で。
今のところはパラメータの内容を表示するだけの処理にしておく。

```text
$ go run github.com/spiegel-im-spiegel/philo -h
Philosophers: 5, Die: 200ms, Eat: 20ms, Sleep: 80ms, Think: 80ms, Eat Count: 3
```

CLI 部分はここまで。
これで本来のシミュレーション・ロジックを組み込む準備は完了である。
次回から実際に問題を解いていく。

[Go]: https://go.dev/
[問題]: {{< ref "./philosophers-1.md" >}} "“Philosophers” を Go で解く（問題編）"
[`spf13/cobra`]: https://github.com/spf13/cobra "spf13/cobra: A Commander for modern Go CLI interactions"

## 参考図書

{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4908686122" %}} <!-- Goならわかるシステムプログラミング 第2版 -->
