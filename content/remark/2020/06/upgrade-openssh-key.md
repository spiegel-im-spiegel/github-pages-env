+++
title = "OpenSSH 鍵をアップグレードする【2020-01-11 改訂】"
date =  "2020-06-01T16:12:32+09:00"
description = "どうせ鍵を新調するのなら楕円曲線暗号で構成するのがいいよね。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "openssh", "ubuntu", "cryptography", "hash", "sha-1", "gnupg", "ecc" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

**【2020-01-11 改訂】**

2020-05-27 に [OpenSSH] 8.3 がリリースされた。

- [OpenSSH 8.3 released (and ssh-rsa deprecation notice) [LWN.net]](https://lwn.net/Articles/821544/)

この中で “Future deprecation notice” として

{{< fig-quote type="markdown" title="OpenSSH 8.3 released (and ssh-rsa deprecation notice)" link="https://lwn.net/Articles/821544/" lang="en" >}}
{{% quote %}}It is now possible to perform chosen-prefix attacks against the SHA-1 algorithm for less than USD$50K. For this reason, we will be disabling the "ssh-rsa" public key signature algorithm by default in a near-future release{{% /quote %}}.
{{< /fig-quote >}}

と書かれていてた[^sha1a]。
といっても，これはサーバ側の設定の話だそうでクライアントには関係ないそうだ（[フィードバック](https://github.com/spiegel-im-spiegel/github-pages-env/discussions/85)感謝）。

[^sha1a]: 実は同様の問題は [GnuPG] でも指摘されていて，半年前にリリースされた 2.2.18 で対応済みである（[GnuPG 2.2.18 リリース： さようなら SHA-1]({{< ref "/release/2019/11/gnupg-2_2_18-is-released.md" >}})）

<!--
クライアント側はそろそろ古い RSA 鍵から交換したほうがよさそうである。
もっとも最近のバージョン[^sha1b] で作った鍵であれば特に問題ないようだ。

[^sha1b]: ちなみに，私は2011年に作った RSA 鍵を使っているが，ハッシュ・アルゴリズムは SHA256 だった。少なくとも10年以内に作った鍵なら問題なさそう？ あとは鍵長かねぇ。

```text
$ ssh-keygen -t rsa
Generating public/private rsa key pair.
Enter file in which to save the key (/home/username/.ssh/id_rsa): 
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in /home/username/.ssh/id_rsa
Your public key has been saved in /home/username/.ssh/id_rsa.pub
The key fingerprint is:
SHA256:qufsNjgco3QZNjE4eupwQiT6mD8fr2a7nXmU3ybxFHo username@hostname
The key's randomart image is:
+---[RSA 3072]----+
|   .             |
|..o o            |
|+. . o           |
|o.. +       .    |
|.* . +  S. . .   |
|* + =  .o o E    |
|o= +.+.. . *     |
| .+ **+o. o +    |
|   =*XOo   o     |
+----[SHA256]-----+
```

ハッシュ・アルゴリズムが SHA256 以上になっているか，がポイント。
-->

もし2048ビット以下の古い RSA 鍵を使ってるなら替えどきかな。
で，どうせ鍵を新調するのなら楕円曲線暗号（Elliptic Curve Cryptography; ECC）で構成するのがいいよね。

ちうわけで，ECC 鍵を作って登録するところまでやってみる。
ちゃんとメモっておかないと忘れるので（笑）

## 前提条件【2021-01-09 変更】

今回は [Ubuntu] 環境下での作業とし鍵管理を [GnuPG] で行うものとする。
[Ubuntu] での設定方法は以下を参考にどうぞ。

- [gpg-agent の設定： GnuPG in Ubuntu]({{< ref "/openpgp/gpg-agent-in-ubuntu.md" >}})

Windows 環境の場合は以下の拙文を参考にどうぞ。

- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}})

ちなみに，今回は鍵の生成を [OpenSSH] で行っているが [GnuPG] でも生成できる。
詳しくは以下を参照のこと。

- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})

## [OpenSSH] 鍵の生成

[OpenSSH] 鍵を生成するには `ssh-keygen` コマンドを使う。
たとえばこんな感じ。

```text
$ ssh-keygen -t ecdsa -b 256 -C "alice@example.com"
Generating public/private ecdsa key pair.
Enter file in which to save the key (/home/username/.ssh/id_ecdsa): 
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in /home/username/.ssh/id_ecdsa
Your public key has been saved in /home/username/.ssh/id_ecdsa.pub
The key fingerprint is:
SHA256:DtXgQm9rz7Dc5M5yWu/CNVo341o1rcfN9UCyYu+SZU4 alice@example.com
The key's randomart image is:
+---[ECDSA 256]---+
|      . .        |
|     . o o       |
|      . = . . .  |
|       + .   +  .|
|      . S + . ..+|
|       = X oE +*=|
|        +.*X.+oo*|
|        .+Bo.... |
|        .+o+=.   |
+----[SHA256]-----+
```

これで作成された `id_ecdsa` および `id_ecdsa.pub` ファイルが鍵ファイルである。
ちなみに `id_ecdsa` ファイルには秘密鍵， `id_ecdsa.pub` ファイルには公開鍵が格納されている。

`-t` オプションで鍵種別を， `-b` で鍵長（ビット数）をセットする。
鍵種別と鍵長の組み合わせと，それぞれに対するセキュリティ強度（ビット数）は以下の通り[^dsa1]。

[^dsa1]: [フィードバック](https://github.com/spiegel-im-spiegel/github-pages-env/discussions/85)で教えていただいたが， [OpenSSH] は FIPS 186-3 以降の DSA に対応していないそうで，1024ビットの鍵長しか選択できないらしい。なので，最近の [OpenSSH] のデフォルト設定では DSA を無効にしているそうな。情報感謝です。あと `ecdsa-sk` では鍵長の指定はできない（無視される）ようだ。256ビット（`nistp256`）固定ってことかな？

{{< div-gen >}}
<figure>
<style>
main table.sshkeys th  {
  vertical-align:middle;
  text-align: center;
}
main table.sshkeys td  {
  vertical-align:middle;
  //text-align: center;
}
</style>
<table class="sshkeys">
<thead>
<tr>
<th>鍵種別</th>
<th>鍵長</th>
<th>強度</th>
<th>楕円曲線名</th>
</tr>
</thead>
<tbody>
<tr>
  <td><strike><code>dsa</code></strike></td>
  <td class="right"><strike>1024</strike></td>
  <td class="right">80</td>
  <td>&mdash;</td>
</tr>
<tr>
  <td rowspan="3"><code>ecdsa</code> / <code>ecdsa-sk</code></td>
  <td class="right">256</td>
  <td class="right">128</td>
  <td><code>nistp256</code></td>
</tr>
<tr>
  <td class="right">384</td>
  <td class="right">192</td>
  <td><code>nistp384</code></td>
</tr>
<tr>
  <td class="right">521</td>
  <td class="right">256</td>
  <td><code>nistp521</code></td>
</tr>
<tr>
<td><code>ed25519</code> / <code>ed25519-sk</code></td>
  <td class="right">&mdash;</td>
  <td class="right">128</td>
  <td><code>ed25519</code></td>
</tr>
<tr>
  <td rowspan="3"><code>rsa</code></td>
  <td class="right">3072</td>
  <td class="right">128</td>
  <td rowspan="3">&mdash;</td>
</tr>
<tr>
  <td class="right">7680</td>
  <td class="right">192</td>
</tr>
<tr>
  <td class="right">15360</td>
  <td class="right">256</td>
</tr>
</tbody>
</table>
</figure>
{{< /div-gen >}}

この組み合わせの（`dsa` 以外の）いずれかであれば2031年以降も問題なく使える。
なお `ecdsa-sk` および `ed25519-sk` は暗号デバイスに登録する際に使うらしい（今回は割愛する）。

余談だが楕円曲線 `ed25519` に対応する電子署名アルゴリズムは EdDSA と呼ばれ [RFC 8032] で規定されている[^ed25519]。
なんで鍵種別を `eddsa` としなかったのかは知らない。
紛らわしかったのかな？

[^ed25519]: `ed25519` (edwards25519) は [Curve25519] と双有理同値な楕円曲線で，鍵長も  [Curve25519] と同じく256ビット（セキュリティ強度128ビット）と見積もられている。ちなみに [Curve25519] は ECDH 用の楕円曲線およびそのライブラリで，公有（public domain）のソフトウェアとして公開されている。また [EdDSA は FIPS 186-5 に組み込まれる予定]({{< ref "/remark/2020/06/eddsa.md" >}} "Edwards-curve Digital Signature Algorithm")になっていて，そうなれば正式に NIST 標準となる。

### 楕円曲線と鍵長

（「[そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}})」より抜粋）

ここでちょっと横道に逸れて鍵長の話をしておく。

たとえば RSA の（平文 $M$ から暗号文 $C$ への）暗号化は以下の式で表される。

{{< fig-math class="box" >}}
\[
    C = M^e \bmod n
\]
{{< /fig-math >}}

$(e, n)$ のセットが公開鍵で， $n$ のサイズが鍵長に相当する。
実際の計算はともかく，感覚としては分かりやすいよね。

ECC の場合は暗号化の前に，まずベースとなる楕円曲線の（素数 $q$ で決められる）素体（prime fields）を決めなければならない[^pf1]。
これは以下の式で表される。

[^pf1]: 素体ではなく「標数2の体（binary fields）」を使う場合もあるが，ここでは割愛する。

{{< fig-math class="box" >}}
\[
    E : y^2 \equiv x^3 + ax + b \pmod{q}
\]
{{< /fig-math >}}

この素体上の有理点の数（位数）を $\\#E$ とした時の $\\#E$ のサイズが鍵長に相当する。
$(a,b,q)$ といったパラメータを眺めただけでは鍵長が分からないというのが面倒臭い感じである[^ecc1]。

[^ecc1]: 楕円曲線と楕円曲線暗号については結城浩さんの『[暗号技術入門 第3版](https://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/)』の付録に比較的分かりやすい解説が載っている。同じく結城浩さんの『[数学ガール ガロア理論](https://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMK4/baldandersinf-22/)』が何となく分かれば楕円曲線についても何となく分かると思う。大丈夫。私も何となくしか分かっていない（笑）

そこで ECC の場合は楕円曲線を表すパラメータのセットが標準化されている。
[OpenSSH] の場合は `nistp256`, `nistp384`, `nistp521`, `ed25519` といった楕円曲線名に対応している。
つまり楕円曲線名から楕円曲線の種類を特定し，そこから鍵長も分かる，というわけだ。

## [OpenSSH] 鍵の登録（クライアント側）

本題に戻る。
作成した [OpenSSH] 鍵ファイル `id_ecdsa` および `id_ecdsa.pub` を登録しよう。

クライアント側は `ssh-add` コマンドを使って秘密鍵ファイル `id_ecdsa` の内容を [GnuPG] の鍵束に永続的に登録できる。

```text
$ ssh-add ./id_ecdsa
Enter passphrase for ./id_ecdsa: 
Identity added: ./id_ecdsa (alice@example.com)
```

この時 `ssh-add` コマンドによるパスフレーズ入力とは別に [GnuPG] の pinentry によるパスフレーズの設定が行われる（確認を含めて2回入力する必要あり）。

{{< fig-img src="./pinentry.png" link="./pinentry.png" title="pinentry" >}}

Pinentry で設定するパスフレーズは `id_ecdsa` ファイルに対するものとは管理が異なるので注意。
というか [GnuPG] の鍵束に登録したら `id_ecdsa` ファイルは不要になる。

[OpenSSH] 秘密鍵が登録できたかどうかは `~/.gnupg/sshcontrol` ファイルで確認できる。
ちゃんと登録できていれば以下のような内容が追記される。

```text
# ECDSA key added on: 2020-06-01 14:05:35
# Fingerprints:  MD5:e4:5b:66:a6:03:9a:a4:0e:f2:1b:a5:04:72:93:f3:f0
#                SHA256:DtXgQm9rz7Dc5M5yWu/CNVo341o1rcfN9UCyYu+SZU4
A5353D587000D820669B0BD55A0B4AD6897458DB 0
```

また `ssh-add -L` コマンドでも登録した鍵を確認できる。

鍵の実体は `~/.gnupg/private-keys-v1.d/` ディレクトリにある。
上述の鍵の場合は

```text
A5353D587000D820669B0BD55A0B4AD6897458DB.key
```

というファイル名で格納されているはずである。

## [OpenSSH] 鍵の登録（サーバ側）

ログイン先のサーバに公開鍵ファイル `id_ecdsa.pub` の中身を登録する。
つっても，どうにかして `id_ecdsa.pub` ファイルをサーバにアップロードして

```text
$ cat ./id_ecdsa.pub >> ~/.ssh/authorized_keys
```

と `authorized_keys` ファイルに追記すればよい。
追記ではなく内容を書き換えるなら

```text
$ cat ./id_ecdsa.pub > ~/.ssh/authorized_keys
```

でもよい。
書き換えるなら以前のファイルのバックアップはとってね。
パーミッションの設定も忘れずに。

```text
$ chmod 700 ~/.ssh
$ chmod 600 ~/.ssh/authorized_keys
```

老婆心ながら，新しい鍵でログインできることを確認するまでは接続中のセッションは切らないこと。
設定を間違えてログイン不能になったら悲惨だからねぇ。

{{< div-box type="markdown" >}}
### 【2021-12-13 追記】 ssh-copy-id を使う方法

[フィードバックで教えてもらった](https://github.com/spiegel-im-spiegel/github-pages-env/discussions/85#discussioncomment-1793423)が（感謝！） `ssh-copy-id` コマンド（中身は shell スクリプト）を使うともっと簡単に登録できるそうだ。

```text
$ ssh-copy-id -h
Usage: /usr/bin/ssh-copy-id [-h|-?|-f|-n] [-i [identity_file]] [-p port] [-F alternative ssh_config file] [[-o <ssh -o options>] ...] [user@]hostname
	-f: force mode -- copy keys without trying to check if they are already installed
	-n: dry run    -- no keys are actually copied
	-h|-?: print this help
```

ふむむ。
`-n` オプションで dry run できるのか。
ありがたい。

`-i` オプションで公開鍵ファイルを指定するが，標準入力からも受け付けるようなので，パイプで繋いで

```text
$ ssh-add -L | ssh-copy-id -n alice@hostname
/usr/bin/ssh-copy-id: INFO: attempting to log in with the new key(s), to filter out any that are already installed
/usr/bin/ssh-copy-id: INFO: 1 key(s) remain to be installed -- if you are prompted now it is to install the new keys
=-=-=-=-=-=-=-=
Would have added the following key(s):

ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFfjejx/Saej929myfZoBQAKgusPi2iiOxdZZfpCLxh5 (none)
=-=-=-=-=-=-=-=
```

てな感じにできる（実際に登録する際は `-n` オプションを外してね）。

あるいは [GnuPG] で作った認証鍵であれば

```text
$ gpg --export-ssh-key alice | ssh-copy-id -n alice@hostname
```

などとすることもできる（実際に登録する際は `-n` オプションを外してね）。

サーバ側に既に `~/.ssh/authorized_keys` ファイルがある場合は（公開鍵認証でログインした上で）ちゃんと追記してくれるし，パーミッションの心配も要らないようだ。

教えていただいて本当にありがとう {{% emoji "ペコン" %}}

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
{{< /div-box >}}

### 各種 [Git] リポジトリ・サービスに公開鍵を登録する

[GitHub] ではリポジトリのアクセスに [OpenSSH] 認証が使える。
[OpenSSH] 認証に使う公開鍵は設定の “[SSH and GPG keys](https://github.com/settings/keys)” で登録する（複数登録可能）。

[Bitbucket] や [GitLab] などの各種 [git][Git] リポジトリ・サービスでも同様に [OpenSSH] 公開鍵を登録できる。

よしゃあ！ これで作業完了。

## ブックマーク

- [ssh-rsa，非推奨のお知らせ](https://orumin.blogspot.com/2020/05/ssh-rsa.html)
- [OpenSSH、将来のリリースでssh-rsa公開鍵の署名アルゴリズムをデフォルトで無効に - ZDNet Japan](https://japan.zdnet.com/article/35154545/)
- [OpenSSHの認証に証明書を使う方法｜ConoHa VPSサポート](https://support.conoha.jp/v/openssh/)
- [SSHのCA認証 - Qiita](https://qiita.com/aat00000/items/a7973b104be9bfd3bb5c)
- [OpenPGP SSH access with Yubikey and GnuPG · GitHub](https://gist.github.com/artizirk/d09ce3570021b0f65469cb450bee5e29)
- [Securing SSH with OpenPGP or PIV](https://developers.yubico.com/PIV/Guides/Securing_SSH_with_OpenPGP_or_PIV.html)
- [Securing SSH with the YubiKey](https://developers.yubico.com/SSH/)
- [セキュリティキー「YubiKey」でEC2へのSSHを2段階認証にしてみた | Developers.IO](https://dev.classmethod.jp/articles/ssh-ubuntu-ec2-with-fido-u2f-security-key/)
- [OpenSSHがSHA-1を使用したRSA署名を廃止、BacklogのGitで発生した問題と解決にいたるまでの道のり | 株式会社ヌーラボ(Nulab inc.)](https://nulab.com/ja/blog/backlog/disables-rsa-sig-using-the-sha-1-in-openssh/)

- [暗号鍵関連の各種変数について]({{< ref "/remark/2017/10/key-parameters.md" >}})
- [（何度目かの）さようなら SHA-1]({{< ref "/remark/2020/01/sayonara-sha-1.md" >}})
- [Edwards-curve Digital Signature Algorithm]({{< ref "/remark/2020/06/eddsa.md" >}})

[OpenSSH]: https://www.openssh.com/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[RFC 8032]: https://tools.ietf.org/html/rfc8032 "RFC 8032 - Edwards-Curve Digital Signature Algorithm (EdDSA)"
[Git]: https://git-scm.com/
[GitHub]: https://github.com/
[Bitbucket]: https://bitbucket.org/product/
[GitLab]: https://gitlab.com/
[Curve25519]: http://cr.yp.to/ecdh.html "Curve25519: high-speed elliptic-curve cryptography"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
