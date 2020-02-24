+++
title = "Hello Rust World"
date =  "2020-02-24T10:02:28+09:00"
description = "description"
image = "/images/attention/kitten.jpg"
tags = [ "remark" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [Rust] の最新バージョンは 1.41.0


```text
$ curl https://sh.rustup.rs -sSf | sh -s -- --help
rustup-init 1.21.0 (fd87b86c6 2019-12-19)
The installer for rustup

USAGE:
    rustup-init [FLAGS] [OPTIONS]

FLAGS:
    -v, --verbose           Enable verbose output
    -q, --quiet             Disable progress output
    -y                      Disable confirmation prompt.
        --no-modify-path    Don't configure the PATH environment variable
    -h, --help              Prints help information
    -V, --version           Prints version information

OPTIONS:
        --default-host <default-host>              Choose a default host triple
        --default-toolchain <default-toolchain>    Choose a default toolchain to install
        --default-toolchain none                   Do not install any toolchains
        --profile [minimal|default|complete]       Choose a profile
    -c, --component <components>...                Component name to also install
    -t, --target <targets>...                      Target name to also install
```

インストール用 shell スクリプト Rustup

```text
$ curl https://sh.rustup.rs -sSf
```

```text
$ curl https://sh.rustup.rs -sSf | sh
info: downloading installer

Welcome to Rust!

This will download and install the official compiler for the Rust
programming language, and its package manager, Cargo.

It will add the cargo, rustc, rustup and other commands to
Cargo's bin directory, located at:

  /home/spiegel/.cargo/bin

This can be modified with the CARGO_HOME environment variable.

Rustup metadata and toolchains will be installed into the Rustup
home directory, located at:

  /home/spiegel/.rustup

This can be modified with the RUSTUP_HOME environment variable.

This path will then be added to your PATH environment variable by
modifying the profile file located at:

  /home/spiegel/.profile

You can uninstall at any time with rustup self uninstall and
these changes will be reverted.

Current installation options:


   default host triple: x86_64-unknown-linux-gnu
     default toolchain: stable
               profile: default
  modify PATH variable: yes

1) Proceed with installation (default)
2) Customize installation
3) Cancel installation
>

```

```text
info: profile set to 'default'
info: default host triple is x86_64-unknown-linux-gnu
info: syncing channel updates for 'stable-x86_64-unknown-linux-gnu'
info: latest update on 2020-01-30, rust version 1.41.0 (5e1a79984 2020-01-27)
info: downloading component 'cargo'
info: downloading component 'clippy'
info: downloading component 'rust-docs'
 12.0 MiB /  12.0 MiB (100 %)   8.0 MiB/s in  1s ETA:  0s
info: downloading component 'rust-std'
 17.5 MiB /  17.5 MiB (100 %)   7.9 MiB/s in  2s ETA:  0s
info: downloading component 'rustc'
 57.9 MiB /  57.9 MiB (100 %)   8.1 MiB/s in  7s ETA:  0s
info: downloading component 'rustfmt'
info: installing component 'cargo'
info: installing component 'clippy'
info: installing component 'rust-docs'
 12.0 MiB /  12.0 MiB (100 %)   6.4 MiB/s in  1s ETA:  0s
info: installing component 'rust-std'
 17.5 MiB /  17.5 MiB (100 %)  15.0 MiB/s in  1s ETA:  0s
info: installing component 'rustc'
 57.9 MiB /  57.9 MiB (100 %)  11.5 MiB/s in  5s ETA:  0s
info: installing component 'rustfmt'
info: default toolchain set to 'stable'

  stable installed - rustc 1.41.0 (5e1a79984 2020-01-27)


Rust is installed now. Great!

To get started you need Cargo's bin directory ($HOME/.cargo/bin) in your PATH
environment variable. Next time you log in this will be done
automatically.

To configure your current shell run source $HOME/.cargo/env
```

`rustup` コマンドインストール後は `rustup update` コマンドでおｋ。

```text
$ rustup update
info: syncing channel updates for 'stable-x86_64-unknown-linux-gnu'
info: checking for self-updates

  stable-x86_64-unknown-linux-gnu unchanged - rustc 1.41.0 (5e1a79984 2020-01-27)

info: cleaning up downloads & tmp directories
```


`$HOME/.prifle` の編集

```bash
export PATH="$HOME/.cargo/bin:$PATH"
```

```bash
# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/.cargo/bin" ] ; then
  PATH="$PATH:$HOME/.cargo/bin"
fi
```

```text
$ rustc --version
rustc 1.41.0 (5e1a79984 2020-01-27)
```


ドキュメント

```text
$ rustup doc --book
```

## Hello World

```rust
fn main() {
    // 世界よ、こんにちは
    println!("Hello, world!");
}
```

https://atom.io/packages/ide-rust

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"
<!-- eof -->
