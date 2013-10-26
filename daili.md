# ssh proxy


## SSH帐号

user@you.sshchina.com 

## WIN 软件

下载 [Myentunnel][Myentunnel-down]

Myentunnel 软件设置如下：

![Myentunnel][Myentunnel]

第一次连接 会自动创建密钥, 选择`YES`即可。

![ssh-key][ssh-key]

链接后就产生了默认的 SOCK5 代理 ，默认为 `127.0.0.1` ,端口为 `7070`

> 第一次生产后需要注销系统，或者重新打开软件，否则检测不到 SOCK5, 具体原因应该和window服务有关，这个可以自己查找下原因。


**使用火狐或CHROME可以跳过此处**  **FOR IE**

查看[使用Privoxy][Privoxy-help]

下载[Privoxy][Privoxy]

使用 Privoxy 将SOCK5代理转HTTP代理，解压后直接点击安装运行，转变后的端口是8118。

![Privoxy-rule][Privoxy-rule]

双击Privoxy托盘图标, 弹出了一个框~

点击菜单栏的 Option —-> Main Configuration ——> 就会自动用记事本打开了 Privoxy 的 配置文件
在最后加上

forward-socks5 / 127.0.0.1:7070 . （请手动输入 扶植粘贴可能有错 输入的时候切换为英文输入法 不要漏了最后一个点 ” . ” ）

保存后, 就重启 Privoxy . 方法是把 Privoxy 退出了. (就是在托盘处看不到了)

然后,重新运行.


## chrome 插件

[Proxy SwitchySharp][chrome-prx-ext]

## firefox 插件

[autoproxy][firefox-auto-ext]


## 在线代理规则

	https://autoproxy-gfwlist.googlecode.com/svn/trunk/gfwlist.txt


[Myentunnel-down]: http://pan.baidu.com/s/1l8GEQ
[Myentunnel]: http://static.oschina.net/uploads/space/2013/0503/225300_rFGH_933643.png
[ssh-key]: http://www.issh.in/upload/MyEnTunnel_ask.gif

[Privoxy]: http://pan.baidu.com/s/1j78np
[Privoxy-help]: http://www.hangssh.info/iedaili.html
[Privoxy-rule]: http://hiphotos.baidu.com/pekdo/pic/item/18e1dd34dc610c9c7c1e71ac.jpg

[chrome-prx-ext]: http://chrome.google.com/webstore/detail/proxy-switchysharp/dpplabbmogkhghncfbfdeeokoefdjegm
[firefox-auto-ext]: https://addons.mozilla.org/en-us/firefox/addon/autoproxy/