#thinkpad 在linux下发热的问题

非常喜欢linux的发行版Ubuntu, 但是有一个困扰了我很久的问题：__右下角托掌处的硬盘温度与使用XP相比要明显高很多__.

原因网上所查到的说法很多, 有的说是因为指纹识别的那个, 也有说是无线网卡的缘故, 但是现在通过试验后, 感觉与硬盘有关的可能性大点.

Ubuntu在交流电模式下, 硬盘处于高性能状态, 磁头不归位, 硬盘不减速, 所以温度比在XP下要高点. 但是在用电池时, 硬盘为了减速节电, 磁头就会平凡归位, 温度自然也会降下来了, 但是磁头归位值就会像WIN7和XP系统下一样非常高, 对硬盘的寿命有一定的影响(a tradeoff).

如果想要降低硬盘温度, 可以采用这个解决办法:

1.安装laptop-mode-tools

>sudo apt-get install laptop-mode-tools

2.修改电源管理配置: laptop-mode.conf

>sudo gedit /etc/laptop-mode/laptop-mode.conf

修改配置文件：

有关参数的说明请参照文件里的注释说明

>ENABLE_LAPTOP_MODE_ON_AC=1 #当笔记本使用交流电时也开启laptop-mode  
ENABLE_LAPTOP_MODE_WHEN_LID_CLOSED=1

>MINIMUM_BATTERY_CHARGE_PERCENT=11

>DISABLE_LAPTOP_MODE_ON_CRITICAL_BATTERY_LEVEL=7  
CONTROL_HD_IDLE_TIMEOUT=0

ArchLinux/Debian 默认 20 ，导致硬盘经常停转起转，Ubuntu 默认 60 ，用了几天没感觉。总之我关掉了这项功能。

>LM_AC_HD_IDLE_TIMEOUT_SECONDS=20  
LM_BATT_HD_IDLE_TIMEOUT_SECONDS=20  
CONTROL_HD_POWERMGMT=1  

ArchLinux/Debian 默认不启用，Ubuntu 默认启用，我也设为启用，但修改下面三个值。

>BATT_HD_POWERMGMT=192  
LM_AC_HD_POWERMGMT=192  
NOLM_AC_HD_POWERMGMT=192  

经过试验(在Thinkpad上), 修改完后, 硬盘温度一直能保持在40度以下.

原文：http://www.java123.net/view-38114-1.html
