#lftp linux下终端ftp工具#
**tags:** ftp,lftp,linux,FlashFXP,FileZilla  
**info:** lftp 是一个功能强大的下载工具，它支持访问文件的协议: ftp, ftps, http, https, hftp, fish.(其中ftps和https需要在编译的时候包含openssl库)。llftp的界面非常想一 个shell: 有命令补全，历史记录，允许多个后台任务执行等功能，使用起来非常方便。 它还有书签、排队、镜像、断点续传、多进程下载等功能。


##登录ftp服务器

>* lftp ftp://user:password@site:port
* lftp user:password@site:port
* lftp site -p port -u user,password
* lftp site:port -u user,password

上面的几种方式都能正常工作，不过密码都是明文，密码会被记录日志。

>lftp user@site:port

系统会提示输入password，密码就回显为******了

另外lftp 自带了 __[bookmark][1]__ （书签）功能，可以储存常用ftp列表。


## 命令行语法
要看lftp的命令行语法，只要在__shell__中输入 `lftp --help`

>     lftp [OPTS] <site>
    `lftp' 是在 rc 文件执行后 lftp 执行的第一个命令
     -f <file>           执行文件中的命令后退出
     -c <cmd>            执行命令后退出
     --help              显示帮助信息后退出
     --version           显示 lftp 版本后退出
     其他的选项同 `open' 命令
     -e <cmd>            在选择后执行命令
     -u <user>[,<pass>]  使用指定的用户名/口令进行验证
     -p <port>           连接指定的端口
     <site>              主机名, URL 或书签的名字
 
如果在命令行中输入的站点名称，lftp将直接登录站点，比如

>     [root@localhost /srv/http]$ lftp ftp://192.168.1.253:2121/incoming/  
     ... ... (此处略去站点登录信息)  
    cd 成功，当前目录=/incoming  
    lftp 192.168.1.253:/incoming>

如果在命令行不输入站点名称，则必须在进入到lftp界面后用open命令打开

>     [root@localhost /srv/http]$ lftp   
    lftp :~> open ftp://192.168.1.253:2121/incoming  
    cd 成功，当前目录=/incoming  
    lftp 192.168.1.253:/incoming>


##常用命令
__下载单个文件和一组文件__，断点续传用-c参数
>     lftp 192.168.1.253:/> get -c ls-lR.txt  
    lftp 192.168.1.253:/> mget *.txt
    
__镜像__(反镜像即上传)一个目录，可以用多个线程并行镜像一个目录(--parallel=N)
>     lftp 192.168.1.253:/> mirror incoming local_name  
    lftp 192.168.1.253:/> mirror -R local_name  
    lftp 192.168.1.253:/> mirror --parallel=3 incoming local_name
    
__多线程下载__，类似网络蚂蚁的功能;缺省是5个线程
>    lftp 192.168.1.253:/> pget -n 4 ls-lR.txt
    
__后台任务管理__  
缺省情况下，按` Ctrl+z`,正在执行的任务将转为后台执行，也可以 在命令行末尾加&符号使任务在后台执行。用`jobs`命令可以查看所 有的后台进程。用`queue`命令可以排队新的任务。如果退出lftp时还有任务在后台执行，lftp将转为后台执行。

__其它用法__
lftp支持类似bash的管道操作，例如用下面的命令可以将ftp服务 器上的特定目录下(也可以是整个站点)所有文件的大小存到本地的 文件ls.txt中

>    lftp 192.168.1.253:/> du incoming > ls.txt 

使用 __man函数__ 或者 `--help` 都可以看到帮助文档。

* __ls__			`#`显示远端文件列表(!ls 显示本地文件列表)。
* __cd__ 		`#`切换远端目录(lcd 切换本地目录)。
* __get__ 		`#`下载远端文件。
* __mget__ 		`#`下载远端文件(可以用通配符也就是 `*`)。
* __pget__ 		`#`使用多个线程来下载远端文件, 预设为五个。
* __mirror__	`#`下载/上传(mirror -R)/同步 整个目录。
* __put__ 		`#`上传文件。
* __mput__ 		`#`上传多个文件(支持通配符)。
* __mv__ 		`#`移动远端文件(远端文件改名)。
* __rm__ 		`#`删除远端文件。 参数-r,递归删除文件夹
* __mrm__ 		`#`删除多个远端文件(支持通配符)。
* __mkdir__ 	`#`建立远端目录。
* __rmdir__ 	`#`删除远端目录。
* __pwd__ 		`#`显示目前远端所在目录(lpwd 显示本地目录)。
* __du__ 		`#`计算远端目录的大小
* __set__ net:limit-rate 10000,10000  	`#`限制上传下载各为10KB/s
* __set__ ftp:charset gbk 					`#`设置远程ftp site用gbk编码
* __!__ 			`#`执行本地 shell的命令(由于lftp 没有 lls, 故可用 !ls 来替代)
* __lcd__ 		`#`切换本地目录
* __lpwd__ 		`#`显示本地目录
* __alias__ 	`#`定义别名
* __bookmark__ `#`设定书签。
* __exit__ 		`#`退出ftp

##快捷书签 lftp bookmark
ftp中的bookmark命令，是将配置写到~/.lftp/bookmarks文件中；
例如添加一行：
>echo 'uftp ftp://user:passwd@ftp.ubuntu.org.cn' >> ~/.lftp/bookmarks

lftp的bookmarks文件书写格式为：
>别名<空格>ftp://user:passwd@site:port

以後要登陆ubuntu-cn的ftp，只要执行：
>lftp uftp

## 相关文件
* __/etc/lftp.conf__  
    全局配置文件，实际位置依赖系统配置文件目录，可能在/etc，也可能在/usr/local/etc
* __~/.lftp/rc, ~/.lftprc__  
    用户配置文件，将在/etc/lftp.conf之后执行，所以这里面的设置会覆盖/etc/lftp.conf中的设置。  
    lftp 缺省不会显示 ftp 服务器的欢迎信息和错误信息，这在很多时候不方便，因为你有可能想知道这个服务器到底是因为没开机连不上，还是连接数已满。如果是这样，你可以在__~/.lftprc__ 里写入一行 `[debug 3][2]` 就可以看到出错信息了。   
    更多的配置选项请查man手册或在lftp界面内用命令 set -a 获得。
* __~/.lftp/log__  
    当lftp转为后台非挂起模式执行时，输出将重定向到这里
* __~/.lftp/bookmarks__  
    这是lftp存储书签的地方，可以lftp查看bookmark命令
* __~/.lftp/cwd_history__  
    这个文件用来存储访问过的站点的工作目录


## 解决乱码 配置 ~/.lftprc

在用lftp访问国内一些ftp服务器时，往往看到的中文是乱码

这是由于服务器和本地编码不一致造成的。我们只要在主目录下新建一个文件~/.lftprc或者~/.lftp/rc

并在其中加入以下内容：

>debug 3  
set ftp:charset GBK  
set file:charset UTF-8  
`#`set ftp:passtive-mode no  
`#`alias utf8 " set ftp:charset UTF-8"  
`#`alias gbk " set ftp:charset GBK"  


##其它客户端
* __kftpgrabber__ KDE下ftp客户端，支持编码选择。对中文支持较好
* __gftp__ gnome下ftp客户端，目前对中文支持尚不太好，受抱怨颇多。
* __fireftp__ firefox的ftp客户端插件，新版对中文支持较好。
* __FileZilla__ 对中文支持较好
* __CrossFTP__ 基于Java的稳定ftp客户端和同步工具。优良的中文/Unicode支持。

<!-- links dict -->
[1]: #4 "bookmark书签记录"
[2]: #6 "乱码解决debug"
