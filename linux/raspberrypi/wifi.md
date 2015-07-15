#Raspberry PI USB wifi 设置


1. 检测usbwifi
	
	lsusb
2、测试wifi信号

	sudo iwlist wlan0 scan
	# 找到自己的SSID
3. 编辑网卡配置信息

	sudo nano /etc/network/interfaces

将wlan0的部份修改为：

	auto lo
	 
	iface lo inet loopback
	iface eth0 inet dhcp
	 
	allow-hotplug wlan0
	#iface wlan0 inet manual
	iface wlan0 inet static
	#	wpa-ssid 你要连接的wifi ssid
	#	wpa-psk 你的wpa连接密码
	address 192.168.1.106   # 设定的静态IP地址
	netmask 255.255.255.0   # 网络掩码
	gateway 192.168.1.1     # 网关
	network 192.168.1.1     # 网络地址
    wpa-roam /etc/wpa_supplicant/wpa_supplicant.conf
	iface default inet dhcp


设成静态的

	ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
	update_config=1

	network={
		ssid="wr700"
		psk="password"
		proto=RSN
		key_mgmt=WPA-PSK
		pairwise=CCMP
		auth_alg=OPEN
	}

重启wifi
sudo /etc/init.d/networking restart

========


Raspberry Pi树莓派无线网卡配置[多重方法备选]

要想让树莓派方便操作，肯定需要配置无线网卡，这样可以大大增强树莓派的移动性和便利性，其实配置无线网卡基本就是和普通linux平台下配置无线网卡一样，几种方法大同小异，具体如下：

一、第一种方法：通过配置 /etc/network/interfaces 文件实现
sudo nano /etc/network/interfaces
修改后文件内容如下：
auto lo

iface lo inet loopback
iface eth0 inet dhcp

auto wlan0
allow-hotplug wlan0
iface wlan0 inet dhcp
wpa-ssid “你的wifi名称”
wpa-psk “你的wifi密码”

具体各行配置的意思如下:
auto lo //表示使用localhost
iface eth0 inet dhcp //表示如果有网卡ech0, 则用dhcp获得IP地址 (这个网卡是本机的网卡，而不是WIFI网卡)
auto wlan0 //表示如果有wlan设备，使用wlan0设备名
allow-hotplug wlan0 //表示wlan设备可以热插拨
iface wlan0 inet dhcp //表示如果有WLAN网卡wlan0 (就是WIFI网卡), 则用dhcp获得IP地址

wpa-ssid “你的wifi名称”//表示连接SSID名
wpa-psk “你的wifi密码”//表示连接WIFI网络时，使用wpa-psk认证方式，认证密码

上述定义后，如果有网线连接，则采取DHCP自动连接获得地址，使用命令
sudo /etc/init.d/networking restart
成功后，用 ifconfig 命令可以看到 wlan0 设备，且有了IP地址(已连接)

二、第二种方法：修改sudo nano /etc/wpa_supplicant/wpa_supplicant.conf实现
ctrl_interface=/var/run/wpa_supplicant
ctrl_interface_group=0
ap_scan=2

network={
ssid=“WIFI名称“
proto=WPA2
key_mgmt=WPA-PSK
pairwise=TKIP
group=TKIP
psk=”WIFI密码“
}

然后修改文件sudo nano /etc/network/interfaces,修改后的文件内容如下：
auto lo
iface lo inet loopback
iface eth0 inet dhcp.

auto wlan0
iface wlan0 inet dhcp
pre-up wpa_supplicant -B -Dwext -iwlan0 -c/etc/wpa_supplicant/wpa_supplicant.conf
post-down killall -q wpa_supplicant

修改完成后，使用以下命令重启网络
sudo /etc/init.d/networking restart
成功后，用 ifconfig 命令可以看到 wlan0 设备，且有了IP地址(已连接)

附注：上述两种方法我们都是使用的DHCP动态IP，如果要设置静态ip方法和以及连接隐藏SSID AP的方法：

（1）设置静态ip：
修改文件sudo nano /etc/network/interfaces
auto lo
iface lo inet loopback
iface eth0 inet dhcp

allow-hotplug wlan0
iface wlan0 inet manual
wpa-roam /etc/wpa_supplicant/wpa_supplicant.conf
iface default inet static
address 192.168.1.2
netmask 255.255.255.0
gateway 192.168.1.1
dns-nameservers x.x.x.x #你的本地dns地址

（2）连接WIFI不广播隐藏SSID：
在ssid=”XXXX”下面加一行scan_ssid=1后重启,具体如下：
sudo nano /etc/wpa_supplicant/wpa_supplicant.conf

ctrl_interface=/var/run/wpa_supplicant
ctrl_interface_group=0
ap_scan=2

network={
ssid=“网络id“

scan_ssid=1
proto=WPA2
key_mgmt=WPA-PSK
pairwise=TKIP
group=TKIP
psk=”密码“
}

重启后就可以连上这个不广播SSID的无线网络。
}
}
