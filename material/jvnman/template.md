# {{ .Info.Title }}

[{{ .Info.ID }}]({{ .Info.URI }} "{{ .Info.Title }}") が発行されたので以下に書き出しておく。

% fig-quote title="{{ .Info.Title }}" link="{{ .Info.URI }}" %
## 脆弱性の概要

{{ .Info.Description }}

## 想定される影響

{{ .Info.Impact }}

### 影響を受ける製品

{{ range .Affects }}- {{ .Name }} / {{ .ProductName }} {{ .VersionNumber }}
{{ end }}

### 深刻度

{{ with .CVSS }}{{ if .Severity }}{{ .Severity }} ({{ .BaseScore }}) : {{ .BaseVector }}

{{ .BaseReport }}{{ else }}CVSSv3 評価なし{{ end }}{{ end }}

## 対策

{{ .Info.Solution }}

## 関連情報

{{ range .Relattions }}- {{ if .Name }}{{ .Name }} {{ end }}[{{ .VulinfoID }}]({{ .URL }}) {{ if .Title }}{{ .Title }}{{ end }}
{{ end }}

(Powerd by [JVN](https://jvn.jp/) & [jvnman](https://github.com/spiegel-im-spiegel/jvnman "spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management"))
% /fig-quote %
