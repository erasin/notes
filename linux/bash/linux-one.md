0001 修改主机名(陈绪)
vi /etc/sysconfig/network，修改HOSTNAME一行为"HOSTNAME=主机名"(没有这行？那就添加这一行吧)，然后运行命令" hostname 主机名"。一般还要修改/etc/hosts文件中的主机名。这样，无论你是否重启，主机名都修改成功。

0002 Red Hat Linux启动到文字界面(不启动xwindow)(陈绪)
vi /etc/inittab
id:x:initdefault:
x=3:文本方式 x=5:图形方式

0003 linux的自动升级更新问题(hutuworm，NetDC，陈绪)
对于redhat，在www.redhat.com/corp/support/errata/找到补丁，6.1以后的版本带有一个工具up2date，它能够测定哪些rpm包需要升级，然后自动从redhat的站点下载并完成安装。
升级除kernel外的rpm: up2date -u
升级包括kernel在内的rpm: up2date -u -f

Gentoo升级方法
更新portage tree：  emerge --sync
更新/安装软件包： emerge [软件包名] （如安装vim:  emerge vim）

Debian跟别的发行版还是有很大的差别的，用Debian做服务器维护更加方便；红帽的升级其实挺麻烦的，当然，如果你交钱给红帽的话，服务是会不一样的。
Debian下升级软件：
apt-get update
apt-get upgrade
前提：配置好网络和/etc/apt/sources.list，也可以用apt-setup设置。

0004 windows下看linux分区的软件(陈绪，hutuworm)
Paragon.Ext2FS.Anywhere.2.5.rar和explore2fs-1.00-pre4.zip
现在不少Linux发行版安装时缺省基于LVM建分区，所以explore2fs也与时俱进地开始支持LVM2：
http://www.chrysocome.net/downloads/explore2fs-1.08beta9.zip

0005 mount用法(sakulagi，sxsfxx，aptkevin)
fat32的分区 mount -o codepage=936,iocharset=cp936 /dev/hda7 /mnt/cdrom
ntfs的分区 mount -o iocharset=cp936 /dev/hda7 /mnt/cdrom
iso文件 mount -o loop /abc.iso /mnt/cdrom
软盘 mount /dev/fd0 /mnt/floppy
USB闪存 mount /dev/sda1 /mnt/cdrom
在有scsi硬盘的计算机上，如果用上面的命令挂载usb闪存，则会mount到/boot分区。这种情况，应该先用fdisk -l /dev/sd? 来看看到底usb闪存盘是在哪个设备下(通常会是sdb或者sdc)。比如某台机器上，就是在sdc1上面。
所有/etc/fstab内容 mount -a
可以指定文件格式"-t 格式", 格式可以为vfat, ext2, ext3等.
访问DVD mount -t iso9660 /dev/dvd /mnt/cdrom或mount -t udf /dev/dvd /mnt/cdrom
注意：dvd的格式一般为iso9660或udf之一

0006 在vmware的LINUX中使用本地硬盘的FAT分区(陈绪)
将本地的FAT分区共享，然后在VMWARE中使用SMBFS挂上。可以将如下的行放到/etc/fstab中：
//win_ip/D$ /mnt/d smbfs defaults,auto,username=win_name,password=win_pass,codepage=936,iocharest=gb2312 0 0
其中win_ip是你的windows的IP地址；
D$是你的windows里面共享的D盘的共享名；
/mnt/d是要将该分区mount到linux的目录；
win_name和win_pass是你的WINDOWS里面可以读取该分区的用户，比如你的管理员名和密码。
如果你运行了/etc/rc.d/init.d/netfs，那么在启动的时候就会自动挂载这个分区。

0007.a 删除名为-a的文件(陈绪)
rm ./-a
rm -- -a  告诉rm这是最后一个选项，参见getopt
ls -i 列出inum，然后用find . -inum inum_of_thisfile -exec rm '{}' /;

0007.b 删除名为/a的文件(陈绪)
rm //a

0007.c 删除名字带的/和‘/0'文件(陈绪)
这些字符是正常文件系统所不允许的字符，但可能在文件名中产生，如unix下的NFS文件系统在Mac系统上使用
1.解决的方法，把NFS文件系统在挂到不过滤'/'字符的系统下删除含特殊文件名的文件。
2.也可将错误文件名的目录其它文件移走，ls -id 显示含该文件目录的inum，umount 文件系统， 
clri清除该目录的inum，fsck，mount，check your lost+found，rename the file in it.
最好是通过WINDOWS FTP过去就可以删除任何文件名的文件了!

0007.d 删除名字带不可见字符的文件(陈绪)
列出文件名并转储到文件：ls -l  >aaa
然后编辑文件的内容加入rm命令使其内容成为删除上述文件的格式：
vi aaa
[rm -r *******
]
把文件加上执行权限 chmod +x aaa
执行 $aaa

0007.e 删除文件大小为零的文件(陈绪)
rm -i `find ./ -size 0`
find ./ -size 0 -exec rm {} /;
或
find ./  -size 0 | xargs rm -f &
或
for file in *   #自己定义需要删除的文件类型
do
    if [ ! -s ${file} ]
    then
        rm ${file}

        echo "rm $file Success!"
    fi
done

0008 redhat设置滚轮鼠标(mc1011)
进入X后，选择鼠标的配置，选择wheel mouse(ps/2)就可以了，
如果鼠标表现异常，重启计算机即可。
(或者su, vi /etc/X11/XF86Config, 把PS/2 改成 ImPS/2) 

0009 加装xwindow(陈绪)
用linux光盘启动，选择升级，然后单独选择包，安装即可

0010 删除linux分区(陈绪)
做一张partition magic的启动软盘,启动后删除. 或者用win2000的启动光盘启动,然后删除.

0011 如何退出man(陈绪)
q

0012 不编译内核，mount ntfs分区(陈绪,hutuworm,qintel)
原装rh8，未升级或编译内核
1. 上google.com搜索并下载 kernel-ntfs-2.4.18-14.i686.rpm
2. rpm -ivh kernel-ntfs-2.4.18-14.i686.rpm
3. mkdir /mnt/c
4. mount -t ntfs /dev/hda1 /mnt/c
或
Read only: http://www.linux-ntfs.org/
Read/Write: http://www.jankratochvil.net/project/captive/

0013 tar 分卷压缩和合并(WongMokin，Waker)
以每卷500M为例
tar分卷压缩：tar cvzpf - somedir | split -d -b 500m
tar多卷合并：cat x* > mytarfile.tar.gz

0014 使用lilo/grub时找回忘记了的root口令(陈绪)
三种办法：
1.在系统进入单用户状态，直接用passwd root去更改
2.用安装光盘引导系统，进行linux rescue状态，将原来/分区挂接上来,作法如下：
cd /mnt
mkdir hd
mount -t auto /dev/hdaX(原来/分区所在的分区号） hd
cd hd
chroot ./

passwd root
这样可以搞定
3.将本机的硬盘拿下来，挂到其他的linux系统上，采用的办法与第二种相同
rh8中
一. lilo
   1. 在出现 lilo: 提示时键入 linux single
      画面显示 lilo:  linux single
   2. 回车可直接进入linux命令行
   3. #vi /etc/shadow
      将第一行，即以root开头的一行中root:后和下一个:前的内容删除，
      第一行将类似于
      root::......
      保存
   4. #reboot重启，root密码为空
二. grub
   1. 在出现grub画面时，用上下键选中你平时启动linux的那一项(别选dos哟)，然后按e键
   2. 再次用上下键选中你平时启动linux的那一项(类似于kernel /boot/vmlinuz-2.4.18-14 ro root=LABEL=/)，然后按e键
   3. 修改你现在见到的命令行，加入single，结果如下：
      kernel /boot/vmlinuz-2.4.18-14 single ro root=LABEL=/
   4. 回车返回，然后按b键启动，即可直接进入linux命令行
   5. #vi /etc/shadow
      将第一行，即以root开头的一行中root:后和下一个:前的内容删除，
      第一行将类似于
      root::......
      保存
   6. #reboot重启，root密码为空

0015 使ctrl + alt + del失效(陈绪)
vi /etc/inittab
将ca::ctrlaltdel:/sbin/shutdown -t3 -r now这行注释掉，就可以了

0016 如何看出redhat的版本是7还是8(hutuworm)

cat /proc/version或者cat /etc/redhat-release或者cat /etc/issue

0017 文件在哪个rpm中(无双)
上www.rpmfind.net上搜，或者rpm -qf 文件名得到

0018 把man或info的信息存为文本文件(陈绪)
以 tcsh 为例：
man tcsh | col -b > tcsh.txt
info tcsh -o tcsh.txt -s

0019 利用现存两个文件，生成一个新的文件(陈绪)
1. 取出两个文件的并集(重复的行只保留一份)
2. 取出两个文件的交集(只留下同时存在于两个文件中的文件)
3. 删除交集，留下其他的行
1. cat file1 file2 | sort | uniq
2. cat file1 file2 | sort | uniq -d
3. cat file1 file2 | sort | uniq -u

0020 设置com1口，让超级终端通过com1口进行登录(陈绪)
确认有/sbin/agetty，编辑/etc/inittab，添加
7:2345:respawn:/sbin/agetty /dev/ttyS0 9600
9600bps是因为联路由器缺省一般都是这种速率，也可以设成
19200、38400、57600、115200
修改/etc/securetty，添加一行：ttyS0，确保root用户能登录
重启机器，就可以拔掉鼠标键盘显示器（启动时最好还是要看看输出信息）了

0021 删除目录下所有文件包括子目录(陈绪)
rm -rf 目录名

0022 查看系统信息(陈绪)
cat /proc/cpuinfo - CPU (i.e. vendor, Mhz, flags like mmx)
cat /proc/interrupts - 中断
cat /proc/ioports - 设备IO端口
cat /proc/meminfo - 内存信息(i.e. mem used, free, swap size)
cat /proc/partitions - 所有设备的所有分区
cat /proc/pci - PCI设备的信息
cat /proc/swaps - 所有Swap分区的信息
cat /proc/version - Linux的版本号 相当于 uname -r
uname -a - 看系统内核等信息

0023 去掉多余的回车符(陈绪)
sed 's/^M//' test.sh > back.sh， 注意^M是敲ctrl_v ctrl-m得到的
或者 dos2unix filename

0024 切换X桌面(lnx3000)
如果你是以图形登录方式登录linux，那么点击登录界面上的session（任务）即可以选择gnome和kde。如果你是以文本方式登录，那执行switchdesk gnome或switchdesk kde，然后再startx就可以进入gnome或kde。
(或者vi ~/.xinitrc，添加或修改成exec gnome-session 或exec startkde，
然后用startx启动X)

0025 通用的声卡驱动程序(lnx3000)
OSS www.opensound.com/   ALSA www.alsa-project.org/

0026 改变redhat的系统语言/字符集(beming/mc1011)
修改 /etc/sysconfig/i18n 文件，如
LANG="en_US"，xwindow会显示英文界面，
LANG="zh_CN.GB18030"，xwindow会显示中文界面。
还有一种方法
cp /etc/sysconfig/i18n $HOME/.i18n
修改 $HOME/.i18n 文件，如
LANG="en_US"，xwindow会显示英文界面，
LANG="zh_CN.GB18030"，xwindow会显示中文界面。
这样就可以改变个人的界面语言，而不影响别的用户
(Debian不支持GB18030(RH的zysong字库是有版权的)
现在好像没有Free的GBK和GB18030字库
vi .bashrc
export LANG=zh_CN.GB2312
export LC_ALL=zh_CN.GB2312) 

0027 把屏幕设置为90列(陈绪)
stty cols 90

0028 使用md5sum文件(陈绪)
md5sum isofile > hashfile, 将 md5sum 档案与 hashfile 档案内容比对, 验证杂凑值
是否一致 md5sum –c hashfile

0029 一次解压多个zip文件(陈绪)
unzip "*"，注意引号不能少

0030 看pdf文件(陈绪)
使用xpdf或者安装acrobat reader for linux

0031 查找权限位为S的文件(陈绪)
find . -type f /( -perm -04000 -o -perm -02000 /) -exec ls -lg {} /;

0032 装中文输入法(陈绪，hutuworm)
以redhat8为例，xwindow及其终端下的不用说了，缺省就安装了，用ctrl-space呼出。
现在讨论纯console，请到http://zhcon.sourceforge.net/下载zhcon-0.2.1.tar.gz，放在任一目录中，tar xvfz zhcon-0.2.1.tar.gz，cd zhcon-0.2.1，./configure，make，make install。安装结束，要想使用，请运行zhcon，想退出，运行exit。

0033 把弹出的光盘收回来(beike)
#eject －t

0034 cd光盘做成iso文件(弱智)
cp /dev/cdrom xxxx.iso

0035 快速观看开机的硬件检测(弱智)
dmesg | more

0036 查看硬盘的使用情况(陈绪)
df -k 以K为单位显示
df -h 以人性化单位显示，可以是b,k,m,g,t..

0037 查看目录的大小(陈绪)
du -sh dirname
-s 仅显示总计
-h 以K、M、G为单位，提高信息的可读性。KB、MB、GB是以1024为换算单 位， -H以1000为换算单位。

0038 查找或删除正在使用某文件的进程(wwwzc)
fuser filename
fuser -k filename

0039 安装软件(陈绪)
rpm -ivh aaa.rpm
tar xvfz aaa.tar.gz; cd aaa; ./configure; make; make install

0040 字符模式下设置/删除环境变量(陈绪)
bash下
设置：export 变量名=变量值
删除：unset 变量名
csh下
设置：setenv 变量名 变量值
删除：unsetenv 变量名

0041 ls如何看到隐藏文件(即以.开头的文件)(双眼皮的猪)
ls -a
l. (适用于redhat)

0042 rpm中的文件安装到哪里去了(陈绪)
rpm -qpl aaa.rpm

0043 使用src.rpm(陈绪)
rpmbuild --rebuild *.src.rpm

0044 vim中显示颜色或不显示颜色(陈绪，sakulagi)
首先确保安装了vim-enhanced包，然后，vi ~/.vimrc;  如果有syntax on，则显示颜色，syntax off，则不显示颜色。
另外，关于vi的syntax color，还有一点是终端类型（环境变量TERM）的设置。比如通常要设置成xterm或xterm-color才能使用syntax color。尤其是从Linux远程登陆到其他的Unix上。

0045 linux是实时还是分时操作系统(陈绪)
分时

0046 make bzImage -j的j是什么意思(wind521)
-j主要是用在当你的系统硬件资源比较大的时候，比较富裕的时候，用这个可以来加快编译的速度，如-j 3      

0047 源码包怎么没有(陈绪)
你没有安装源代码，你把你光盘上rpm -i *kernel*source*.rpm装上，就可以看到你的源代码了。

0048 修改系统时间(陈绪，laixi781211，hutuworm)
date -s “2003-04-14 cst”，cst指时区，时间设定用date -s 18:10
修改后执行clock -w 写到CMOS
hwclock --systohc
set the hardware clock to the current system time

0049 开机就mount上windows下的分区(陈绪)
自动将windows的d盘挂到/mnt/d上，用vi打开/etc/fstab，加入以下一行
/dev/hda5 /mnt/d vfat defaults,codepage=936,iocharset=cp936 0 0
注意，先得手工建立一个/mnt/d目录

0050 linux怎么用这么多内存(陈绪)
为了提高系统性能和不浪费内存，linux把多的内存做了cache，以提高io速度

0051 FSTAB 最后的配置项里边最后两个数字是什么意思(lnx3000)
第一个叫fs_freq,用来决定哪一个文件系统需要执行dump操作，0就是不需要；
第二个叫fs_passno,是系统重启时fsck程序检测磁盘的顺序号
1 是root文件系统，2 是别的文件系统。fsck按序号检测磁盘，0表示该文件系统不被检测
dump 执行ext2的文件系统的备份操作
fsck 检测和修复文件系统 

0052 linux中让用户的密码必须有一定的长度,并且符合复杂度(eapass)
vi /etc/login.defs，改PASS_MIN_LEN

0053 linux中的翻译软件(陈绪，hutuworm)
星际译王 xdict
console下还有个dict工具，通过DICT协议到dict.org上查11本字典，例如：dict RTFM

0054 不让显示器休眠(陈绪)
setterm -blank 0
setterm -blank n (n为等待时间)

0055 用dat查询昨天的日期(gadfly)
date --date='yesterday'

0056 xwindow下如何截屏(陈绪)
Ksnapshot或者gimp

0057 解压小全(陈绪,noclouds)
tar -I或者bunzip2命令都可以解压.bz2文件
tar xvfj example.tar.bz2
tar xvfz example.tar.gz
tar xvfz example.tgz
tar xvf example.tar
unzip example.zip
tar -jvxf some.bz，就是把tar的zvxf 改成jvxf
zip/tar rh8下有一个图形界面的软件file-roller可以做这件事。另外可以用unzip *.zip解开zip文件，unrar *.rar解开rar文件，不过unrar一般系统不自带，要到网上下载。
# rpm2cpio example.rpm │ cpio -div
# ar p example.deb data.tar.gz | tar zxf -
Alien提供了.tgz, .rpm, .slp和.deb等压缩格式之间的相互转换：
http://sourceforge.net/projects/alien
sEx提供了几乎所有可见的压缩格式的解压接口：
http://sourceforge.net/projects/sex

0057-2 tar压缩、解压用法（platinum）
解压：x
压缩：c
针对gz：z
针对bz2：j
用于显示：v

解压实例
gz文件：tar xzvf xxx.tar.gz
bz2文件：tar xjvf xxx.tar.bz2

压缩实例
gz文件：tar czvf xxx.tar.gz /path
bz2文件：tar cjvf xxx.tar.bz2 /path

0058 在多级目录中查找某个文件的方法(青海湖)
find /dir -name filename.ext 
du -a | grep filename.ext 
locate filename.ext

0059 不让普通用户自己改密码(myxfc)
[root@xin_fc etc]# chmod 511 /usr/bin/passwd 
又想让普通用户自己改密码
[root@xin_fc etc]# chmod 4511 /usr/bin/passwd 

0060 显卡实在配不上怎么办(win_bigboy)
去http://www.redflag-linux.com/ ，下了xfree86 4.3安装就可以了.

0061 超强删除格式化工具(弱智)
比PQMagic安全的、建立删除格式化的小工具：sfdisk.exe for msdos
http://www.wushuang.net/soft/sfdisk.zip

0062 如何让xmms播放列表里显示正确的中文(myxfc)
-*-*-*-*-*-iso8859-1,-misc-simsun-medium-r-normal--12-*-*-*-*-*-gbk-0,*-r-
把这个东西完全拷贝到你的字体里面
操作方法:
右键单击xmms播放工具的任何地方
会看到一个"选项",然后选择"功能设定"选择"fonts"
然后把上面的字体完整的拷贝到"播放清单"和 "user x font

0063 redhat linux中播放mp3文件(hehhb)
原带的xmms不能播放MP3(无声)，要安装一个RPM包：rpm -ivh xmms-mp3-1.2.7-13.p.i386.rpm。打开xmms，ctl-p，在font栏中先在上半部的小框内打勾，再选择 “fixed(misc) gbk-0 13”号字体即可显示中文歌曲名。在音频输出插件中选择 "开放音频系统驱动程序 1.2.7 [lioOSS.so]，即可正常播放MP3文件。

0064 安装中文字体(hehhb)
先下载 http://freshair.netchina.com.cn/~George/sm.sh
(参考文献: http://www.linuxeden.com/edu/doctext.php?docid=2679)
SimSun18030.ttc在微软网站可下载，http://www.microsoft.com/china/windows2000/downloads/18
030.asp　它是个msi文件，在 mswindows中安装用的，装好后在windows目录下的fonts
目录里面就可以找到它。把simsun.ttc，SimSun18030.ttc，tahoma.ttf，tahomabd.ttf
拷贝到/usr/local/temp，然后下载的shell文件也放到这个目录里，然后打开终端
cd /usr/local/temp
chmod 755 sm.sh
./sm.sh

0065 装载windows分区的FAT32、FAT16文件系统(hehhb，NetDC)
以root身份进入KDE，点击桌面上的“起点”图标，在/mnt目录下建立如下文件夹：c,d,e,f,g,usb.分别用作windows下各分区和usb闪盘。
用文本编辑器打开/etc/fstab 文件.加入如下:
/dev/hda1 /mnt/c vfat iocharset=gb2312,umask=0,codepage=936 0 0
/dev/hda5 /mnt/d vfat iocharset=gb2312,umask=0,codepage=936 0 0
/dev/hda6 /mnt/e vfat iocharset=gb2312,umask=0,codepage=936 0 0
/dev/hda7 /mnt/f vfat iocharset=gb2312,umask=0,codepage=936 0 0
/dev/hda8 /mnt/g vfat iocharset=gb2312,umask=0,codepage=936 0 0
/dev/cdrom /mnt/cdrom udf,iso9660 noauto,iocharset=gb2312,owner,kudzu,ro 0 0
/dev/sda1 /mnt/usb vfat iocharset=gb2312,umask=0,codepage=936 0 0
存盘退出. 重新启动后即可正常访问FAT32或FAT16格式分区,解决显示WINDOWS分区下和光盘中文文件名乱码
问题。其中共六列，每列用Tab键分开。注意此方法只能mount上Fat 分区格式，sda1是闪盘。
另外，如果还出现乱码，可以改为iocharset=utf8。 

0066 在X下使用五笔和拼音,区位输入法(hmkart)
从http://www.fcitx.org/上下载fcitx的rpm包安装即可

0067 在Linux下如何解压rar文件(hmkart)
http://www.linuxeden.com/download/softdetail.php?softid=883
下载rar for Linux 3.2.0，解压开后make
然后可以用unrar e youfilename.rar解压rar文件

0068 硬盘安装后怎么添加/删除rpm包(sakulagi)
redhat-config-packages --isodir=<PATH>
可以指定iso文件所在的目录

0069 字符下控制音量(grub007，天外闲云)
使用aumix。另外，要保存oss的音量大小，步骤为：
1、用aumix将音量调整为你们满意的音量
2、用root用户进入/usr/lib/oss下(oss的默认安装目录)
3、执行./savemixer ./mixer.map
4、ok，以后oss开启之后就是你在第一步调整的音量了。
ps:阅读该目录下的README可以得到更多的有用信息。

0070 用dd做iso(grub007)
dd if=/dev/cdrom of=/tmp/aaa.iso

0071 删除几天以前的所有东西(包括目录名和目录中的文件)(shally5)
find . -ctime +3 -exec rm -rf {} /;
或
find ./ -mtime +3 -print|xargs rm -f -r

0072 用户的crontab在哪里(hutuworm)
/var/spool/cron/下以用户名命名的文件

0073 以不同的用户身份运行程序(陈绪)
su - username -c "/path/to/command"
有时候需要运行特殊身份的程序, 就可以让su来做

0074 如何清空一个文件(陈绪)
> filename

0075 为什么OpenOffice下不能显示中文(allen1970)
更改字体设置
tools->options->font replacement
Andale Sans UI -> simsun

0076 如何备份Linux系统(Purge)
Symantec Ghost 7.5以后的版本支持Ext3 native复制 

0077 linux上的partition magic(wwwzc)
Linux下一个有用的分区工具: parted
可以实时修改分区大小, 删除/建立分区.     

0078 /proc/sys/sem中每项代表什么意思? (sakulagi)
/proc/sys/sem内容如下
250 32000 32 128
这4个参数依次为SEMMSL(每个用户拥有信号量最大数量),SEMMNS(系统信号量最大数量),SEMOPM(每次semop系统调用操作数),SEMMNI(系统信号量集最大数量) 

0079 Grub 引导菜单里 bigmem smp up 都是什么意思？(lnx3000)
smp: （symmetric multiple processor）对称多处理器模式
bigmem: 支持1G 以上内存的优化内核
up:（Uni processor） 单处理器的模式

0080 Oracle的安装程序为什么显示乱码？(lnx3000)
现在Oracle的安装程序对中文的支持有问题，只能使用英文界面来安装，在执行runinstaller之前，执行：export LANG=C;export LC_ALL=C     

0081 linux下文件和目录的颜色代表什么(sakulagi,弱智)
蓝色表示目录；绿色表示可执行文件；红色表示压缩文件；浅蓝
色表示链接文件；灰色表示其它文件；红色闪烁表示链接的文件有问题了；黄色是设备文件，包括block, char, fifo。
用dircolors -p看到缺省的颜色设置，包括各种颜色和“粗体”，下划线，闪烁等定义。     

0082 查看有多少活动httpd的脚本(陈绪)
#!/bin/sh
while (true)
do
pstree |grep "*/[httpd/]$"|sed 's/.*-/([0-9][0-9]*/)/*/[httpd/]$//1/'
sleep 3
done

0083 如何新增一块硬盘(好好先生)
一、关机，物理连接硬盘
如果是IDE硬盘，注意主、从盘的设置；如果是SCSI硬盘，注意选择一个没有被使用的ID号。
二、开机，检查硬盘有没有被linux检测到
dmesg |grep hd*(ide硬盘)
dmesg |grep sd*(SCSI硬盘)
或者 less /var/log/dmesg
如果你没有检测到你的新硬盘，重启，检查连线，看看bios有没有认出它来。
三、分区
你可以使用fdisk，Sfdisk或者parted（GNU分区工具,linux下的partition magic)
四、格式化
mkfs
五、修改fstab
vi /etc/fstab

0084 linux下怎么看分区的卷标啊 (q1208c)
e2label /dev/hdxn, where x=a,b,c,d....; n=1,2,3...     

0085 RH8,9中安装后如何添加新的语言包(好好先生)
一.8.0中
1.放入第一张光盘
2.cd /mnt/cdrom/Redhat/RPMS
3.rpm -ivh ttfonts-ZH_CN-2.11-29.noarch.rpm(简体中文,你可以用tab键来补齐后面的部分,以免输入有误)
4.rpm -ivh ttfonts-ZH_TW-2.11-15.noarch.rpm(繁体中文)
如果你还想装日文、韩文,试试第二张光盘上的ttfonts*.rpm.
二.9.0中
9.0不在第一张盘上,在第三张盘上.rpm包名分别为:
ttfonts-zh_CN-2.12-1.noarch.rpm(简体中文)
ttfonts-zh_TW-2.11-19.noarch.rpm (繁体中文)

0086 终端下抓屏(tsgx)
cat /dev/vcsX >screenshot 其中，X表示第X个终端
还可以运行script screen.log，记录屏幕信息到screen.log里。一会记录到你exit为此。这也是抓屏的好方法。
这是在debian的cookbook上看到的。在RH9上能用。没有在其它的系统上测试过。

0087 让一个程序在退出登陆后继续运行(NetDC，双眼皮的猪)
#nohup 程序名 &
或者使用disown命令也可以

0088 man命令不在路径中，如何查看非标准的man文件(陈绪)
nroff -man /usr/man/man1/cscope.1 | more

0089 cp时显示进度(陈绪)
cp -r -v dir1 dir2
cp -a -d -v dir1 dir2

0090 编辑/etc/inittab后直接生效(陈绪)
#init q

0091 让linux连续执行几个命令，出错停止(陈绪)
command1 && command2 && command3

0092 如何将grub安装到mbr(陈绪, NetDC)
grub> root (hd0, 0)
grub> setup (hd0)
也可以用#grub-install /dev/hda来安装grub。

0093 安装时把grub(lilo)写到linux分区的引导区还是主引导扇区(MBR)(陈绪)
如果你想电脑一启动就直接进入操作系统启动菜单就把grub(lilo)写到MBR上，如果写到linux分区的引导区则要用引导盘引导。建议写到 MBR，方便点，至于说写到MBR不安全，该怎么解释呢？每装一次win98，MBR都会被修改一次，大家觉得有什么不安全的吗？

0094 如何让多系统共存(陈绪)
98系统的话用lilo(grub)引导，2k/nt则使用osloader引导多系统

0095 如何在图形界面和控制台（字符界面）之间来回切换(陈绪)
a.图形界面到控制台：Ctr+Alt+Fn(n=1,2,3,4,5,6)。
b.各控制台之间切换：Alt+Fn(n=1,2,3,4,5,6)。
c.控制台到图形：Alt+F7

0096 Redhat linux常用的命令(陈绪)
<1>ls：列目录。
用法：ls或ls dirName，参数：-a显示所有文件，-l详细列出文件。
<2>mkdir：建目录。
用法：mkdir dirName，参数：-p建多级目录，如：mkdir a/b/c/d/e/f -p
<3>mount：挂载分区或镜像文件(.iso,.img)文件。
用法：
a.磁盘分区：mount deviceName mountPoint -o options，其中deviceName是磁盘分区的设备名，比如/dev/hda1,/dev/cdrom,/dev/fd0，mountPoint 是挂载点，它是一个目录，options是参数，如果分区是linux分区，一般不用-o options，如果是windows分区那options可以是iocharset=cp936，这样windows分区里的中文文件名就能显示出来了。用例：比如/dev/hda5是linux分区，我要把它挂到目录a上（如没目录a那就先mkdir a），mount /dev/hda5 a，这样目录a里的东西就是分区hda5里的东西了，比如hda1是windows分区，要把它挂到b上，mount /dev/hda1 b -o iocharset=cp936。
b.镜像文件：mount fileName mountPoint -o loop，fileName是镜像文件名(*.iso,*.img)，其它的不用说了，跟上面一样。用例：如我有一个a.iso光盘镜像文件，mount a.iso a -o loop，这样进入目录a你就能浏览a.iso的内容了，*.img文件的用法一样。
<4>find：查找文件。
用法：find inDir -name filename，inDir是你要在哪个目录找，filename是你要找的文件名(可以用通配符)，用通配符时filename最好用单引号引起来，否则有时会出错，用例：find . -name test*，在当前目录查找以test开头的文件。
<5>grep：在文件里查找指定的字符串。
用法：grep string filename，在filename(可用通配符)里查找string(最好用双引号引起来)。参数：-r在所有子目录里的filename里找。用例：grep hello *.c -r在当前目录下（包括子目录）的所有.c文件里查找hello。
<5>vi：编辑器。
用法：vi filename。filename就是你要编辑的文本文件。用了执行vi filename后，你可能会发现你无法编辑文本内容，不要着急，这是因为vi还没进入编辑状态，按a或i就可以进入编辑状态了，进入编辑状态后你就可以编辑文本了。要退出编辑状态按Esc键就可以了。以下操作均要在非编辑状态下。查找文本：输入/和你要查找的文本并回车。退出：输入: 和q并回车，如果你修改了文本，那么你要用:q!回车才能退出。保存：输入: w回车，如果是只读文件要用: w!。保存退出：输入: wq回车，如果是只读就: wq!回车。取消：按u就可以了，按一次就取消一步，可按多次取消多步。复制粘贴一行文本：把光标移到要复制的行上的任何地方，按yy（就是连按两次 y），把光标移到要粘贴地方的上一行，按p，刚才那行文本就会被插入到光标所在行的下一行，原来光标所在行后面所有行会自动下移一行。复制粘贴多行文本：跟复制一行差不多，只是yy改成先输入要复制的行数紧接着按yy，后面的操作一样。把光标移到指定行：输入:和行号并回车，比如移到123行:123回车，移到结尾:$回车。

0097 linux文本界面下如何关闭pc喇叭(labrun)
将/etc/inputrc中的set bell-style none 前的＃去掉，或echo "set bell-style none" >> ~/.bashrc

0098 重装windows导致linux不能引导的解决办法(好好先生)
如果没有重新分区，拿linux启动盘(或者第一张安装光盘)引导，进入rescue模式。首先找到原来的/分区mount在什么地方。redhat通常是/mnt/sysimage. 执行"chroot /mnt/sysimage". 如果是grub，输入grub-install /dev/hd*(根据实际情况)；如果是lilo，输入lilo -v，然后重新启动。如果分区有所改变，对应修改/etc/lilo.conf和/boot/grub/grub.conf然后再执行上述命令。

0099 为什么装了LINUX后win2K很慢(lnx3000，好好先生)
老问题了，你在2000是不是能看见Linux的逻辑盘，但不能访问？
在磁盘管理里，选中这个盘，右击->更改"驱动器名和路径"->"删除"就可以了，注意不是删除这个盘!

0100 将linux发布版的iso文件刻录到光盘的方法(陈绪)
借用windows中的nero软件，选择映象文件刻录，选择iso文件，刻录即可！123[url=111]111[/url]


1101 linux中刻录iso的方法(hutuworm)
方法一：使用xcdroast，选择制作光碟，选择ISO文件，刻录!
参见http://www.xcdroast.org/xcdr098/faq-a15.html#17
方法二：找刻录机的命令：
cdrecord --scanbus
输出结果为：
0,0,0 0) 'ATAPI ' 'CD-R/RW 8X4X32 ' '5.EZ' Removable CD-ROM
刻录的命令：
cdrecord -v speed=8 dev=0,0,0 hutuworm.iso
方法三：使用k3b可以刻录CD/DVD 
k3b主页：http://www.k3b.org/ 
(实际上k3b是个图形界面，刻录CD利用了cdrecord，刻录DVD利用了dvd+rw-tools http://fy.chalmers.se/~appro/linux/DVD+RW/ )

1102 屏幕变花时怎么办(双眼皮的猪)
当您一不小心cat了一个并不是文本的文件的时候，这时屏幕会变花，那么您可以按两下"Enter"键，再敲"reset"，那么屏幕就恢复正常了....

1103 卸载软件包时如何得知具体包名(diablocom)
大家知道删除软件包的命令是rpm -e XXX，但是当我们不知道这个XXX的确切拼写时，可以用rpm -q -a查询所有安装的软件包或者用rpm -qa |grep xxxx查询出名字

1104 使用内存作linux下的/tmp文件夹(yulc)
在/etc/fstab中加入一行：
none /tmp tmpfs default 0 0
或者在/etc/rc.local中加入
mount tmpfs /tmp -t tmpfs -o size=128m
注：size=128m 表示/tmp最大能用128m
不管哪种方式，只要linux重启，/tmp下的文件全部消失

1105 用ls只列出目录(yulc)
ls -lF | grep ^d
ls -lF | grep /$
ls -F | grep /$

1106 在命令行下列出本机IP地址，而不是得到网卡信息(yulc)
ifconfig |grep "inet" |cut -c 0-36|sed -e 's/[a-zA-Z: ]//g'
hostname -i

1107 修改/etc/profile或者$HOME/.profile文件后如何立即生效(peter333)
#source /etc/profile (或者source .profile)

1108 bg和fg的使用(陈绪)
输入ctrl+z，当前一个任务会被挂起并暂停， 同时屏幕上返回进程号，此时用 "bg %进程号"，会把这个进程放到后台执行，而用" fg %进程号 "就能让这个进程放到前台来执行。另外，job命令用来查看当前的被bg的进程

1109 ctrl+s与ctrl+q(陈绪)
ctrl-s用来暂停向终端发送数据的，屏幕就象死了一样，可以用ctrl-q来恢复

1110 目录统计脚本(陈绪)
保存成total.sh，然后用total.sh 绝对路径，就能统计路径下目录的大小了
代码:
#!/bin/sh
du $1 --max-depth=1 | sort -n|awk '{printf "%7.2fM ----> %s/n",$1/1024,$2}'|sed 's:/.*//([^/]/{1,/}/)$:/1:g'

1111 grep不显示本身进程(陈绪)
#ps -aux|grep httpd|grep -v grep
grep -v grep可以取消显示你所执行的grep本身这个进程，-v参数是不显示所列出的进程名

1112 删除目录中含输入关键字的文件(WongMokin)
find /mnt/ebook/ -type f -exec grep "在此输入关键字" {} /; -print -exec rm {} /;

1113 让cron中的任务不回馈信息, 本例5分钟检查一次邮件(WongMokin)
0-59/5 * * * * /usr/local/bin/fetchmail > /dev/null 2>&1

1114 在当前目录下解压rpm文件(陈绪)
cat kernel-ntfs-2.4.20-8.i686.rpm | rpm2cpio | pax -r

1115 合并两个Postscript或PDF文件(noclouds)
$ gs -q -dNOPAUSE -dBATCH -sDEVICE=pswrite /
-sOutputFile=bar.ps -f foo1.ps foo2.ps
$ gs -q -dNOPAUSE -dBATCH -sDEVICE=pdfwrite /
-sOutputFile=bar.pdf -f foo1.pdf foo2.pdf

1116 去掉apache的manual目录中的所有.en的后缀名(陈绪)
进入到manual目录
代码:find ./ -regex .*/.en|awk -F. '{ printf "mv %s.%s.%s.%s %s.%s.%s/n",$1,$2,$3,$4,$1,$2,$3}'|sh

1117 如何起多个X(noclouds)
startx默认以display :0.0起第一个X，通过传递参数给Xserver可以起多个X：
# startx -- :1.0
# startx -- :2.0
...
然后用Ctrl-Alt-F7/F8...切换。

1118 让一个程序在退出登陆后继续运行(noclouds,陈绪)
# <cmd>
# disown
或者是
nohup command &

1119 看Linux启动时屏幕的显示信息(陈绪)
在启动完后用命令dmesg查看

1120 让vi不响铃(sakulagi)
echo "set vb t_vb=" >> ~/.vimrc

1121 让fedora开机后自动login(dzho002)
1) rpm -ihv autologin-1.0.0-7mdk.i586 rpm
2) 建立文件 /etc/sysconfig/autologin
在里面加上一行.
USER = root

1122 如何配置让哪些服务启动(天外闲云，q1208c)
方法1 运行ntsysv或者setup命令，进入菜单进行配置
方法2 chkconfig --list 显示服务
chkconfig name on/off 打开/关闭“name”服务

1123 安全删除linux(天外闲云)
步骤1 Dos下使用fdisk /mbr或者用win2000/xp的光盘启动进入故障恢复控制台，使用命令fixmbr
步骤2 格式化linux分区为windows分区即可。

1124 用grub引导进文本界面(天外闲云)
进入grub之后，按a，输入 空格 3 就可以引导进入文本界面，但是不修改系统的运行级，只在当次有效。

1125 先测试patch是否运行正常，暂不将更改应用到kernel(jiadingjun)
patch --dry-run

1126 redhat和debian上的文件安装删除用法(NetDC)
删除一个软件包：
rpm -e <package-name>
dpkg -r <package-name>
显示一个软件包的内容：
rpm -qvl <package-name.rpm>
dpkg -c <package-name.deb>
显示所有已经安装的软件包：
rpm -qvia
dpkg -l
打印一个包的信息：
rpm -qpi <package-name.rpm>
dpkg -I <package-name.deb>
检验包characteristics：
rpm -Va
debsums -a
检验一个文件属于哪个包：
rpm -qf </path/to/file>
dpkg -S </path/to/file>
安装新软件包：
rpm -Uvh <package-name.rpm>
dpkg -i <package-name.deb>

1127 如何使新用户首次登陆后强制修改密码(猫小)
#useradd -p '' testuser; chage -d 0 testuser

1128 日志维护工具logrotate(hotbox)
在/etc/logrotate.conf中配置，作用：定义log文件达到预定的大小或时间时，自动压缩log文件

1129 Linux中默认的管理员叫什么(陈绪)
root

1130 如何产生一个长度固定（例如文件长度为1M）字节的空文件，即每个字节的值全为0x00(sakulagi)
dd if=/dev/zero of=/tmp/zero_file bs=1024 count=1024

1131 RedHat Linux里修改时间的步骤(hutuworm)
1. 设置你的时区： timeconfig里选择Asia/Shanghai （如果你位于GMT+8中国区域）
2. 与标准时间服务器校准： ntpdate time.nist.gov
2.5 当然，如果你是李嘉诚，也可以跟自己的手表校准： date -s STRING （STRING格式见man date）
3. 写回硬件时钟： hwclock --systohc

1132 查找当前目录下文件并更改扩展名(零二年的夏天)
更改所有.ss文件为.aa
# find ./ -name "*.ss" -exec rename .ss .aa '{}' /;

1133 patch的使用(天才※樱木)
语法是patch [options] [originalfile] [patchfile]
例如：
patch -p[num] <patchfile
-p参数决定了是否使用读出的源文件名的前缀目录信息，不提供-p参数，则忽略所有目录信息，-p0（或者-p 0）表示使用全部的路径信息，-p1将忽略第一个"/"以前的目录，依此类推。如/usr/src/linux-2.4.16/Makefile这样的文件名，在提供-p3参数时将使用linux-2.4.16/Makefile作为所要patch的文件。
对于刚才举的Linux内核源码2.4.16升级包的例子，假定源码目录位于/usr/src/linux中，则在当前目录为/usr/src时使用"patch -p0 <patch-2.4.16"可以工作，在当前目录为/usr/src/linux时，"patch -p1<patch-2.4.16"也可以正常工作。

1134 将file.txt里的123改为456(hutuworm)
方法1 
sed 's/123/456/g' file.txt > file.txt.new
mv -f file.txt.new file.txt
方法2
vi file.txt
输入命令：
:%s/123/456/g


1135 将一个分区格式化为ext3日志文件系统(hutuworm)
mkfs -j /dev/xxxx

1136 开启硬盘ATA66 (laixi781211)
/sbin/hdparm -d1 -X68 -c3 -m16 /dev/hda

1137 查看当前运行级别(双眼皮的猪)
runlevel

1138 查看当前登陆身份(双眼皮的猪)
(1)who am i
(2)whoami
(3)id
注意(1)跟(2)的小区别

1139 删除rpm -e删除不了的包(wwwzc)
1、如果在删除包之前删除了包的目录
rpm -e --noscripts
2、如果系统里一个包被装两次（由于某些异常引起的）
rpm -e multi-installed-pkgs --allmatches

1140 如何定制用户登录时显示的信息(jiadingjun)
在/etc目录下放一个名字叫motd的文本文件实现的，例如，建立自己的/etc/motd: 
$cat /etc/motd 
welcome to my server ! 
那么，当用户登录系统的时候会出现这样的信息： 
Last login: Thu Mar 23 15:45:43 from *.*.*.* 
welcome to my server !

1141 用命令清空Root回收站中的文件(dtedu)
cd /var/.Trash-root 
rm -rf *

1142 在Red Hat上加Simsun.ttc字体(陈绪)
以Red Hat 7.3为例，安装时选取简体中文安装，先复制一个simsun.ttc到/usr/X11R6/lib/X11/font/TrueType，改名为simsun.ttf；然后进入/usr/X11R6/lib/X11/font/TrueType目录下，运行ttmkfdir > fonts.dir命令；接着用vi编辑fonts.dir文件，把有simsun.ttf行修改如下:
simsun.ttf -misc-SimSun-medium-r-normal--0-0-0-0-c-0-ascii-0
simsun.ttf -misc-SimSun-medium-r-normal--0-0-0-0-c-0-iso10646-1
simsun.ttf -misc-SimSun-medium-r-normal--0-0-0-0-p-0-iso8859-15
simsun.ttf -misc-SimSun-medium-r-normal--0-0-0-0-p-0-iso8859-1
simsun.ttf -misc-SimSun-medium-r-normal--0-0-0-0-c-0-gb2312.1980-0 
simsun.ttf -misc-SimSun-medium-r-normal--0-0-0-0-p-0-gb2312.1980-0 
simsun.ttf -misc-SimSun-medium-r-normal--0-0-0-0-m-0-gb2312.1980-0 
simsun.ttf -misc-SimSun-medium-r-normal--0-0-0-0-p-0-gbk-0
接着运行cat fonts.dir > fonts.scale命令，修改/etc/X11/XF86config-4, 在Section“Files”加上下面这一行：
FontPath “/usr/X11R6/lib/X11/fonts/TrueType”
最后回到KDE桌面里, 在“开始”→“选项”→“观感”→“字体”，将所有字体改为Simsun。

1143 Unicon和Zhcon的区别和作用(陈绪)
Unicon是内核态的中文平台，基于修改Linux FrameBuffer和Virtual Console（fbcon）实现的。由于是在系统底层实现的，所以兼容性极好，可以直接支持gpm鼠标。但是相对比较危险，稍有漏洞就可能会危及系统安全。Zhcon是用户态的中文平台，有点像UCDOS。

1144 如何卸载tar格式安装的软件(陈绪)
进入安装该软件的原代码目录，运行make uninstall。如果不行，也可以查看一下Makefile文件，主要是看install部分，从其中找出tar格式的文件被复制到了什么路径，然后进入相应的目录进行删除即可。

1145 定制linux提示符 (陈绪)
在bash中提示符是通过一个环境变量$PS1指定的。用export $PS1查看现在的值，比较直观常用的提示符可以设定为export PS1=“[/u@/h /W]/$”。其中/u代表用户名，/h代表主机名，/W代表当前工作目录的最后一层，如果是普通用户/$则显示$，root用户显示#。

1146 在vi中搜索了一个单词，该单词以高亮显示，看起来很不舒服，怎么能将它去掉(陈绪)
在vi的命令模式下输入:nohlsearch就可以了。另外可以在~/.vimrc中写上下面的语句就会有高亮显示：
set hlsearch
加上下面的语句就不会有高亮显示：
set nohlsearch

1147 如何找出系统中所有的*.cpp、*.h文件(陈绪)
用find命令就可以了。不过如果从根目录查找消耗资源较高，使用下面的命令就可以：
find / -name "*.cpp" -o -name "*.h"

1148 如安装Debian需要几张盘就够了？7张盘全部都要下载吗？(陈绪)
如果经常有网络环境的话，下载第一张就可以了。要是没有网络环境的话不推荐使用Debian，因为Debian主要依赖网络来更新软件。实在要安装的话，要下载全部7张盘，否则可能会出现需要的软件包找不到的问题。

1149 Debian第一张光盘为什么有两个版本？debian-30r1-i386-binary-1.iso和debian-30r1-i386-binary-1_NONUS.iso该下载哪一个呢？它们有什么区别？(陈绪)
因为含有“non-US”（不属美国）的软件不能合法地存放在架设于美国境内的服务器中。以前，其原因通常是因为软件含有严密的密码编码，而今天，则是因为程序使用了美国专利保护的演算法。每个人应该取用“non-US”来供私人用途所用；而没有这个标识的iso则只对架设在美国的镜像及供应商才有用处。其它二进制的光盘则不会含有任何“US-sensitive”（与美国相关的）软件，它们和其它种binary-1光盘一样运作得很好。因此，个人使用还是下载debian-30r1-i386-binary-1_NONUS.iso版本。

1150 为何我使用umount /mnt/cdrom命令的时候出现device is busy这样的语句，不能umount(陈绪)
在使用umount的时候一定要确保已退出/mnt/cdrom这个目录，退出这个目录就可以使用umount /mnt/cdrom了。

1151 我使用的是笔记本电脑，怎么才能在控制台下显示现在还剩多少电量呢？ (陈绪)
使用apm -m就可以看到还有多少分钟了，具体参数可以用man apm查看。

1152 为什么我进入Linux的终端窗口时，man一条命令出来的都是乱码呢？ (陈绪)
这是因为你的字符集设置有问题。临时解决办法可以使用export LANG=“en_US”。要想不必每次都修改的话，在/etc/sysconfig/i18n文件里面修改LANG=“en_US”就可以了。也可以针对某个用户来做，这样就可以改变个人的界面语言，而不影响别的用户。命令如下：# cp /etc/sysconfig/i18n $HOME/.i18n。

1153 编译内核的时候出错，提示“Too many open files”，请问怎么处理 (陈绪)
这是因为file-max默认值（8096）太小。要解决这个问题，可以root身份执行下列命令（或将它们加入/etc/rcS.d/*下的init脚本）：
# echo "65536"  > /proc/sys/
最后进入解压后的目录，运行安装命令。
# cd vmware-linux-tools
# ./install.pl

1154 本来装有Linux与Windows XP，一次将Windows XP重装后，发现找不到Linux与Windows XP的启动选单，请问如何解决(陈绪)
首先光盘启动，进入rescue模式，运行GRUB，进入grub提示符grub>，然后敲入下面的语句，重启就好了。
root (hd0,2)，setup (hd0)

1155 安装了一台Linux服务器，想自己编译内核，一步一步做下来，GRUB也添加进去了，但出现“kernel Panic:VFS:Unable to mount root fs on 0:00”的错误，请问是怎么回事？(陈绪)
一般情况下initrd这个文件在台式机上不是必须的，但是在有SCSI设备的服务器上却是必须的。有可能因为编译内核的时候没有产生initrd那个文件，所以会有上面的错误提示。用户可以使用mkinitrd命令来生成一个initrd.img文件，然后加入GRUB，重启试一试。

1156 如何设置用户登录后的欢迎信息？(陈绪)
修改/etc/motd文件，往里面写入文本，就能使用户通过Telnet正确登录后，执行Shell之前得到相应的提示信息。
motd就是“messages of the day”，也就是当日信息的意思。管理员可以往里面写一些需要注意的事项或通知等来提醒正式用户。

1157 我下载了rcs5.7，用./configure && make && make install时报错如下：./conf.sh: testing permissions ... ./conf.sh: This command should not be run with superuser permissions. 我是以root用户身份登录编译安装的，为什么会这样？(陈绪)
有些软件确实因为考虑到安全等其它原因不能用root用户编译。这时只要用其它用户编译，到make install这步时，如果该软件安装在不属于编译时的用户的主目录下时，需要使用su命令转换为root用户再执行make install。

1158 我在安装USBView时失败，具体情况如下： #rpm -ivh usbview-1.0-9.src.rpm warning:usbview-1.0-9.src.rpm:V3 DSAsignature:NOKEY,key IDab42a60e (陈绪)
这行代码说明安装失败是因为你的系统上没有安装合适的钥匙来校验签名。要使该软件包通过校验，可以通过导入Red Hat的公匙来解决，具体的方式是在Shell下运行如下命令：
#rpm -import /usr/share/rhn/RPM-GPG-KEY
（注意大小写）

1159 如何防止某个关键文件被修改？(陈绪)
在Linux下，有些配置文件是不允许任何人（包括root）修改的。为了防止被误删除或修改，可以设定该文件的“不可修改位(immutable) ”。命令如下：
# chattr +i /etc/fstab
如果需要修改文件则采用下面的命令：
# chattr -i /etc/fstab

1160 怎样限制一个用户可以启动的进程数？(陈绪)
先确定一下/etc/pam.d/login文件中下面一行的存在：
session required /lib/security/pam_limits.so
然后编辑/etc/security/limits.conf，在里面可以设置限制用户的进程数、CPU占用率和内存使用率等，如hard nproc 20就是指限制20个进程，具体可以看man。

1161 如何限制Shell命令记录大小 ？(陈绪)
默认情况下，bash会在文件$HOME/.bash_history中存放多达500条命令记录。有时根据具体的系统不同，默认记录条数不同。系统中每个用户的主目录下都有一个这样的文件。为了系统的安全，在此强烈建议用户限制该文件的大小。用户可以编辑/etc/profile文件，修改其中的选项如下：
HISTFILESIZE=30 或 HISTSIZE=30
这样就将记录的命令条数减少到30条。

1162 我想将开机时显示的信息保留下来，以检查电脑出了问题的地方，请问怎么办？(陈绪)
可输入下面的命令:
#dmesg > bootmessage
该命令将把开机时显示的信息重定向输出到一个文件bootmessage中。

1163 我想在注销时删除命令记录，请问怎么做？(陈绪)
编辑/etc/skel/.bash_logout文件，增加如下行:
rm -f $HOME/.bash_history
这样，系统中的所有用户在注销时都会删除其命令记录。
如果只需要针对某个特定用户，如root用户进行设置，则可只在该用户的主目录下修改/$HOME/.bash_history文件，增加相同的一行即可。

1164 编译内核，支持ntfs的步骤(platinum，陈绪)
1. # cd /usr/src/linux-2.4
2. # make menuconfig
3. 选中File System下的NTFS file system support (read only)为M
4. # uname -a
2.4.21-27.0.2.EL
5. # vi Makefile
确保前几行为
VERSION = 2
PATCHLEVEL = 4
SUBLEVEL = 21
EXTRAVERSION = -27.0.2.EL
6. # make dep
7. # make modules SUBDIRS=fs/ntfs
8. # mkdir /lib/moduels/2.4.21-27.0.2.EL/kernel/fs/ntfs
9. # cp -f fs/ntfs/*.o /lib/moduels/2.4.21-27.0.2.EL/kernel/fs/ntfs/
10. # depmod -a
11. # modprobe ntfs
12. # lsmod
确保有ntfs在里面

1165 如何使用ssh通道技术(陈绪)
本文讨论所有机器均为Linux操作系统。
比如说我的机器是A，中间服务器为B，目标服务器是C。
从A可以ssh到B，从B可以ssh到C，但是A不能直接ssh到C。
现在展示利用ssh通道技术从A直接传输文件到C。
1. ssh -L1234:C:22 root@B
input B's password
2. scp -P1234 filename root@localhost:
input C's password

1166 使用rpm命令时没有任何响应，如何解决 (初学摄影)
rm -rf /var/lib/rpm/__db.*

1167 向登陆到同一台服务器上的所有用户发一条信息(陈绪)
1)输入wall并回车
2)输入要发送的消息
3)结束时按“Control-d”键,消息即在用户的控制窗口中显示

1168 输入短消息到单个用户(陈绪)
1)输入write username，当用户名出现在多个终端时，在用户名后可加tty,以表示在哪个tty下的用户。
2)输入要发送的消息。
3)结束时按“Control-d”键,消息即在用户的控制窗口中显示。
4）对于接收消息方，可以设定是否允许别人送消息给你。
指令格式为：mesg n[y]
%write liuxhello! Everybody, I’llcome.
%
用户控制窗口中显示的消息:Message from liux on ttyp1 at 10:00…hello! Everybody, I’llcome.EOF
当使用CDE或OpenWindows等窗口系统时，每个窗口被看成是一次单独的登录；如果用户登录次数超过一次则消息直接发送到控制窗口。

1169 发送文件中的消息到单个用户(陈绪)
如果有一个较长的消息要发送给几个用户，用文件方式：
1)创建要发送的消息文本的文件filename.
2)输入write username<filename回车，用cat命令创建包含短消息的文件：
% cat >messagehello! Everybody, I’llcome.
% write liux<messagewrite:liux logged in more than once…write to console
% 用户在一个以上窗口登录，消息显示在控制窗口中Message from liux on ttyp1 at 10:00…hello! Everybody, I’llcome.EOF 

1170 向远程机器上的所有用户发送消息(陈绪)
使用rwall(向所有人远程写)命令同时发送消息到网络中的所有用户。
rwall hostname file
当使用CDE或OpenWindows等窗口系统时,每个窗口被看成是一次单个的登录;
如果用户登录次数超过一次则消息直接发送到控制窗口。

1171 向网络中的所有用户发送消息(陈绪)
发送消息到网络中的所有用户
1)输入rwall -n netgroup并回车
2)输入要发送的消息
3)结束时按“Control-d”键，消息即在系统每个用户的控制窗口中显示，下面是系统管理员发消息到网络组Eng每个用户的例子：
% rwall -n EngSystem will be rebooted at 11:00.(Control-d)
%
用户控制窗口中的消息:Broadcast message from root on console…System will be rebooted at 11:00.EOF
注意：也可以通过rwall hostname（主机名）命令到系统的所有用户。

1172 我需要编译内核，内核源码在哪里？(platinum)
1、一般在发行版的盘里都有，比如 RedHat，一般在第二、第三张上
　　2.4 内核的叫 kernel-source-2.4.xx-xx.rpm
　　2.6 内核的叫 kernel-devel-2.6.xx-xx.rpm
2、去 www.kernel.org 下载一份你喜欢的

1173 将top的结果输出到文件中(bjweiqiong)
top -d 2 -n 3 -b >test.txt
可以把top的结果每隔2秒，打印3次，这样后面页的进程也能够看见了

1174 vim中改变全文大小写的方法(陈绪)
光标放在全文开头
gUG 所有字母变大写
guG 所有字母变小写
g~G 所有字母，大写变小写，小写变大写
2001 让apache的默认字符集变为中文(陈绪)
vi httpd.conf，找到 AddDefaultCharset ISO-8859-1 一行
apache版本如果是1.*，改为 AddDefaultCharset GB2312
如果是2.0.1-2.0.52，改为 AddDefaultCharset off
然后运行/etc/init.d/httpd restart重启apache即可生效。
注意：对于2.0.53以上版本，不需要修改任何配置，即可支持中文。

2002 永久更改ip(陈绪)
ifconfig eth0 新ip
然后编辑/etc/sysconfig/network-scripts/ifcfg-eth0，修改ip

2003 从Linux上远程显示Windows桌面(lnx3000)
安装rdesktop包

2004 手动添加默认网关(陈绪)
以root用户，执行: route add default gw 网关的IP
想更改网关
1 vi /etc/sysconfig/network-scripts/ifcfg-eth0
更改GATEWAY
2 /etc/init.d/network restart

2005 redhat 8.0上msn和qq(陈绪)
下载Gaim 0.58版：
gaim-0.58-2.i386.rpm
下载QQ插件 for gcc2.9版：
libqq-0.0.3-ft-0.58-gcc296.so.gz
将下载的文件放入/temp目录，然后将系统中已有的Gaim删除，即在终端仿真器中键入命令：rpm -e gaim。
开始安装
打开终端仿真器，继续执行下列命令安装Gaim 0.58版，即：
cd /temp　　　　　　　　　(进入temp目录)
rpm -ivh gaim-0.58-2.i386.rpm　(安装软件)
当安装成功后，你就可以在GNOME或KDE桌面建立Gaim图标了。
继续安装QQ插件，即键入命令：
gunzip libqq-0.0.3-ft-0.58-gcc296.so.gz (解压缩文件）
cp libqq-0.0.3-ft-0.58-gcc296.so /usr/lib/gaim (复制插件到gaim库目录中)
软件设置
首次启动Gaim 0.85版时，会出现的登录界面。先选择“插件”，在插件对话框中点击“加载”，分别将libmsn.so和libqq-0.0.3-ft-0.58-gcc296.so文件装入，确认后关闭。然后再选择“所有帐号”，在出现的帐号编辑器中继续点击“增加”，当出现的修改帐号页面时，我们就可以输入自己的QQ或MSN号了，登录名填写QQ号码或MSN邮箱，密码填写对应的QQ或MSN密码，Alias填写自己的昵称，协议选择相应的QQ或MSN，其他的设置按默认的即可。当全部设置完成后就可以登录使用了。
由于MS对msn的协议经常升级，导致linux上的gaim和msn插件必须升级，目前尚无万无一失的解决方案，请见谅

2006 查出22端口现在运行什么程序(陈绪)
lsof -i

2007 查看本机的IP，gateway, dns(陈绪)
IP：
以root用户登录，执行ifconfig。其中eth0是第一块网卡，lo是默认的设备
Gateway:
以root用户登录，执行netstat -rn，以0.0.0.0开头的一行的Gateway即为默认网关
也可以查看/etc/sysconfig/network文件，里面有指定的地址！
DNS：
more /etc/resolv.conf，内容指定如下：
nameserver 202.96.69.38
nameserver 202.96.64.38

2008 RH8.0命令行下改变ping 的TTL值(cgweb，lnx)
方法1(重启后有效)：
#sysctl -w net.ipv4.ip_default_ttl=N
(N=0~255),若N>255,则ttl=0
方法2(重启后无效)：
#echo N(N为0～255) > /proc/sys/net/ipv4/ip_default_ttl

2009 开启LINUX的IP转发(houaq)
编辑/etc/sysctl.conf, 例如，将
net.ipv4.ip_forward = 0
变为
net.ipv4.ip_forward = 1
重启后生效，用sysctl -a查看可知

2010 mount局域网上其他windows机器共享出的目录(陈绪)
mount -t smbfs -o username=guest,password=guest //machine/path /mnt/cdrom

2011 允许｜禁止root通过SSH登陆(Fun-FreeBSD)
修改sshd_config:PermitRootLogin no|yes

2012 让root直接telnet登陆(陈绪，platinum)
方法1：
编辑/etc/pam.d/login，去掉
auth required /lib/security/pam_securetty.so 这句话
方法2：
vi /etc/securetty
添加
pts/0
pts/1
...

2013 在linux接adsl设备(wind521)
需要一个运转正常的Linux + 至少一块网卡 + 宽带设备已经申请完毕，同时已经开通。目前市场上大概有几种ADSL设备，他们工作的方式有一些细微的差别。
就是通过虚拟拨号来完成上网的这一过程，也就是利用pppoe设备来进行虚拟拨号的叫作全向猫，就是一种加电后自动的进行拨号的工作，然后留给我们的接口是RJ45，大连地区一般留给我们的网关都是10.0.0.2,这种设备最容易对付，最后是直接分配给用户一个固定的IP，相对大家来说也比较容易对付
1.第一种需要进行拨号：
这几种设备都是通过eth接口与计算机进行通讯的，所以先将硬件设备的连接作好，尤其是宽带猫的，一定要确认无误（否则一会儿要不去可不算我的事情）
然后启动系统，确认系统上是否安装rp-pppoe这个软件（通过rpm -qa|grep pppoe来查找），如没有安装的用户，在光盘里或是到网上去down一个来，安装上后，以root用户执行adsl-setup，这样就进入了adsl的资料的设定状态，要求输入申请宽带的用户名以及其他一些信息，确认没有问题，接受直至最后（里面都是E文，但是一看即能懂，比较简单，有关一个防火墙的设置，我一般都不用，选0，大家可以具体考虑）。
配置完成后，以root用户执行adsl-start，这样将进行adsl的拨号工作，正常就会一下上线，如有什么具体问题，去看一下日志（/var/log/messages）里面告诉你什么了。
停掉adsl，执行adsl-stop就可以了（很简单的）
2.另外两种比较容易对付：
  全向猫：只要将你的网卡的IP设置成一个10网段的IP，然后网关指到全向猫的IP，上（10.0.0.2)，基本上不有太大的问题
　固定IP：就像配置本地儿的网卡一样，将IP，网关，DNS都按申请来的填写上就可以搞定了

2014 让linux自动同步时间(shunz)
vi /etc/crontab
加上一句：
00 0 1 * * root rdate -s time.nist.gov

2015 linux的网上资源有哪些(陈绪)
国外
http://lwn.net/
http://www.tldp.org/
http://www.yolinux.com/(flying-dance big big pig)
http://www.justlinux.com/
http://www.linuxtoday.com/
http://www.linuxquestions.org/
http://www.fokus.gmd.de/linux/
http://www.linux-tutorial.info/
[url]http://public.www.planetmirror.com/[/url]
http://www.freebsdforums.org/forums/
http://www.netfilter.org/documentation/
http://www-106.ibm.com/developerworks/linux/

国内
http://www.linuxmine.com/
http://www.fanqiang.com/
http://www.linuxsir.com/
http://www.chinaunix.net/
http://www.linuxfans.org/(deadcat)
http://www.linuxeden.com/
http://www.linuxforum.net/
http://www.linuxaid.com.cn/
http://freesoft.online.sh.cn/
http://www-900.ibm.com/developerWorks/cn/linux/index.shtml
http://www.neweasier.com/software.html
http://www.blueidea.com/bbs/archivecontent.asp?id=635906(sqh)
http://westlinux.ywzc.net/(onesun)

2016 改变sshd的端口(陈绪)
在/etc/ssh/sshd_config中加入一行：Port 2222，/etc/init.d/sshd restart重启守护进程

2017 改变telnet的端口(陈绪)
将/etc/services文件中telnet对应的端口号21改为你想要的值，/etc/init.d/xinetd restart重启守护进程

2018 终端模式有问题(sakulagi)
export TERM=vt100

2019 模仿超级终端，LINUX里什么程序连接路由器和交换机(alstone)
minicom

2020 ssh上来能不能不自动断线(wind521，双眼皮的猪)
修改自己HOME目录下的.bash_profile文件，加上
export TMOUT=1000000  (以秒为单位)
然后运行source .bash_profile

2021 用什么工具做入侵检测(陈绪)
snort

2022 Linux下检测程序内存泄漏的工具(陈绪)
cchecker或是efence库都可以

2023 linux下如何监视所有通过本机网卡的数据(陈绪)
tcpdump或者iptraf

2024 为什么root执行好多命令都说command not found(陈绪)
你是telnet上来，然后su成root的吧，改改你的su命令格式，应该是su - root

2025 关闭用户的POP3权限(tiansgx)
把POP3的端口关了就可以了。 在文件/etc/services中找到这一行 pop-3 110/tcp 把这一行前加个'#',把它注释掉就可以了。

2026 linux下播放flash动画(myxfc)
linux下播放flash动画用这个东西，不会造成浏览器的关闭(其他的插件不好用）
首先下载flash播放动画在linux的插件
http://www.collaborium.org/onsit ... /flash_linux.tar.gz
tar zxvf flash_linux.tar.gz
打开包之后,会看到Linux文件夹
在linux文件颊里有两个文件libflashplayer.so 和shockwaveflash.class,把这两个文件拷贝到你的浏览器里的插件里(浏览器不一样,插件的位置可能也不一样)
/usr/lib/mozilla-1.0.1/plugins,就可以了 

2027 锁定wu-ftp用户目录(wangla)
编辑ftpaccess文件
restricted-uid *
这一句很重要，限制了ftp用户在自己的目录里。

2028 服务器怎么不让telnet(知秋一叶)
服务器上必须启动telnet服务 && 服务器的防火墙优先级应该设为低

2029 防止任何人使用su命令成为root(xiaohu0)
1.vi /etc/pam.d/su
auth sufficient /lib/security/pam_rootok.so debug
auth required /lib/security/pam_wheel.so group=wheel
2.在/etc/pam. d/su配置文件中定义了wheel组.

2030 如何使lynx浏览器能够浏览中文网页(Ghost_Vale)
浏览简体中文网页就的修改如下设置
Save options to disk: [X]
Display and Character Set
Display character set : [Chinese________________________]
Assumed document character set(!): [iso-8859-1______]
CJK mode (!) : [ON_]
然后移到最下面的 Accept Changes 按下 Enter 保存就可以了
当然你的系统要支持简体中文才可以

2031 网卡激活了，却上不了网，怎么办？(Slock，双眼皮的猪)
traceroute，看看到底是在那一块被阻住的。
1.ping自己
2.ping网关
3.ping DNS
4.traceroute DNS
如果一切正常
nslookup www.sina.com.cn
ping sina的address
traceroute sina的address
基本上就可以知道结果了

2032 在redhat9下配samba,win2000能访问，win98不能访问？(squall2003)
如果是wind98必需修改注册表：HKEY_LOCAL_MACHINE/system/correntcontrolset/services/Vxd/VNETSUP下建个D值：EnablePlainTextpasswd，键值1

2033 如何得到网卡的MAC地址(陈绪，hutuworm)
arp -a | awk '{print $4}'
ifconfig eth0 | head -1 | awk '{print $5}' 

2034 如何得到网卡的IP地址(mb)
ifconfig eth0 |awk '/inet addr/ {split($2,x,":");print x[2]}'

2035 如何修改Linux机器所在的工作组(hutuworm)
vi /etc/samba/smb.conf，修改workgroup = 一行，将组名写在后面。

2036 一块网卡如何绑定两个ip(linuxloveu)
#cd /etc/sysconfig/network-scripts
#cp ifcfg-eth0 ifcfg-eth0:1
#vi ifcfg-eth0:1
修改IP和设备名
Debian下一个网卡绑定多个ip的方法(NetDC)
修改/etc/network/interfaces
auto eth0
iface eth0 inet static
address 172.16.3.123
netmask 255.255.255.0
network 172.16.3.0
broadcast 172.16.3.255
gateway 172.16.3.1

auto eth0:1
iface eth0:1 inet static
address 10.16.3.123
netmask 255.255.0.0
network 10.16.0.0
broadcast 10.16.255.255
修改/etc/network/ifstate
lo=lo
eth0=eth0
eth0:1=eth0:1
然后/etc/init.d/networking restart就可以了。
一个网卡绑定多ip另一法(hotbox)
在/etc/sysconfig/network-scripts/下创建一个文件：ifcfg-ethX-rangeX （"X"为网卡号）
文件内容：
IPADDR_START=<start ip>
IPADDR_END=<end ip>
CLONENUM=0
可以有256个ip

2037 一个ip如何绑定两块网卡(hutuworm)
假设192.168.0.88是ip,192.168.0.1是网关:
/sbin/modprobe bonding miimon=100 mode=1
/sbin/ifdown eth0
/sbin/ifdown eth1
/sbin/ifconfig bond0 192.168.0.88
/sbin/ifenslave bond0 eth0 eth1
/sbin/route add default gw 192.168.0.1

2038 192.168.1.0/24(双眼皮的猪)
它与192.168.1.0/255.255.255.0是等价的，只是表示方式不同....

2039 linux下清空arp表的命令(NetDC)
#arp -d -a(适用于bsd)
for HOST in `arp | sed '/Address/d' | awk '{ print $1}'` ; do arp -d $HOST; done

2040 使用ntp协议从服务器同步时间(NetDC)
ntpdate NTP-SERVER 例：ntpdate 172.16.2.1 

2041 host命令的用法(陈绪)
host能够用来查询域名，它还能得到更多的信息
host -t mx example.com可以查询出example.com的MX记录，以及处理mail的host的名字
host -l example.com会返回所有注册在example.com下的域名
host -a example.com则会显示这个主机的所有域名信息.

2042 立刻让LINUX支持NAT(platinum)
echo 1 > /proc/sys/net/ipv4/ip_forward
iptables -t nat -I POSTROUTING -j MASQUERADE

2043 rh8.0下rcp的用法设置(zhqh1)
只对root用户生效
1、在双方root用户根目录下建立.rhosts文件,并将双方的hostname加进去.在此之前应在双方的/etc/hosts文件中加入对方的IP和hostname
2、把rsh服务启动起来，redhat默认是不启动的。方法：用执行ntsysv命令，在rsh选项前用空格键选中，确定退出。 然后执行：service xinetd restart即可。
3、到/etc/pam.d/目录下，把rsh文件中的auth required /lib/security/pam_securetty.so一行用“#”封掉即可。

2044 在ethX设备上，使LINUX支持网络广播功能（默认是不支持的）(platinum)
ip route add 255.255.255.255 dev ethX

2045 路由设置手册(NetDC)
查看路由信息：
netstat -rn
route -n
手工增加一条路由：
route add -net 192.168.0.0 netmask 255.255.255.0 gw 172.16.0.1
手工删除一条路由：
route del -net 192.168.0.0 netmask 255.255.255.0 gw 172.16.0.1
好了，下面到了重要的了，让系统启动的时候自动启用路由设置。
在redhat中添加一条路由，修改文件/etc/sysconfig/static-routes
any net 192.168.0.0 netmask 255.255.255.0 gw 172.16.0.1
在debian中添加一条路由，

方法一：修改/etc/network/interfaces
代码:
auto eth0
iface eth0 inet static
        address 172.16.3.222
        netmask 255.255.0.0
        network 172.16.0.0
        broadcast 172.16.255.255
        gateway 172.16.2.1
   up route add -net 192.168.0.0 netmask 255.255.255.0 gw 172.16.0.1
   down route del -net 192.168.0.0 netmask 255.255.255.0 gw 172.16.0.1
方法二：在/etc/network/if-up.d目录下建立一个简单的脚本文件，如static-route$（记得以$符号结尾，要不有个run-parts会跑出来告诉你一些东西）脚本最简单的就好啦，如：
代码:
#!/bin/bash
route add -net 192.168.0.0 netmask 255.255.255.0 gw 172.16.0.1
嘿嘿，你也可以猜到/etc/network/目录下的其他目录的作用了吧。
发觉在debian中这个route的设置其实只是它的那些配置文件的一个比较简单的应用而已，你完全可以做更复杂的应用。

2046 利用ssh复制文件(platinum)
假如A、B都有SSH服务，现在在A的SSH里
1、从A复制B（推过去）
scp -rp /path/filename username@remoteIP:/path
2、从B复制到A（拉过来）
scp -rp username@remoteIP:/path/filename /path
如果其中一个不是LINUX系统，可以在WINDOWS上用SecureFX软件

2047 samba3.0中文显示问题的解决办法(linuxzfp, jiadingjun)
在samba 3.0的配置文件中(/etc/samba/smb.conf)的[global]中加入下面两句：
unix charset=cp936 
重启服务
service smb restart

2048 临时修改网卡MAC地址的方法
关闭网卡：/sbin/ifconfig eth0 down 
然后改地址：/sbin/ifconfig eth0 hw ether 00:AA:BB:CCD:EE 
然后启动网卡:/sbin/ifconfig eth0 up

2049 conntrack 表满的处理方法(cgweb)
前段时间配置的iptables+squid做的proxy server ,一直工作正常。今天我上控制台上发现 
Jun 18 12:43:36 red-hat kernel: ip_conntrack: table full, dropping packet. 
Jun 18 12:49:51 red-hat kernel: ip_conntrack: table full, dropping packet. 
Jun 18 12:50:57 red-hat kernel: ip_conntrack: table full, dropping packet. 
Jun 18 12:57:38 red-hat kernel: ip_conntrack: table full, dropping packet. 

IP_conntrack表示连接跟踪数据库(conntrack database)，代表NAT机器跟踪连接的数目，连接跟踪表能容纳多少记录是被一个变量控制的，它可由内核中的ip- sysctl函数设置。每一个跟踪连接表会占用350字节的内核存储空间，时间一长就会把默认的空间填满，那么默认空间时多少？我以redhat为例在内存为64MB的机器上时4096,内存为128MB是 8192,内存为256MB是16376，那末就能在/proc/sys/net/ipv4/ip_conntrack_max里查看、设置。 
例如：增加到81920，可以用以下命令: 
echo "81920" > /proc/sys/net/ipv4/ip_conntrack_max 

那样设置是不会保存的，要重启后保存可以在/etc/sysctl.conf中加： 
net.ipv4.ip_conntract_max =81920 
按照此方法改变后一切正常，要是在满了可以加大其值.

2050 Linux下怎么使用BT(atz0001)
azureus，http://azureus.sourceforge.net/

2051 Linux下查看光纤网卡的工作模式(sakulagi)
主板上PCI—X插槽中插入一块64位的光纤网卡，在LINUX9.0的环境下，要知道它是否工作在64位模式下，可使用getconf WORD_BIT

2052 在线更新RHEL的另类途径(hutuworm)
1.安装相应的APT包： 
Red Hat EL 2.1 - i386 
rpm -ihv http://dag.wieers.com/packages/a ... .0.el2.dag.i386.rpm 
Red Hat EL 3 - i386 
rpm -ihv http://dag.wieers.com/packages/a ... .1.el3.dag.i386.rpm 
Red Hat EL 3 - x86_64 
rpm -ihv http://dag.wieers.com/packages/a ... .el3.dag.x86_64.rpm 
2.在线更新 
apt-get update 
apt-get upgrade

2053 SOCKS5启动后一段时间停止工作。用命令ps auxw | grep socks5查看，发现有很多SOCKS defunct进程，为什么(陈绪)
主要是打补丁的问题。如果socks5-tar.gz是没打过补丁的版本，必须下一个带补丁的v1.0-r11版本，重新安装、运行问题就可以解决了。

2054 在VMware WorkStation 4.0.5中安装Debian 3.0时，提示找不到硬盘，需要SCSI的驱动。但是我用的是IDE硬盘，请问该怎么办？ (陈绪)
由于VMware将用户划分的硬盘空间虚拟成SCSI硬盘，而Debian安装盘中没有对应的驱动，而安装其它Linux版本时，有的在一开始会加载SCSI驱动，所以没有这个问题。用户可以修改VMware的配置，将其改为模拟IDE硬盘就可以了。

2055 如何让Linux网关后面的WIN32下的用户直接点击FTP连接下载？(platinum)
modprobe ip_nat_ftp

2056 请问用户的IP是动态的，如何在Squid中限定在同一时间内同一账户在线的数量？(陈绪)
例如限制单个用户只能打开12个HTTP连接，采用下面的方法：
acl all src 0.0.0.0/0.0.0.0
acl limit maxconn 12
acl localnet src 192.168.0.0/24
http_access deny localnet maxconn
http_access allow localnet
http_access deny all

2057 如果我用Squid代理的代理服务器在192.168.1.0这个网段里，例如它的IP是192.168.1.1，我有一些客户端在192.168.2.0这个网段内，怎样设置才能通过这个代理服务器出去？(陈绪)
如果不用透明代理，直接在浏览器的代理选项里设置就可以了。否则首先是在代理服务器的网卡上再挂一个IP为192.168.2.1，添加相应的路由，再修改Squid的squid.conf文件里的监听地址和端口等，最后在192.168.2.0网段的客户端设置其网关为 192.168.2.1，再直接在浏览器的代理选项里设置一下就可以了。

2058 如何使用netrc文件进行自动FTP？(陈绪)
在自己的home目录下建立一个权限为600，后缀名为.netrc的文件，内容如下：
machine 172.168.15.1 login admin password admin
这样用户以后每次登录FTP服务器172.168.15.1的时候，系统都会帮用户以用户名admin、密码admin登录。用户利用这个特征可以实现自动FTP。例如用户想要每天6:00到172.168.15.1机器上面获得/admin目录下的文件admin.txt，可以按如下方法做。
建立一个文件ftp_cmd，内容如下：
cd admin
get amin.txt
bye
然后使用crontab -e设置定时任务：
0 6 * * * ftp 172.168.15.1 < ftp_cmd

2059 怎样得到ipchains的日志？(陈绪)
用户设置规则的时候必须加入-l参数才会在/etc/messages里面做记录。不过建议还是不加的好，不然用户的/etc/messages会变得非常大。

2060 如何不显示其它用户的消息？(陈绪)
用户可以使用mesg n来禁止别人给自己发送信息，其实就是禁止别人往自己的终端上面的写的权限。当别人试图再使用write给自己发送信息时，发送者将会看见提示如下：
write: user has messages disabled on pts/n

2061 minicom彩色显示(双眼皮的猪)
minicom -s进行serial port配置,然后配置好以后, 
minicom -o -c on 
-o表示不初始化 
-c on表示color on

2062 启用SELinux的Apache的配置文件httpd.conf里面修改DocumentRoot无用或者出现403 Forbidden错误(arbor)
# chcon -u system_u -t httpd_sys_content_t -R website目录 

2063 apache2 的log文件位置如何自定义目录(tomi)
编辑httpd.conf里的 
ErrorLog /var/log/http/error_log          <== 这是管errorlog的 
CustomLog /var/log/http/access_log common        <== 这是管accesslog的

2064 更改eth0是否混杂模式(wwy)
网卡eth0改成混杂模式：
ifconfig eth0 promisc
关闭混杂模式：
ifconfig eth0 -promisc

2065 字符界面下的ftp中，下载整个文件夹(陈绪)
1. lftp IP
2. > user username
password
3. > mirror -c --parallel=number remotedir localdir
3a. > help mirror

2066 如何让ssh只允许指定的用户登录(xinyv，好好先生，wolfg，我爱钓鱼)
方法1：在/etc/pam.d/sshd文件中加入
auth   required   pam_listfile.so  item=user  sense=allow  file=/etc/sshusers  onerr=fail
然后在/etc下建立sshusers文件,编辑这个文件,加入你允许使用ssh服务的用户名,重新起动sshd服务即可。
方法2：pam规则也可以写成deny的
auth   required   pam_listfile.so  item=user  sense=deny  file=/etc/sshusers  onerr=succeed
方法3：在sshd_config中设置AllowUsers，格式如
AllowUsers a b c
重启sshd服务，则只有a/b/c3个用户可以登陆。

2067 在Linux下如何绑定IP地址和硬件地址(陈绪)
可以编辑一个地址对应文件，里面记录了IP地址和硬件地址的对应关系，然后执行“arp –f 地址对应文件”。如果没有指定地址对应文件，则通常情况下一默认文件/etc/ethers为准。地址对应文件的格式如下：
192.168.0.1 00:0D:61:27:58:93
192.168.0.2 00:40:F4:2A:2E:5C
192.168.0.3 00:0A:EB:5E:BA:8E

2068 已知网络中一个机器的硬件地址，如何知道它所对应的IP地址(陈绪)
在Linux下，假定要查“00:0A:EB:27:17:B9”这样一个硬件地址所对应的IP地址，可以使用以下命令：
# cat /proc/net/arp |grep 00:0A:EB:27:17:B9
192.168.2.54 0x1 0x6 00:0A:EB:27:17:B9 *eth2
另外，还可以用“arp -a”命令查询：
# arp –a|grep 00:0A:EB:27:17:B9
（192.168.2.54）at 00:0A:EB:27:17:B9[ether] on eth2

2069 基于Apache的HTTPD或Sendmail服务在启动时被挂起了，如何解决此问题(陈绪)
遇到此类问题，请确认/etc/hosts文件中是否包含如下一行：
127.0.0.1 localhost.localdomain localhost
127.0.0.1 是网络的回路地址。

2070 如何使Linux系统对ping不反应(陈绪)
要使Linux对ping没反应，也就是使Linux系统忽略I CMP包。用如下命令可以达到此目的：
# echo 1 > /proc/sys/net/ipv4/icmp-echo-ignore-all
若想恢复，可用如下命令：
# echo 0 > /proc/sys/net/ipv4/icmp-echo-ignore-all

2071 压缩传输文件或目录(FunBSD)
传输到远程：tar czf - www | ssh server "tar zxf -"
压缩到远程：tar czf - www | ssh server "cat > www.tar.gz"
解压到远程：ssh server "tar zxf -" < www.tar.gz
解压到本地：ssh server "cat www.tar.gz" | tar zxf -

2072 rsync同步压缩传输文件或目录(FunBSD)
rsync -aze ssh --delete sample_dir/ remote_host:remote_dir/sample_dir/
目录最后的/不能少

2073 无需输入密码使用ssh密钥登录 (FunBSD)
ssh-keygen -b 1024 -t rsa
ssh server "mkdir .ssh; chmod 0700 .ssh"
scp ~/.ssh/id_rsa.pub server:~/.ssh/authorized_keys
这样就不在提示密码，直接可以登录server了
对文件复制、同步等操作都比较方便
在ssh_config里加入这两句就更方便了
ForwardAgent yes
StrictHostKeyChecking no

2074 wget下载整个网站(陈绪)
wget -t0 -c -nH -np -b -m -P /localdir http://freesoft.online.sh.cn/mirrors/ftp.redhat.com -o wget.log

2075 命令行下发送带附件的邮件(陈绪)
方法1.    uuencode <in_file> <remote_file> | mail -s "title" mail@address
<in_file> 本地需要作为附件的文件名。
<remote_file> 邮件中的附件文件名，可以和<in_file>不同，其实内容一样。
方法2.    cat <mailcontent.txt> | mutt -s "title" -a <attachfile> mail@address
<mailcontent.txt>邮件正文内容。
<attachfile>本地需要作为附件的文件名。

2076 高效率使用1000兆网卡(陈绪)
系统加载模块时，可以根据实际情况调节参数，使网卡工作在最佳状态。驱动重新提供的可选择参数有速率、工作模式、自适应和流控等
在Linux下，可以定义合法速率参数为0、10、100和1000。却省为0，表示网卡工作在自适应状态下，其他值分别为10Mb、100Mb和1000Mb。
工作模式有全、半双工方式。0表示适应；1表示半双工；2表示全双工。
自适应方式的有效期值范围0~3。0表示不设置流控；1表示仅对Rx流控；2表示仅对Tz流控；3表示对Rx/Tx双向流控。缺省为3

2077 管理SSH监听端口(陈绪)
从安全角度考虑，SSH应当取代Telnet。目前在Linux上使用广泛的SSH服务器软件sshd-config（默认路径是 /etc/ssh/sshd-config）文件中，Port 22是sshd监听的端口，即为连接到主机时需要使用的端口。使用以下代码可以指定sshd监听的接口地址：
ListenAddress 192.168.0.254
这样，就可以避免向未知的用户提供登录服务

----------------------------程序开发篇--------------------------
3001 linux下调试core文件(陈绪)
gdb <progname> <core>
<progname>:出错产生core dump的可执行程序。
<core>: core dump的文件名，缺省是“core”

3002 gcc abc.c得到的a.out不能运行(陈绪)
./a.out

3003 c++ 编译时为什么出错信息说cout没定义(陈绪)
include头文件完后加入 using namespace std;

3004 新编译生成的gcc ，使用的标准连接库都在/usr/local/lib 下了，但使用的缺省的连接路径是 /usr/lib 怎样添加？（除了在每次编译时 增加 -L /usr/local/lib 以外)(sakulagi, hutuworm)
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib
写到~/.bash_profile里面。
增加一种简便办法：
将/usr/local/lib加入/etc/ld.so.conf，然后运行一次ldconfig

3005 RH9下GCC的安装(一起走过的日子，hutuworm)
三种方法选一：
(1)利用CD上rpm安装
CD-1:compat-gcc-7.3-2.96.118.i386.rpm
CD-1:compat-gcc-c++-7.3-2.96.118.i386.rpm
CD-1:libgcc-3.2.2-5.i386.rpm
CD-2:compat-gcc-g77-7.3-2.96.118.i386.rpm
CD-2:compat-gcc-java-7.3-2.96.118.i386.rpm
CD-2:compat-gcc-objc-7.3-2.96.118.i386.rpm
CD-2:gcc-3.2.2-5.i386.rpm
CD-2:gcc-c++-3.2.2-5.i386.rpm
CD-2:gcc-g77-3.2.2-5.i386.rpm
CD-2:gcc-gnat-3.2.2-5.i386.rpm
CD-2:gcc-java-3.2.2-5.i386.rpm
CD-2:gcc-objc-3.2.2-5.i386.rpm
比如碰到系统提示：
warning : gcc-3.2.2-5.i386.rpm : V3 DSA signature :MOKEY key ID db42a60e
error : Failed dependencies :
binutils >=2.13.90.0.18-9 is needed by gcc-3.2.2-5
glibc-devel >=2.3.2-11.9 is needed by gcc-3.2.2-5...
就先安裝glibc-devel包，依此类推
(2)更好的方法就是在X-window下选“主菜单”──>“系统设置”──>“添加/删除应用程序”──>“开发工具”中的gcc并安装它
(3) up2date gcc便可自动解决dependency问题

3006 shell脚本为何无法运行(GOD_Father)
第一，脚本权限要为可执行 #chmod +x test.sh
第二，脚本所在的目录在环境变量PATH中，或者直接执行 #./test.sh

3007 查看某个文件被哪些进程在读写(bjweiqiong)
lsof 文件名

3008  查看某个进程打开了哪些文件(bjweiqiong)
lsof –c 进程名
lsof –p 进程号

3009  lsof是什么意思(bjweiqiong)
list open files

3010 lsof用法小全(bjweiqiong)
lsof abc.txt 显示开启文件abc.txt的进程
lsof -i :22 知道22端口现在运行什么程序
lsof -c nsd 显示nsd进程现在打开的文件
lsof -g gid 显示归属gid的进程情况
lsof +d /usr/local/ 显示目录下被进程开启的文件
lsof +D /usr/local/ 同上，但是会搜索目录下的目录，时间较长
lsof -d 4  显示使用fd为4的进程
lsof -i 用以显示符合条件的进程情况
语法: lsof -i[46] [protocol][@hostname|hostaddr][:service|port]
46 --> IPv4 or IPv6
protocol --> TCP or UDP
hostname --> Internet host name
hostaddr --> IPv4位置
service --> /etc/service中的 service name (可以不只一個)
port --> 埠號 (可以不只一個)
例子: TCP:25 - TCP and port 25
@1.2.3.4 - Internet IPv4 host address 1.2.3.4
tcp@ohaha.ks.edu.tw:ftp - TCP protocol host:ohaha.ks.edu.tw service name:ftp
lsof -n 不将IP转换为hostname，预设是不加上-n参数
例子: lsof -i tcp@ohaha.ks.edu.tw:ftp -n
lsof -p 12  看进程号为12的进程打开了哪些文件    
lsof +|-r [t] 控制lsof不断重复执行，缺省是15s刷新
-r，lsof会永远不断的执行，直到收到中断讯号
+r，lsof会一直执行，直到没有档案被显示
例子：不断查看目前ftp连接的情况：lsof -i tcp@ohaha.ks.edu.tw:ftp -r
lsof -s 列出打开文件的大小，如果没有大小，则留下空白
lsof -u username  以UID，列出打开的文件

3011 让某用户只能ftp，不能ssh/telnet(bjweiqiong)
vi /etc/passwd
将用户行中的/bin/bash改为/bin/false即可

----------------------------经典图书篇--------------------------
4001 GNU/Linux高级网络应用服务指南(陈绪)
linuxaid网站
机械工业出版社
优点：又全又精，全都是实战之作
缺点：针对版本较低，为redhat 6.2

4002 Linux Apache Web Server管理指南(Linux Apache Web Server Administration)(陈绪)
Charles Aulds 马树奇/金燕译
电子工业出版社
优点：目前我还没有发现哪个关于apache的问题这本书没有讲过
缺点：针对1.3.x，最新的针对2.0.*的英文版已出，中文版待出

4003 Linux内核情景分析(陈绪)
毛德操/胡希明
浙江大学出版社
优点：太透彻了，没法不懂
缺点：还是版本问题，内核更新太快了，不过还是必读

4004 Unix环境高级编程(陈绪)
Richard Stevens
机械工业出版社
优点：博大精深
缺点：初学者是很难理解的，否则怎么叫《高级编程》呢？

4005 编程精粹--Microsoft编写优质无错c程序秘诀(陈绪)
Steve Maguire
电子工业出版社
优点：不说了，作者是微软的资深工程师
缺点：很难找了，1994年出的

4006 Understanding the Linux Kernel, 2nd Edition(hutuworm)     
Daniel P. Bovet & Marco Cesati
O'Reilly出版社
读了这本书之后，你就会明白在什么情况下Linux具有最佳的性能，以及它如何面对挑战，在各种环境中提供进程调度、文件访问和内存管理时的优良的系统响应。作者通过解释其重要性来引入每一个题目，并将内核操作与Unix程序员和用户熟悉的系统调用或实用程序联系起来。

4007 UNIX操作系统教程(英文版)(弱智)     
Syed Mansoor Sarwar等
机械工业出版社
特点：浅显易懂，着重unix基础概念和整体理解，顺便复习英语。
另外：机械工业出版社已经出版了中文版，名称：UNIX 教程

4008 UNIX编程环境(弱智)     
Brian W.Kernighan, Rob Pike  陈向群等译
机械工业出版社
特点：浅显，深入浅出讲解如何使用UNIX及各种工具，简单介绍Unix编程环境；对比“UNIX环境高级编程”，此书适合新手入门。

4009 The Art of UNIX Programming(hutuworm)     
Eric Steven Raymond
http://catb.org/~esr/writings/taoup/html/
优点： E.S. Raymond的经典著作

4010 unix网络编程--卷一 套接口API和X/Open传输接口API(slg1972)     
Richard Stevens
清华大学出版社
优点：详细地讲解unix网络的编程

4011 unix网络编程--卷二 进程间通讯(slg1972)
Richard Stevens
清华大学出版社
优点：详细讲解unix的进程之间，线程之间的关系，及各种不同标准的进程编程的异同

4012 unix网络编程--卷三 应用程序(slg1972, hutuworm)     
未出，因为Richard Stevens大师英年早逝，再也不可能完成这计划中的第三卷了。据说其未竟稿可能由Gary R. Wright整理续写出来，但是自大师驾鹤以来一直杳无音信

4013 基于C++ CORBA高级编程(slg1972) 
Michi Henning，Steve Vinoski
清华大学出版社
优点：中间件的好书，通向corba应用的必备资料。

4014 unix linux网管通鉴(odin_free) 
电子版的，包括本版精华
优点：我见过关于unix知识最全面、最实用的chm文档，相当于一个小型网站，里面支持全文检索，推荐所有还没有的兄弟姐妹们下载

4015 www.chinaoy.com(aomin5555) 
不错，挺全的，图书下载的好网址：
redhat linux9.0 官方入门指南
·redhat linux9.0 官方安装指南
·redhat linux9.0 官方定制设置手册
·redhat linux基础教程
·Linux 参考大全
·清华论坛linux精华
·Linux系统管理员指南中文手册
·Linux网站建设和维护全攻略
·redhat linux8.0 安装手册
·Linux环境database管理员指南 

4016 Linux Advanced Routing & Traffic Control(hutuworm) 
专门讲LINUX IPROUTE2的书，大概100页左右，www.lartc.org
中文版在：http://www.lartc.org/LARTC-zh_CN.GB2312.pdf

4017 Debian User强烈推荐看的书(NetDC) 
Debian Reference （Debian参考手册）
http://qref.sourceforge.net/
简体中文版的pdf文档。
http://qref.sourceforge.net/Debian/reference/reference.zh-cn.pdf

4018 Advanced Bash-Scripting Guide(hutuworm) 
An in-depth exploration of the art of shell scripting.
Mendel Cooper. <thegrendel@theriver.com>.
http://www.tldp.org/LDP/abs/abs-guide.pdf
优点：是Bash编程的圣经，而且该书作者不断在更新其内容，一两个月就会翻新一个版本，值得一读，一读再读。

4019 JAVA完美经典(陈绪) 
江义华 编著 林彩瑜 文编 
中国铁道出版社
定价：65元
优点：不愧是台湾同胞的力作，讲解清楚，知识全面，我看了之后，感到很有收获


----------------------------mysql相关篇--------------------------
5001 mysql的数据库存放在什么地方(陈绪) 
1. 如果使用rpm包安装，应该在/var/lib/mysql目录下，以数据库名为目录名
2. 如果源码安装在/usr/local/mysql中，应该在/usr/local/mysql/var中，以数据库名为目录名

5002 从mysql中导出和导入数据(陈绪) 
导出数据库
mysqldump 数据库名 > 文件名
导入数据库
mysqladmin create 数据库名
mysql 数据库名 < 文件名

5003 忘了mysql的root口令怎么办(陈绪) 
# service mysql stop
# mysqld_safe --skip-grant-tables &
# mysqladmin -u user password 'newpassword''
# mysqladmin flush-privileges

5004 快速安装php/mysql(陈绪)
确保使用系统自带的apache，从安装光盘中找出所有以mysql及php-mysql开头的rpm包，然后运行#rpm -ivh mysql*.rpm php-mysql*.rpm; mysql_install_db; service mysql start

5005 修改mysql的root口令(陈绪，yejr) 
大致有2种方法：
1、mysql>mysql -uroot -pxxx mysql
mysql>update user set password=password('new_password') where user='user';
mysql>flush privileges;
2、格式：mysqladmin -u用户名 -p旧密码 password 新密码
#mysqladmin -uroot -password ab12
注：因为开始时root没有密码，所以-p旧密码一项就可以省略了

5006 如何使用rpm方式安装mysql(yejr) 
首先下载合适的rpm包，例如下载了文件 MySQL-5.0.19-0.i386.rpm
用一下方法安装：
#rpm -ivhU MySQL-5.0.19-0.i386.rpm
通常情况下，安装完这个rpm包后，只具备有mysqld服务功能，其它相关的client程序和开发包还需要另外安装
#rpm -ivhU MySQL-devel-5.0.19-0.i386.rpm
#rpm -ivhU MySQL-client-5.0.19-0.i386.rpm

5007 如何安装已经编译好了的mysql二进制包(yejr) 
首先下载合适的二进制包，例如下载了文件 mysql-standard-4.1.13-pc-linux-gnu-i686.tar.gz
#groupadd mysql
#useradd -g mysql mysql
#cd /usr/local
#tar zxf mysql-standard-4.1.13-pc-linux-gnu-i686.tar.gz
#ln -s mysql-standard-4.1.13-pc-linux-gnu-i686 mysql
#cd mysql
#scripts/mysql_install_db --user=mysql
#chgrp -R mysql *
#bin/mysqld_safe --user=mysql &
有什么个性化的配置，可以通过创建 /etc/my.cnf 或者 /usr/local/mysql/data/my.cnf，增加相关的参数来实现

5008 如何自己编译mysql(yejr) 
以redhat linux 9.0为例：
下载文件 mysql-4.1.13.tar.gz
#tar zxf mysql-4.1.13.tar.gz
#cd mysql-4.1.13
#./configure --prefix=/usr/local/mysql --enable-assembler /
--with-mysqld-ldflags=-all-static --localstatedir=/usr/local/mysql/data /
--with-unix-socket-path=/tmp/mysql.sock --enable-assembler /
--with-charset=complex --with-low-memory --with-mit-threads
#make
#make install
#groupadd mysql
#useradd -g mysql mysql
#chgrp -R mysql /usr/local/mysql/
#/usr/local/mysql/bin/mysqld_safe --user=mysql &
有什么个性化的配置，可以通过创建 /etc/my.cnf 或者 /usr/local/mysql/data/my.cnf，增加相关的参数来实现

5009 如何登录mysql(yejr)
使用mysql提供的客户端工具登录
#PATH_TO_MYSQL/bin/mysql -uuser -ppassword dateabase

5010 mysqld起来了，却无法登录，提示"/var/lib/mysql/mysql.sock"不存在(yejr)
这种情况大多数是因为你的mysql是使用rpm方式安装的，它会自动寻找 /var/lib/mysql/mysql.sock 这个文件，
通过unix socket登录mysql。
常见解决办法如下：
1、创建/修改文件 /etc/my.cnf，至少增加/修改一行
[mysql]
[client]
socket = /tmp/mysql.sock
#在这里写上你的mysql.sock的正确位置，通常不是在 /tmp/ 下就是在 /var/lib/mysql/ 下
2、指定IP地址，使用tcp方式连接mysql，而不使用本地sock方式
#mysql -h127.0.0.1 -uuser -ppassword
3、为 mysql.sock 加个连接，比如说实际的mysql.sock在 /tmp/ 下，则
# ln -s /tmp/mysql.sock /var/lib/mysql/mysql.sock即可

5011 如何新增一个mysql用户(yejr)
格式：grant select on 数据库.* to 用户名@登录主机 identified by "密码"
例1、增加一个用户test1密码为abc，让他可以在任何主机上登录，并对所有数据库有查询、插入、修改、删除的权限。首先用以root用户连入MYSQL，然后键入以下命令：
mysql>grant select,insert,update,delete on *.* to test1@"%" Identified by "abc";
但例1增加的用户是十分危险的，你想如某个人知道test1的密码，那么他就可以在internet上的任何一台电脑上登录你的mysql数据库并对你的数据可以为所欲为了，解决办法见例2。
例2、增加一个用户test2密码为abc,让他只可以在localhost上登录，并可以对数据库mydb进行查询、插入、修改、删除的操作（localhost指本地主机，即MYSQL数据库所在的那台主机），这样用户即使用知道test2的密码，他也无法从internet上直接访问数据库，只能通过MYSQL主机上的web页来访问了。
mysql>grant select,insert,update,delete on mydb.* to test2@localhost identified by "abc";
如果你不想test2有密码，可以再打一个命令将密码消掉。
mysql>grant select,insert,update,delete on mydb.* to test2@localhost identified by "";
另外，也可以通过直接往user表中插入新纪录的方式来实现

5012 如何查看mysql有什么数据库(yejr)
mysql>show databases;

5013 如何查看数据库下有什么表(yejr)
mysql>show tables;

5014 导出数据的几种常用方法(yejr) 
1、使用mysqldump
#mysqldump -uuser -ppassword -B database --tables table1 --tables table2 > dump_data_20051206.sql
详细的参数
2、backup to语法
mysql>BACKUP TABLE tbl_name[,tbl_name...] TO '/path/to/backup/directory';
详细请查看mysql手册
3、mysqlhotcopy
#mysqlhotcopy db_name [/path/to/new_directory]
或
#mysqlhotcopy db_name_1 ... db_name_n /path/to/new_directory
或
#mysqlhotcopy db_name./regex/
详细请查看mysql手册
4、select into outfile
详细请查看mysql手册
5、客户端命令行
#mysql -uuser -ppassword -e "sql statements" database > result.txt
以上各种方法中，以mysqldump最常用

5015 如何在命令行上执行sql语句(yejr)
#mysql -uuser -ppassword -e "sql statements" database

5016 导入备份出来文件的常见方法(yejr)
1、由mysqldump出来的文件
#mysql -uuser -ppassword [database] < dump.sql
2、文件类型同上，使用source语法
mysql>source /path_to_file/dump.sql;
3、按照一定格式存储的文本文件或csv等文件
#mysqlimport [options] database file1 [file2....]
详细请查看mysql手册
4、文件类型同上，也可以使用load data语法导入
详细请查看mysql手册

5017 让mysql以大内存方式启动(陈绪)
将/usr/share/mysql下的某个mysql-*.cnf(如1G内存时为mysql-huge.cnf)拷贝为/etc/mysql.cnf文件，并重启mysql

-------------------------------杂项篇--------------------------------
感谢allan0909指正
请不要做浮躁的人
http://www.chinaunix.net/forum/viewtopic.php?t=93113
欢迎转载本文，请注明来自www.chinaunix.net和www.linuxmine.com，转载本文的网址太多，以至无法列举，请参见google。