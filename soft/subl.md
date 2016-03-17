# subl3

<http://www.sublimetext.com/3>

## Package control 安装

**subl3** 

下载地址： <https://sublime.wbond.net/installation#st3>

<!-- https://packagecontrol.io/installation -->
按`Ctrl+\``调出console执行

```
import urllib.request,os,hashlib; h = 'eb2297e1a458f27d836c04bb0cbaf282' + 'd0e7a3098092775ccb37ca9d6b2e4b7d'; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); by = urllib.request.urlopen( 'http://packagecontrol.io/' + pf.replace(' ', '%20')).read(); dh = hashlib.sha256(by).hexdigest(); print('Error validating download (got %s instead of %s), please try manual install' % (dh, h)) if dh != h else open(os.path.join( ipp, pf), 'wb' ).write(by)
```

vim 模式
```
"ignored_packages":
	[
		"Markdown",
		// "Vintage"
	],
```

## Plugins packages

**SublimeLinter**
> 错误提示检查错误

**Sublime CodeIntel**
> 它提供了很多IDE提供的功能，例如代码自动补齐，快速跳转到变量定义，在状态栏显示函数快捷信息等。

**DocBlockr**
> 提供类PHP的注释 以 `/** tab` 开始

**Bracket Highlighter**
> 标签高亮，括号双引号等

**Sublime Alignment**
>  对齐

**Theme-Spacegray**
>  主题

**CTags**
>  ctags 还要说么，win下就不要装了，麻烦

### Git

**Git**

git 命令集合

**GitGutter**

> 显示修改部分  这个有压力,对电脑配置有要求

	cd ./config/sublime\ text\ 3/Packages/
	git clone git://github.com/jisaacks/GitGutter.git

**gist**  

> 可以显示Github的gist代码片段

**filediff**

> 文件差异比较

### PHP

**SublimeLinter-php**  
>   PHP 检查错误，windows需要给出php.exe所在位置配置**系统变量**

**CodeIgniter Utilities**
>  CI 代码片段

**phpcs**
>  php class 格式自检测

### JavaScript

**Emmet(ZenCoding)**
> 前端必备，快速开发HTML/CSS，现已更名为Emmet。

**SublimeLinter-jscs**
>  js 检错 for nodejs

**JQuery**
> jQuery 必备

**javascript completions**
自动提示

**javascript Patterns**
> js处理

**JsFormat**
> javascript 格式化

**AutoPrefixr**
> Prefixr插件能将CSS3代码自动生成针对不同的浏览器写一堆的CSS3代码前缀; 按下control + command + X（Mac）或者 ctrl + Alt + X（Windows），会转换

**AngularJS Snippets**
Ionic Framework Snippets

### CSS

**Css format**

**ColorPicker**
> 色彩调试

### nodejs

**nodejs**
>   nodejs

**JavaScript & NodeJS Snippets**

### golang

**GoSublime**
>  golang 必备,需要gocode 

设定环境变量

	{
		"env":{"GOPATH":"$HOME/golang:$GS_GOPATH"}
	}

### C

**c99**

**clang Format**

## markdown

**MarkdownEditting**

> MarkdownEditing 从视觉和便捷性上针对 Markdown 文档的编辑进行了一系列的优化。具体如下：  
> 安装后针对 md\mdown\mmd\txt 格式文件启用插件。颜色方案仿 Byword 及 iA writer。  
> 自动匹配星号（*）、下划线（_）及反引号（`），选中文本按下以上符号能自动在所选文本前后添加配对的符号，方便粗体、斜体和代码框的输入。  
> 直接输入配对的符号后按下退格键（backspace），则两个符号都会被删除；直接输入配对的符号后按下空格键，则会自动删除后一个。  
> 对“选中文字后输入左括号”这一动作进行了调整，以便插入 markdown 链接。  
> 拷贝一个链接，选中文本后按下 ⌘⌥V 会自动插入内联链接。  
> 拷贝一个链接，选中文本后按下 ⌘⌥R 会自动插入引用链接。  
> ⌘⌥K 插入链接；⌘⇧K 插入图片。  
> ⌘⌥B 和 ⌘⌥I 分别用于加粗体和斜体。  
> 选中文本后按下 # 会自动在文本前后进行配对，可重复按下来定义标题级别，还可用 ⌘⇧空格 来增加 # 与所选文本之间的空格（也是自动配对的）。  

**AcademicMarkdown** 

> 一种写作方式来书写markdown

## linux下中文输入fcitx解决方案

项目 <https://github.com/pavelhurt/sublime2-fcitx-fix/blob/master/sublime-imfix.c> 已经给出解决方案.   
可以保存 [sublime-imfix.c](https://raw.github.com/pavelhurt/sublime2-fcitx-fix/master/sublime-imfix.c).

编译 

	gcc -shared -o libsublime-imfix.so sublime_imfix.c  `pkg-config --libs --cflags gtk+-2.0` -fPI

为subl启动开启此插件. __注意__将_下_面的`pwd`修改为自己的路径.

	LD_PRELOAD=/pwd/libsublime-imfix.so sublime_text

可以为其定义 `alias` ,另外快捷方式也可以加上该前缀.

编译环境

	sudo apt-get install pkg-config
	sudo apt-get install build-essential
	sudo apt-get install libgtk2.0-dev

gcc -shared -o libsublime-imfix.so sublime_imfix.c  `pkg-config --libs --cflags gtk+-2.0` -fPIC

参看[ubuntu sublime3 chinese input](http://www.apkdv.com/ubuntu-sublime-text-3-chinese-input/)


SideBarEnhancements 是一款很实用的右键菜单增强插件，有以 diff 形式显示未保存的修改、在文件管理器中显示该文件、复制文件路径、在侧边栏中定位该文件等功能，也有基础的诸如新建文件/目录，编辑，打开/运行，显示，在选择中/上级目录/项目中查找，剪切，复制，粘贴，重命名，删除，刷新等常见功能。

