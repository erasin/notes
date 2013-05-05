# raspberry 下 transmission BT 下载机 
**title:** raspberry pi 下 transmission-daemon bt下载使用详解  
**tags:** transmission,raspberry  

在raspberry 使用Bt下载主要推荐 transmission 。

## transmission-gtk 最为简便的使用

    apt-get install transmission-gtk

打开transmission-gtk后， 在设置中的web页面启用web服务器，最好将下面的白名单取消掉并设置账号和密码，这样就可以使用任何浏览器访问 `Your-IP:9091` 就可以远程控制了。

不过这个要同 [vnc服务](vnc-server)来控制了。

## transmission-daemon 服务

    apt-get install transmission-daemon

起安装开始，已加入开机启动项

将用户 pi 加入到组，赋予权限

    sudo usermod -a -G debian-transmission pi

修改配置 `sudo vim /etc/transmission-daemon/settings.json`

将里面 "umask": 18, 的18改成0 --不然后面会出现下一下就Permission denied的权限问题

常用的修改

| name | value | info
|------|-------|-------------------------------------------------------
|unask                          | 0         | 修改权限，18 to 0
|rpc-authentication-required    | true      | --远程管理认证需要:是
|rpc-enabled                    | true      | -- 远程管理功能打开
|rpc-password                   | "username"| --远程管理的账号(自定义输入)) ]
|rpc-username                   | "password"| --远程管理的密码(自定义输入) , 明文修改，自动加密
|rpc-whitelist                  | `"*.*.*.*"`| --白名单IP,全部改为 `*` 
|rpc-whitelist-enabled          | false     |--使用白名单:否
|download-dir                   | "/mount/usb/dir"    | 默认下载地址
|incomplete-dir                 | "/mount/usb/dir2 "  | 结束后转移文件的地址
|incomplete-dir-enabled         | true/false          | 是否使用 incomplete-dir

其他的配置可按照需求,自行修改.

修改好配置后可以对 transmission-daemon 进行`start/stop/restart`操作

    sudo service transmission-daemon start 

在浏览器中访问'http://Your-IP:9091/' 开始使用bt下载机 

最好开启[ssh服务](ssh-server)，配合使用


## USB disk 挂载问题

如果挂载的事有没有写入权限，重新挂载添加权限

    sudo mount -n -o remount,rw /mount/usb

也可以通过参数来添加权限， uid=pi,gid=pi表示用户"pi"有读写权限(FAT32文件系统本身不支持权限,只能通过这种指派的方式控制读取权限)

    sudo mount -o uid=pi,gid=pi /dev/sda1 /home/pi/usb



