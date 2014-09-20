# subl3

<http://www.sublimetext.com/3>

## Package control 安装

**subl3** 
<https://sublime.wbond.net/installation#st3>

按Ctrl+`调出console执行

	import urllib.request,os,hashlib; h = '7183a2d3e96f11eeadd761d777e62404' + 'e330c659d4bb41d3bdf022e94cab3cd0'; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); by = urllib.request.urlopen( 'http://sublime.wbond.net/' + pf.replace(' ', '%20')).read(); dh = hashlib.sha256(by).hexdigest(); print('Error validating download (got %s instead of %s), please try manual install' % (dh, h)) if dh != h else open(os.path.join( ipp, pf), 'wb' ).write(by)


## Plugins packages


SublimeLinter
:	错误提示检查错误

Sublime CodeIntel
:	 为部分语言增强自动完成功能，包括了Python。这个插件同时也可以让你跳转到符号定义的地方，通过按住alt并点击符号。非常方便。


DocBlockr
: 	提供类PHP的注释 以 `/** tab` 开始



BracketHighlighter
: 	标签高亮，括号双引号等

Sublime Alignment
:   对齐

Theme-Spacegray
:    主题


CTags
: 	ctags 还要说么，win下就不要装了，麻烦

### Git

Git
:	git cmd


GitGutter
:	显示修改部分  这个有压力
	cd ./config/sublime\ text\ 3/Packages/
	git clone git://github.com/jisaacks/GitGutter.git


gist  


### PHP

SublimeLinter-php  
:   PHP 检查，windows需要给出php.exe所在位置位系统变量


CodeIgniter Utilities
:    CI 

phpcs
:    php class



### 前端

SublimeLinter-jscs
:    js 检错

Emmet(ZenCoding)
: 	前端必备，快速开发HTML/CSS，现已更名为Emmet。

JsFormat
:	javascript 格式化

JQuery
:	jQuery 必备



javascript Patterns
:    js处理


AngularJS Snippets   
Ionic Framework Snippets

### golang

GoSublime
:	golang 必备,需要gocode



## linux下中文输入fcitx解决方案

项目 <https://github.com/pavelhurt/sublime2-fcitx-fix/blob/master/sublime-imfix.c> 已经给出解决方案.   
可以保存 [sublime-imfix.c](https://raw.github.com/pavelhurt/sublime2-fcitx-fix/master/sublime-imfix.c).

编译

	gcc -shared -o libsublime-imfix.so sublime_imfix.c  `pkg-config --libs --cflags gtk+-2.0` -fPI

为subl启动开启此插件. 注意将下面的`pwd`修改为自己的路径.

	LD_PRELOAD=/pwd/libsublime-imfix.so sublime_text

可以为其定义 `alias` ,另外快捷方式也可以加上该前缀.

编译环境

sudo apt-get install pkg-config
sudo apt-get install build-essential
sudo apt-get install libgtk2.0-dev

gcc -shared -o libsublime-imfix.so sublime_imfix.c  `pkg-config --libs --cflags gtk+-2.0` -fPIC


参看<http://www.apkdv.com/ubuntu-sublime-text-3-chinese-input/>

