#e4rat 加速器

e4rat ("Ext4 - Reducing Access Times") 是一款将 Ext4 文件系统上文件块优化排序以加快系统和应用程序启动速度的程序。

对于一般的机械硬盘来讲，将需要顺序载入的文件在硬盘上按照顺序存储可以大大减少寻道时间和旋转延迟，e4rat 利用 2.6.31 内核的引入的 EXT4*IOC*MOVE_EXT 在线碎片整理功能，来完成并行化载入。顺序读取和高缓存命中的结合使用将可以将**系统启动时间缩短至原先的三分之一**。

**系统需求：** 

*   要求是**原生的 Ext4 文件系统**，若是从老版本 Ext 升级上来的不可(仅限 0.20 版本)。
*   需要是传统的机械磁盘，对于 SSD 固态硬盘无效。
*   内核必须启用 FS*IOC*FIEMAP（目前默认 Debian 内核未启用）。
*   需要停用功能类似的 readahead： 
    *   Ubuntu: `dpkg --purge ureadahead ubuntu-minimal`
    *   Fedora 15+: `systemctl disable systemd-readahead-collect.service systemd-readahead-replay.service`
*   和 auditd 服务冲突。



**简洁使用步骤：** 

1.  如果使用的不是 `/sbin/init` 的初始化程序的话（比如 upstart 和 systemd），请配置 `/etc/e4rat.conf`;
2.  设置 `init=/sbin/e4rat-collect` 初始化进程，收集系统启动情况，写入到 `/var/lib/e4rat/startup.log` 文件中（仅需一次）；
3.  依据 `/var/lib/e4rat/startup.log` 中的内容，在 `single` 用户模式下重新排序文件位置 `e4rat-realloc  /var/lib/e4rat/startup.log`；
4.  将 `init=/sbin/e4rat-preload` 永久性的加入 GRUB 的内核引导行中。

具体使用步骤请参照
[官方 Wiki](http://e4rat.sourceforge.net/wiki/index.php/Main_Pag)和[Arch Wiki](https://wiki.archlinux.org/index.php/E4rat)。

[e4rat 官方主页及下载](http://e4rat.sourceforge.net/)
