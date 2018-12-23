+++
title = "2018-06-03 のリリース情報"
date =  "2018-06-03T17:40:07+09:00"
description = "「「サクラエディタ」プロジェクトが“SourceForge.net”から“GitHub”への移行を検討」他"
image = "/images/attention/tools.png"

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近リリースされたツールやサービスについて挙げておく。
私個人が気になっているものなので全く網羅的ではないが悪しからずご了承の程を。

- [「サクラエディタ」プロジェクトが“SourceForge.net”から“GitHub”への移行を検討 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1122960.html)
- [「VMware Workstation 14」がWindows 10 April 2018 Update/Ubuntu 18.04をサポート - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1123186.html)
- [「Wireshark」v2.6.1 / 2.4.7 / 2.2.15 が公開、脆弱性を修正したメンテナンスリリース - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1123429.html)
- [「Google Chrome 67」が正式版に ～“Spectre”脆弱性の緩和策“サイト分離”をテスト - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1124593.html)
- [Apple security updates - Apple Support](https://support.apple.com/en-us/HT201222)
    - [About the security content of tvOS 11.4 - Apple Support](https://support.apple.com/en-us/HT208850)
    - [About the security content of iCloud for Windows 7.5 - Apple Support](https://support.apple.com/en-us/HT208853)
    - [About the security content of Safari 11.1.1 - Apple Support](https://support.apple.com/en-us/HT208854)
    - [About the security content of macOS High Sierra 10.13.5, Security Update 2018-003 Sierra, Security Update 2018-003 El Capitan - Apple サポート](https://support.apple.com/ja-jp/HT208849)
    - [About the security content of iOS 11.4 - Apple Support](https://support.apple.com/en-us/HT208848)
    - [About the security content of watchOS 4.3.1 - Apple Support](https://support.apple.com/en-us/HT208851)
    - [About the security content of iTunes 12.7.5 for Windows - Apple サポート](https://support.apple.com/ja-jp/HT208852)
    - [Apple、「iOS 11.4」を公開 ～“AirPlay 2”やメッセージの“iCloud”保存をサポート - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1124592.html)
- [MariaDB Server 10.3/MariaDB TX 3.0リリース。Oracle Database互換機能を搭載し、同じデータ型やPL/SQLのストアドプロシジャをサポート － Publickey](https://www.publickey1.jp/blog/18/mariadb_server_103mariadb_tx_30oracle_databaseplsql.html)
- [パスワード管理ツール「1Password 7 for Windows」が正式版に ～“Windows Hello”に対応 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1124908.html)

## 最近の脆弱性情報

最近2週間に公開・更新された脆弱性情報のうち深刻度が高いものを列挙してみる。

| ID  | タイトル | 深刻度 | 　　発見日　　 | 　最終更新日　 |
| --- | -------- | ------ | ------ | ---------- |
| [JVNDB-2018-003716](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003716.html) | Apple macOS のメモコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003715](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003715.html) | Apple iOS および macOS の iCloud Drive コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003714](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003714.html) | Apple macOS の kext ツールコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003712](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003712.html) | Apple macOS の IOFireWireFamily コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003711](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003711.html) | Apple macOS の Intel Graphics Driver コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003710](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003710.html) | 複数の Apple 製品などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003709](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003709.html) | 複数の Apple 製品などで使用される WebKit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003708](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003708.html) | 複数の Apple 製品などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003707](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003707.html) | 複数の Apple 製品などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003706](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003706.html) | 複数の Apple 製品などで使用される WebKit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003705](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003705.html) | 複数の Apple 製品などで使用される WebKit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003704](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003704.html) | 複数の Apple 製品などで使用される WebKit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003703](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003703.html) | 複数の Apple 製品などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003702](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003702.html) | 複数の Apple 製品などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003701](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003701.html) | 複数の Apple 製品などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003700](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003700.html) | 複数の Apple 製品のシステム環境設定コンポーネントにおけるアクセス制限を回避される脆弱性 | 緊急 (9.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003699](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003699.html) | 複数の Apple 製品などで使用される WebKit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003697](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003697.html) | 複数の Apple 製品のグラフィックドライバコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年1月23日 | 2018年6月1日 |
| [JVNDB-2018-003695](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003695.html) | 複数の Apple 製品などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2017-013151](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013151.html) | Apple macOS の CoreTypes コンポーネントにおけるディスクイメージのマウントを誘発される脆弱性 | 重要 (7.4) | 2017年9月25日 | 2018年6月1日 |
| [JVNDB-2017-013150](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013150.html) | Apple iOS および macOS のセキュリティコンポーネントにおける XPC メッセージの送信の制限を回避される脆弱性 | 重要 (7.0) | 2017年5月15日 | 2018年6月1日 |
| [JVNDB-2017-013148](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013148.html) | Apple iOS および macOS の SQLite コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2017年5月15日 | 2018年6月1日 |
| [JVNDB-2017-013147](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013147.html) | Apple iOS および macOS の SQLite コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2017年5月15日 | 2018年6月1日 |
| [JVNDB-2017-013146](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013146.html) | Apple iOS および macOS の SQLite コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2017年5月15日 | 2018年6月1日 |
| [JVNDB-2018-003691](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003691.html) | 複数の Apple 製品のファイルシステムイベントコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003690](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003690.html) | 複数の Apple 製品の NSURLSession コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003689](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003689.html) | 複数の Apple 製品などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003688](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003688.html) | 複数の Apple 製品などで使用される WebKit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003687](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003687.html) | 複数の Apple 製品などで使用される WebKit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003686](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003686.html) | 複数の Apple 製品などで使用される WebKit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003685](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003685.html) | 複数の Apple 製品のクイックルックコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003684](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003684.html) | Apple iOS および macOS の PluginKit コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003683](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003683.html) | 複数の Apple 製品の CoreFoundation コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003682](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003682.html) | Apple iOS および macOS のストレージコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003681](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003681.html) | 複数の Apple 製品のカーネルコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003680](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003680.html) | 複数の Apple 製品のカーネルコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003679](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003679.html) | 複数の Apple 製品の CoreText コンポーネントにおけるサービス運用妨害 (DoS) の脆弱性 | 重要 (7.5) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003678](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003678.html) | Apple iOS の Web App コンポーネントにおける Cookie の永続性に関連する制限を回避される脆弱性 | 緊急 (9.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003676](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003676.html) | 複数の Apple 製品の LinkPresentation コンポーネントにおけるサービス運用妨害 (DoS) の脆弱性 | 重要 (7.5) | 2018年1月23日 | 2018年6月1日 |
| [JVNDB-2018-003675](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003675.html) | 複数の Apple 製品のオーディオコンポーネントにおける任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年1月23日 | 2018年6月1日 |
| [JVNDB-2018-003673](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003673.html) | Apple macOS の Sandbox コンポーネントにおけるサンドボックス保護メカニズムを回避される脆弱性 | 緊急 (10.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2017-013145](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013145.html) | Apple Xcode の ld64 コンポーネントにおけるバッファオーバーフローの脆弱性 | 重要 (7.8) | 2017年12月4日 | 2018年6月1日 |
| [JVNDB-2017-013143](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013143.html) | Apple Safari などで使用される WebKit Web インスペクタコンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2017年12月6日 | 2018年6月1日 |
| [JVNDB-2017-013140](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013140.html) | 複数の Apple 製品の Wi-Fi コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2017年7月19日 | 2018年6月1日 |
| [JVNDB-2017-013132](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013132.html) | Apple macOS の Font Importer コンポーネントにおけるサービス運用妨害 (DoS) の脆弱性 | 重要 (7.1) | 2017年7月19日 | 2018年6月1日 |
| [JVNDB-2017-013130](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013130.html) | Apple macOS の Installer コンポーネントにおける FileVault のロック解除キーにアクセスされる脆弱性 | 重要 (7.5) | 2017年9月25日 | 2018年6月1日 |
| [JVNDB-2017-013129](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013129.html) | Apple Mac OS X の kext ツールコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2017年9月25日 | 2018年6月1日 |
| [JVNDB-2018-003671](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003671.html) | Apple macOS の IOHIDFamily コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年1月23日 | 2018年6月1日 |
| [JVNDB-2018-003670](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003670.html) | Apple macOS のカーネルコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年1月23日 | 2018年6月1日 |
| [JVNDB-2018-003668](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003668.html) | Apple macOS の Touch Bar Support コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年1月23日 | 2018年6月1日 |
| [JVNDB-2017-013126](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013126.html) | Apple macOS の セキュリティコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2017年10月31日 | 2018年6月1日 |
| [JVNDB-2017-013125](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013125.html) | Apple Safari などで使用される Webkit コンポーネントにおける任意のコードを実行される脆弱性 | 重要 (8.8) | 2017年3月27日 | 2018年6月1日 |
| [JVNDB-2017-013124](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013124.html) | Apple macOS の AppleGraphicsControl コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2017年7月19日 | 2018年6月1日 |
| [JVNDB-2018-003667](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003667.html) | 複数の Apple 製品の CoreFoundation コンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.0) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2018-003666](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003666.html) | 複数の Apple 製品の Security コンポーネントにおけるバッファオーバーフローの脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年6月1日 |
| [JVNDB-2017-013123](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013123.html) | 複数の Apple 製品のカーネルコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2017年9月19日 | 2018年6月1日 |
| [JVNDB-2018-003664](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003664.html) | 複数の Qualcomm 製品上で稼動する Android におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年6月1日 |
| [JVNDB-2018-003662](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003662.html) | Apple macOS の LaunchServices コンポーネントにおけるコード署名の保護メカニズムを回避される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003659](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003659.html) | Apple macOS の Admin Framework コンポーネントにおけるパスワードを取得される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003657](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003657.html) | Apple Xcode の LLVM コンポーネントにおける脆弱性 | 緊急 (9.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003656](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003656.html) | Apple macOS のカーネルコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003655](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003655.html) | Apple iOS の SafariViewController コンポーネントにおけるユーザインターフェースを偽装される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003654](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003654.html) | Apple iOS の Telephony コンポーネントにおけるバッファオーバーフローの脆弱性 | 緊急 (9.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003653](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003653.html) | Apple iOS の Telephony コンポーネントにおけるサービス運用妨害 (DoS) の脆弱性 | 重要 (7.5) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003652](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003652.html) | Apple iOS および Safari の Safari Login AutoFill コンポーネントにおける自動補完されたデータを読まれる脆弱性 | 重要 (7.5) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003651](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003651.html) | Apple macOS のカーネルコンポーネントにおける特権付きコンテキスト内で任意のコードを実行される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003650](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003650.html) | Apple iOS の Safari コンポーネントにおけるユーザインターフェースを偽装される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003649](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003649.html) | Apple iOS および macOS の WindowServer コンポーネントにおけるセキュリティ入力モードの保護メカニズムを回避される脆弱性 | 重要 (7.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003646](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003646.html) | Apple macOS のディスク管理コンポーネントにおける APFS ボリュームパスワードの切り捨てを誘発される脆弱性 | 緊急 (9.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003644](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003644.html) | Apple Mac OS X のターミナルコンポーネントの Bracketed Paste Mode における貼り付けられたコンテンツ内に任意のコマンドを挿入される脆弱性 | 重要 (8.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2018-003643](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003643.html) | Apple Mac OS X の APFS コンポーネントにおける APFS ボリュームパスワードの切り捨てを誘発される脆弱性 | 緊急 (9.8) | 2018年3月29日 | 2018年5月31日 |
| [JVNDB-2015-008177](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008177.html) | Docker Notary における危険なタイプのファイルの無制限アップロードに関する脆弱性 | 緊急 (9.8) | 2015年7月31日 | 2018年5月30日 |
| [JVNDB-2015-008176](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008176.html) | Docker Notary における暗号に関する脆弱性 | 重要 (7.5) | 2015年7月31日 | 2018年5月30日 |
| [JVNDB-2018-003642](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003642.html) | 複数の Qualcomm 製品上で稼動する Android における整数オーバーフローの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003641](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003641.html) | 複数の Qualcomm 製品上で稼動する Android における NULL ポインタデリファレンスに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003640](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003640.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003639](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003639.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003638](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003638.html) | 複数の Qualcomm 製品上で稼動する Android における NULL ポインタデリファレンスに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003637](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003637.html) | 複数の Qualcomm 製品上で稼動する Android における整数オーバーフローの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003636](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003636.html) | 複数の Qualcomm 製品上で稼動する Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003635](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003635.html) | 複数の Qualcomm 製品上で稼動する Android におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003634](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003634.html) | 複数の Qualcomm 製品上で稼動する Android における鍵管理のエラーに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003633](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003633.html) | 複数の Qualcomm 製品上で稼動する Android における入力確認に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003632](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003632.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003631](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003631.html) | 複数の Qualcomm 製品上で稼動する Android における整数オーバーフローの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003630](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003630.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003629](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003629.html) | 複数の Qualcomm 製品上で稼動する Android における境界外書き込みに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003628](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003628.html) | 複数の Qualcomm 製品上で稼動する Android における入力確認に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003627](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003627.html) | 複数の Qualcomm 製品上で稼動する Android における入力確認に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003626](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003626.html) | 複数の Qualcomm 製品上で稼動する Android における配列インデックスの検証に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003625](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003625.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003624](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003624.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003623](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003623.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003622](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003622.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003621](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003621.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003620](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003620.html) | 複数の Qualcomm 製品上で稼動する Android における NULL ポインタデリファレンスに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003619](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003619.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003618](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003618.html) | 複数の Qualcomm 製品上で稼動する Android における脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003617](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003617.html) | 複数の Qualcomm 製品上で稼動する Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003616](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003616.html) | 複数の Qualcomm 製品上で稼動する Android におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003615](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003615.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003614](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003614.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003613](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003613.html) | 複数の Qualcomm 製品上で稼動する Android における NULL ポインタデリファレンスに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003612](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003612.html) | 複数の Qualcomm 製品上で稼動する Android における競合状態に関する脆弱性 | 重要 (8.1) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003611](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003611.html) | 複数の Qualcomm 製品上で稼動する Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003610](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003610.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003609](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003609.html) | 複数の Qualcomm 製品上で稼動する Android におけるリソース管理に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003608](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003608.html) | 複数の Qualcomm 製品上で稼動する Android における整数オーバーフローの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003607](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003607.html) | 複数の Qualcomm 製品上で稼動する Android におけるアクセス制御に関する脆弱性 | 重要 (8.1) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003606](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003606.html) | 複数の Qualcomm 製品上で稼動する Android におけるアクセス制御に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月30日 |
| [JVNDB-2018-003605](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003605.html) | 複数の Qualcomm 製品上で稼動する Android における競合状態に関する脆弱性 | 重要 (8.1) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003604](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003604.html) | 複数の Qualcomm 製品上で稼動する Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003603](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003603.html) | 複数の Qualcomm 製品上で稼動する Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003602](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003602.html) | 複数の Qualcomm 製品上で稼動する Android におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003601](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003601.html) | 複数の Qualcomm 製品上で稼動する Android におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003600](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003600.html) | 複数の Qualcomm 製品上で稼動する Android における競合状態に関する脆弱性 | 重要 (8.1) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003599](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003599.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 重要 (8.1) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003598](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003598.html) | 複数の Qualcomm 製品上で稼動する Android における入力確認に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003597](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003597.html) | 複数の Qualcomm 製品上で稼動する Android における情報漏えいに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003596](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003596.html) | 複数の Qualcomm 製品上で稼動する Android におけるデータ処理に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003595](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003595.html) | 複数の Qualcomm 製品上で稼動する Android における認証に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003594](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003594.html) | 複数の Qualcomm 製品上で稼動する Android における競合状態に関する脆弱性 | 重要 (8.1) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003593](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003593.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003592](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003592.html) | 複数の Qualcomm 製品上で稼動する Android におけるインジェクションに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003591](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003591.html) | Qualcomm Snapdragon Mobile 上で稼動する Android における範囲エラーに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003590](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003590.html) | 複数の Qualcomm 製品上で稼動する Android における整数オーバーフローの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003589](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003589.html) | 複数の Qualcomm 製品上で稼動する Android における数値処理に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003588](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003588.html) | Qualcomm Snapdragon Mobile における NULL ポインタデリファレンスに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2015-008175](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008175.html) | Android用の MyScript SDK における信頼性のないデータのデシリアライゼーションに関する脆弱性 | 緊急 (9.8) | 2015年6月16日 | 2018年5月29日 |
| [JVNDB-2018-003586](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003586.html) | 複数の Qualcomm 製品上で稼動する Android におけるエラー処理に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003585](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003585.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003584](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003584.html) | 複数の Qualcomm 製品上で稼動する Android における鍵管理のエラーに関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003583](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003583.html) | 複数の Qualcomm 製品上で稼動する Android における入力確認に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003582](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003582.html) | 複数の Qualcomm 製品上で稼動する Android における環境設定に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-003581](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003581.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月29日 |
| [JVNDB-2018-000057](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000057.html) | 「フレッツ・ウイルスクリア 申込・設定ツール」および「フレッツ・ウイルスクリアv6 申込・設定ツール」のインストーラにおける実行ファイル呼び出しに関する脆弱性 | 重要 (7.8) | 2018年5月29日 | 2018年5月29日 |
| [JVNDB-2017-013121](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013121.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2017年9月28日 | 2018年5月28日 |
| [JVNDB-2017-013120](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013120.html) | 複数の Qualcomm 製品上で稼動する Android における認証に関する脆弱性 | 緊急 (9.8) | 2017年9月28日 | 2018年5月28日 |
| [JVNDB-2016-009003](https://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-009003.html) | MySQL for PCF における証明書・パスワードの管理に関する脆弱性 | 緊急 (10.0) | 2016年12月28日 | 2018年5月28日 |
| [JVNDB-2018-003580](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003580.html) | 複数の Unisys 製品における SQL インジェクションの脆弱性 | 重要 (8.1) | 2018年3月26日 | 2018年5月28日 |
| [JVNDB-2018-003577](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003577.html) | Cloud Foundry Garden-runC における情報漏えいに関する脆弱性 | 重要 (8.8) | 2018年3月28日 | 2018年5月28日 |
| [JVNDB-2018-003576](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003576.html) | TNLSoftSolutions Sentry Vision デバイスにおける証明書・パスワードの管理に関する脆弱性 | 緊急 (9.8) | 2018年3月29日 | 2018年5月28日 |
| [JVNDB-2018-003575](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003575.html) | screen-resolution-extra における競合状態に関する脆弱性 | 重要 (7.0) | 2018年3月26日 | 2018年5月28日 |
| [JVNDB-2018-003574](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003574.html) | TIM 1531 IRC における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年3月27日 | 2018年5月28日 |
| [JVNDB-2014-008573](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008573.html) | Drupal 用 Storage API モジュールにおける入力確認に関する脆弱性 | 緊急 (9.8) | 2014年7月30日 | 2018年5月28日 |
| [JVNDB-2017-013119](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013119.html) | Android におけるバッファエラーの脆弱性 | 重要 (7.8) | 2017年6月5日 | 2018年5月28日 |
| [JVNDB-2017-013118](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013118.html) | 複数の Qualcomm 製品上で稼動する Android における解放済みメモリの使用に関する脆弱性 | 緊急 (9.8) | 2017年9月28日 | 2018年5月28日 |
| [JVNDB-2017-013117](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013117.html) | 複数の Qualcomm 製品上で稼動する Android における入力確認に関する脆弱性 | 緊急 (9.8) | 2017年9月28日 | 2018年5月28日 |
| [JVNDB-2015-008174](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008174.html) | Android 用 GraceNote GNSDK における範囲エラーに関する脆弱性 | 緊急 (9.8) | 2015年6月14日 | 2018年5月28日 |
| [JVNDB-2015-008173](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008173.html) | Android 用 PJSIP PJSUA2 SDK における範囲エラーに関する脆弱性 | 緊急 (9.8) | 2015年6月14日 | 2018年5月28日 |
| [JVNDB-2015-008172](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008172.html) | Android 用 MetaIO SDK における範囲エラーに関する脆弱性 | 緊急 (9.8) | 2015年6月11日 | 2018年5月28日 |
| [JVNDB-2018-003572](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003572.html) | zsh におけるバッファエラーの脆弱性 | 重要 (7.8) | 2018年3月24日 | 2018年5月28日 |
| [JVNDB-2018-003570](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003570.html) | Cisco IOS XE ソフトウェアにおける解放済みメモリの使用に関する脆弱性 | 重要 (7.5) | 2018年3月28日 | 2018年5月28日 |
| [JVNDB-2018-003569](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003569.html) | 複数の Cisco IOS 製品におけるバッファエラーの脆弱性 | 重要 (8.8) | 2018年3月28日 | 2018年5月28日 |
| [JVNDB-2018-003568](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003568.html) | Cisco IOS XE ソフトウェアにおけるリソース管理に関する脆弱性 | 重要 (8.6) | 2018年3月28日 | 2018年5月28日 |
| [JVNDB-2018-003566](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003566.html) | Cisco IOS XE ソフトウェアにおけるデータ処理に関する脆弱性 | 重要 (8.6) | 2018年3月28日 | 2018年5月28日 |
| [JVNDB-2018-003565](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003565.html) | Cisco Catalyst 4500 シリーズおよび 4500-X シリーズスイッチにおけるエラー処理に関する脆弱性 | 重要 (8.6) | 2018年3月28日 | 2018年5月28日 |
| [JVNDB-2017-012698](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-012698.html) | HPE Network Node Manager i ソフトウェアにおけるセキュリティ機能に関する脆弱性 | 緊急 (9.8) | 2017年6月29日 | 2018年5月28日 |
| [JVNDB-2018-003564](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003564.html) | Dell EMC ScaleIO におけるバッファエラーの脆弱性 | 重要 (7.5) | 2018年3月26日 | 2018年5月28日 |
| [JVNDB-2018-003562](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003562.html) | Creditwest Bank CMS Project におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月24日 | 2018年5月28日 |
| [JVNDB-2017-013114](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013114.html) | NVIDIA Tegra カーネルにおける認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2017年2月23日 | 2018年5月28日 |
| [JVNDB-2018-003561](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003561.html) | Jenkins Mailer プラグインにおけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.0) | 2018年3月26日 | 2018年5月28日 |
| [JVNDB-2018-003560](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003560.html) | Swisscom MySwisscomAssistant におけるデータ処理に関する脆弱性 | 重要 (7.8) | 2018年3月22日 | 2018年5月28日 |
| [JVNDB-2018-003559](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003559.html) | Swisscom TVMediaHelper におけるデータ処理に関する脆弱性 | 重要 (7.8) | 2018年3月22日 | 2018年5月28日 |
| [JVNDB-2018-003557](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003557.html) | Dell EMC ScaleIO におけるコマンドインジェクションの脆弱性 | 重要 (7.5) | 2018年3月26日 | 2018年5月28日 |
| [JVNDB-2018-003556](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003556.html) | Dell EMC ScaleIO における認証に関する脆弱性 | 緊急 (9.8) | 2018年3月26日 | 2018年5月28日 |
| [JVNDB-2018-003555](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003555.html) | NordVPN におけるセキュリティ機能に関する脆弱性 | 重要 (8.8) | 2018年2月28日 | 2018年5月28日 |
| [JVNDB-2018-003552](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003552.html) | X-Pack におけるパストラバーサルの脆弱性 | 緊急 (9.8) | 2018年3月21日 | 2018年5月28日 |
| [JVNDB-2018-003551](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003551.html) | D-Link DIR-601 における証明書・パスワードの管理に関する脆弱性 | 重要 (8.0) | 2018年3月28日 | 2018年5月28日 |
| [JVNDB-2018-003550](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003550.html) | Android における解放済みメモリの使用に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2018-003549](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003549.html) | Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2018-003548](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003548.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2018-003547](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003547.html) | Android における解放済みメモリの使用に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2018-003546](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003546.html) | Android におけるバッファエラーの脆弱性 | 重要 (7.8) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2018-003545](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003545.html) | Android における NULL ポインタデリファレンスに関する脆弱性 | 重要 (7.8) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2017-013112](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013112.html) | Android における入力確認に関する脆弱性 | 緊急 (9.8) | 2017年12月19日 | 2018年5月28日 |
| [JVNDB-2017-013111](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013111.html) | Android における整数オーバーフローの脆弱性 | 緊急 (9.8) | 2017年10月3日 | 2018年5月28日 |
| [JVNDB-2017-013110](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013110.html) | Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2017年8月11日 | 2018年5月28日 |
| [JVNDB-2017-013109](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013109.html) | Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2017年10月18日 | 2018年5月28日 |
| [JVNDB-2017-013108](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013108.html) | Android における解放済みメモリの使用に関する脆弱性 | 緊急 (9.8) | 2017年10月13日 | 2018年5月28日 |
| [JVNDB-2017-013107](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013107.html) | Android における解放済みメモリの使用に関する脆弱性 | 緊急 (9.8) | 2017年8月31日 | 2018年5月28日 |
| [JVNDB-2017-013106](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013106.html) | Android における境界外書き込みに関する脆弱性 | 緊急 (9.8) | 2017年6月21日 | 2018年5月28日 |
| [JVNDB-2017-013105](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013105.html) | Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2017年7月7日 | 2018年5月28日 |
| [JVNDB-2018-003544](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003544.html) | 複数の Qualcomm 製品上で稼動する Android における入力確認に関する脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2018-003543](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003543.html) | 複数の Qualcomm 製品上で稼動する Android におけるリソース管理に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2018-003542](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003542.html) | 複数の Qualcomm 製品上で稼動する Android におけるバッファエラーの脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月28日 |
| [JVNDB-2015-008171](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008171.html) | Android 用 Jumio SDK における範囲エラーに関する脆弱性 | 緊急 (9.8) | 2015年6月11日 | 2018年5月28日 |
| [JVNDB-2014-008571](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008571.html) | Zikula Application Framework におけるコードインジェクションの脆弱性 | 緊急 (9.8) | 2014年2月17日 | 2018年5月28日 |
| [JVNDB-2018-003541](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003541.html) | Drupal における入力確認に関する脆弱性 | 緊急 (9.8) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003540](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003540.html) | Atlassian Fisheye および Crucible における入力確認に関する脆弱性 | 重要 (7.2) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003539](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003539.html) | Atlassian Bamboo における入力確認に関する脆弱性 | 重要 (8.8) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003538](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003538.html) | 複数の Qualcomm Snapdragon 製品におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月25日 |
| [JVNDB-2018-003537](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003537.html) | 複数の Qualcomm Snapdragon 製品におけるデータ処理に関する脆弱性 | 重要 (7.5) | 2018年4月2日 | 2018年5月25日 |
| [JVNDB-2018-003535](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003535.html) | WireMock における XML 外部エンティティの脆弱性 | 緊急 (9.1) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2015-008170](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008170.html) | ESRI ArcGIS Runtime SDK における範囲エラーに関する脆弱性 | 緊急 (9.8) | 2015年8月10日 | 2018年5月25日 |
| [JVNDB-2015-008169](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008169.html) | IBM QRadar SIEM におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2015年2月19日 | 2018年5月25日 |
| [JVNDB-2015-008168](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008168.html) | IBM Endpoint Manager for Remote Control における脆弱性 | 重要 (8.8) | 2015年6月24日 | 2018年5月25日 |
| [JVNDB-2016-009002](https://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-009002.html) | cf-release における情報漏えいに関する脆弱性 | 緊急 (9.6) | 2016年11月2日 | 2018年5月25日 |
| [JVNDB-2017-013104](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013104.html) | Android のカメラドライバにおける認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2017年10月24日 | 2018年5月25日 |
| [JVNDB-2017-013103](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013103.html) | Android のカメラドライバにおける NULL ポインタデリファレンスに関する脆弱性 | 重要 (7.8) | 2017年9月16日 | 2018年5月25日 |
| [JVNDB-2017-013102](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013102.html) | Android におけるバッファエラーの脆弱性 | 重要 (7.5) | 2017年5月23日 | 2018年5月25日 |
| [JVNDB-2015-008165](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008165.html) | IBM Rational ClearCase における暗号に関する脆弱性 | 重要 (7.4) | 2015年6月24日 | 2018年5月25日 |
| [JVNDB-2014-008569](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008569.html) | ownCloud Server におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2014年7月3日 | 2018年5月25日 |
| [JVNDB-2018-003529](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003529.html) | Zoho ManageEngine Desktop Central における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.2) | 2018年4月17日 | 2018年5月25日 |
| [JVNDB-2018-003528](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003528.html) | Zoho ManageEngine Desktop Central における入力確認に関する脆弱性 | 緊急 (9.8) | 2018年3月27日 | 2018年5月25日 |
| [JVNDB-2018-003527](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003527.html) | Zoho ManageEngine Desktop Central におけるアクセス制御に関する脆弱性 | 重要 (7.2) | 2018年4月24日 | 2018年5月25日 |
| [JVNDB-2018-003526](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003526.html) | Zoho ManageEngine Desktop Central におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年4月24日 | 2018年5月25日 |
| [JVNDB-2018-003525](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003525.html) | Zoho ManageEngine Desktop Central における重要な機能に対する認証の欠如に関する脆弱性 | 緊急 (9.8) | 2018年3月27日 | 2018年5月25日 |
| [JVNDB-2018-003524](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003524.html) | Zoho ManageEngine Desktop Central におけるパストラバーサルの脆弱性 | 緊急 (9.1) | 2018年3月27日 | 2018年5月25日 |
| [JVNDB-2018-003521](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003521.html) | Cisco IOS ソフトウェアおよび Cisco IOS XE ソフトウェアにおける入力確認に関する脆弱性 | 重要 (8.6) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003520](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003520.html) | Cisco IOS XE ソフトウェアにおけるリソース管理に関する脆弱性 | 重要 (7.4) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2014-008568](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008568.html) | Android の SQLi API における SQL インジェクションの脆弱性 | 緊急 (9.8) | 2014年7月26日 | 2018年5月25日 |
| [JVNDB-2017-013100](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013100.html) | Firebird SQL Server における SQL インジェクションの脆弱性 | 重要 (8.8) | 2017年11月21日 | 2018年5月25日 |
| [JVNDB-2018-003516](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003516.html) | Square 9 GlobalForms における SQL インジェクションの脆弱性 | 重要 (7.5) | 2018年3月27日 | 2018年5月25日 |
| [JVNDB-2018-003515](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003515.html) | Philips Alice 6 System における暗号に関する脆弱性 | 緊急 (9.8) | 2018年3月27日 | 2018年5月25日 |
| [JVNDB-2018-003514](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003514.html) | Philips Alice 6 System における認証に関する脆弱性 | 緊急 (9.8) | 2018年3月27日 | 2018年5月25日 |
| [JVNDB-2017-013099](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013099.html) | Wanscam HW0021 ネットワークカメラにおける証明書・パスワードの管理に関する脆弱性 | 緊急 (9.8) | 2017年11月10日 | 2018年5月25日 |
| [JVNDB-2018-003513](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003513.html) | Cisco IOS XE ソフトウェアにおける OS コマンドインジェクションの脆弱性 | 重要 (7.8) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003511](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003511.html) | Cisco IOS ソフトウェアおよび Cisco IOS XE ソフトウェアにおける入力確認に関する脆弱性 | 重要 (8.6) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003510](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003510.html) | Cisco IOS XE ソフトウェアにおけるデータ処理に関する脆弱性 | 重要 (7.5) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003509](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003509.html) | Cisco IOS XE ソフトウェアにおける OS コマンドインジェクションの脆弱性 | 重要 (7.8) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003508](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003508.html) | 複数の Cisco IOS 製品における書式文字列に関する脆弱性 | 重要 (8.0) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003507](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003507.html) | Cisco IOS ソフトウェアおよび Cisco IOS XE ソフトウェアにおける入力確認に関する脆弱性 | 重要 (8.6) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003506](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003506.html) | Cisco IOS ソフトウェアおよび Cisco IOS XE ソフトウェアにおける入力確認に関する脆弱性 | 重要 (8.6) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003505](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003505.html) | Cisco IOS XE ソフトウェアにおける認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2018年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003503](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003503.html) | DedeCMS におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月30日 | 2018年5月25日 |
| [JVNDB-2017-013098](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013098.html) | Android における脆弱性 | 緊急 (9.8) | 2017年9月28日 | 2018年5月25日 |
| [JVNDB-2017-013097](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013097.html) | Android におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年7月7日 | 2018年5月25日 |
| [JVNDB-2018-003502](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003502.html) | 複数の Cloud Foundry Foundation 製品におけるアクセス制御に関する脆弱性 | 重要 (8.8) | 2018年3月5日 | 2018年5月25日 |
| [JVNDB-2018-003501](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003501.html) | Yxcms における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.5) | 2018年3月18日 | 2018年5月25日 |
| [JVNDB-2018-003500](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003500.html) | libevt におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年3月17日 | 2018年5月25日 |
| [JVNDB-2017-013096](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013096.html) | authentikat-jwt における時間とステータスに関する脆弱性 | 緊急 (9.8) | 2017年12月20日 | 2018年5月25日 |
| [JVNDB-2017-013095](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013095.html) | Android における NULL ポインタデリファレンスに関する脆弱性 | 重要 (7.8) | 2017年5月16日 | 2018年5月25日 |
| [JVNDB-2017-013093](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013093.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2017年2月12日 | 2018年5月25日 |
| [JVNDB-2018-003498](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003498.html) | LDAP Account Manager におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年4月3日 | 2018年5月25日 |
| [JVNDB-2018-003495](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003495.html) | Ruby 用 Sanitize gem における入力確認に関する脆弱性 | 重要 (7.5) | 2018年3月20日 | 2018年5月25日 |
| [JVNDB-2017-013091](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013091.html) | TPshop におけるサーバサイドのリクエストフォージェリの脆弱性 | 緊急 (9.8) | 2017年11月6日 | 2018年5月25日 |
| [JVNDB-2017-013090](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013090.html) | Android における競合状態に関する脆弱性 | 重要 (7.8) | 2017年10月20日 | 2018年5月25日 |
| [JVNDB-2017-013089](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013089.html) | Android におけるバッファエラーの脆弱性 | 重要 (7.8) | 2017年10月11日 | 2018年5月25日 |
| [JVNDB-2017-013088](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013088.html) | Android における入力確認に関する脆弱性 | 重要 (7.8) | 2017年9月19日 | 2018年5月25日 |
| [JVNDB-2014-008567](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008567.html) | Seafile Server および Server Professional Edition における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2014年8月7日 | 2018年5月25日 |
| [JVNDB-2014-008566](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008566.html) | OpenScape Deployment Service における SQL インジェクションの脆弱性 | 緊急 (9.8) | 2014年3月28日 | 2018年5月25日 |
| [JVNDB-2018-003492](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003492.html) | cf-deployment および routing-release における入力確認に関する脆弱性 | 重要 (8.1) | 2018年2月13日 | 2018年5月25日 |
| [JVNDB-2018-003491](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003491.html) | Windows Stemcells における認可・権限・アクセス制御に関する脆弱性 | 重要 (8.5) | 2018年2月22日 | 2018年5月25日 |
| [JVNDB-2018-003490](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003490.html) | 複数の F5 BIG-IP 製品における入力確認に関する脆弱性 | 重要 (7.5) | 2018年3月22日 | 2018年5月24日 |
| [JVNDB-2018-003488](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003488.html) | 複数の F5 BIG-IP 製品におけるデータ処理に関する脆弱性 | 重要 (8.1) | 2018年3月22日 | 2018年5月24日 |
| [JVNDB-2018-003487](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003487.html) | F5 BIG-IP Policy Enforcement Manager における入力確認に関する脆弱性 | 重要 (7.5) | 2018年3月22日 | 2018年5月24日 |
| [JVNDB-2018-003486](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003486.html) | 複数の F5 BIG-IP 製品における証明書検証に関する脆弱性 | 重要 (7.5) | 2018年3月22日 | 2018年5月24日 |
| [JVNDB-2018-003484](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003484.html) | Atlassian Bitbucket Server におけるリンク解釈に関する脆弱性 | 緊急 (9.9) | 2018年3月21日 | 2018年5月24日 |
| [JVNDB-2014-008564](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008564.html) | TrueCrypt における整数オーバーフローの脆弱性 | 重要 (7.1) | 2014年3月4日 | 2018年5月24日 |
| [JVNDB-2018-003481](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003481.html) | elfutils におけるバッファエラーの脆弱性 | 重要 (7.8) | 2018年3月16日 | 2018年5月24日 |
| [JVNDB-2018-003480](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003480.html) | Jupyter Notebook における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2018年3月15日 | 2018年5月24日 |
| [JVNDB-2017-013083](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013083.html) | Android システムにおける入力確認に関する脆弱性 | 重要 (7.5) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013082](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013082.html) | Android システムにおける入力確認に関する脆弱性 | 重要 (7.5) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013081](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013081.html) | Android のメディアフレームワークにおける入力確認に関する脆弱性 | 重要 (7.5) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013080](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013080.html) | Android のメディアフレームワークにおける脆弱性 | 重要 (7.5) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013079](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013079.html) | Android の Upstream カーネルにおける認可・権限・アクセス制御に関する脆弱性 | 重要 (7.3) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013078](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013078.html) | Android の Upstream カーネルの mnh ドライバにおける認可・権限・アクセス制御に関する脆弱性 | 重要 (7.3) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013075](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013075.html) | Android における境界外書き込みに関する脆弱性 | 重要 (7.8) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2018-003478](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003478.html) | NCR S2 Dispenser における境界外書き込みに関する脆弱性 | 重要 (7.5) | 2018年3月20日 | 2018年5月24日 |
| [JVNDB-2018-003477](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003477.html) | Yii におけるコードインジェクションの脆弱性 | 重要 (8.1) | 2018年3月20日 | 2018年5月24日 |
| [JVNDB-2018-003476](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003476.html) | Yii における SQL インジェクションの脆弱性 | 緊急 (9.8) | 2018年3月20日 | 2018年5月24日 |
| [JVNDB-2014-008562](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008562.html) | OpenCart における XML 外部エンティティの脆弱性 | 緊急 (9.8) | 2014年6月6日 | 2018年5月24日 |
| [JVNDB-2014-008561](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008561.html) | Core FTP Server におけるバッファエラーの脆弱性 | 重要 (7.8) | 2014年1月9日 | 2018年5月24日 |
| [JVNDB-2018-003475](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003475.html) | Kamailio におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年2月11日 | 2018年5月24日 |
| [JVNDB-2018-003474](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003474.html) | Linux Kernel におけるバッファエラーの脆弱性 | 重要 (7.8) | 2018年3月19日 | 2018年5月24日 |
| [JVNDB-2018-003473](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003473.html) | Gitlab Community および Enterprise Editions におけるパストラバーサルの脆弱性 | 重要 (7.8) | 2018年1月16日 | 2018年5月24日 |
| [JVNDB-2018-003471](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003471.html) | GNOME NetworkManager における情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年5月14日 | 2018年5月24日 |
| [JVNDB-2018-003470](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003470.html) | Pivotal Spring Batch Admin におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月16日 | 2018年5月24日 |
| [JVNDB-2018-003469](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003469.html) | Gitlab Community Edition におけるパストラバーサルの脆弱性 | 重要 (8.8) | 2018年1月16日 | 2018年5月24日 |
| [JVNDB-2018-003468](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003468.html) | QOS.CH SLF4J における信頼性のないデータのデシリアライゼーションに関する脆弱性 | 緊急 (9.8) | 2018年3月14日 | 2018年5月24日 |
| [JVNDB-2018-003467](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003467.html) | Western Digital WD My Cloud における認証に関する脆弱性 | 緊急 (9.8) | 2018年3月27日 | 2018年5月24日 |
| [JVNDB-2017-013072](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013072.html) | Android における境界外書き込みに関する脆弱性 | 重要 (8.8) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013071](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013071.html) | Android における境界外書き込みに関する脆弱性 | 重要 (8.8) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013070](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013070.html) | Android における境界外書き込みに関する脆弱性 | 重要 (7.8) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013069](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013069.html) | Android における入力確認に関する脆弱性 | 重要 (7.8) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2018-003466](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003466.html) | PrestaShop 用 Responsive Mega Menu Pro モジュールにおけるコードインジェクションの脆弱性 | 緊急 (9.8) | 2018年3月6日 | 2018年5月24日 |
| [JVNDB-2018-003465](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003465.html) | Studio 42 elFinder におけるパストラバーサルの脆弱性 | 重要 (7.5) | 2018年3月28日 | 2018年5月24日 |
| [JVNDB-2011-005419](https://jvndb.jvn.jp/ja/contents/2011/JVNDB-2011-005419.html) | openbuildservice におけるコードインジェクションの脆弱性 | 重要 (8.8) | 2011年9月2日 | 2018年5月24日 |
| [JVNDB-2017-013068](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013068.html) | Intel Software Guard Extensions Platform Software Component における認可・権限・アクセス制御に関する脆弱性 | 重要 (8.8) | 2017年2月1日 | 2018年5月24日 |
| [JVNDB-2018-003464](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003464.html) | Android 用 Wire アプリケーションにおけるパストラバーサルの脆弱性 | 重要 (7.5) | 2018年3月7日 | 2018年5月24日 |
| [JVNDB-2017-013067](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013067.html) | qcacld-2.0 における入力確認に関する脆弱性 | 重要 (7.8) | 2017年6月4日 | 2018年5月24日 |
| [JVNDB-2017-013065](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013065.html) | Android における境界外読み取りに関する脆弱性 | 重要 (7.5) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013064](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013064.html) | Android における境界外読み取りに関する脆弱性 | 重要 (7.5) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013063](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013063.html) | Android における境界外読み取りに関する脆弱性 | 重要 (7.5) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2017-013062](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013062.html) | Android における境界外読み取りに関する脆弱性 | 重要 (7.5) | 2017年8月23日 | 2018年5月24日 |
| [JVNDB-2018-003463](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003463.html) | SickRage における証明書・パスワードの管理に関する脆弱性 | 緊急 (9.8) | 2018年3月10日 | 2018年5月24日 |
| [JVNDB-2018-000056](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000056.html) | Susie プラグイン axpdfium における DLL 読み込みに関する脆弱性 | 重要 (7.0) | 2018年5月24日 | 2018年5月24日 |
| [JVNDB-2018-000046](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-000046.html) | Windows版 PlayMemories Home のインストーラにおける DLL 読み込みに関する脆弱性 | 重要 (7.8) | 2018年5月24日 | 2018年5月24日 |
| [JVNDB-2018-003462](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003462.html) | Softros Network Time System における入力確認に関する脆弱性 | 重要 (7.5) | 2018年3月5日 | 2018年5月24日 |
| [JVNDB-2018-003461](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003461.html) | Heimdal PRO における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2018年3月12日 | 2018年5月24日 |
| [JVNDB-2018-003459](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003459.html) | DVD X Player Standard におけるバッファエラーの脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月24日 |
| [JVNDB-2018-003456](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003456.html) | Qualcomm Snapdragon Mobile における整数オーバーフローの脆弱性 | 緊急 (9.8) | 2018年4月2日 | 2018年5月24日 |
| [JVNDB-2018-003455](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003455.html) | HashiCorp Terraform Amazon Web Services プロバイダにおける PRNG の不十分なエントロピーに関する脆弱性 | 緊急 (9.8) | 2018年3月30日 | 2018年5月24日 |
| [JVNDB-2018-003453](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003453.html) | Heimdal PRO における入力確認に関する脆弱性 | 重要 (7.0) | 2018年3月12日 | 2018年5月24日 |
| [JVNDB-2018-003452](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003452.html) | QuickAppsCMS におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月27日 | 2018年5月23日 |
| [JVNDB-2018-003451](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003451.html) | GitLab における入力確認に関する脆弱性 | 緊急 (9.8) | 2018年3月20日 | 2018年5月23日 |
| [JVNDB-2018-003449](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003449.html) | Beckhoff TwinCAT における入力確認に関する脆弱性 | 重要 (7.8) | 2018年3月13日 | 2018年5月23日 |
| [JVNDB-2018-003448](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003448.html) | I, Librarian におけるアクセス制御に関する脆弱性 | 緊急 (9.1) | 2018年4月2日 | 2018年5月23日 |
| [JVNDB-2018-003441](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003441.html) | libvirt におけるリソースの枯渇に関する脆弱性 | 重要 (7.5) | 2018年3月14日 | 2018年5月23日 |
| [JVNDB-2018-003439](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003439.html) | Twonky Server におけるパストラバーサルの脆弱性 | 重要 (7.5) | 2018年3月27日 | 2018年5月23日 |
| [JVNDB-2018-003436](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003436.html) | RSA Authentication Agent for Web におけるバッファエラーの脆弱性 | 重要 (7.5) | 2018年3月26日 | 2018年5月23日 |
| [JVNDB-2018-003431](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003431.html) | Cisco IOS XE ソフトウェアにおける OS コマンドインジェクションの脆弱性 | 重要 (7.8) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2018-003428](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003428.html) | Cisco IOS XE ソフトウェアにおける OS コマンドインジェクションの脆弱性 | 重要 (7.8) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2018-003427](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003427.html) | Cisco IOS ソフトウェア Cisco IOS XE ソフトウェアにおける入力確認に関する脆弱性 | 重要 (7.5) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2018-003426](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003426.html) | Cisco IOS ソフトウェアおよび Cisco IOS XE ソフトウェアにおける入力確認に関する脆弱性 | 重要 (7.5) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2018-003425](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003425.html) | Cisco IOS ソフトウェアにおけるリソース管理に関する脆弱性 | 重要 (7.5) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2018-003424](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003424.html) | Cisco IOS XE ソフトウェア における認可・権限・アクセス制御に関する脆弱性 | 重要 (8.8) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2018-003423](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003423.html) | Cisco IOS ソフトウェアおよび Cisco IOS XE ソフトウェア におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2018-003422](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003422.html) | Cisco IOS XE ソフトウェアにおけるハードコードされた認証情報の使用に関する脆弱性 | 緊急 (9.8) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2017-013060](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013060.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013059](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013059.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013058](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013058.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013057](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013057.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013056](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013056.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013055](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013055.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013054](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013054.html) | Android の alarm.cc における解放済みメモリの使用に関する脆弱性 | 緊急 (9.8) | 2017年8月23日 | 2018年5月23日 |
| [JVNDB-2017-013053](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013053.html) | Android の Upstream カーネルの mnh_sm ドライバにおける認可・権限・アクセス制御に関する脆弱性 | 重要 (7.3) | 2017年8月23日 | 2018年5月23日 |
| [JVNDB-2017-013052](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013052.html) | Android の Upstream カーネルの mnh_sm ドライバにおける認可・権限・アクセス制御に関する脆弱性 | 重要 (7.3) | 2017年8月23日 | 2018年5月23日 |
| [JVNDB-2018-003421](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003421.html) | Joomla! 用 Acyba AcyMailing エクステンションにおける入力確認に関する脆弱性 | 重要 (8.8) | 2018年3月26日 | 2018年5月23日 |
| [JVNDB-2018-003420](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003420.html) | Joomla! 用 Acyba AcySMS エクステンションにおける入力確認に関する脆弱性 | 重要 (8.8) | 2018年3月27日 | 2018年5月23日 |
| [JVNDB-2018-003419](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003419.html) | MiniCMS におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月20日 | 2018年5月23日 |
| [JVNDB-2017-013051](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013051.html) | Bose SoundTouch デバイスにおけるアクセス制御に関する脆弱性 | 重要 (8.8) | 2017年12月20日 | 2018年5月23日 |
| [JVNDB-2017-013047](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013047.html) | Android の Qualcomm 製品における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013046](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013046.html) | Android の Qualcomm 製品における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013045](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013045.html) | Android の Qualcomm 製品における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2018-003418](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003418.html) | Intelbras TELEFONE IP におけるパストラバーサルの脆弱性 | 緊急 (9.8) | 2018年3月16日 | 2018年5月23日 |
| [JVNDB-2018-003417](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003417.html) | dsmall における情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年3月25日 | 2018年5月23日 |
| [JVNDB-2018-003416](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003416.html) | rsyslog librelp におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年3月26日 | 2018年5月23日 |
| [JVNDB-2018-003415](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003415.html) | Electron におけるデータ処理に関する脆弱性 | 重要 (8.1) | 2018年3月21日 | 2018年5月23日 |
| [JVNDB-2018-003414](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003414.html) | PHPOK における危険なタイプのファイルの無制限アップロードに関する脆弱性 | 緊急 (9.8) | 2018年2月28日 | 2018年5月23日 |
| [JVNDB-2017-013044](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013044.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年4月3日 | 2018年5月23日 |
| [JVNDB-2017-013043](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013043.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2017年4月3日 | 2018年5月23日 |
| [JVNDB-2017-013041](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013041.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2017年4月3日 | 2018年5月23日 |
| [JVNDB-2017-013040](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013040.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年4月3日 | 2018年5月23日 |
| [JVNDB-2017-013039](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013039.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2018-003413](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003413.html) | DedeCMS におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月7日 | 2018年5月23日 |
| [JVNDB-2018-003412](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003412.html) | Dell EMC Isilon OneFS におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年2月14日 | 2018年5月23日 |
| [JVNDB-2014-008559](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008559.html) | Knot DNS における入力確認に関する脆弱性 | 重要 (7.5) | 2014年9月9日 | 2018年5月23日 |
| [JVNDB-2017-013038](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013038.html) | Bomgar Remote Support におけるパストラバーサルの脆弱性 | 緊急 (10.0) | 2017年8月11日 | 2018年5月23日 |
| [JVNDB-2017-013037](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013037.html) | Kaseya Virtual System Administrator における競合状態に関する脆弱性 | 重要 (7.4) | 2017年8月3日 | 2018年5月23日 |
| [JVNDB-2017-013034](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013034.html) | Android におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2017年8月23日 | 2018年5月23日 |
| [JVNDB-2017-013033](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013033.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.3) | 2017年8月23日 | 2018年5月23日 |
| [JVNDB-2017-013032](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013032.html) | Android における脆弱性 | 重要 (7.3) | 2017年8月23日 | 2018年5月23日 |
| [JVNDB-2017-013031](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013031.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.3) | 2017年8月23日 | 2018年5月23日 |
| [JVNDB-2018-003407](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003407.html) | Samsung モバイルデバイスのソフトウェアにおけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年3月30日 | 2018年5月23日 |
| [JVNDB-2018-003406](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003406.html) | Samsung モバイルデバイスのソフトウェアにおける入力確認に関する脆弱性 | 重要 (7.0) | 2018年3月30日 | 2018年5月23日 |
| [JVNDB-2018-003405](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003405.html) | Samsung モバイルデバイスのソフトウェアにおける入力確認に関する脆弱性 | 重要 (7.8) | 2018年3月30日 | 2018年5月23日 |
| [JVNDB-2018-003403](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003403.html) | Samsung モバイルデバイスのソフトウェアにおけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年3月30日 | 2018年5月23日 |
| [JVNDB-2017-013030](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013030.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013029](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013029.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013028](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013028.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013027](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013027.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013026](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013026.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2018-003397](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003397.html) | Cisco IOS XE ソフトウェアにおける認証に関する脆弱性 | 重要 (8.8) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2017-013025](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013025.html) | HashiCorp Vagrant VMware Fusion における証明書・パスワードの管理に関する脆弱性 | 重要 (7.0) | 2017年11月15日 | 2018年5月23日 |
| [JVNDB-2017-013024](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013024.html) | HashiCorp Vagrant VMware Fusion における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2017年11月3日 | 2018年5月23日 |
| [JVNDB-2017-013023](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013023.html) | HashiCorp Vagrant VMware Fusion における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.8) | 2017年11月17日 | 2018年5月23日 |
| [JVNDB-2018-003390](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003390.html) | DedeCMS におけるコードインジェクションの脆弱性 | 緊急 (9.8) | 2018年4月1日 | 2018年5月23日 |
| [JVNDB-2016-009001](https://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-009001.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2016年10月5日 | 2018年5月23日 |
| [JVNDB-2016-009000](https://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-009000.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2016年10月5日 | 2018年5月23日 |
| [JVNDB-2016-008999](https://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-008999.html) | Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2016年10月5日 | 2018年5月23日 |
| [JVNDB-2016-008998](https://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-008998.html) | Android における情報漏えいに関する脆弱性 | 重要 (7.5) | 2016年10月5日 | 2018年5月23日 |
| [JVNDB-2016-008997](https://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-008997.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2016年10月5日 | 2018年5月23日 |
| [JVNDB-2017-013021](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013021.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2017-013020](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013020.html) | Android における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年6月5日 | 2018年5月23日 |
| [JVNDB-2018-003385](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003385.html) | DedeCMS におけるコードインジェクションの脆弱性 | 緊急 (9.8) | 2018年4月1日 | 2018年5月23日 |
| [JVNDB-2018-003381](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003381.html) | Dell EMC NetWorker におけるバッファエラーの脆弱性 | 重要 (7.5) | 2018年3月16日 | 2018年5月23日 |
| [JVNDB-2018-003378](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003378.html) | YzmCMS におけるコマンドインジェクションの脆弱性 | 重要 (7.2) | 2018年3月17日 | 2018年5月23日 |
| [JVNDB-2018-003377](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003377.html) | Joyent SmartOS における入力確認に関する脆弱性 | 重要 (7.0) | 2018年3月8日 | 2018年5月23日 |
| [JVNDB-2015-008159](https://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-008159.html) | Garden におけるアクセス制御に関する脆弱性 | 重要 (7.5) | 2015年12月15日 | 2018年5月23日 |
| [JVNDB-2014-008558](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008558.html) | Grails Resource プラグインにおけるパストラバーサルの脆弱性 | 重要 (7.5) | 2014年5月14日 | 2018年5月23日 |
| [JVNDB-2018-003375](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003375.html) | LibreSSL における証明書検証に関する脆弱性 | 重要 (7.4) | 2018年3月23日 | 2018年5月23日 |
| [JVNDB-2018-003374](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003374.html) | Nessus における認可・権限・アクセス制御に関する脆弱性 | 重要 (7.0) | 2018年3月19日 | 2018年5月23日 |
| [JVNDB-2018-003373](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003373.html) | MikroTik RouterOS におけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年3月15日 | 2018年5月23日 |
| [JVNDB-2018-003372](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003372.html) | Cisco IOS XE ソフトウェアにおけるOS コマンドインジェクションの脆弱性 | 重要 (7.8) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2018-003371](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003371.html) | Z-BlogPHP におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月27日 | 2018年5月23日 |
| [JVNDB-2017-013018](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013018.html) | Linux kernel における整数オーバーフローの脆弱性 | 重要 (7.8) | 2017年2月23日 | 2018年5月23日 |
| [JVNDB-2017-013017](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013017.html) | Huawei Smart Phone のソフトウェアにおける整数オーバーフローの脆弱性 | 重要 (7.8) | 2017年10月14日 | 2018年5月23日 |
| [JVNDB-2017-013016](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013016.html) | Kentico における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2017年12月18日 | 2018年5月23日 |
| [JVNDB-2017-013015](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013015.html) | Cisco Spark Hybrid Calendar Service における情報漏えいに関する脆弱性 | 重要 (7.5) | 2017年10月23日 | 2018年5月23日 |
| [JVNDB-2018-003370](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003370.html) | Cloud Foundry Silk CNI プラグインにおけるアクセス制御に関する脆弱性 | 重要 (8.1) | 2018年3月26日 | 2018年5月23日 |
| [JVNDB-2018-003369](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003369.html) | Cloud Foundry Cloud Controller におけるパストラバーサルの脆弱性 | 重要 (8.1) | 2018年3月26日 | 2018年5月23日 |
| [JVNDB-2018-003368](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003368.html) | Cloud Foundry BOSH CLI におけるアクセス制御に関する脆弱性 | 重要 (8.8) | 2018年3月26日 | 2018年5月23日 |
| [JVNDB-2018-003365](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003365.html) | D-Link DIR-850L デバイスにおける認証に関する脆弱性 | 緊急 (9.8) | 2018年3月30日 | 2018年5月23日 |
| [JVNDB-2018-003364](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003364.html) | Studio 42 elFinder におけるパストラバーサルの脆弱性 | 重要 (7.5) | 2018年3月28日 | 2018年5月23日 |
| [JVNDB-2017-013014](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013014.html) | Huawei HG532 における入力確認に関する脆弱性 | 重要 (8.8) | 2017年11月30日 | 2018年5月23日 |
| [JVNDB-2017-013013](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013013.html) | GE Centricity PACS RA1000 におけるハードコードされた認証情報の使用に関する脆弱性 | 緊急 (9.8) | 2017年8月30日 | 2018年5月23日 |
| [JVNDB-2017-013012](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013012.html) | GE Xeleris におけるハードコードされた認証情報の使用に関する脆弱性 | 緊急 (9.8) | 2017年8月30日 | 2018年5月23日 |
| [JVNDB-2017-013011](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013011.html) | GE GEMNet License server におけるハードコードされた認証情報の使用に関する脆弱性 | 緊急 (9.8) | 2017年8月30日 | 2018年5月23日 |
| [JVNDB-2017-013010](https://jvndb.jvn.jp/ja/contents/2017/JVNDB-2017-013010.html) | GE Infinia/Infinia with Hawkeye 4 におけるハードコードされた認証情報の使用に関する脆弱性 | 緊急 (9.8) | 2017年8月30日 | 2018年5月23日 |
| [JVNDB-2018-003362](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003362.html) | Tenda AC15 ルータにおけるハードコードされた認証情報の使用に関する脆弱性 | 緊急 (9.8) | 2018年3月19日 | 2018年5月23日 |
| [JVNDB-2018-003359](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003359.html) | 複数の AMD 製品における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.0) | 2018年3月21日 | 2018年5月22日 |
| [JVNDB-2018-003358](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003358.html) | AMD Ryzen および Ryzen Pro プラットフォームにおける認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.0) | 2018年3月21日 | 2018年5月22日 |
| [JVNDB-2018-003357](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003357.html) | AMD Ryzen および Ryzen Pro プラットフォームにおける認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.0) | 2018年3月21日 | 2018年5月22日 |
| [JVNDB-2018-003356](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003356.html) | AMD EPYC Server におけるアクセス制御に関する脆弱性 | 緊急 (9.0) | 2018年3月21日 | 2018年5月22日 |
| [JVNDB-2018-003355](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003355.html) | AMD Ryzen および Ryzen Pro におけるアクセス制御に関する脆弱性 | 緊急 (9.0) | 2018年3月21日 | 2018年5月22日 |
| [JVNDB-2018-003354](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003354.html) | 複数の AMD 製品におけるアクセス制御に関する脆弱性 | 緊急 (9.0) | 2018年3月21日 | 2018年5月22日 |
| [JVNDB-2018-003353](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003353.html) | 複数の AMD 製品における入力確認に関する脆弱性 | 緊急 (9.0) | 2018年3月21日 | 2018年5月22日 |
| [JVNDB-2018-003352](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003352.html) | Dell EMC iDRAC7 および iDRAC8 におけるパストラバーサルの脆弱性 | 重要 (7.5) | 2018年3月20日 | 2018年5月22日 |
| [JVNDB-2018-003351](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003351.html) | Dell EMC iDRAC7 および iDRAC8 におけるインジェクションに関する脆弱性 | 緊急 (9.8) | 2018年3月20日 | 2018年5月22日 |
| [JVNDB-2018-003349](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003349.html) | WordPress 用 Site Editor プラグインにおける情報漏えいに関する脆弱性 | 重要 (7.5) | 2018年3月15日 | 2018年5月22日 |
| [JVNDB-2018-003348](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003348.html) | Ceph における NULL ポインタデリファレンスに関する脆弱性 | 重要 (7.5) | 2018年4月6日 | 2018年5月22日 |
| [JVNDB-2014-008555](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008555.html) | WordPress 用 Ajax Pagination プラグインにおけるパストラバーサルの脆弱性 | 重要 (7.5) | 2014年3月28日 | 2018年5月22日 |
| [JVNDB-2014-008553](https://jvndb.jvn.jp/ja/contents/2014/JVNDB-2014-008553.html) | WordPress 用 Subscribe To Comments Reloaded プラグインにおけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2014年3月10日 | 2018年5月22日 |
| [JVNDB-2018-003346](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003346.html) | Geutebruck G-Cam/EFD-2250 および Topline TopFD-2125 における認証に関する脆弱性 | 緊急 (9.8) | 2018年3月20日 | 2018年5月22日 |
| [JVNDB-2018-003345](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003345.html) | Geutebruck G-Cam/EFD-2250 および Topline TopFD-2125 における SQL インジェクションの脆弱性 | 緊急 (9.1) | 2018年3月20日 | 2018年5月22日 |
| [JVNDB-2018-003344](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003344.html) | Geutebruck G-Cam/EFD-2250 および Topline TopFD-2125 におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月20日 | 2018年5月22日 |
| [JVNDB-2018-003343](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003343.html) | Geutebruck G-Cam/EFD-2250 および Topline TopFD-2125 におけるアクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年3月20日 | 2018年5月22日 |
| [JVNDB-2018-003342](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003342.html) | Geutebruck G-Cam/EFD-2250 および Topline TopFD-2125 におけるサーバサイドのリクエストフォージェリの脆弱性 | 重要 (7.3) | 2018年3月20日 | 2018年5月22日 |
| [JVNDB-2018-003338](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003338.html) | PHPSHE における SQL インジェクションの脆弱性 | 緊急 (9.8) | 2018年2月2日 | 2018年5月22日 |
| [JVNDB-2018-003336](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003336.html) | Open-AudIT Professional におけるクロスサイトリクエストフォージェリの脆弱性 | 重要 (8.8) | 2018年3月27日 | 2018年5月22日 |
| [JVNDB-2018-003332](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003332.html) | Philips Intellispace Portal における入力確認に関する脆弱性 | 緊急 (9.8) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2018-003331](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003331.html) | Philips Intellispace Portal における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2018-003330](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003330.html) | Philips Intellispace Portal における信頼性のない検索パスに関する脆弱性 | 重要 (7.8) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2018-003329](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003329.html) | Philips Intellispace Portal における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2018-003328](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003328.html) | Philips Intellispace Portal における暗号に関する脆弱性 | 重要 (7.5) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2018-003327](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003327.html) | Philips Intellispace Portal における暗号に関する脆弱性 | 重要 (7.5) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2018-003326](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003326.html) | Philips Intellispace Portal における暗号に関する脆弱性 | 重要 (7.5) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2018-003325](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003325.html) | Philips Intellispace Portal における暗号に関する脆弱性 | 重要 (7.5) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2018-003324](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003324.html) | Philips Intellispace Portal における認可・権限・アクセス制御に関する脆弱性 | 重要 (8.1) | 2018年2月26日 | 2018年5月22日 |
| [JVNDB-2016-008996](https://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-008996.html) | Malwarebytes Anti-Malware consumer におけるセキュリティ機能に関する脆弱性 | 重要 (7.8) | 2016年8月28日 | 2018年5月22日 |
| [JVNDB-2018-003321](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003321.html) | Apache Syncope における脆弱性 | 重要 (7.2) | 2018年3月20日 | 2018年5月22日 |
| [JVNDB-2018-003320](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003320.html) | Commons-Email における入力確認に関する脆弱性 | 重要 (7.5) | 2018年1月26日 | 2018年5月22日 |
| [JVNDB-2018-003318](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003318.html) | Tenda AC15 における認可・権限・アクセス制御に関する脆弱性 | 緊急 (9.8) | 2018年3月19日 | 2018年5月22日 |
| [JVNDB-2018-003315](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003315.html) | 複数の Microsoft 製品におけるリモートでコードを実行される脆弱性 | 重要 (8.8) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003314](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003314.html) | 複数の Microsoft Windows 製品におけるリモートでコードを実行される脆弱性 | 重要 (8.8) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003313](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003313.html) | 複数の Microsoft Windows 製品におけるリモートでコードを実行される脆弱性 | 重要 (8.8) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003311](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003311.html) | 複数の Microsoft Windows 製品における権限を昇格される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003308](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003308.html) | 複数の Microsoft Windows 製品の Microsoft JET データベースエンジンにおけるバッファオーバーフローの脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003307](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003307.html) | 複数の Microsoft 製品におけるリモートでコードを実行される脆弱性 | 重要 (8.8) | 2018年4月3日 | 2018年5月22日 |
| [JVNDB-2018-003305](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003305.html) | Microsoft Internet Explorer 11 におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003304](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003304.html) | Microsoft Excel 2010 におけるリモートでコードを実行される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003303](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003303.html) | Microsoft Internet Explorer 9 から 11 のスクリプトエンジンにおけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003302](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003302.html) | Internet Explorer におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003301](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003301.html) | Internet Explorer のスクリプトエンジンにおけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003300](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003300.html) | Microsoft Internet Explorer におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003299](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003299.html) | Microsoft Internet Explorer 9 から 11 のスクリプトエンジンにおけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003298](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003298.html) | 複数の Microsoft Windows 製品の VBScript エンジンにおけるリモートでコードを実行される脆弱性 | 重要 (8.8) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003297](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003297.html) | Microsoft Excel におけるリモートでコードを実行される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月22日 |
| [JVNDB-2018-003296](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003296.html) | Cisco IOS ソフトウェアおよび Cisco IOS XE ソフトウェアにおけるバッファエラーの脆弱性 | 緊急 (9.8) | 2018年3月28日 | 2018年5月22日 |
| [JVNDB-2018-003288](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003288.html) | Microsoft Office におけるリモートでコードを実行される脆弱性 | 重要 (8.8) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003287](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003287.html) | Microsoft Internet Explorer 11 におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003286](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003286.html) | Microsoft Internet Explorer 9 から 11 におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003285](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003285.html) | Microsoft Edge および ChakraCore におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003284](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003284.html) | 複数の Microsoft Excel および Office 製品におけるリモートでコードを実行される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003283](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003283.html) | 複数の Microsoft Excel および Office 製品におけるリモートでコードを実行される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003282](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003282.html) | Microsoft Office におけるリモートでコードを実行される脆弱性 | 重要 (8.8) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003281](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003281.html) | 複数の Microsoft Windows 製品の HTTP 2.0 プロトコルスタックにおけるサービス運用妨害 (DoS) の脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003273](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003273.html) | 複数の Microsoft Windows 製品の Windows カーネルにおける権限を昇格される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003265](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003265.html) | 複数の Microsoft 製品の Windows フォントライブラリにおけるリモートでコードを実行される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003264](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003264.html) | 複数の Microsoft 製品の Windows フォントライブラリにおけるリモートでコードを実行される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003263](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003263.html) | 複数の Microsoft 製品の Windows フォントライブラリにおけるリモートでコードを実行される脆弱性 | 重要 (7.8) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003261](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003261.html) | ChakraCore におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003260](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003260.html) | ChakraCore におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003259](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003259.html) | ChakraCore におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003258](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003258.html) | ChakraCore におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003257](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003257.html) | ChakraCore におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003256](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003256.html) | ChakraCore におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |
| [JVNDB-2018-003255](https://jvndb.jvn.jp/ja/contents/2018/JVNDB-2018-003255.html) | ChakraCore におけるリモートでコードを実行される脆弱性 | 重要 (7.5) | 2018年4月10日 | 2018年5月21日 |

(Powerd by [JVN](https://jvn.jp/) & [jvnman](https://github.com/spiegel-im-spiegel/jvnman "spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management"))
