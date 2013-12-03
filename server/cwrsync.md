# windows下使用同步工具cwRsync

下载软件 <https://www.itefix.no/i2/content/cwrsync-free-edition>


## SERVER

IP: 192.168.0.199

安装 server 端，比如安装到`D:\Program Files\ICW`，之后修改配置文件`rsyncd.conf`配置文件。

添加 

	uid = 0
    gid = 0
    port = 8888   					# 端口

修改 `[test]`

	# 映射路径
	[test]
	# 对应test的 d:\rsyncdir 文件夹
	path = /cygdrive/d/rsyncdir
	# 权限	
	read only = false 				
	transfer logging = yes
	# 授予用户
	auth users = username
	# 密钥文件也就是密码
	secrets file = pwd.conf

创建`pwd.conf`存储用户和密码

	username:password

修改`pwd.conf`文件权限其他用户不可读写


在安装服务端时候，已经创建了用户`SvcCWRSYNC`和服务`RsyncServer`在OS中，

对同步文件夹`d:\rsyncdir`进行权限修改为`SvcCWRSYNC`用户拥有全部的权限。

服务管理器中，`RsyncServer` 服务默认是手动开启，可以修改为自动开启。

## CLIENT

测试连通

	telnet 192.168.0.199:8888

若安装客户端到`D:\Program Files\cwRsync`后，修改OS环境变量**PATH**,添加`D:\Program Files\cwRsync\bin`

拷贝服务端的 `chown.exe` 到安装目录的 `bin`目录中去。


创建 脚本文件夹 `d:/bat` ，其中创建 脚本文件 `do.bat` 和 `pwd.conf`

`pwd.conf`文件用于存储用户的密码

	password

对`pwd.conf` 文件权限修改,打开CMD

	d:
	cd bat
	chmod -c 600 pwd.conf
	chown administrator pwd.conf

测试单向Client向Server传输，若是双向传输，则将下面的语句复制之后转换路径顺序即可。编辑`do.bat`文件。

	rsync -avzP --progress --password-file=/cygdrive/d/bat/pwd.conf  --delete  /cygdrive/d/clientdir rsync://username@192.168.0.199:8888/test

执行do.bat 即可同步。

查看效果！

设定OS定时任务，进行同步