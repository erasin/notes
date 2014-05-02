1）在正文的css处加入竖排设定：
body {
writing-mode: vertical-rl;
-webkit-writing-mode:vertical-rl;
-epub-writing-mode: vertical-rl;
-epub-line-break: auto;
}
（2）opf文件处的meta部分加入
<meta property="page-progression-direction">rtl</meta>
spine部分加入page-progression-direction，如
<spine toc="ncx" page-progression-direction="rtl">
以支持从右到左翻页。