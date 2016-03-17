# firewall

配置防火墙防止syn，ddos攻击

	# vim /etc/sysconfig/iptables
	# 在iptables中加入下面几行

	#anti syn，ddos
	iptables -A FORWARD -p tcp --syn -m limit --limit 1/s --limit-burst 30 -j ACCEPT
	iptables -A FORWARD -p tcp --tcp-flags SYN,ACK,FIN,RST RST -m limit --limit 1/s -j ACCEPT
	iptables -A FORWARD -p icmp --icmp-type echo-request -m limit --limit 1/s -j ACCEPT

说明：
* 第一行：每秒中最多允许5个新连接
* 第二行：防止各种端口扫描 
* 第三行：Ping洪水攻击（Ping of Death）

可以根据需要调整或关闭

重启防火墙

	# service iptables save
	# /etc/init.d/iptables restart

iptables 

屏蔽一个IP

	iptables -I INPUT -s 192.168.0.1 -j DROP

怎么防止别人ping我？？

	iptables -A INPUT -p icmp -j DROP

防止同步包洪水（Sync Flood） 

	iptables -A FORWARD -p tcp --syn -m limit --limit 1/s -j ACCEPT

防止各种端口扫描 

	iptables -A FORWARD -p tcp --tcp-flags SYN,ACK,FIN,RST RST -m limit --limit 1/s -j ACCEPT

Ping洪水攻击（Ping of Death） 

	iptables -A FORWARD -p icmp --icmp-type echo-request -m limit --limit 1/s -j ACCEPT

 

# NMAP FIN/URG/PSH

	iptables -A INPUT -i eth0 -p tcp --tcp-flags ALL FIN,URG,PSH -j DROP

# Xmas Tree

	iptables -A INPUT -i eth0 -p tcp --tcp-flags ALL ALL -j DROP

# Another Xmas Tree

	iptables -A INPUT -i eth0 -p tcp --tcp-flags ALL SYN,RST,ACK,FIN,URG -j DROP

# Null Scan(possibly)

	iptables -A INPUT -i eth0 -p tcp --tcp-flags ALL NONE -j DROP

# SYN/RST

	iptables -A INPUT -i eth0 -p tcp --tcp-flags SYN,RST SYN,RST -j DROP

# SYN/FIN -- Scan(possibly)

	iptables -A INPUT -i eth0 -p tcp --tcp-flags SYN,FIN SYN,FIN -j DROP

##限制对内部封包的发送速度

	iptables -A INPUT -f -m limit --limit 100/s --limit-burst 100 -j ACCEPT

##限制建立联机的转

	iptables -A FORWARD -f -m limit --limit 100/s --limit-burst 100 -j ACCEPT

 



防范DDOS攻击脚本
#防止SYN攻击 轻量级预防 
iptables -N syn-flood 
iptables -A INPUT -p tcp --syn -j syn-flood 
iptables -I syn-flood -p tcp -m limit --limit 3/s --limit-burst 6 -j RETURN 
iptables -A syn-flood -j REJECT
#防止DOS太多连接进来,可以允许外网网卡每个IP最多15个初始连接,超过的丢弃 
iptables -A INPUT -i eth0 -p tcp --syn -m connlimit --connlimit-above 15 -j DROP 
iptables -A INPUT -p tcp -m state --state ESTABLISHED,RELATED -j ACCEPT

#用Iptables抵御DDOS (参数与上相同)
iptables -A INPUT  -p tcp --syn -m limit --limit 12/s --limit-burst 24 -j ACCEPT
iptables -A FORWARD -p tcp --syn -m limit --limit 1/s -j ACCEPT
##########################################################
防范CC攻击
当apache站点受到严重的cc攻击，我们可以用iptables来防止web服务器被CC攻击，实现自动屏蔽IP的功能。
1．系统要求
(1)LINUX 内核版本：2.6.9-42ELsmp或2.6.9-55ELsmp（其它内核版本需要重新编译内核，比较麻烦，但是也是可以实现的）。
(2)iptables版本：1.3.7
2. 安装
安装iptables1.3.7和系统内核版本对应的内核模块kernel-smp-modules-connlimit
3. 配置相应的iptables规则
示例如下：
(1)控制单个IP的最大并发连接数
iptables -I INPUT -p tcp --dport 80 -m connlimit  --connlimit-above 50 -j REJECT #允许单个IP的最大连接数为 30
#默认iptables模块不包含connlimit,需要自己单独编译加载，请参考该地址
http://sookk8.blog.51cto.com/455855/280372不编译内核加载connlimit模块

(2)控制单个IP在一定的时间（比如60秒）内允许新建立的连接数
iptables -A INPUT -p tcp --dport 80 -m recent --name BAD_HTTP_ACCESS --update --seconds 60 --hitcount 30 -j REJECT iptables -A INPUT -p tcp --dport 80 -m recent --name BAD_HTTP_ACCESS --set -j ACCEPT
#单个IP在60秒内只允许最多新建30个连接

4. 验证
（1）工具：flood_connect.c（用来模拟攻击)
（2）查看效果：
使用
watch 'netstat -an | grep:21 | grep<模拟攻击客户机的IP>| wc -l'

实时查看模拟攻击客户机建立起来的连接数，
使用
watch 'iptables -L -n -v | \grep<模拟攻击客户机的IP>'

查看模拟攻击客户机被 DROP 的数据包数。
5．注意
为了增强iptables防止CC攻击的能力，最好调整一下ipt_recent的参数如下：
#cat/etc/modprobe.conf options ipt_recent ip_list_tot=1000 ip_pkt_list_tot=60
#记录1000个IP地址，每个地址记录60个数据包 #modprobe ipt_recent


 

一个不错的防火墙代码
#####################################################

-A INPUT -f -m limit --limit 100/sec --limit-burst 100 -j ACCEPT

-A INPUT -p tcp -m tcp --tcp-flags SYN,RST,ACK SYN -m limit --limit 20/sec --limit-burst 200 -j

ACCEPT

-A INPUT -p udp -m udp --dport 138 -j DROP

-A INPUT -p udp -m udp --dport 137 -j DROP

-A INPUT -p tcp -m tcp --dport 1068 -j DROP

-A INPUT -p icmp -m limit --limit 12/min --limit-burst 2 -j DROP

-A FORWARD -f -m limit --limit 100/sec --limit-burst 100 -j ACCEPT

-A FORWARD -p tcp -m tcp --tcp-flags SYN,RST,ACK SYN -m limit --limit 20/sec --limit-burst 200

-j ACCEPT

-A FORWARD -p tcp -m tcp --dport 445 -j DROP

-A FORWARD -p udp -m udp --dport 138 -j DROP

-A FORWARD -p udp -m udp --dport 137 -j DROP

-A FORWARD -p tcp -m tcp --dport 1068 -j DROP

-A FORWARD -p tcp -m tcp --dport 5554 -j DROP

-A FORWARD -p icmp -j DROP

:PREROUTING ACCEPT [986908:53126959]

:POSTROUTING ACCEPT [31401:2008714]

:OUTPUT ACCEPT [30070:1952143]

-A POSTROUTING -p tcp -m tcp --dport 445 -j DROP

#####################################################

 

iptables 防火墙例子
 

#!/bin/bash
#
# The interface that connect Internet

# echo
echo "Enable IP Forwarding..."
echo 1 > /proc/sys/net/ipv4/ip_forward
echo "Starting iptables rules..."

IFACE="eth0"

# include module
modprobe ip_tables
modprobe iptable_nat
modprobe ip_nat_ftp
modprobe ip_nat_irc
modprobe ip_conntrack
modprobe ip_conntrack_ftp
modprobe ip_conntrack_irc
modprobe ipt_MASQUERADE

# init
/sbin/iptables -F 
/sbin/iptables -X
/sbin/iptables -Z
/sbin/iptables -F -t nat
/sbin/iptables -X -t nat
/sbin/iptables -Z -t nat

/sbin/iptables -X -t mangle

# drop all
/sbin/iptables -P INPUT DROP
/sbin/iptables -P FORWARD ACCEPT
/sbin/iptables -P OUTPUT ACCEPT
/sbin/iptables -t nat -P PREROUTING ACCEPT
/sbin/iptables -t nat -P POSTROUTING ACCEPT
/sbin/iptables -t nat -P OUTPUT ACCEPT

/sbin/iptables -A INPUT -f -m limit --limit 100/sec --limit-burst 100 -j ACCEPT
/sbin/iptables -A INPUT -p tcp -m tcp --tcp-flags SYN,RST,ACK SYN -m limit --limit 20/sec --limit-burst 200 -j ACCEPT

/sbin/iptables -A INPUT -p icmp -m limit --limit 12/min --limit-burst 2 -j DROP

/sbin/iptables -A FORWARD -f -m limit --limit 100/sec --limit-burst 100 -j ACCEPT
/sbin/iptables -A FORWARD -p tcp -m tcp --tcp-flags SYN,RST,ACK SYN -m limit --limit 20/sec --limit-burst 200 -j ACCEPT

# open ports
/sbin/iptables -A INPUT -i $IFACE -p tcp --dport 21 -j ACCEPT
/sbin/iptables -A INPUT -i $IFACE -p tcp --dport 22 -j ACCEPT
/sbin/iptables -A INPUT -i $IFACE -p tcp --dport 25 -j ACCEPT
/sbin/iptables -A INPUT -i $IFACE -p tcp --dport 53 -j ACCEPT
/sbin/iptables -A INPUT -i $IFACE -p udp --dport 53 -j ACCEPT
/sbin/iptables -A INPUT -i $IFACE -p tcp --dport 80 -j ACCEPT
/sbin/iptables -A INPUT -i $IFACE -p tcp --dport 100 -j ACCEPT
/sbin/iptables -A INPUT -i $IFACE -p tcp --dport 113 -j ACCEPT

# close ports
iptables -I INPUT -p udp --dport 69 -j DROP
iptables -I INPUT -p tcp --dport 135 -j DROP
iptables -I INPUT -p udp --dport 135 -j DROP
iptables -I INPUT -p tcp --dport 136 -j DROP
iptables -I INPUT -p udp --dport 136 -j DROP
iptables -I INPUT -p tcp --dport 137 -j DROP
iptables -I INPUT -p udp --dport 137 -j DROP
iptables -I INPUT -p tcp --dport 138 -j DROP
iptables -I INPUT -p udp --dport 138 -j DROP
iptables -I INPUT -p tcp --dport 139 -j DROP
iptables -I INPUT -p udp --dport 139 -j DROP
iptables -I INPUT -p tcp --dport 445 -j DROP
iptables -I INPUT -p udp --dport 445 -j DROP
iptables -I INPUT -p tcp --dport 593 -j DROP
iptables -I INPUT -p udp --dport 593 -j DROP
iptables -I INPUT -p tcp --dport 1068 -j DROP
iptables -I INPUT -p udp --dport 1068 -j DROP
iptables -I INPUT -p tcp --dport 4444 -j DROP
iptables -I INPUT -p udp --dport 4444 -j DROP
iptables -I INPUT -p tcp --dport 5554 -j DROP
iptables -I INPUT -p tcp --dport 1434 -j DROP
iptables -I INPUT -p udp --dport 1434 -j DROP
iptables -I INPUT -p tcp --dport 2500 -j DROP
iptables -I INPUT -p tcp --dport 5800 -j DROP
iptables -I INPUT -p tcp --dport 5900 -j DROP
iptables -I INPUT -p tcp --dport 6346 -j DROP
iptables -I INPUT -p tcp --dport 6667 -j DROP
iptables -I INPUT -p tcp --dport 9393 -j DROP

iptables -I FORWARD -p udp --dport 69 -j DROP
iptables -I FORWARD -p tcp --dport 135 -j DROP
iptables -I FORWARD -p udp --dport 135 -j DROP
iptables -I FORWARD -p tcp --dport 136 -j DROP
iptables -I FORWARD -p udp --dport 136 -j DROP
iptables -I FORWARD -p tcp --dport 137 -j DROP
iptables -I FORWARD -p udp --dport 137 -j DROP
iptables -I FORWARD -p tcp --dport 138 -j DROP
iptables -I FORWARD -p udp --dport 138 -j DROP
iptables -I FORWARD -p tcp --dport 139 -j DROP
iptables -I FORWARD -p udp --dport 139 -j DROP
iptables -I FORWARD -p tcp --dport 445 -j DROP
iptables -I FORWARD -p udp --dport 445 -j DROP
iptables -I FORWARD -p tcp --dport 593 -j DROP
iptables -I FORWARD -p udp --dport 593 -j DROP
iptables -I FORWARD -p tcp --dport 1068 -j DROP
iptables -I FORWARD -p udp --dport 1068 -j DROP
iptables -I FORWARD -p tcp --dport 4444 -j DROP
iptables -I FORWARD -p udp --dport 4444 -j DROP
iptables -I FORWARD -p tcp --dport 5554 -j DROP
iptables -I FORWARD -p tcp --dport 1434 -j DROP
iptables -I FORWARD -p udp --dport 1434 -j DROP
iptables -I FORWARD -p tcp --dport 2500 -j DROP
iptables -I FORWARD -p tcp --dport 5800 -j DROP
iptables -I FORWARD -p tcp --dport 5900 -j DROP
iptables -I FORWARD -p tcp --dport 6346 -j DROP
iptables -I FORWARD -p tcp --dport 6667 -j DROP
iptables -I FORWARD -p tcp --dport 9393 -j DROP

/sbin/iptables -A INPUT -i $IFACE -m state --state RELATED,ESTABLISHED -j ACCEPT
/sbin/iptables -A INPUT -i $IFACE -m state --state NEW,INVALID -j DROP

# drop ping
/sbin/iptables -A INPUT -p icmp -j DROP

/sbin/iptables -I INPUT -s 222.182.40.241 -j DROP
