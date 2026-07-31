| ID  | タイトル | 深刻度 | 　　発見日　　 | 　最終更新日　 |
| --- | -------- | ------ | ------ | ---------- |
{{ range . }}| [{{ .ID }}]({{ .URI }}) | {{ .Title }} | {{ .Severity }} | {{ .DatePublic }} | {{ .DateUpdate }} |
{{ end }}

(Powerd by [JVN](https://jvn.jp/) & [jvnman](https://github.com/spiegel-im-spiegel/jvnman "spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management"))
