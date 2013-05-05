#ssh 服务安装
**title:** Raspberry Pi SSH远程控制服务器  
**tags:** raspberrypi,ssh,raspbian,树莓派 

**OS:** Raspbian “wheezy”

1.安装ssh 默认有安装

    apt-get install openssh-server

2.安装成功后，启动ssh

    sudo /etc/init.d/ssh start

3.默认启动

    sudo update-rc.d ssh defaults

## 客户端

    ssh pi@ip 开始访问

参考 [scp sftp 文件管理](../soft/sftp)用法
