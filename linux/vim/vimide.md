# vim ide 

github ： <http://github.com/erasin/vimide>

使用 

	git clone git@github.com:erasin/vimide.git
	mv vimide ~/.vim
	ln -s ~/.vim/vimrc ~/.vimrc


## NERDtree 

下载地址 <http://www.vim.org/scripts/script.php?script_id=1658>

vimrc 定制快捷键 F3

## MinibuferExplorer
<http://www.vim.org/scripts/script.php?script_id=159>

file tabs buffer

	:MiniBufExplorer

## powerLine
<https://github.com/Lokaltog/vim-powerline>

如果效果, 需要 to use a patched font

	let g:Powerline_symbols = 'fancy'

清理缓存
	:PowerlineClearCache

状态栏宽度
	set laststatus=2

## doxygentoolkit
Dox 注释生成

<http://www.vim.org/scripts/script.php?script_id=987>

## NERD_commenter 
注释插件

<https://github.com/scrooloose/nerdcommenter>

使用
	
* `<leader>`ca，在可选的注释方式之间切换，比如C/C++ 的块注释/* */和行注释//
* `<leader>`cc，注释当前行
* `<leader>`c，切换注释/非注释状态
* `<leader>`cs，以”性感”的方式注释
* `<leader>`cA，在当前行尾添加注释符，并进入Insert模式
* `<leader>`cu，取消注释
* Normal模式下，几乎所有命令前面都可以指定行数
* Visual模式下执行命令，会对选中的特定区块进行注释/反注释

## align
Align 对齐 

<http://www.vim.org/scripts/script.php?script_id=294>

对其方向

	:AlignCtrl =lp1P1I 

对其规则

	:Align =
	:5,10Align = 

## taglist

<http://www.vim.org/scripts/script.php?script_id=273>

快捷键 t

需要安装 ctags

## OmniCppComplete
弹出框 使用ctags生成的文件来补全类和方法

<http://www.vim.org/scripts/script.php?script_id=1520>

需要 ctags 生成标签库

C-x C-o 打开

## superTab
省去Ctrl-n或Ctrl-p快捷键，通过按tab键快速显示补全代码

* <http://www.vim.org/scripts/script.php?script_id=1643>
* <https://github.com/ervandew/supertab>

## ultisnips
自动补全

<git://github.com/SirVer/ultisnips.git>

需要 `python`支持


## zendcodding

快速html,css 生成

<https://github.com/mattn/zencoding-vim>

使用 `c+y ,` 来将规则生成code

[zendcodding 使用教程](https://raw.github.com/mattn/zencoding-vim/master/TUTORIAL)

## golang

<https://github.com/jnwhiteh/vim-golang>

需要 `gocode`支持

未解决：
* taglist的支持

## ctrlp
使用 Ctrl+P 来快速定位文件

<http://kien.github.io/ctrlp.vim/>

## Indent Guides
使用 space 替代 tab时说使用的缩进线

<https://github.com/nathanaelkane/vim-indent-guides>

## css color
色彩显示

<http://www.vim.org/scripts/script.php?script_id=2150>

