# rsync 同步




rsync远程同步的两种模式：

rsync连接远程主机进行同步或备份时有两种途径：使用远程shell程序（如ssh或rsh）进行连接，或使用TCP直接连接rsync daemon。

我是使用了一下rsync的ssh方式连接进行同步的。

当源路径或目的路径的主机名后面包含一个冒号分隔符时，rsync使用远程shell传输；当源路径或目的路径的主机名后面包含两个冒号，或使用 `rsync://URL`时，rsync使用TCP直接连接rsync daemon。

特别的，如果只指定了源路径，而没有指定目的路径，rsync将会显示源路径中的文件列表，类似于使用命令ls -l。rsync把本地端看作client，把远程端当成server。


本人目前常用的rsync远程同步命令如下：

	rsync vmm-qa.smilejay.com:/home/ /share/home/ -avz –delete

**–delete8** 从目的目录中删除不必要的文件（在server端没有的文件）  
**-a** 相当于 `-rlptgoD`，所以这个 `-a` 是最常用的参数了。

当上面的命令在client端进行同步时，在server端，可以看到如下进程：

	[root@myserver ~]# ps -ef | grep rsync | grep -v grep

	root      7502  7500 66 21:36 ?        00:00:04 rsync –server –sender -vlogDtprz . /home/


命令

	rsync vgt.smilejay.com:/etc/ /share/temp/ -avz -F –exclude yum

**-F –exclude abc**  #同步时，排除abc这个文件或目录


关于rsync命令的一些选项意义如下：

	[root@linux ~]# rsync [-avrlptgoD] [-e ssh] [user@host:/dir] [/local/path]

参数：

* **-v** ：观察模式，可以列出更多的资讯；
* **-q** ：与 -v  相反，安静模式，输出的资讯比较少；
* **-r** ：递回复制！可以针对”目录”来处理！很重要！
* **-u** ：仅更新 (update)，不会覆盖目标的新档案；
* **-l** ：复制连结档的属性，而非连结的目标原始档案内容；
* **-p** ：复制时，连同属性 (permission) 也保存不变！
* **-g** ：保存原始档案的拥有群组；
* **-o** ：保存原始档案的拥有人；
* **-D** ：保存原始档案的装置属性 (device)
* **-t** ：保存原始档案的时间参数；
* **-I** ：忽略更新时间 (mtime) 的属性，档案比对上会比较快速；
* **-z** ：加上压缩的参数！
* **-e** ：使用的通道协定，例如使用 ssh 通道，则 -e ssh
* **-a** ：相当于 -rlptgoD ，所以这个 -a 是最常用的参数了！

更多说明请参考 man rsync！

rsync shell模式同步不需要密码的方法：

使用ssh + key_gen 来进行免密码同步。

如果还是需要密码则可能为 sshd 禁止了root的ssh远程登录，检查修改  `/etc/ssh/sshd.conf` ,找到

	PermitRootLogin yes

然后可以ssh自动登录了，才能实现rsync免密码同步。

	rsync -azv --delete -e ssh pi@ssh2:/pwd/ /pwd/


