# ssh 的文件传输 scp 和 sftp
**title:**  ssh 的文件传输 scp 和 sftp详解  
**tags:** ssh,scp,sftp,linux

scp 命令是 SSH 中最方便有用的命令了，试想，在两台服务器之间直接传送文件，仅仅用 scp 一个命令就完全解决了。 你可以在一台服务器上以 root 身份运行 #scp servername:/home/ftp/pub/file1 . 这样就把另一台服务器上的文件 /home/ftp/pub/file1 直接传到本机器的当前目录下，当然你也可以用 #scp /tmp/file2 servername:/boot 把本机上的文件 /tmp/file2 送到另一台机器的 /boot 目录下。而且整个传送过程仍然是用 SSH 加密的。


## scp 命令

scp 就是 secure copy, 是用来进行远程文件拷贝的 . 数据传输使用 ssh1, 并且和 ssh1 使用相同的认证方式 , 提供相同的安全保证。与 rcp 不同的是 ,scp 会要求你输入密码，如果需要的话。

命令基本格式：

    scp [可选参数] file_source file_target
    scp 本地用户名 @IP 地址 : 文件名 1 远程用户名 @IP 地址 : 文件名 2

常用的参数:

* **-v** 和大多数 linux 命令中的 -v 意思一样 , 用来显示进度 . 可以用来查看连接 , 认证 , 或是配置错误 .
* **-C** 使能压缩选项 .
* **-P** 选择端口 . 注意 -p 已经被 rcp 使用 .
* **-4** 强行使用 IPV4 地址 .
* **-6** 强行使用 IPV6 地址 .
* **-r** Recursively copy entire directories.

###从 本地 复制到 远程

**复制文件：**   
命令格式：

    scp local_file remote_username@remote_ip:remote_folder
    scp local_file remote_username@remote_ip:remote_file
    scp local_file remote_ip:remote_folder
    scp local_file remote_ip:remote_file

第1,2个指定了用户名，命令执行后需要再输入密码，第1个仅指定了远程的目录，文件名字不变，第2个指定了文件名；  
第3,4个没有指定用户名，命令执行后需要输入用户名和密码，第3个仅指定了远程的目录，文件名字不变，第4个指定了文件名；

例程：

    scp /home/space/music/1.mp3 root@www.cumt.edu.cn:/home/root/others/music
    scp /home/space/music/1.mp3 root@www.cumt.edu.cn:/home/root/others/music/002.mp3
    scp /home/space/music/1.mp3 www.cumt.edu.cn:/home/root/others/music
    scp /home/space/music/1.mp3 www.cumt.edu.cn:/home/root/others/music/002.mp3

**复制目录：** 
命令格式：

    scp -r local_folder remote_username@remote_ip:remote_folder
    scp -r local_folder remote_ip:remote_folder

第1个指定了用户名，命令执行后需要再输入密码；   
第2个没有指定用户名，命令执行后需要输入用户名和密码；

例程：

    scp -r /home/space/music/ root@www.cumt.edu.cn:/home/root/others/
    scp -r /home/space/music/ www.cumt.edu.cn:/home/root/others/

上面 命令 将 本地 music 目录 复制 到 远程 others 目录下，即复制后有 远程 有 ../others/music/ 目录

    scp -r /home/space/music/.* www.cumt.edu.cn:/home/root/others/musc/

拷贝目录,**-r**是将目录下的目录递归拷贝。`.*` 是将隐藏文件也拷贝过去。需要先在远端创建好相应的目录。

###从 远程 复制到 本地

从 远程 复制到 本地，只要将 从 本地 复制到 远程 的命令 的 后2个参数 调换顺序 即可；

例如：

    scp root@www.cumt.edu.cn:/home/root/others/music /home/space/music/i.mp3
    scp -r www.cumt.edu.cn:/home/root/others/ /home/space/music/

scp的优点是使用简单，缺点是无法列出远端目录和改变目录。

##sftp {#sftp} 

sftp：

    sftp -o port=60066 user@serverip:/home/user/

其中-o port选项指定非缺省的ssh端口。
不过在使用上基本和lftp 保持

|命令                              |解释
|----------------------------------|----------------------------------------
|cd path                           |改变服务器路径
|lcd path                          |改变本地路径
|ls [-1afhlnrSt] [path]            |列出服务器文件列表
|lls [ls-options [path]]           |列出本地文件列表
|get [-Ppr] remote [local]         |下载文件
|put [-Ppr] local [remote]         |上传文件
|mkdir path                        |创建文件夹
|lmkdir path                       |创建本地文件夹
|rename oldpath newpath            |Rename 
|ln [-s] oldpath newpath           |Link remote file (-s for symlink)
|rm path                           |Delete remote file
|rmdir path                        |Remove remote di
|chgrp grp path                    |相当于本地的mv
|chmod mode path                   |chmod 修改权限
|chown own path                    |chown 修改用户与用户组
|df [-hi] [path]                   |df 查看占用空间
|pwd                               |服务器当前路径
|lpwd                              |本地当前路径
|exit                              |退出

通配符对于ls,lls,get和put是支持的。格式在sshregex手册中有描述。从sftp使用加密技术以来，一直有一个障碍：连接速度慢（以我的经验有2-3倍），但是这一点对于非常好的安全性来讲只能放在一边了。在一个测试中，在我们局域网上的Sniffer可以在一个小时中捉住ftp连接上的4个password。sftp的使用可以从网络上传送文件并且除去这些安全问题。 
