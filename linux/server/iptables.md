# iptables

## ip tcp

```
# 开启hi
service iptables start 
## 新版本 centos 7
systemctl status iptables.service
```





```
iptables -I INPUT -p tcp --dport 80 -j ACCEPT
```

- - - - - - - - 


但是我们在配置服务器时候往往是只打开某个端口，其他的端口全部
关闭来提高我们服务器的安全性。
下面我就用端口22来举个例子如何实现，只允许端口22的访问 其他端口全部
都不能访问的， 端口22就是我们通过ssh来进行远程访问Linux默认端口。
看我现在服务器的情况：
通过命令 netstat -tnl 可以查看当前服务器打开了哪些端口

```
[root@localhost ~]# netstat -tnl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address Foreign Address State
tcp 0 0 0.0.0.0:3306 0.0.0.0:* LISTEN
tcp 0 0 0.0.0.0:111 0.0.0.0:* LISTEN
tcp 0 0 0.0.0.0:10000 0.0.0.0:* LISTEN
tcp 0 0 127.0.0.1:631 0.0.0.0:* LISTEN
tcp 0 0 127.0.0.1:25 0.0.0.0:* LISTEN
tcp 0 0 :::80 :::* LISTEN
tcp 0 0 :::22 :::* LISTEN
tcp 0 0 :::443 :::* LISTEN
```

我这里的端口挺多的，比较常用80(web服务) 22(ssh) 3306(mysql数据库)
再看看我的防火墙设置如何，通过命令:iptables -L -n 可以查看
看到我的防火墙 INPUT FORWARD OUTPUT 三个部分全部是 ACCEPT 的，也就是没有做任何限制
下面我通过 putty.exe 这个软件来链接我的服务器，这个软件通过端口22来访问我们的Linux.
成功链接上去。这是理所当然成功的，以为Linux打开了22端口，而且Linux防火墙iptables 没有
做任何限制。
下面我们关闭所有的端口

```
]# iptables -P INPUT DROP
]# iptables -P FORWARD DROP
]# iptables -P OUTPUT DROP
```

再查看一下 iptables -L -n

好成功关闭了所有端口。
再使用 putty.exe 链接上去，验证一下是否真的关闭了。
一直等待状态。好链接失败 Connection timed out 超时，链接不上。
说明已经关闭了端口。
下面我只打开22端口，看我是如何操作的，就是下面2个语句
]# iptables -A INPUT -p tcp --dport 22 -j ACCEPT
]# iptables -A OUTPUT -p tcp --sport 22 -j ACCEPT
再查看下 iptables -L -n 是否添加上去, 看到添加了
Chain INPUT (policy DROP)
target prot opt source destination
ACCEPT tcp -- 0.0.0.0/0 0.0.0.0/0 tcp dpt:22
Chain FORWARD (policy DROP)
target prot opt source destination
Chain OUTPUT (policy DROP)
target prot opt source destination
ACCEPT tcp -- 0.0.0.0/0 0.0.0.0/0 tcp spt:22
现在Linux服务器只打开了22端口，用putty.exe测试一下是否可以链接上去。
可以链接上去了，说明没有问题。
最后别忘记了保存 对防火墙的设置
通过命令：service iptables save 进行保存
]# iptables -A INPUT -p tcp --dport 22 -j ACCEPT
]# iptables -A OUTPUT -p tcp --sport 22 -j ACCEPT
针对这2条命令进行一些讲解吧
-A 参数就看成是添加一条 INPUT 的规则
-p 指定是什么协议 我们常用的tcp 协议，当然也有udp 例如53端口的DNS
到时我们要配置DNS用到53端口 大家就会发现使用udp协议的
而 --dport 就是目标端口 当数据从外部进入服务器为目标端口
反之 数据从服务器出去 则为数据源端口 使用 --sport
-j 就是指定是 ACCEPT 接收 或者 DROP 不接收

#允许目的端口为125的连接
iptables -t filter -A INPUT -p tcp --dport 125 -j ACCEPT
#允许本地的回环
iptables -A INPUT -i lo -j ACCEPT
iptables -A INPUT -o lo -j ACCEPT
#防止ip欺骗，因为出入eth1的包的ip应该是公网ip
iptables -A INPUT -d 192.168.25.0/24 -i eth1 -j DROP
iptables -A INPUT -s 192.168.20.0/24 -i eth1 -j DROP
iptables -A OUTPUT -d 192.168.20.0/24 -o eth1 -j DROP
iptables -A OUTPUT -s 192.168.20.0/24 -o eth1 -j DROP
#屏蔽端口 5000
iptables -A INPUT -p tcp -m tcp --sport 5000 -j DROP
iptables -A INPUT -p udp -m udp --sport 5000 -j DROP
iptables -A OUTPUT -p tcp -m tcp --dport 5000 =j DROP
iptables -A OUTPUT -p udp -m udp --dport 5000 =j DROP
#禁止 Internet网的用户访问 mysql 服务器 (就是 3306 端口)
iptables -A INPUT -i eth1 -p tcp -m tcp --dport 3306 -j DROP
iptables -A INPUT -s 192.168.20.0/24 -i eth1 -p tcp -m tcp --dport 3306 -j ACCEPT
iptables -A INPUT -s 116.255.172.110/24 -i eth0 -p tcp -m tcp --dport 3306 -j ACCEPT
iptables -A INPUT -p tcp -m tcp --dport 3306 -j DROP
#防止syn-flood攻击
iptables -N syn-flood
iptables -A INPUT - i ppp0 -p tcp --syn -j syn-flood
iptables -A syn-flood -m limit --limit 1/s --limit-burst 4 -j RETURN
iptables -A syn-flood -j DROP
#允许http的规则
iptables -A INPUT -i ppp0 -p tcp -s 0/0 --sport 80 -m state --state ESTABLISHED,RELATED -j ACCEPT
iptables -A INPUT -i ppp0 -p tcp -s 0/0 --sport 443 -m state --state ESTABLISHED,RELATED -j ACCEPT
iptables -A INPUT -i ppp0 -p tcp -d 0/0 --dport 80 -j ACCEPT
iptables -A INPUT -i PPP0 -P tcp -d 0/0 --dport 442 -j ACCEPT
#允许DNS 的规则
iptables -A INPUT -i ppp0 -p udp -s0/0 --sport 53 -m state --state ESTABLISHED -j ACCEPT
iptables -A INPUT -i ppp0 -p udp -s0/0 --sport 53 -j ACCEPT
#允许IP1和ip2 远程ssh连接 
iptables -A INPUT -p tcp -s ip1/32 --dport 22 -j ACCEPT
iptables -A INPUT -p tcp -s ip2/32 --dport 22 -j ACCEPT
#允许外网的vpn连接
iptables -A INPUT -p tcp --dport 1723 -j ACCEPT
iptables -A INPUT -p gre -j ACCEPT
#不允许无知不明的数据包
iptables -A INPUT -i eth0 -m state --state INVALID -j DROP
#禁止周一到周六的8:00-20:30禁止内网上QQ网页
iptables -I FORWARD -s 192.168.0.0/24 -m string --string "qq.com" -m time --timestart 8:00 --timestop 20:30 --days Mon,Tue,Wed,Thu,Fri,Sat -j DROP
#是自己用的机器统一放行
iptables -I INPUT -s 192.168.55.250 -j ACCEPT
iptables -I FORWARD -s 192.168.55.250 -j ACCEPT
#禁止BT连接
iptables -A FORWARD -m ipp2p --edk --kazaa --bit -j DROP 
iptables -A FORWARD -p tcp -m ipp2p --ares -j DROP 
iptables -A FORWARD -p udp -m ipp2p --kazaa -j DROP 
#状态是ESTABLISHED,RELATED的直接允许通过。
iptables -A INPUT -i ppp0 -p tcp -s 0/0 -m state --state ESTABLISHED,RELATED -j ACCEPT
#其他情况不允许
iptables -A INPUT -i eth0 -j DROP

