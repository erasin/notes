#清除DNS污染

__windows 下IE浏览器dns污染清空方法__

使用命令 查看dns缓存列表

> ipconfig /displaydns    

使用命令 清空dns缓存记录

> ipconfig /flushdns 

__linux 下清空DNS缓存记录__

> sudo /etc/init.d/nscd restart

Archlinux下 
> sudo /etc/rc.d/nscd restart
