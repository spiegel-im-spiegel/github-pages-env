# 本家サイトリニューアルのメモ

## リダイレクト設定

```
# HTTP Redirect
Redirect permanent /policy.shtml /site-policy/
Redirect permanent /sitemap.shtml /sitemap/
Redirect permanent /archives.shtml /blog/
Redirect permanent /spiegel/log2/ /blog/
Redirect permanent /spiegel/remark/archives.shtml /blog/
Redirect permanent /spiegel/remark/archives/ /blog/
Redirect permanent /spiegel/remark/ /blog/
Redirect permanent /spiegel/profile/ /profile/
Redirect permanent /spiegel/pubkeys/ /pubkeys/
Redirect permanent /spiegel/docs/ /spiegel/archive/docs/
Redirect permanent /spiegel/maps/ /spiegel/archive/maps/
Redirect permanent /spiegel/archive/rdfa/ /spiegel/rdfa/
Redirect permanent /spiegel/archive/cc-license/ /spiegel/cc-license/
Redirect permanent /spiegel/archive/pgpdump/openpgp.shtml /spiegel/openpgp/
Redirect permanent /spiegel/archive/pgpdump/ /spiegel/pgpdump/
Redirect permanent /atom.xml https://text.baldanders.info/index.xml
Redirect permanent /rss.rdf.xml https://text.baldanders.info/index.xml
```

## ブログ URL の置換（正規表現）

検索文字列1: `https://baldanders.info/spiegel/log2/([0-9]*).shtml`
検索文字列2: `https://baldanders.info/remark/archives/([0-9]*).shtml`

置換文字列: `https://baldanders.info/blog/$1/`
