# php-mcrypt 

在 centos 6（php5.3.3版本） 无法安装 php-mcrypt , 因为默认的 centos6 是没有添加 mcrypt 扩展，php （5.3.3） 无法设定此项设置。

使用 `yum install php-mcrypt` 命令安装无效。

检索源库文件,在 Fedora 提供的 RHEL 源中有 mcrypt 包。

下载[阿里源](http://mirrors.aliyun.com/)

```
cd /etc/yum.repos/
wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-6.repo
```

如果使用下载中需要认证

```
wget http://mirrors.aliyun.com/epel/RPM-GPG-KEY-EPEL-6
rpm --import RPM-GPG-KEY-EPEL-6 
```

之后更新源缓存文件

```
yum --enablerepo=epel install php-mcrypt
# 或者
yum makecache 
yum install php-mcrypt
```

之后使用 `php -i` 检查PHP。

参考文案 <http://qiita.com/soundws/items/cc84da42419f1ab3443b>

日本源 <http://ftp.riken.jp/>
