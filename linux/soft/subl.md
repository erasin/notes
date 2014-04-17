#subl3

## Package control 安装

**subl3** 

	cd Packages
	git clone https://github.com/wbond/sublime_package_control.git "Package Control"

	import urllib.request,os; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); open(os.path.join(ipp, pf), 'wb').write(urllib.request.urlopen( 'http://sublime.wbond.net/' + pf.replace(' ','%20')).read())

**subl2**

按Ctrl+`调出console执行

	import urllib2,os;pf='Package Control.sublime-package';ipp=sublime.installed_packages_path();os.makedirs(ipp) if not os.path.exists(ipp) else None;open(os.path.join(ipp,pf),'wb').write(urllib2.urlopen('http://sublime.wbond.net/'+pf.replace(' ','%20')).read())

## Plugins packages

SublimeLinter
:	错误提示

Emmet(ZenCoding)
: 	前端必备，快速开发HTML/CSS，现已更名为Emmet。

JsFormat
:	javascript 格式化

JQuery
:	jQuery 必备

GoSublime
:	golang 必备

Soda themes
:	可以直接用 Package install来安装.   
	主题介绍: <https://github.com/buymeasoda/soda-theme/>  
	对应的高亮方案: <http://buymeasoda.github.com/soda-theme/extras/colour-schemes.zip> ,内部文件到 Packages->User 中即可.
上图的扁平化 Theme： Flatland

CTags
: 	ctags 还要说么

Git
:	git cmd

DocBlockr
: 	提供类PHP的注释 以 `/** tab` 开始

BracketHighlighter
: 	标签高亮，括号双引号等

Sublime Alignment
:   对齐


javascript Patterns  
CodeIgniter Utilities  
gist  




## linux下中文输入fcitx解决方案

项目 <https://github.com/pavelhurt/sublime2-fcitx-fix/blob/master/sublime-imfix.c> 已经给出解决方案.   
可以保存 [sublime-imfix.c](https://raw.github.com/pavelhurt/sublime2-fcitx-fix/master/sublime-imfix.c).

编译

	gcc -shared -o libsublime-imfix.so sublime_imfix.c  `pkg-config --libs --cflags gtk+-2.0` -fPI

为subl启动开启此插件. 注意将下面的`pwd`修改为自己的路径.

	LD_PRELOAD=/pwd/libsublime-imfix.so sublime_text

可以为其定义 `alias` ,另外快捷方式也可以加上该前缀.

