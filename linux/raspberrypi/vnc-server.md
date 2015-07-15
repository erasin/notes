# VNC-SERVER
**title:** Raspberry Pi VNC远程控制服务器  
**tags:** raspberrypi,vnc,raspbian,树莓派     

**OS:** Raspbian “wheezy”

## 安装服务
安装  tightvncserver

    # apt-get install tightvncserver

使用编辑器创建 启动脚本 `/etc/init.d/tightvncserver`  
脚本内容:
```sh
### BEGIN INIT INFO
# Provides: tightvnc
# Required-Start: $remote_fs $syslog
# Required-Stop: $remote_fs $syslog
# Default-Start: 2 3 4 5
# Default-Stop: 0 1 6
# Short-Description: Start VNC Server as a service
# Description: Start VNC Server as a service.
### END INIT INFO
#!/bin/sh
# /etc/init.d/tightvncserver
# Customised by Stewart Watkiss

# Set the VNCUSER variable to the name of the user to start tightvncserver under
VNCUSER='pi'
eval cd ~$VNCUSER
case "$1" in
start)
  su $VNCUSER -c '/usr/bin/tightvncserver :1'
  echo "Starting TightVNC server for $VNCUSER "
  ;;
stop)
  pkill Xtightvnc
  echo "Tightvncserver stopped"
  ;;
*)
echo "Usage: /etc/init.d/tightvncserver {start|stop}"
exit 1
;;
esac
exit 0
#
```

修改脚本所属权限

    sudo chmod 755 /etc/init.d/tightvncserver

或者

    sudo chmod +x /etc/init.d/tightvncserver

设定开机启动

    sudo update-rc.d tightvncserver defaults

执行

    vncserver

设定访问密码,会出现两次，之后会询问一个只读密码，n跳过

##使用vnc客户端访问

默认  ip:5901


本文来源：<http://www.shumeipai.net/thread-1394-1-1.html>


