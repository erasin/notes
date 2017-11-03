mdbook
===========

[github](https://github.com/azerupi/mdBook) [doc](http://azerupi.github.io/mdBook/)

<!-- TOC -->

- [1. 安装](#1-安装)
- [2. 使用](#2-使用)
- [3. 高阶](#3-高阶)
    - [3.1. 为页面追加 js 或者 css](#31-为页面追加-js-或者-css)
    - [3.2. 修改 index.hbs 模板](#32-修改-indexhbs-模板)

<!-- /TOC -->

# 1. 安装

建议到[mdbook 项目地址](https://github.com/azerupi/mdBook/releases)，下载与系统对应的最新版本。

如果安装了 RUST 环境，可以使用 cargo 安装

```bash
cargo install mdbook
```

# 2. 使用

**mdbook** 命令

- help 查看帮助
- init 创建和初始化文档
- build 生成book
- serve 建立本地服务，默认地址为 <http://localhost:3000>
- watch 监控文件变动

使用 init 创建文档

```bash
# 1. 创建 bookname 
mdbook init bookname  # 直接生成文件夹以及初始文件
# 1.2 将现有文件夹转换为 book, 和上面等价
mkdir  bookname
cd bookname
mdbook init
```

使用 serve 建立本地服务，实时预览

```bash
mdbook serve
```

在发布 html book 之前 ，需要使用 build 生成正式的文件

> 在使用 serve 命令的时候，默认开启了 WebSocket 来自动更新，所以最后需要使用 build 来生成正式的文件。 


# 3. 高阶

在 mdbook 中可以追加 `book.toml` 和 `theme` 文件夹来实现高级的功能。

book.toml， 内部参数可以选择使用，仅需添加需要的配置

```toml
title = ""
author = ""
description = ""

[output.html]
# theme 路径
theme = "./theme"
mathjax-support = true
additional-js = ["./theme/xxx.js"]
additional-css = ["./theme/xxx.css"]
google_analytics = "" # 这里为统计ID

# 使用编辑器
[output.html.playpen]
#editor = ""  #自定义编辑器的时候输入
editable = true 

```

## 3.1. 为页面追加 js 或者 css

比如想要将页面中的 link 中特定的字符串转换为图标

1. 修改 `book.toml` 追加 js, 让页面上加载 ` Custom JS script `

    ```toml
    ...

    [output.html]
    additional-js = ["./theme/icons.js"]

    ...
    ```

2. 创建 `./theme/icons.js` 追加内容

    ```js
    $( document ).ready(function() {
        $('a').each(function(element) {
            var link = $(this);
            switch (link.text()) {
                case "github":
                    link.html('<i class="fa fa-github"></i>');
                    break;
                case "doc":
                    link.html('<i class="fa fa-book"></i>');
                    break;
                case "site":
                    link.html('<i class="fa fa-home"></i>');
                    break;
                case "crate":
                    link.html('<i class="fa fa-cubes"></i>');
                    break;
                default:
                    break;
            }
        });
    });
    ```

>  参照 [fontawesome-icon](http://fontawesome.io/icons/) 进行图标处理

## 3.2. 修改 index.hbs 模板

mdbook 提供了可以直接修改文章模板来直接修改布局， 可以参考和拷贝[默认模板](https://github.com/azerupi/mdBook/blob/master/src/theme/index.hbs) 来修改出自己的默认布局。

在项目文件夹目录内创建 theme 文件夹，可以在内部替换默认的一些文件

- index.hbs  文档模板
- book.js  主要 js 处理
- book.css 页面样式
- favicon.png 页面图标
- ...

> 这里不建议替换，而是选择配置 `additional-js|css` 来扩展定义。当然需要进行大量的定制的情况除外。

index.hbs 使用的是 handlebars 模板系统， 内部变量统一使用 `{{ name_of_property }}` 来表示。
模板提供了一些默认的配置和变量。  建议阅读[官方文档](http://azerupi.github.io/mdBook/format/theme/index-hbs.html) 和 [默认模板](https://github.com/azerupi/mdBook/blob/master/src/theme/index.hbs) 进行解析。
