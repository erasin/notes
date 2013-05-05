#ArchLinux离线wiki包的使用

## 安装 Arch wiki

>	sudo pacman -S arch-wiki-docs arch-wiki-lite

## 使用方法
[bash]
	wiki-search arch     # 输出带有编号的 arch相关的wiki信息
	wiki-search 1	      # 在终端中查看上面所查询中条目为1的页面
	wiki-search-html 1 # 查看html页面,需要设置浏览器
[/bash]

## 配置语言
配置脚本可以放在 ~/.bashrc 中

设置默认为中文输出
[bash]
	export wiki_lang=&amp;quot;简体中文&amp;quot;
	export wiki_browser=&amp;quot;/usr/bin/firefox&amp;quot;   # 设置wiki-search-html所需的浏览器
[/bash]

使用命令

>	wiki-search --lang  #列出当前可以使用的语言版本,默认为英语
