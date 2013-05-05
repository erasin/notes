#screen的用法

简单来说，Screen是一个可以在多个进程之间多路复用一个物理终端的窗口管理器,这意味着你能够使用一个单一的终端窗口运行多终端的应用。Screen中有会话的概念，用户可以在一个screen会话中创建多个screen窗口，在每一个screen窗口中就像操作一个真实的telnet/SSH连接窗口那样。

## 安装screen  
Archlinux 下 pacman -S screen  

##  screen

screen

###Screen命令语法：

    screen [-AmRvx -ls -wipe][-d <作业名称>][-h <行数>][-r <作业名称>][-s ][-S <作业名称>]

Screen命令参数：

* **-A -[r|R]**          将所有的视窗都调整为目前终端机的大小。
* **-c filename**        用指定的filename文件替代screen的配置文件’.screenrc’.
* **-d [pid.tty.host]**  断开screen进程(使用该命令时，screen的状态一定要是Attached，也就是说有用户连在screen里)。一般进程的名字是以pid.tty.host这种形式表示(用screen -list命令可以看出状态)。
* **-D [pid.tty.host]**  与-d命令实现一样的功能，区别就是如果执行成功，会踢掉原来在screen里的用户并让他logout。
* **-h <行数>** 　       指定视窗的缓冲区行数。

* **-ls或–list**         显示目前所有的screen作业。
* **-m**                 即使目前已在作业中的screen作业，仍强制建立新的screen作业。
* **-p number or name**  预先选择一个窗口。
* **-r [pid.tty.host]**  恢复离线的screen进程，如果有多个断开的进程，需要指定[pid.tty.host]
* **-R**                 先试图恢复离线的作业。若找不到离线的作业，即建立新的screen作业。
* **-s shell**           指定建立新视窗时，所要执行的shell。
* **-S <作业名称>**      指定screen作业的名称。(用来替代[pid.tty.host]的命名方式,可以简化操作).
* **-v**                 显示版本信息。
* **-wipe**              检查目前所有的screen作业，并删除已经无法使用的screen作业。
* **-x**                 恢复之前离线的screen作业。

Screen命令的常规用法:
* **screen -d -r**:     连接一个screen进程，如果该进程是attached，就先踢掉远端用户再连接。
* **screen -D -r**:     连接一个screen进程，如果该进程是attached，就先踢掉远端用户并让他logout再连接
* **screen -ls/-list**: 显示存在的screen进程，常用命令
* **screen -m**:        如果在一个Screen进程里，用快捷键crtl+a c或者直接打screen可以创建一个新窗口,screen -m可以新建一个screen进程。
* **screen -dm**:       新建一个screen，并默认是detached模式，也就是建好之后不会连上去。
* **screen -p number or name**:预先选择一个窗口。

## 快捷键用法
* **Ctrl+a** 然后按`c`     建立一个新的screen 会话
* **Ctrl+a** 然后按`n`     跳转到下一个screen 会话
* **Ctrl+a** 然后按`p`     返回到上一个screen 会话
* **Ctrl+a** 然后按`d`     将当前的screen 会话放在背景执行
* **Ctrl+a** 然后按(大写)`S`    分离一个screen 会话出来，分离后用Ctrl+a 然后按tab键 在分离出来的各screen间跳转。
* **screen -ls**    列出当前所有的screen会话
* **screen -r 进程号**    之前Ctrl+a 然后按d  放在背景执行的会话 呼叫回来。

Screen 命令 (在 screen 中) 所有命令都以Ctrl-a 开始。

* Ctrl-a ? 各功能的帮助摘要
* Ctrl-a c 创建一个新的 window (终端)
* Ctrl-a Ctrl-n 和 Ctrl-a Ctrl-p 切换到下一个或前一个 window
* Ctrl-a Ctrl-N N 为 0 到 9 的数字，用来切换到相对应的 window
* Ctrl-a " 获取所有正在运行的 window 的可导航的列表
* Ctrl-a a 清楚错误的 Ctrl-a
* Ctrl-a Ctrl-d 断开所有会话，会话中所有任务运行于后台
* Ctrl-a x 用密码锁柱 screen 终端

当程序内部运行终端关闭并且你登出该终端时，该 screen 会话就会被终止。


ssh中如果发生了突然断线 那么你重新登陆后 screen -ls 会发现 有screen的状态是处于(Attached)状态 此刻我们使用  screen -d  将他强行放到背景，然后再用screen -r  进程号将他呼叫回来。
如果 screen -ls 看到有死亡的会话  可以用screen -wipe 进程号  将他杀掉。

eg: #screen -list可以看到正在运行的screen实例

    #screen -list
    There is a screen on:
    80338.ttyp3.chh (Detached)
    1 Socket in /tmp/screens/S-chh.
    screen -r 80338 # 返回
    exit #退出

Screen 提供了两个主要功能：

* 在一个终端内运行多个终端会话(terminal session)。  
* 一个已启动的程序与运行它的真实终端分离的，因此可运行于后台。真实的终端可以被关闭，还可以在稍后再重新接上(reattached)。

简短实例

开启 screen：

    # screen

在screen 会话中，我们可以开启一个长时间运行的程序(如top)。Detach 这个终端，之后可以从其他机器reattach 这个相同的终端(比如通过 ssh)。

    # top

现在用Ctrl-a Ctrl-d 来 detach。Reattach 终端：

    # screen -r

或更好的：

    # screen -R -D
现在attach 到这里。具体意思是：先试图恢复离线的screen 会话。若找不到离线的screen 会话，即建立新的 screen 会话给用户。


