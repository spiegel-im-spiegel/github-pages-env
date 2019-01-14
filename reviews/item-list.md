| ASIN | Title | Author | Binding | URL |
| ---- | ----- | ------ | ------- | --- |
{{ range .Items }}| {{ .ASIN }} | {{ .ItemAttributes.Title }} |{{ range .ItemAttributes.Author }} {{ . }}{{ end }} | {{ .ItemAttributes.Binding }} | {{ .URL }} |
{{ end }}
