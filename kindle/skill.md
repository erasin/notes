将epub用压缩软件打开修改其中的opf文件
将<spine toc="ncx">改成<spine toc="ncx" page-progression-direction="rtl"搜索>
打开CSS文件，在body的大括号中添加一行：-webkit-writing-mode : vertical-rl ;
在CSS文件中，将p的大括号里的text-align句子改成text-align: left;