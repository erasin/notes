#linux 查看硬件信息
**title:**linux 查看硬件信息  
**tags:**linux,hwinfo,硬件查看   
**info:**使用linux命令来查看硬件信息

##linux 查看硬件信息：proc##

使用/proc查看的硬件信息丰富

>	 cat /proc/cpuinfo
	#查看CPU信息，内容很全哦！

	cat /proc/meminfo
	#查看内存信息。

	cat /proc/ioports
	#查看IO端口

	cat /proc/swaps
	#查看交换分区信息(/proc)

	cat /proc/interrupts
	#中断信息

	cat /proc/partitions
	#查看磁盘分区

	cat /proc/bus/usb/devices
	#查看USB设备

	cat /proc/bus/input/devices
	#查看输入设备：键盘鼠标

	cat /proc/bus/pci/devices
	#查看PCI设备

	cat /proc/loadavg
	#查看系统负载

	cat /var/log/demsg
	#查看开机检查的硬件，可以使用grep过虑：eth,cpu,mem,pci,usb,vga,sda……

##linux 查看硬件信息：终端命令行##

使用终端命令行查看的硬件信息可读性好很多了，而且也比较丰富。

>	 lscpu
	#查看CPU信息

	lspci
	#查看PCI设备

	lsusb
	#查看USB设备

	vmstat
	#报告虚拟内存统计信息

	fdisk -l
	#查看分区信息

	hdparm -i /dev/sda
	#查看磁盘参数

	df -h
	#查看磁盘分信息

	dmidecode
	#读取系统DMI表来显示硬件和BIOS信息。

	lsmod
	#当前加载的驱动

	dmesg
	#查看开机检查的硬件，可以使用grep过虑：eth,cpu,mem,pci,usb,vga,sda……

	uptime
	#查看系统负载
	
__NOTE:__

也可以安装其它软件查看硬件信息：aptitude install lshw hwinfo，这些也很强大的。

##linux 查看硬件信息：脚本##

wow ubuntu 提供下载 [hwconfig脚本](http://www.itwhy.org/wp-content/uploads/software/hwconfig) ，这个脚本还有参数哦：hwconfig -h。

it form [http://wowubuntu.com/hardware-info.html]
