#  FFMPEG-PHP 

os:centos/radhat

编译环境

	# yum install -y automake autoconf libtool gcc gcc-c++

## 安装 ffmpeg

添加` /etc/yum.repos.d/dag.repo`

> 注意修改 $releasever 为 `5` or `6`

	[dag]
	name=Dag RPM Repository for Red Hat Enterprise Linux
	baseurl=http://apt.sw.be/redhat/el$releasever/en/$basearch/dag
	gpgcheck=1
	enabled=1

出错时添加 GPG-KEY

	rpm -Uhv http://apt.sw.be/redhat/el5/en/i386/rpmforge/RPMS/rpmforge-release-0.3.6-1.el5.rf.i386.rpm

or

	rpm -Uhv http://apt.sw.be/redhat/el5/en/x86_64/rpmforge/RPMS//rpmforge-release-0.3.6-1.el5.rf.x86_64.rpm


安装 Install ffmpeg 等模块 

	yum -y install ffmpeg ffmpeg-devel

## 安装 ffmpeg-php

	cd /usr/local/src 
	wget http://garr.dl.sourceforge.net/sourceforge/ffmpeg-php/ffmpeg-php-0.6.0.tbz2 
	tar jxvf ffmpeg-php-0.6.0.tbz2 
	cd ffmpeg-php-0.6.0 
	phpize 
	./configure
	make 
	make install

生成文件`/usr/lib64/php/modules/ffmpeg.so`

在php.ini文件加上这句 

	extension=ffmpeg.so

## Q&A

错误:

	ffmpeg-php – error: ‘PIX_FMT_RGBA32’ undeclared (first use in this function) | HOW GEEK!.

解决,make前执行

	rpl -R PIX_FMT_RGBA32 PIX_FMT_RGB32 *

如果没有 rpl 安装

	# yum install rpl

## 参考

* <http://www.mysql-apache-php.com/ffmpeg-install.htm>
* <http://www.nixgurus.com/fedora/installing-ffmpeg-using-yum-on-centos-redhat-fedora/>
* <http://www.cnblogs.com/phphuaibei/archive/2011/09/09/2172589.html> 
* <http://www.fuchaoqun.com/2008/07/ffmpeg-php-install-on-centos/>
