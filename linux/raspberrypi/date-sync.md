# date sync

date -R

树莓派没有电池，断电后无法保存时间。树莓派默认安装了NTP(Network Time Protocol)服务来获取互联网上ntp服务器提供的时间。如果这个时间不准，可以用这个命令校准一下。

```sh
sudo ntpd -s -d
```

修改配置文件`etc/ntp.conf` 添加复旦大学提供服务器时间

```sh
server ntp.fudan.edu.cn iburst perfer
server time.asia.apple.com iburst
server asia.pool.ntp.org iburst
server ntp.nict.jp iburst
server time.nist.gov iburst
```

重启服务

    sudo /etc/init.d/ntp restart



可以使用 `ntpdate`来校准时间



sudo apt-get install ntpdate

ntpdate -s ntp.sjtu.edu.cn
