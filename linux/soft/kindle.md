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
5. use FBReader review mobi file.
6. copy mobi file to Kindle Documents