#subl3

## Package control 安装

```bash
cd Packages
git clone https://github.com/wbond/sublime_package_control.git "Package Control"
```
## 主要快捷键列表
Ctrl+L 选择整行（按住-继续选择下行）
Ctrl+KK 从光标处删除至行尾
Ctrl+Shift+K 删除整行
Ctrl+Shift+D 复制光标所在整行，插入在该行之前
Ctrl+J 合并行（已选择需要合并的多行时）
Ctrl+KU 改为大写
Ctrl+KL 改为小写
Ctrl+D 选词 （按住-继续选择下个相同的字符串）
Ctrl+M 光标移动至括号内开始或结束的位置
Ctrl+Shift+M 选择括号内的内容（按住-继续选择父括号）
Ctrl+/ 注释整行（如已选择内容，同“Ctrl+Shift+/”效果）
Ctrl+Shift+/ 注释已选择内容
Ctrl+Z 撤销
Ctrl+Y 恢复撤销
Ctrl+M 光标跳至对应的括号
Alt+. 闭合当前标签
Ctrl+Shift+A 选择光标位置父标签对儿
Ctrl+Shift+[ 折叠代码
Ctrl+Shift+] 展开代码
Ctrl+KT 折叠属性
Ctrl+K0 展开所有
Ctrl+U 软撤销
Ctrl+T 词互换
Tab 缩进 自动完成
Shift+Tab 去除缩进
Ctrl+Shift+↑ 与上行互换
Ctrl+Shift+↓ 与下行互换
Ctrl+K Backspace 从光标处删除至行首
Ctrl+Enter 光标后插入行
Ctrl+Shift+Enter 光标前插入行
Ctrl+F2 设置书签
F2 下一个书签
Shift+F2 上一个书签


# Plugins packages

SublimeLinter:
	错误提示

Alignment:
	Ctrl+Alt+A，可以将凌乱的代码以等号为准左右对其

Emmet(ZenCoding):
	前端必备，快速开发HTML/CSS，现已更名为Emmet。

JsFormat:
	javascript 格式化

JQuery:
	jQuery 必备

GoSublime:
	golang 必备

colorpicker:
	调色板

Soda themes:
	可以直接用 Package install来安装. 
	主题介绍: <https://github.com/buymeasoda/soda-theme/>
	对应的高亮方案: <http://buymeasoda.github.com/soda-theme/extras/colour-schemes.zip> ,内部文件到 Packages->User 中即可.


# linux下中文输入fcitx解决方案

项目 <https://github.com/pavelhurt/sublime2-fcitx-fix/blob/master/sublime-imfix.c> 已经给出解决方案.   
可以保存 [sublime-imfix.c](https://raw.github.com/pavelhurt/sublime2-fcitx-fix/master/sublime-imfix.c).

编译

	gcc -shared -o libsublime-imfix.so sublime_imfix.c  `pkg-config --libs --cflags gtk+-2.0` -fPI

为subl启动开启此插件. 注意将下面的`pwd`修改为自己的路径.

	LD_PRELOAD=/pwd/libsublime-imfix.so sublime_text

可以为其定义 `alias` ,另外快捷方式也可以加上该前缀.


