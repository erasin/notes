#清除DNS污染

## Windows ( XP, ME, 2000, 98)系统：

1. 开始——>运行——>输入cmd并回车
2. 在打开的命令行窗口，输入
	ipconfig /flushdns 
3. 完成！你的Windows DNS 缓存已经得到刷新。

## Windows Vista 或 Windows 7 系统：

1. 单击屏幕左下角的Microsoft Windows Vista或Windows 7 的开始标志 
2. 单击所有程序 
3. 单击附件
4. 右键单击命令提示符 
5. 选择以管理员身份运行
6. 在打开的命令行窗口，输入

	ipconfig /flushdns

你将会看到如下的确认信息： Windows IP 配置

已成功刷新 DNS 解析缓存。


## Linux系统：

刷新 DNS 缓存需要重起 nscd daemon： 

1. 要重起nscd daemon，在命令行窗口（terminal）输入

	/etc/rc.d/init.d/nscd restart

2. 命令执行完毕，你的DNS缓存就被刷新了。
如果是比较新的Linux版本，你可能需要使用下面的命令： 

	/etc/init.d/nscd restart


## Mac OS X苹果系统：
1. 在命令行窗口（terminal）输入

	lookupd -flushcache

2. 命令执行完毕，你的DNS缓存就得到了更新。
较新的苹果Mac OS X系统应该使用下面的命令： 

	type dscacheutil -flushcache
