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
	
	[test] 							# 映射文件
	path = /cygdrive/d/rsyncdir   	# 对应test的 d:\rsyncdir 文件夹
	read only = false 				# 权限
	transfer logging = yes 			
	auth users = username 			# 创建用户
	secrets file = pwd.conf 		# 密钥也就是密码

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

测试单向Client向Server传输，若是双向传输，则将下面的语句复制之后转换路径顺序即可。创建 do.bat 文件，编辑文件。

	set curdir=%~dp0
	cd /d %curdir%

	rsync -avzP --progress --password-file=bin/pwd.conf  --delete  /cygdrive/d/cwRsync/a/ rsync://username@192.168.0.199:8888/test

创建 `pwd.conf`(这个文件随着你的do.bat文件位置改变)。放置密码（对应服务器中对应用户密码）

修改pwd.conf文件权限其他用户不可读写

	passwd

执行 cmd ，执行do.bat 即可同步。

设定OS定时任务，进行同步