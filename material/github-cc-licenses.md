# GitHub リポジトリに CC Licenses を設定したい

Web 上で GitHub リポジトリを作成する際に “Add a license” で各種オープンソース・ライセンスを設定できるが，オープンソース・ライセンス以外の「自由のライセンス」は設定できない。この記事では GitHub リポジトリに [CC Licenses] を設定する方法を紹介する。

GitHub リポジトリのライセンスはリポジトリ直下の `LICENSE` ファイルで指定できる。そこで [CC Licenses] の legal code を `LICENSE` ファイルとしてセットする。

[CC Licenses] の legal code のテキスト版は以下から取得できる。

- [BY 4.0 (plaintext)](https://creativecommons.org/licenses/by/4.0/legalcode.txt)
- [BY-SA 4.0 (plaintext)](https://creativecommons.org/licenses/by-sa/4.0/legalcode.txt)
- [BY-NC 4.0 (plaintext)](https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt)
- [BY-NC-SA 4.0 (plaintext)](https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode.txt)
- [BY-ND 4.0 (plaintext)](https://creativecommons.org/licenses/by-nd/4.0/legalcode.txt)
- [BY-NC-ND 4.0 (plaintext)](https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode.txt)
- [CC0 1.0 legalcode.txt](https://creativecommons.org/publicdomain/zero/1.0/legalcode.txt)

取得したテキストファイルを `LICENSE` にリネームしてリポジトリ直下に置き commit & push すれば OK。これでリポジトリの Web 表示も

![license.png](https://qiita-image-store.s3.amazonaws.com/0/68318/b3c24838-d135-74c5-5e42-5d18b1674ce0.png)

のような感じになる。

なお， [CC Licenses] はプログラムコード用には調整されていないので，プログラムコードの利用を許諾するのであればオープンソース・ライセンスを適用することを強くお勧めする。

- [Githubによる、オープンソースライセンスの選び方 | オープンソース・ライセンスの談話室](http://www.catch.jp/oss-license/2013/09/10/github/)
- [たくさんあるオープンソースライセンスのそれぞれの特徴のまとめ | コリス](https://coliss.com/articles/build-websites/operation/work/choose-a-license-by-github.html)

## ブックマーク

- [Plaintext versions of Creative Commons licenses and CC0 - Creative Commons](https://creativecommons.org/2011/04/15/plaintext-versions-of-creative-commons-licenses-and-cc0/)
- [Plaintext versions of Creative Commons 4.0 licenses - Creative Commons](https://creativecommons.org/2014/01/07/plaintext-versions-of-creative-commons-4-0-licenses/)
- [改訂3版： CC Licenses について | text.Baldanders.info](http://text.baldanders.info/cc-licenses/)

[CC Licenses]: https://creativecommons.org/licenses/
