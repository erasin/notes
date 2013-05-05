# Debian init 开机启动管理
**title:** Debian update-rc.d 命令开机启动管理
**tags:** Debian,update-rc.d,raspberrypi,raspbian,树莓派  

**OS:** Raspbian “wheezy”

Debian系统中`update-rc.d`命令，是用来更新系统启动项的脚本。这些脚本的链接位于`/etc/rcN.d/`目录，对应脚本位于`/etc/init.d/`目录

## Linux系统主要启动步骤

读取 MBR 的信息，启动 Boot Manager。

加载系统内核，启动 init 进程， init 进程是 Linux 的根进程，所有的系统进程都是它的子进程。

init 进程读取 /etc/inittab 文件中的信息，并进入预设的运行级别。通常情况下 /etc/rcS.d/ 目录下的启动脚本首先被执行，然后是/etc/rcN.d/[^1] 目录。

根据 /etc/rcN.d/ 文件夹中对应的脚本启动 Xwindow 服务器 xorg，Xwindow 为 Linux 下的图形用户界面系统。

启动登录管理器，等待用户登录。

## 运行级别

Debian中的运行级别

* **0** - 停机（千万不要把initdefault设置为0 ）
* **1** - 单用户模式(单用户模式，只允许root用户对系统进行维护。)
* **2** - 多用户，但是没有NFS
* **3** - 完全多用户模式(字符界面)
* **4** - 基本不用 
* **5** - X11(图形界面)
* **6** - 重新启动（千万不要把initdefault设置为6 ） 

可修通过修改 `/etc/inittab` 来修改启动级别

切换运行级别

    init [0123456Ss]

例如：

    init 0 命令关机; 
    init 6 命令重新启动

启动项管理工具

sysv-rc-conf

## update-rc.d命令详解

从所有的运行级别中删除指定启动项

    update-rc.d -f remove

按指定顺序、在指定运行级别中启动或关闭

    update-rc.d start|stop

Insert links using the defaults:

    update-rc.d foobar defaults

Equivalent command using explicit argument sets:

    update-rc.d foobar start 20 2 3 4 5 . stop 20 0 1 6 .

More typical command using explicit argument sets:

    update-rc.d foobar start 30 2 3 4 5 . stop 70 0 1 6 .

Remove  all  links  for  a  script  (assuming  foobar  has been deletedalready):

    update-rc.d foobar remove

Example of disabling a service:

    update-rc.d -f foobar remove
    update-rc.d foobar stop 20 2 3 4 5 .

Example of a command for installing a  system  initialization-and-shut‐down script:

    update-rc.d foobar start 45 S . start 31 0 6 .

Example of a command for disabling a system initialization-and-shutdown script:

    update-rc.d -f foobar remove
    update-rc.d foobar stop 45 S .

实例：

    update-rc.d apachectl start 20 2 3 4 5 . stop 20 0 1 6 .

表示在2、3、4、5这五个运行级别中，由小到大，第20个开始运行apachectl;在 0 1 6这3个运行级别中，第20个关闭nginx。这是合并起来的写法，注意它有2个点号，效果等于下面方法：

    update-rc.d nginx defaults

A启动后B才能启动，B关闭后A才关闭

    update-rc.d A defaults 80 20
    update-rc.d B defaults 90 10

启动和关闭顺序为90，级别默认

    update-rc.d defaults 90

[^1]:
    这里 N 是指表示级别的数字  
    用户启动项目一般在 /etc/rc2.d 中。  
    可以使用 ls |grep item来查找item是否在启动项中。

