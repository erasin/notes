rust 是一门很有趣的语言, 最近正在试着玩它. 有一些链接感觉很有用, 所以找个地方保存一下.

## 综述

这些链接提供了一些比较通用的介绍，手册和列表，以供在大块的时间中阅览。

- [The Rust Reference](https://doc.rust-lang.org/stable/reference.html) -- 这个链接提供了一组 rust 的语言语法规范. 理论上说粗略一看, 然后以后需要的时候查询会是一个比较经济的做法. 不过个人发现里面讲得蛮好玩的, 坐地铁看起来很舒服. 只是比较长, 不太容易一次看完.
- [Rust Guidelines](https://aturon.github.io/README.html) -- 这个链接定义了很多 rust 编写时候的规范. 代码编写的时候, 建议在形式上写得和别人一样, 这样会比较好管理.
- [Rust by Example](http://rustbyexample.com/) -- 这个链接通过一些例子介绍了 rust 的用法.
- [The Rust Programming Language](https://doc.rust-lang.org/book/README.html) -- 另一个通过例子介绍 rust 的书. 有人翻译了中文版，[在此](https://kaisery.gitbooks.io/rust-book-chinese/content/)。
- [Awesome Rust](https://github.com/kud1ing/awesome-rust) -- 好事者列出的当前不错的库列表.

## 讨论

本节列出一些有趣的博文，介绍 rust 的一些特别的机制。

- [Wrapper Types in Rust: Choosing Your Guarantees](http://manishearth.github.io/blog/2015/05/27/wrapper-types-in-rust-choosing-your-guarantees/) -- 讨论了类型定义（授权等）
- [Error Handling in Rust](http://blog.burntsushi.net/rust-error-handling/) -- 讨论了错误处理
- [A Practical Intro to Macros in Rust 1.0](https://danielkeep.github.io/practical-intro-to-macros.html) -- 讨论了宏
- [Some notes on Send and Sync](http://huonw.github.io/blog/2015/02/some-notes-on-send-and-sync/) -- 讨论了 Send/Sync

## 工具

本节列出一些工具。

- [cargo](https://crates.io/) -- 存放各种库的地方, [此处](https://argcv.com/articles/4171.c)有个详细介绍.
- [Rust Playground](https://play.rust-lang.org/) -- 在线试着跑一些代码. 还可以生成一些外链的代码以展示. 不过无法引用 crates 的库, 稍有局限

## 第三方库

本节列出一些可能有用的第三方库。

- [rustc-serialize](https://crates.io/crates/rustc-serialize) -- 这个库提供了 json 的序列化和解析，提供了普通文本和 base64，普通文本和 hex 的互转两个简单功能
- [nalgebra](http://nalgebra.org/) -- 一个线性计算库，比如物理计算 etc.
- [glium](https://github.com/tomaka/glium) -- 绑定 opengl 的一种解决方法.
- [glfw](https://crates.io/crates/glfw) -- 绑定 glfw 使用 opengl 的一种解决方法，需要自行安装 glfw.
- [mio](https://crates.io/crates/mio) -- 网络库，支持 kqueue, epoll etc.
- [coio](https://github.com/zonyitoo/coio-rs) -- 另一个网络库，强调了好久 Scheduler.
- [hyper](https://github.com/hyperium/hyper) -- hyper is a fast, modern HTTP implementation written in and for Rust. It is a low-level typesafe abstraction over raw HTTP, providing an elegant layer over "stringly-typed" HTTP.
- [tangle](https://github.com/thehydroimpulse/tangle) -- 线程库，实现了一个 "Future".
- [rust-jwt](https://github.com/keats/rust-jwt) -- JWT 的实现.
- [RustCMake](https://github.com/SiegeLord/RustCMake) -- 一个例子，介绍怎么样以 CMake 为环境，构建一个基于 cargo 的工程, 尝试了下蛮厉害的，这人是勇士啊.
- [bmemcached](https://github.com/jaysonsantos/bmemcached-rs) -- client for Memcached
- [es](https://github.com/benashford/rs-es) -- client for ElasticSearch
- [MongoDB Rust Driver Prototype](https://github.com/mongodb-labs/mongo-rust-driver-prototype) -- client for MongoDB
- [clap](https://github.com/kbknapp/clap-rs) -- Command-line argument parsing
- [rust-bindgen](https://github.com/crabtw/rust-bindgen) -- auto generate rs from c headers.
- [rusty-cheddar](https://github.com/Sean1708/rusty-cheddar) -- auto generate c headers from rust code.
- [nickel.rs](http://nickel.rs/) -- web application framework for rust.
- [IRON](http://ironframework.io/) -- extensible web framework for rust.
- [Rustless](http://rustless.org/) -- Use Rust for web with Rustless
- [pencil](http://fengsp.github.io/blog/2016/3/introducing-pencil/) -- a flask like http server.