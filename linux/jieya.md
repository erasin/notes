
#linux下 tar解压 gz解压 bz2等各种解压文件使用方法

    博客分类：
    Linux

LinuxJDKF# 
大致总结了一下linux下各种格式的压缩包的压缩、解压方法。但是部分方法我没有用到，也就不全，希望大家帮我补充，我将随时修改完善，谢谢！
　　
　　.tar
　　解包：tar xvf FileName.tar
　　打包：tar cvf FileName.tar DirName
　　（注：tar是打包，不是压缩！）
　　———————————————
　　.gz
　　解压 1：gunzip FileName.gz
　　解压2：gzip -d FileName.gz
　　压缩：gzip FileName
　　.tar.gz 和 .tgz
　　解压：tar zxvf FileName.tar.gz
　　压缩：tar zcvf FileName.tar.gz DirName
　　———————————————
　　.bz2
　　解压1：bzip2 -d FileName.bz2
　　解压2：bunzip2 FileName.bz2
　　压缩： bzip2 -z FileName
　　.tar.bz2
　　解压：tar jxvf FileName.tar.bz2        或tar --bzip xvf FileName.tar.bz2
　　压缩：tar jcvf FileName.tar.bz2 DirName
　　 ———————————————
　　.bz
　　解压1：bzip2 -d FileName.bz
　　解压2：bunzip2 FileName.bz
　　压缩：未知
　　.tar.bz
　　解压：tar jxvf FileName.tar.bz
　　压缩：未知
　　———————————————
　　.Z
　　解压：uncompress FileName.Z
　　压缩：compress FileName
　　.tar.Z
　　解压：tar Zxvf FileName.tar.Z
　　压缩：tar Zcvf FileName.tar.Z DirName
　　———————————————
　　.zip
　　解压：unzip FileName.zip
　　压缩：zip FileName.zip DirName
　　压缩一个目录使用 -r 参数，-r 递归。例： $ zip -r FileName.zip DirName
　　———————————————
　　.rar
　　解压：rar x FileName.rar
　　压缩：rar a FileName.rar DirName
　　
　　rar 请到：http://www.rarsoft.com/download.htm 下载！
　　解压后请将rar_static拷贝到/usr /bin目录（其他由$PATH环境变量指定的目录也可以）：
　　[root@www2 tmp]# cp rar_static /usr/bin/rar
　　———————————————
　　.lha
　　解压：lha -e FileName.lha
　　压缩：lha -a FileName.lha FileName
　　
　　lha请到：http://www.infor.kanazawa-it.ac.jp/~ishii/lhaunix/下载！
　　>解压后请将 lha拷贝到/usr/bin目录（其他由$PATH环境变量指定的目录也可以）：
　　[root@www2 tmp]# cp lha /usr/bin/
　　———————————————
　　.rpm
　　解包：rpm2cpio FileName.rpm | cpio -div
　　———————————————
　　.deb
　　解包：ar p FileName.deb data.tar.gz | tar zxf -
　　———————————————
　　.tar .tgz .tar.gz .tar.Z .tar.bz .tar.bz2 .zip .cpio .rpm .deb .slp .arj .rar .ace .lha .lzh .lzx .lzs .arc .sda .sfx .lnx .zoo .cab .kar .cpt .pit .sit .sea
　　解压：sEx x FileName.*
　　压缩：sEx a FileName.* FileName
　　
　　sEx只是调用相关程序，本身并无压缩、解压功能，请注意！
　　sEx请到： http://sourceforge.net/projects/sex下载！
　　解压后请将sEx拷贝到/usr/bin目录（其他由$PATH环境变量指定的目录也可以）：
　　[root@www2 tmp]# cp sEx /usr/bin/　　Linux下常见文件解压方法及命令
　　系统·System
　　
　　1.以.a为扩展名的文件:
　　#tar xv file.a
　　2.以.z为扩展名的文件:
　　#uncompress file.Z
　　3.以.gz为扩展名的文件:
　　#gunzip file.gz
　　4.以.bz2为扩展名的文件:
　　#bunzip2 file.bz2
　　5.以.tar.Z为扩展名的文件:
　　#tar xvZf file.tar.Z
　　或 #compress -dc file.tar.Z | tar xvf
　　6.以.tar.gz/.tgz为扩展名的文件:
　　#tar xvzf file.tar.gz
　　或 gzip -dc file.tar.gz | tar xvf -
　　7.以.tar.bz2为扩展名的文件:
　　#tar xvIf file.tar.bz2
　　或 bzip2 -dc file.tar.bz2 | xvf -
　　8.以.cpio.gz/.cgz为扩展名的文件:
　　#gzip -dc file.cgz | cpio -div
　　9. 以.cpio/cpio为扩展名的文件:
　　#cpio -div file.cpio
　　或cpio -divc file.cpio
　　10.以.rpm为扩展名的文件安装:
　　#rpm -i file.rpm
　　11.以.rpm为扩展名的文件解压缩：
　　 #rpm2cpio file.rpm | cpio -div
　　12.以.deb为扩展名的文件安装：
　　#dpkg -i file.deb
　　13.以.deb为扩展名的文件解压缩:
　　#dpkg-deb -fsys-tarfile file.deb | tar xvf - ar p
　　file.deb data.tar.gz | tar xvzf -
　　14.以.zip为扩展名的文件:
　　#unzip file.zip
　　在linux下解压Winzip格式的文件
　　要是装了jdk的话，可以用 jar命令；还可以使用unzip命令。
　　直接解压.tar.gz文件
　　xxxx.tar.gz文件使用tar带zxvf参数，可以一次解压开。XXXX为文件名。 例如：
　　$tar zxvf xxxx.tar.gz 各种压缩文件的解压（安装方法）
　　
　　文件扩展名 解压（安装方法）
　　　
　　.a ar xv file.a
　　.Z uncompress file.Z
　　.gz gunzip file.gz
　　.bz2 bunzip2 file.bz2
　　.tar.Z tar xvZf file.tar.Z
　　compress -dc file.tar.Z | tar xvf -
　　.tar.gz/.tgz tar xvzf file.tar.gz
　　gzip -dc file.tar.gz | tar xvf -
　　.tar.bz2 tar xvIf file.tar.bz2
　　bzip2 -dc file.tar.bz2 | xvf -
　　.cpio.gz/.cgz gzip -dc file.cgz | cpio -div
　　.cpio/cpio cpio -div file.cpio
　　cpio -divc file.cpio
　　.rpm/install rpm -i file.rpm
　　.rpm/extract rpm2cpio file.rpm | cpio -div
　　.deb/install dpkg -i file.deb
　　.deb/exrtact dpkg-deb -fsys-tarfile file.deb | tar xvf -
　　ar p file.deb data.tar.gz | tar xvzf -
　　.zip unzip file.zip
　　
　　
　　bzip2 -d myfile.tar.bz2 | tar xvf
　　
　　
　　tar xvfz myfile.tar.bz2
　　
　　
　　x 是解压
　　v 是复杂输出
　　f 是指定文件
　　z gz格式
　　
　　
　　gzip
　　gzip[选项]要压缩（或解压缩）的文件名
　　-c将输出写到标准输出上，并保留原有文件。
　　-d将压缩文件压缩。
　　-l对每个压缩文件，显示下列字段：压缩文件的大小，未压缩文件的大小、压缩比、未压缩文件的名字
　　-r递归式地查找指定目录并压缩或压缩其中的所有文件。
　　-t测试压缩文件是正完整。
　　-v对每一个压缩和解压缩的文件，显示其文件名和压缩比。
　　-num-用指定的数字调整压缩的速度。
　　举例：
　　把/usr目录并包括它的子目录在内的全部文件做一备份，备份文件名为usr.tar
　　tar cvf usr.tar /home
　　把/usr 目录并包括它的子目录在内的全部文件做一备份并进行压缩，备份文件名是usr.tar.gz
　　tar czvf usr.tar.gz /usr
　　压缩一组文件，文件的后缀为tar.gz
　　#tar cvf back.tar /back/
　　#gzip -q back.tar
　　or
　　#tar cvfz back.tar.gz /back/
　　释放一个后缀为tar.gz 的文件。
　　#tar zxvf back.tar.gz
　　#gzip back.tar.gz
　　#tar xvf back.tar 
