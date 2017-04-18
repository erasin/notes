mdbook 
==============

mdbook 是用 rust 仿照 gitbook 写的，下载地址为 <https://github.com/azerupi/mdBook>，也可以直接使用 `cargo install mdbook` 命令下载。


### 在mdbook 中使用 mermaid

[meraid](https://github.com/knsv/mermaid) 是利用 js 来制作流程图，参考[文档](http://knsv.github.io/mermaid/)。

建立 `src/theme/index.hbs` 来追加对 meraid 的处理

```js
<!-- mermaid.js -->
<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/mermaid/6.0.0/mermaid.min.js"></script>
<script>
    $('pre .language-mermaid').each(function () {
        $(this).html('<div class="mermaid">' + this.textContent + '</div>');
    });
</script>
```

 追加 js，并解析 code。
