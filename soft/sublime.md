# sublime Text 2

## plugin 
也可以安装package control组件，然后直接在线安装：

按Ctrl+`调出console

粘贴以下代码到底部命令行并回车：

	import urllib2,os;pf='Package Control.sublime-package';ipp=sublime.installed_packages_path();os.makedirs(ipp) if not os.path.exists(ipp) else None;open(os.path.join(ipp,pf),'wb').write(urllib2.urlopen('http://sublime.wbond.net/'+pf.replace(' ','%20')).read())

重启Sublime Text 2。

如果在Perferences->package settings中看到package control这一项，则安装成功。

[package control](http://wbond.net/sublime_packages/package_control/installation)

用Package Control安装插件的方法：

按下Ctrl+Shift+P调出命令面板,输入install 调出 Install Package 选项并回车，然后在列表中选中要安装的插件。


### plugin:
[sublime packages](http://wbond.net/sublime_packages/)

* [git](https://github.com/kemayo/sublime-text-2-git/wiki)
* [lint](https://github.com/SublimeLinter/SublimeLinter)
	*[lint use](http://www.avrw.com/article/art_109_2102.htm)
* [SFTP](http://wbond.net/sublime_packages/sftp/usage)


* [Zen Codeing]
	[Zen Coding: 一种快速编写HTML/CSS代码的方法](http://www.qianduan.net/zen-coding-a-new-way-to-write-html-code.html)
* [Sublime Prefixr]
	Prefixr，CSS3 私有前缀自动补全插件，显然也很有用哇

* [GoSublime] golang 


Bracket Highligter:
	突出了括号，引号和HTML标签
