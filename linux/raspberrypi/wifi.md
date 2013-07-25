#wifi 


# lsusb

、测试wifi信号
sudo iwlist wlan0 scan

细心看，找到自己无线路由器的ssid(即用手机搜wifi，显示的wifi名字，路由器中可以配置)
笔者ssid是：TP-LINK_FB5906

3、编辑网卡配置信息
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

笔者把IP设在了静态的了，你总不想每次SSH登陆还要换IP吧！建议设成静态的

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
