# kindle paperwhite

## 5.4.5 降级 

一、5.4.5 版固件降级步骤

因为新版本的越狱方法还没出现，所以需要降级到上一个固件版本才行，但又发现按原来的方法却无法直接降级和越狱了，为了解决这个看似棘手的问题，Kindle 伴侣为您介绍一种特别简单的 KPW2 降级方法，可以成功将 KPW2 的固件从 5.4.5 或 5.4.5.1 降级到 5.4.3.2，步骤如下：

1、下载亚马逊官方 KPW2 5.4.3.2 固件：<http://pan.baidu.com/s/1hqtai9Y>

2、用 USB 数据线将您的 KPW2 连接到电脑；

3、把下载的官方 5.4.3.2 版固件文件（update_kindle_5.4.3.2.bin）复制到 kpw2 的根目录；

4、不要拔出数据线，直接长按电源键重启，重启后就可以得到5.4.3.2版的kpw2了。

二、禁止 Kindle 自动升级

如果遇到把固件降级到 5.4.3.2 版后又被自动升级到 5.4.5.1 版的情况，这就需要禁用 Kindle 固件自动升级，可以试试下面网友提供的一个方法（此方法未验证）：

1、用 USB 数据线将您的 kindle 连接到电脑，打开 kindle 盘，如 h:。在根目录下新建一个文件夹，命名为 a（感觉这个文件么应该是随便取就成）。

2、打开电脑的开始-附件-命令提示符，输入 h: 回车。在命令提示符里输入命令 rename a update.bin.tmp.partial 回车确认。

3、关闭命令提示符框，打开 kindle 盘 h:，发现刚才新建立的那个 a 的文件夹变成了名为 update.bin.tmp.partial 的文件夹。右键单击这个文件夹，在属性里把它改成“只读”。完成。

> 需要注意的是，当你复制完官方固件到 KPW2 的根目录后，不要马上按电源键开始重启，需要等待一段时间以使Windows的磁盘高速缓存刷新，保证数据不会丢失。


这个比下面的好
<http://kindlefere.com/post/33.html>

## 越狱

原帖是在 http://www.douban.com/group/topic/46545085/ ;

总的来说是下载资源
http://pan.baidu.com/s/1gdqoa91

文件基本都在`kindle越狱`文件.

1、越狱

找到`kindle-5.4-jailbreak.zip` ,解压其，将解压出来的文件拷贝到kindle根目录（当然usb连接PC电脑）.断开usb连接后，到设置中更新kindle。 当屏幕下方出现下面的字样就安装成功了。
    
    **** JAILBREAK **** 

2、启动器

破解Kindlet：

找到 `prerequisites-for-k5-kpw-all.zip`，解压后将`update_kindlet-dev-20130710-k5-ALL_install.bin`拷贝到kindle根目录，断开usb，到设置中更新，可能会自动重启。

安装KUAL 

解压`KUAL.V.2.2.zip`,将`KindleLauncher-2.0.azw2`拷贝到 kindle的`documents`中,在目录中打开 kual 即可使用。

3、插件 extensions

在打开 kual后会自动布局目录。连接PC电脑，将其他插件解压后的文件拷贝到kindle中的`extensions`即可。断开usb，打开 kual 即可看到插件了。

koreader安装

解压`koreader-kindle-arm...zip`,解压出的文件，直接拷贝到kindle根目录即可。kual就可以查看了

4、USBNetwork

解压`update_usbnet_0.14.N_install.zip`将bin文件拷贝到kindle根目录，然后到kindle设置中更新即可。自动重启。

ssh的wifi传输，去原帖子看吧。


