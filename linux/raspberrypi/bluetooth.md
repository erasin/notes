
lsusb 

sudo apt-get -y update

sudo apt-get install bluetooth blueman bluez-compat bluez-audio  bluez-alsa bluez-cups bluez-gstreamer

sudo aptitude install bluez 

sudo service bluetooth start
sudo service bluetooth stauts

sudo hciconfig hci0 up
hciconfig - a

获取蓝牙地址
sudo hcitool scan


bluez-simple-agent hci0 aa:bb:cc:dd:ee:ff

用这个命令（别忘了替换成正确的地址）将这个新设备加入到信任列表：

bluez-test-device trusted aa:bb:cc:dd:ee:ff yes

用这个命令手工开始连接：

bluez-test-input connect aa:bb:cc:dd:ee:ff


bluez-simple-agent hci0 74:A3:4A:01:F5:4A
bluez-test-device trusted 74:A3:4A:01:F5:4A yes
bluez-test-input connect 74:A3:4A:01:F5:4A



---------------
https://wiki.archlinux.org/index.php/Bluetooth_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)#Bluetoothctl
modprobe btusb


Bluetoothctl
通过命令行配对是最可靠的选择。准确的配对过程依不同设备类型及其所提供的输入功能而各不相同。下面提供使用/usr/bin/bluetoothctl配对的一般过程：
启动 bluetoothctl 交互命令。可以输入 help 列出所有有效的命令。
输入power on 命令打开控制器电源。默认是关闭的。
输入devices 命令获取要配对设备的 MAC 地址。
如果设备未在清单中列出，输入 scan on 命令设置设备发现模式。
输入agent on 命令打开代理。
输入 pair MAC Address 开始配对（支持 tab 键补全）。
如果使用无 PIN 码设备，再次连接可能需要手工认证。输入 trust MAC Address 命令。
最后，用 connect MAC_address 命令建立连接。


http://wangye.org/blog/archives/921/



https://www.raspberrypi.org/forums/viewtopic.php?t=68779
