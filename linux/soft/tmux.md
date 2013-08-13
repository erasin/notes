#tmux 

tmux是一个优秀的终端复用软件，类似GNU Screen，但来自于OpenBSD，采用BSD授权。使用它最直观的好处就是，通过一个终端登录远程主机并运行tmux后，在其中可以开启多个控制台而无需再使用更多的SSH会话来连接这台远程主机；其功能远不止于此。

1、安装
在freebsd中可以直接使用ports工具安装，位置在：/usr/ports/sysutils/tmux/，ubuntu系统下默认自带byou，与tmux很像，甚至快捷键都是一样的。这里只详细说明在Centos6.3下如何安装和使用tmux的。

Centos6.3的软件库里没有tmux，只有screen，所以要想使用tmux需要自己编译安装。

(1) 下载tmux：

wget   http://sourceforge.net/projects/tmux/files/tmux/tmux-1.6/tmux-1.6.tar.gz/download

(2) 编译安装：

tar zxvf  tmux-1.6.tar.gz

cd tmux-1.6

./configure

make;make install

2、启动tmux

安装完成后输入命令tmux即可打开软件，界面十分简单，类似一个下方带有状态栏的终端控制台；但根据tmux的定义，在开启了tmux服务器后， 会首先创建一个会话，而这个会话则会首先创建一个窗口，其中仅包含一个面板；也就是说，这里看到的所谓终端控制台应该称作tmux的一个面板，虽然其使用 方法与终端控制台完全相同。

tmux使用C/S模型构建，主要包括以下单元模块：

一个tmux命令执行后启动一个tmux服务
一个tmux服务可以拥有多个session，一个session可以看作是tmux管理下的伪终端的一个集合
一个session可能会有多个window与之关联，每个window都是一个伪终端，会占据整个屏幕
一个window可以被分割成多个pane
多个pane的编号规则，以3个pane为例:


3、tmux快捷键
tmux在会话中使用大量的快捷键来控制多个窗口、多个会话等。

    Ctrl+b  #激活控制台；此时以下按键生效   
    系统操作   
        ?   #列出所有快捷键；按q返回   
        d   #脱离当前会话；这样可以暂时返回Shell界面，输入tmux attach能够重新进入之前的会话   
        D   #选择要脱离的会话；在同时开启了多个会话时使用   
        Ctrl+z  #挂起当前会话   
        r   #强制重绘未脱离的会话   
        s   #选择并切换会话；在同时开启了多个会话时使用   
        :   #进入命令行模式；此时可以输入支持的命令，例如kill-server可以关闭服务器   
        [   #进入复制模式；此时的操作与vi/emacs相同，按q/Esc退出   
        ~   #列出提示信息缓存；其中包含了之前tmux返回的各种提示信息   
    窗口操作   
        c   #创建新窗口   
        &   #关闭当前窗口   
        数字键 #切换至指定窗口   
        p   #切换至上一窗口   
        n   #切换至下一窗口   
        l   #在前后两个窗口间互相切换   
        w   #通过窗口列表切换窗口   
        ,   #重命名当前窗口；这样便于识别   
        .   #修改当前窗口编号；相当于窗口重新排序   
        f   #在所有窗口中查找指定文本   
    面板操作   
        ”   #将当前面板平分为上下两块   
        %   #将当前面板平分为左右两块   
        x   #关闭当前面板   
        !   #将当前面板置于新窗口；即新建一个窗口，其中仅包含当前面板   
        Ctrl+方向键    #以1个单元格为单位移动边缘以调整当前面板大小   
        Alt+方向键 #以5个单元格为单位移动边缘以调整当前面板大小   
        Space   #在预置的面板布局中循环切换；依次包括even-horizontal、even-vertical、main-horizontal、main-vertical、tiled   
        q   #显示面板编号   
        o   #在当前窗口中选择下一面板   
        方向键 #移动光标以选择面板   
        {   #向前置换当前面板   
        }   #向后置换当前面板   
        Alt+o   #逆时针旋转当前窗口的面板   
        Ctrl+o  #顺时针旋转当前窗口的面板   

4、配置文件

tmux配置文件在~/.tmux.conf和/etc/tmux.conf中，配置文件中可以修改默认绑定的快捷键

配置文件示例：

	#此类配置可以在命令行模式中输入show-options -g查询   
	set-option -g base-index 1                        #窗口的初始序号；默认为0，这里设置为1   
	set-option -g display-time 5000                   #提示信息的持续时间；设置足够的时间以避免看不清提示，单位为毫秒   
	set-option -g repeat-time 1000                    #控制台激活后的持续时间；设置合适的时间以避免每次操作都要先激活控制台，单位为毫秒   
	set-option -g status-keys vi                      #操作状态栏时的默认键盘布局；可以设置为vi或emacs   
	set-option -g status-right "#(date +%H:%M' ')"    #状态栏右方的内容；这里的设置将得到类似23:59的显示   
	set-option -g status-right-length 10              #状态栏右方的内容长度；建议把更多的空间留给状态栏左方（用于列出当前窗口）   
	set-option -g status-utf8 on                      #开启状态栏的UTF-8支持   
	  
	#此类设置可以在命令行模式中输入show-window-options -g查询   
	set-window-option -g mode-keys vi    #复制模式中的默认键盘布局；可以设置为vi或emacs   
	set-window-option -g utf8 on         #开启窗口的UTF-8支持   
	  
	#将激活控制台的快捷键由Ctrl+b修改为Ctrl+a，Ctrl+a是Screen的快捷键
	set-option -g prefix C-a   
	unbind-key C-b   
	bind-key C-a send-prefix   
	  
	#添加自定义快捷键   
	bind-key z kill-session                     #按z结束当前会话；相当于进入命令行模式后输入kill-session   
	bind-key h select-layout even-horizontal    #按h将当前面板布局切换为even-horizontal；相当于进入命令行模式后输入select-layout even-horizontal   
	bind-key v select-layout even-vertical      #按v将当前面板布局切换为even-vertical；相当于进入命令行模式后输入select-layout even-vertical   


6、从screen到tmux无缝转化

（1）Screen 和 Tmux 命令对照表

Tmux 的指令是需要另外学习的, 剩下的进入 Tmux 后的操作, 只要配置文件配置跟 Screen 相同即可.

* 返回窗口
	* screen -r
	* tmux attach
* 根据session名称返回窗口
	* screen -r session-name
	* tmux attach -t session-id
* 新建session名称的窗口
	* screen -S session-name
	* tmux new -s session-name # ex: tmux new -s irc
* 列出所有 session
	* screen -ls
	* tmux ls # tmux list-sessions
* Tmux 配置 - .tmux.conf

如果, Screen 已经非常习惯了,现在要转换到 Tmux, 还的重新适应, 为此只要把 Tmux 配置成跟 Screen 一样的操作键, 从而可以实现无缝转换。

Debian / Ubuntu 安裝 Tmux 后, 设置的范例中, 就有设置跟 Screen 一样的配置文件, 还有吧 Vim 的快捷键操作方式(切割窗口)加入 的配置例子, 位置如:

/usr/share/doc/tmux/examples/screen-keys.conf
/usr/share/doc/tmux/examples/vim-keys.conf

vim ~/.tmux.conf # 将以上两个配置计入, 再加上我自己习惯的 Status 颜色的设定, 详细配置如下。 (把文件下载, 改名为~/.tmux.conf 即可)

配置文件见 [Github: shell/tmux.conf](https://raw.github.com/tsung/config/master/shell/tmux.conf)

而 Tmux 有特殊功能部份, 如下述: (设: 下述都是按照上面的 .tmux.conf, 若配置文件不同, 可能快捷键有可能会不同)

* Tmux 功能
	* C-a ? 看说明 (C 代表要先按一下 CTRL鍵)
	* C-a t 看当下时间
	* C-a q 可以看到窗口数字 - 在视窗切割的时候比较好区分
* 分割 / 切割窗口相关
	* C-a Ctrl按着 + 上下左右, 可以调整窗口大小
	* C-a s 横切
	* C-a v 纵切 (或 C-a %)
	* C-a C-o 兑换切割窗口的位置
	* C-a 上下左右 跳到上下左右的分割窗口
* 切割窗口有些 Default Layout 可以用, 比如已经切了4个窗口.
	* C-a Esc 1 直的切割窗口排列 (M-1)
	* C-a Esc 2 横的切割窗口排列 (M-2)
	* C-a Esc 3 上面一个大的横窗口 + 下面直的三个直窗口(M-3)
	* C-a Esc 4 左面一个大的横窗口 + 右面直的三个直窗口 (M-4)
	* C-a Esc 5 四个窗口各 1/4 (M-5)
* 多人共同操作/显示同一界面
	* Screen 有多人可以一起登录看同一个界面的功能, 作法: Linux 文字模式 螢幕畫面共享 - 使用Screen
	* 对Tmux 来说, 不需要做任何事情, 只要每个人登录后, 直接 tmux attach 进去, 就是多人共同看到的界面, 大家都可以操作同一界面面。

Tmux 测试 256 色支持

vim .bashrc # 若 tmux 沒有 256 色, 在 .bashrc 加入下面语句试试看。

	alias tmux='TERM=xterm-256color tmux -2'

	$ echo -e "My favorite color is \033[38;5;148mYellow-Green\033[39m" # 应该看到黄绿色的的颜色, 若不支持 256 色, 會看到绿色
	$ tput colors
	256


## 参考文件：

[screen 教學](http://blog.longwin.com.tw/2005/11/screen_teach/)
[screen 常用教學筆記(問答)](http://blog.longwin.com.tw/2005/11/screen_teach_for_use/)
[本文](http://my.oschina.net/cshell/blog/135261)
