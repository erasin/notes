# vsftpd

## vsftpd 服务

检测软件是否安装

	$ rpm -qa |grep vsftpd

vsftpd 服务

	# service vsftpd status 
	# service vsftpd start
	# service vsftpd stop
	# service vsftpd restart

开机启动

	# chkconfig vsftpd on


## vsftpd的配置

`cnteos` 下配置文件 `/etc/vsftpd/`

	ftpusers    该文件用来指定那些用户不能访问ftp服务器。
	user_list   指定的用户是否可以访问ftp服务器由vsftpd.conf文件中的userlist_deny的取值来决定。
	vsftpd.conf vsftpd的主配置文件

其他的配置文件

	/etc/hosts.allow
	/etc/hosts.deny

### vsftpd.conf

	vim /etc/vsftpd/vsftpd.conf

普通情况下，使用默认配置即可。

	local_enable=YES
	deny_email_enable=YES
	banned_email_file=/etc/vsftpd/banned_emails

	write_enable=YES
	local_umask=022
	chroot_local_user=YES
	chroot_list_enable=YES
	chroot_list_file=/etc/vsftpd/chroot_list
	nopriv_user=ftpsecure 

	idle_session_timeout=600
	data_connection_timeout=120

	dirmessage_enable=YES
	xferlog_enable=YES 
	xferlog_file=/var/log/xferlog 

	listen_port=2121 

参数                                                            | 说明
----------------------------------------------------------------|---------------------------------
listen_address=ip address                                       | 指定侦听IP
listen_port=port_value                                          | 指定侦听端口，默认21
anonymous_enable=YES                                            | 是否允许使用匿名帐户
local_enable=YES                                                | 是否允许本地用户登录
nopriv_user=ftp                                                 | 指定vsftpd服务的运行帐户，不指定时使用ftp
write_enable=YES                                                | 是否允许写入
anon_upload_enable=YES                                          | 匿名用户是否可上传文件
anon_mkdir_write_enable=YES                                     | 匿名用户是否建立目录
dirmessage_enable=YES                                           | 进入每个目录是显示欢迎信息，在每个目录下建立.message文件在里面写欢迎信息
xferlog_enable=YES                                              | 上传/下载文件时记录日志
connect_from_port_20=YES                                        | 是否使用20端口传输数据(是否使用主动模式)
chown_uploads=YES、chown_username=whoever                        | 修改匿名用户上传文件的拥有者
xferlog_file=/var/log/vsftpd.log                                | 日志文件
xferlog_std_format=YES                                          | 使用标准文件日志
idle_session_timeout=600                                        | 会话超时，客户端连接到ftp但未操作
data_connection_timeout=120                                     | 数据传输超时
async_abor_enable=YES                                           | 是否允许客户端使用sync等命令
ascii_upload_enable=YES、ascii_download_enable=YES               | 是否允许上传/下载二进制文件
chroot_local_user=YES                                           | 限制所有的本地用户在自家目录
chroot_list_enable=YES、chroot_list_file=/etc/vsftpd/chroot_list | 指定不能离开家目录的用户，将用户名一个一行写在/etc/vsftpd/chroot_list文件里，使用此方法时必须chroot_local_user=NO
ls_recurse_enable=YES                                           | 是否允许使用ls -R等命令
listen=YES                                                      | 开启ipv4监听
listen_ipv6=YES                                                 | 开启ipv6监听
pam_service_name=vsftpd                                         | 使用pam模块控制，vsftpd文件在/etc/pam.d目录下
userlist_enable=YES                                             | 此选项被激活后，vsftpd将读取userlist_file参数所指定的文件中的用户列表。当列表中的用户登录FTP服务器时，该用户在提示输入密码之前就被禁止了。即该用户名输入后，vsftpd查到该用户名在列表中，vsftpd就直接禁止掉该用户，不会再进行询问密码等后续步聚
userlist_deny=YES                                               | 决定禁止还是只允许由userlist_file指定文件中的用户登录FTP服务器。此选项在userlist_enable 选项启动后才生效。YES，默认值，禁止文件中的用户登录，同时也不向这些用户发出输入密码的提示。NO，只允许在文件中的用户登录FTP服务器
tcp_wrappers=YES                                                | 是否允许tcp_wrappers管理
local_root=/home/ftp                                            | 所有用户的根目录，，对匿名用户无效
anon_max_rate                                                   | 匿名用户的最大传输速度，单位是Byts/s
local_max_rate                                                  | 本地用户的最大传输速度，单位是Byts/s
download_enable= YES                                            | 是否允许下载


## 限制

IP限制：`/etc/hosts.allow`

	vsftpd:192.168.5.128:DENY 

设置该IP地址不可以访问ftp服务


访问时间限制：`/etc/vsftpd/vsftpd.xinetd`

修改 disable = no

	access_time = hour:min-hour:min (添加配置访问的时间限制(注：与vsftpd.conf中listen=NO相对应)

例: access_time = 8:30-11:30 17:30-21:30 表示只有这两个时间段可以访问ftp


## 帐号

添加帐号： _注意修改_ 默认目录 `-d /pwd/dir` ,用户组 `-G http`

~~~{.bash}
	useradd -s /sbin/nologin -d /pwd/dir -G groupname username  
	passwd username
~~~

修改文件权限 `chmod` `chown`

### 虚拟帐号


## 最为简单的使用

###  系统
**Centos/RedHat**

查看系统版本

	find /etc/ -iname "*-release"

找到后`cat`或直接使用 `cat /etc/*-release` 来查看系统版本

检查 开启 ，随机启动

	service vsftpd status
	service vsftpd start
	chkconfig vsftpd on

### 安装

如果没有安装`vsftpd`


Centos / Radhat / fedora 使用 yum 包

    yum install vsftpd

GUN/Linux Debian / ubuntu 使用牛力安装

    apt-get install vsftpd

### 配置

在默认的基础上添加禁止用户访问其他的目录

	chroot_list_enable=YES
	chroot_list_file=/etc/vsftpd/chroot_list

去创建 `/etc/vsftpd/chroot_list`,并将要使用的ftp帐号添加其中。

    touch /etc/vsftpd/chroot_list
    echo username >> /etc/vsftpd/chroot_list
    service vsftpd restart

### 帐号

创建帐号,不给予系统登录权限，设定该用户指向的目录地址, 给予一个组权限

	useradd -s /sbin/nologin -d /pwd/dir -G groupname username  

一般在服务器上，站点文件都在nginx或apache下， 

* nginx 的默认组和帐号 有 `www:www` 或 `http:http`,`nginx:nginx` 具体查看`nginx.conf` 配置文件 
* apache 默认组和帐号 `apache:apache` , `web:web`, `http:http`

查看系统上所有的帐号

	cat /etc/passwd

查看系统上组
 
	cat /etc/group


设定密码

	passwd username

### 权限

给予改用户的文件组权限,以`http:http`为例

	chown -R http:http /pwd/dir
	chmod -R 775 /pwd/dir

文件权限为 `664` .


## 关于 selinux

**CentOS** 和 **RedHat/Fedora** 有开启 selinux.


如出现 `500 OOPS`... 错误

    setsebool allow_ftpd_full_access 1 

或

    setsebool allow_ftpd_full_access on

查看可用

    getsebool -a|grep ftp



- <http://www.ha97.com/4113.html>
- <http://wiki.ubuntu.org.cn/Vsftpd%E8%99%9A%E6%8B%9F%E7%94%A8%E6%88%B7%E8%AE%BE%E7%BD%AE>
- <http://blog.csdn.net/tianlesoftware/article/details/6151317>
