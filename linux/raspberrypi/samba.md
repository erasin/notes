# samba service

架设samba文件共享服务，

```
sudo apt-get install samba samba-common-bin
```

samba 包含 smb 协议框架，支持微软系统服务，加入网络域组。

下载OK后，进行编辑配置文件处理。

```
sudo vim /etc/samba/smb.conf
```

设定所在域

```
workgroup = you workgroup name
wins support = yes
```

添加配置片段到 smb.conf

```
[pihome]
	comment = PI Home
	path = /home/pi 
	browseable = Yes
	writeable = Yes 
	only guest = no 
	create mask =0777
	directory mask = 0777 
	public = no
```

设定密码，花澤の一人できるかな、

```
smbpassd -a pi
```
