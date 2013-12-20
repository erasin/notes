# kindle soft

软件：

* kindlegen
* FBReader
* markdown 

使用markdown -》 mobi

1. format markdown file. eg: a.md

2. convert markdown file to html file . eg: markdown a.md > a.html
3. Add meta of html file. and you can add style in file.

        <html>
        <head>
            <meta http-equiv="content-language" content="zh-CN" />
            <meta http-equiv="Content-type" content="text/html; charset=utf-8">
            <meta name="Author" content="era">
            <title> Book title </title>
        </head>
        <body>
        ... markdown content
        </body>
        </html>

4. convert html file to mobi. 

    kindlegen a.html 

5. use FBReader to preview mobi file.

6. copy mobi file to Kindle Documents

[官方教程和示例](http://www.amazon.com/gp/feature.html?ie=UTF8&docId=1000234621)

## 使用 ebook-convert

calibre 组件 ebook-convert

    ebook-convert progit.zh.html progit.cn.mobi --cover ebooks/cover.png --authors 'pro' --comments 'progit 中文' --level1-toc '//h:h1' --level2-toc '//h:h2' --level3-toc '//h:h3' --language 'zh_CN'
