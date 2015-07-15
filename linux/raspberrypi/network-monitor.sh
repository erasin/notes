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
