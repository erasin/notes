# net tool

## hostname

* hostname 没有选项，显示主机名字
* hostname –d 显示机器所属域名
* hostname –f 显示完整的主机名和域名
* hostname –i 显示当前机器的ip地址

## ping

ping 将数据包发向用户指定地址。当包被接收。目标机器发送返回数据包. ping 主要有两个作用

* 用来确认网络连接是畅通的。
* 用来查看连接的速度信息。

如果你 ping www.yahoo.com 它将返回它的ip地址 。你可以通过 ctrl+C 来停止命令。

## ifconfig

查看用户网络配置。它显示当前网络设备配置。对于需要接收或者发送数据错误查找，这个工具极为好用。

## iwconfig

iwconfig 工具与 ifconfig 和ethtool类似。是用于无线网卡的 . 你可以用他查看设置基本的Wi-Fi 网络信息,例如 SSID, channel和encryption.还有其他很多配置你也可以查看和修改，, 包括 接收灵敏度, RTS/CTS, 发送数据包的分片大小,以及无线网卡的重传机制

## nslookup

nslookup 这个命令在 有ip地址时，可以用这个命令来显示主机名，可以找到给定域名的所有ip地址。而你必须连接到互联网才能使用这个命令

例子. nslookup blogger.com

你也可以使用 nslookup 从ip获得主机名或从主机名获得ip。

## traceroute

一个方便的工具。可用来查看数据包在提交到远程系统或者网站时候所经过的路由器的IP地址、跳数和响应时间。同样你必须链接到互联网才能使用这个命令

## finger

查看用户信息。显示用户的登录名字、真实名字以及登录终端的名字和登录权限。这是unix一个很老的命令，现在已很少使用了

## telnet

通过telnet协议连接目标主机，如果telnet连接可以在任一端口上完成即代表着两台主机间的连接良好。
telnet hostname port - 使用指定的端口telnet主机名。这通常用来测试主机是否在线或者网络是否正常。

## ethtool

ethtool允许你查看和更改网卡的许多设置（不包括Wi-Fi网卡）。你可以管理许多高级设置，包括tx/rx、校验及网络唤醒功能。下面是一些你可能感兴趣的基本命令：
显示一个特定网卡的驱动信息，检查软件兼容性时尤其有用。

* ethtool -i 启动一个适配器的指定行为，比如让适配器的LED灯闪烁，以帮助你在多个适配器或接口中标识接口名称：
* ethtool -p 显示网络统计信息：
* ethtool -s 设置适配器的连接速度，单位是Mbps：

ethtool speed <10|100|1000>

## netstat

发现主机连接最有用最通用的Linux命令。你可以使用"netstat -g"查询该主机订阅的所有多播组（网络）

* netstat -nap | grep port 将会显示使用该端口的应用程序的进程id
* netstat -a  or netstat –all 将会显示包括TCP和UDP的所有连接  
* netstat --tcp  or netstat –t 将会显示TCP连接
* netstat --udp or netstat –u 将会显示UDP连接
* netstat -g 将会显示该主机订阅的所有多播网络。
