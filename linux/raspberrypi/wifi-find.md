
http://blog.csdn.net/bailyzheng/article/details/33336709


1、检查USB无线网卡是否已经正确识别
将无线USB网卡插入树莓派后启动树莓派，比较不建议热插拔，因为插入的一瞬间会有比较高的电流，如果电源输出不够可能导致树莓派重启。用自己的方法进入shell界面后输入命令：

	lsusb

如果树莓派已经正常识别，在显示类似于如下的信息中可以看到你的USB无线网卡设备ID和芯片型号

最新的raspbian已经有了wifi必要的包，直接插上就可以用了。不过最好还是可以看看iwconfig确认一下，输入

	iwconfig

如果出现了wlan0，那说明网卡已经正常工作了。（这里的示例是已经用usb无线网卡连接上网络了，所以会显示ESSID。）如果这里的显示不正常，请安装连接wifi必要的包

	sudo apt-get install wireless-tools wpasupplicant firmware-realtek  

2、设置wifi

输入如下命令可以搜索附近所有可连接的Wifi AP：

wlist wlan0 scan  








