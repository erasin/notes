
# wifi reconnect

Raspberry Pi 使用USB无线网卡的时候不会因为路由重启而掉线。

```sh
#!/bin/bash
while true ; do
   if ifconfig wlan0 | grep -q "inet addr:" ; then
      sleep 60
   else
      echo "Network connection down! Attempting reconnection."
      ifup --force wlan0
      sleep 10
   fi
done
------邪恶的分割线-----
#!/bin/bash
if ifconfig wlan0 | grep -q "inet addr:" ; then
      exit
   else
      echo "Network connection down! Attempting reconnection."
      ifup --force wlan0
   fi
#crontab -e
#*/10 * * * * bash /home/network-monitor.sh #每十分钟执行一次
```







将代码复制到你的根目录中保存为network-monitor.sh然后运行命令。
sudo chmod +x ./network-monitor.sh
把它设置为可执行文件,使用命令可在后台运行。
sudo ./network-monitor.sh &
它会每60秒检查，如果你的无线网络具有网络连接。如果它发现它没有网络地址将尝试强制重新连接并继续这样执行，直到重新建立连接。如果你想阻止它在后台运行，首先使用以下命令.
fg
将迫使它为前台运行，然后你可以使用ctrl-c停止.
测试了它几个方面。
首先关闭我的WiFi接入点，该脚本没有检测到网络连接并开始试图强迫一个连接。几分钟后，重新建立连接供电接入点回来后。
另一个试验，我删除了MAC地址从地址列表中允许接入点的MAC地址过滤。网络连接了下来，重新启用的MAC地址和它一、两分钟后回来了。
