## Rust

https://github.com/mozilla/sccache

cargo 加速和缓存编译，`vim ~/.cargo/config`

```shell
[source.crates-io]
registry = "https://github.com/rust-lang/crates.io-index"
replace-with = 'ustc'
[source.ustc]
registry = "git://mirrors.ustc.edu.cn/crates.io-index"


[build]
rustc-wrapper = "/home/lv/.cargo/sccache"
```

### Cargo

Cargo toml

```toml
[package]
name = "hello_cargo"
version = "0.1.0"
authors = ["wuqz <wuqz@siccs.cn>"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]

```

### macro

```rust

```

