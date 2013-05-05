#Windows平台下android ADB的安装与使用
**title:**Windows平台下android ADB的安装与使用  
**tags:**window,android,adb,sdk
**info:**window平台安装Android adb工具和学习工具的使用。

##准备##
Android SDK（如觉SDK文件过大，也可以只准备G1的USB for Windows驱动和ADB工具包）    
WinRAR    
G7（HTC Dream）    
数据线     

- - - - - - - - - 

##设置手机##
点击“设置（Settings）”-->“应用程序（Applications）”-->“开发（Development）”-->勾选“USB 调试（USB Debugging）”

- - - - - - - - - 

##安装驱动##
用数据线将手机和电脑连接起来，但千万不要选择“挂载（装载）”！    
当电脑发现新硬件后，安装好[USB驱动](http://driver.zol.com.cn/detail/3/26995.shtml)（如果你不会装驱动，请移步这里）。     
如果你使用的Android SDK包，驱动在SDK包下的usb_driver目录下。

- - - - - - - - - 

##安装ADB##

[ Android SDK 2.2 核心包 for Windows版本](http://dl-ssl.google.com/android/repository/android-2.2_r01-windows.zip)      
使用官方的<http://dl.google.com/android/android-sdk_r08-windows.zip> 工具可以下载到Android 2.3 SDK，包含了对Nexus S的驱动支持，同时本次Google更新了ADT插件到0.8.0版本，自带ProGurad默认发布版本中自动实行扰码加密你的Java应用。

将Android SDK包中tools目录下的 adb.exe 和 AdbWinApi.dll 复制到 Windows下的 system32 目录下。     
如果你使用ADB工具包，可直接解压缩至 system32 目录。  

附：ADB文件浏览器工具（下载）      
下载解压缩后可直接使用，使用前确保你的手机与电脑相连，并且未“挂载”。        
双击“ADB File Explorer v03.exe”文件，你就可以看到你手机里的文件了（窗口右侧）。

在 cmd 中输入 adb 可以获得 adb 使用文档

###常用命令###
>adb shell - 登录设备shell，后面也可直接跟运行命令。如：adb shell rm -r /system/sd/app       
adb pull - 从手机中下载文件到电脑上。如：adb pull /data/app_s/Stock.apk C:\\Stock.apk       
adb push - 从电脑中上传文件到手机上。如：adb push C:\\Stock.apk /data/app_s/Stock.apk       
adb install - 安装软件到手机上。如：adb install C:\\apps2sd.apk       
adb uninstall - 卸载手机上的软件。如：adb uninstall linda.apk    
adb devices 　-　　察看手机是否连接(手机需要打开 USB debug)    
adb remount  - 得到手机的系统文件读写权      

###例子：刷官方高版本后重新获取ROOT和降级方法（部份操作） ###

>adb shell        进行adb运行环境    
su                设置超级用户模式     
mount -o remount,rw -t yaffs2 /dev/block/mtdblock3 /system　映射出目录      
exit        退出su                                              
exit        退出shell准备接受下一个abd命令    
adb push flash_image.zip /system/bin/flash_image　增加flash_image文件到手机上    
adb push recovery-RA-heroc-v1.6.2.img /sdcard        recovery先复制到卡上    
adb shell                                再次进入adb环境    
chmod 0755 /system/bin/flash_image        修改文件属性    
reboot                                        重启系统      
