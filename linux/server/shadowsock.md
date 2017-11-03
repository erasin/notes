shadowsock 
================

## ssserver

到 <https://github.com/shadowsocks/shadowsocks-rust/releases> 下载 ssserver 等命令

```
wget https://github.com/shadowsocks/shadowsocks-rust/releases
tar xvf *.gz
```

创建配置

```
touch ss.conf
vim ss.conf
```

配置文件内容

```
{
    "servers": [
        {
            "address": "广域 IP 或域名",
            "port": 8082,
            "password": "YOU-PASSWD",
            "method": "aes-256-cfb",
            "timeout": 600
        },
        {
            "address": "广域 IP 或域名",
            "port": 8082,
            "password": "YOU-PASSWD",
            "method": "aes-256-cfb",
            "timeout": 600
        }
    ],
    "local_port": 1080,
    "local_address": "127.0.0.1"
}
```

```
执行 
ssserver -c ss.conf
```


## bbr

利用网络加速，提高宽带利用率。

```
wget –no-check-certificate https://www.xiaoweigod.com/shell/bbr.sh
chmod +x bbr.sh
./bbr.sh
# 查看内核 4.9.0
uname -r
sysctl net.ipv4.tcp_available_congestion_control
# 检查
lsmod | grep bbr
```

* bbr: http://blog.csdn.net/dog250/article/details/52830576
