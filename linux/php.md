
Centos搭建PHP5.3.8+Nginx1.0.9+Mysql5.5.17

算不上经验,只能说是个人总结.
在搭建过程中遇到的问题.并记录下来.
刚开学习始搭建环境的朋友少走一些弯路.
这没有过多的参数,并不是什么高性能网站的部署.
本人经验有限,在此分享一些遇到的问题.


操作环境


　　操作系统:Mac Os Lion

　　虚拟主机:VMware Fusion

　　虚拟系统:Centos 5.5+

　　操作用户:Root

　　实现目的:搭建LNMP环境.



安装依赖库和开发环境


复制代码
#依赖库和开发工具
yum -y install gcc gcc-c++ autoconf libjpeg libjpeg-devel libpng libpng-devel freetype freetype-devel libxml2 libxml2-devel zlib zlib-devel glibc glibc-devel glib2 glib2-devel bzip2 bzip2-devel ncurses ncurses-devel curl curl-devel e2fsprogs e2fsprogs-devel krb5 krb5-devel libidn libidn-devel openssl openssl-devel openldap openldap-devel nss_ldap openldap-clients openldap-servers

#Nginx
yum -y install pcre-devel  zlib-devel

#Php
yum -y install gd-devel libjpeg-devel libpng-devel freetype-devel libxml2-devel curl-devel freetype-devel

#Mysql
yum -y install bison gcc gcc-c++ autoconf automake zlib* libxml* ncurses-devel libtool-ltdl-devel* mysql-devel

复制代码


下载软件包


复制代码
#创建目录
mkdir /web
cd /web

#PHP5.3.7
wget http://cn.php.net/distributions/php-5.3.8.tar.bz2

#PHP库文件
wget http://ncu.dl.sourceforge.net/project/mcrypt/MCrypt/2.6.8/mcrypt-2.6.8.tar.gz
wget http://ncu.dl.sourceforge.net/project/mhash/mhash/0.9.9.9/mhash-0.9.9.9.tar.gz
wget http://ncu.dl.sourceforge.net/project/mcrypt/Libmcrypt/2.5.8/libmcrypt-2.5.8.tar.gz
wget http://ftp.gnu.org/pub/gnu/libiconv/libiconv-1.14.tar.gz

#Nginx1.0.9
wget http://www.nginx.org/download/nginx-1.0.9.tar.gz

#Nginx(pcre)
wget ftp://ftp.csx.cam.ac.uk/pub/software/programming/pcre/pcre-8.13.tar.gz

#Mysql5.5.17
wget http://dev.mysql.com/get/Downloads/MySQL-5.5/mysql-5.5.17.tar.gz/from/http://mysql.ntu.edu.tw/

#Mysql(cmake)
wget http://www.cmake.org/files/v2.8/cmake-2.8.6.tar.gz

复制代码


安装Mysql


复制代码
#安装cmake
tar -zxvf cmake-2.8.6.tar.gz
cd cmake-2.8.6/
./configure
gmake && gmake install  && cd ../

#添加mysql用户
/usr/sbin/groupadd mysql
/usr/sbin/useradd -g mysql mysql
mkdir -p /data/mysql
chown -R mysql:mysql /data/mysql

#安装Mysql
tar -zxvf mysql-5.5.17.tar.gz
cd mysql-5.5.17
cmake . -DCMAKE_INSTALL_PREFIX=/usr/local/mysql -DMYSQL_DATADIR=/data/mysql -DSYSCONFDIR=/etc/
make && make install

#设置Mysql
#在support-files目录中有五个配置信息文件：
#my-small.cnf (内存<=64M)
#my-medium.cnf (内存 128M)
#my-large.cnf (内存 512M)
#my-huge.cnf (内存 1G-2G)
#my-innodb-heavy-4G.cnf (内存 4GB)
cd /usr/local/mysql
cp ./support-files/my-medium.cnf /etc/my.cnf
vi /etc/my.cnf
#在 [mysqld] 段增加
datadir = /data/mysql
wait-timeout = 30
max_connections = 512
default-storage-engine = MyISAM
#在 [mysqld] 段修改
max_allowed_packet = 16M

#生成授权表
cd /usr/local/mysql
./scripts/mysql_install_db --user=mysql

#更改密码
/usr/local/mysql/bin/mysqladmin -u root password 123456

#开启mysql
/usr/local/mysql/bin/mysqld_safe &

#测试连接mysql
/usr/local/mysql/bin/mysql -u root -p 123456
show databases;
exit;

#设置开机启动
vi /etc/rc.d/rc.local

#加入
/usr/local/mysql/bin/mysqld_safe &
复制代码


安装PHP


复制代码
    #1
    tar -zxvf libiconv-1.14.tar.gz && cd libiconv-1.14/
    ./configure --prefix=/usr/local
    make && make install && cd ../

    #2
    tar -zxvf libmcrypt-2.5.8.tar.gz && cd libmcrypt-2.5.8/
    ./configure &&  make && make install
    /sbin/ldconfig && cd libltdl/ && ./configure --enable-ltdl-install
    make && make install && cd ../

    #3
    tar -zxvf mhash-0.9.9.9.tar.gz && cd mhash-0.9.9.9/ && ./configure
    make && make install && cd ../

    #4
    ln -s /usr/local/lib/libmcrypt.la /usr/lib/libmcrypt.la
    ln -s /usr/local/lib/libmcrypt.so /usr/lib/libmcrypt.so
    ln -s /usr/local/lib/libmcrypt.so.4 /usr/lib/libmcrypt.so.4
    ln -s /usr/local/lib/libmcrypt.so.4.4.8 /usr/lib/libmcrypt.so.4.4.8
    ln -s /usr/local/lib/libmhash.a /usr/lib/libmhash.a
    ln -s /usr/local/lib/libmhash.la /usr/lib/libmhash.la
    ln -s /usr/local/lib/libmhash.so /usr/lib/libmhash.so
    ln -s /usr/local/lib/libmhash.so.2 /usr/lib/libmhash.so.2
    ln -s /usr/local/lib/libmhash.so.2.0.1 /usr/lib/libmhash.so.2.0.1
    ln -s /usr/local/bin/libmcrypt-config /usr/bin/libmcrypt-config

    #5
    tar -zxvf mcrypt-2.6.8.tar.gz &&cd mcrypt-2.6.8/
    /sbin/ldconfig
    ./configure
    make && make install && cd ../

    #6
    tar -xjvf php-5.3.8.tar.bz2
    cd php-5.3.8

    ./configure --prefix=/usr/local/php \
    --with-config-file-path=/usr/local/php/etc \
    --with-iconv-dir=/usr/local/ --with-freetype-dir \
    --with-mysql=/usr/local/mysql \
    --with-mysqli=/usr/local/mysql/bin/mysql_config \
    --with-jpeg-dir --with-png-dir --with-zlib \
    --with-mhash --enable-sockets --enable-ftp \
    --with-libxml-dir --enable-xml --disable-rpath \
    --enable-safe-mode --enable-bcmath \
    --enable-shmop --enable-sysvsem \
    --enable-inline-optimization --with-curl \
    --with-curlwrappers \
    --enable-mbregex \
    --enable-mbstring --with-mcrypt --with-gd \
    --enable-gd-native-ttf --with-openssl --with-mhash \
    --enable-pcntl --enable-sockets --with-ldap --with-ldap-sasl \
    --enable-fpm \
    --with-xmlrpc --enable-zip --enable-soap \
    --without-pear \


    ./configure --enable-fpm --enable-mbstring --with-mysqli=


    make ZEND_EXTRA_LIBS='-liconv'

#注意这里容易出现 make: *** [ext/phar/phar.php] 错误 127

#出现mysql client解决方法
#ln -s /usr/local/mysql/lib/libmysqlclient.so /usr/lib/
#ln -s /usr/local/mysql/lib/libmysqlclient.so.18 /usr/lib/libmysqlclient.so.18

#或者chmod: 无法访问 “ext/phar/phar.phar”: 没有那个文件或目录
#make: [ext/phar/phar.phar] 错误 1 (忽略)
#解决方法在编译的时候加--without-pear参数

#如果还不行,make的时候不添加 ZEND_EXTRA_LIBS='-liconv' 参数

make install

#选择PHP.ini配置文件
cp php.ini-production /usr/local/php/etc/php.ini
复制代码


更改PHP-FPM

    #添加WWW用户
    /usr/sbin/groupadd www && /usr/sbin/useradd -g www www
    mkdir -p /var/log/nginx && chmod +w /var/log/nginx &&chown -R www:www /var/log/nginx
    mkdir -p /data/www && chmod +w /data/www && chown -R www:www /data/www

    cp /usr/local/php/etc/php-fpm.conf.default /usr/local/php/etc/php-fpm.conf
    vi /usr/local/php/etc/php-fpm.conf

    #去掉/更改 配置文件中的;
    pm.max_children = 64
    pm.start_servers = 20
    pm.min_spare_servers = 5
    pm.max_spare_servers = 35
    pm.max_requests = 1024
    user = www
    group = www

    #检查语法是否正确
    /usr/local/php/sbin/php-fpm -t
    #出现NOTICE: configuration file /usr/local/php/etc/php-fpm.conf test is successful 测试成功
    /usr/local/php/sbin/php-fpm &
    #设置开机启动
    vi /etc/rc.d/rc.local
    #在行末加入
    /usr/local/php/sbin/php-fpm &

    #返回安装包目录
    cd /web



安装Nginx

    #安装pcre库
    tar -zxvf pcre-8.13.tar.gz && cd pcre-8.13/ && ./configure
    make && make install && cd ../

    #安装Nginx
    tar -zxvf nginx-1.0.9.tar.gz && cd nginx-1.0.9 &&
    ./configure --user=www --group=www \
    --prefix=/usr/local/nginx \
    --sbin-path=/usr/local/nginx/sbin/nginx \
    --conf-path=/usr/local/nginx/conf/nginx.conf \
    --with-http_stub_status_module \
    --with-http_ssl_module \
    --with-pcre \
    --lock-path=/var/run/nginx.lock \
    --pid-path=/var/run/nginx.pid

    make && make install && cd ../

    #更改配置
    vi /usr/local/nginx/conf/nginx.conf

    #修改一些参数,别直接替换文件,这只是一部分
    user www

    events {
        use epoll;
        worker_connections  1024;
    }

    location ~ \.php$ {
                root           html;
                fastcgi_pass   127.0.0.1:9000;
                fastcgi_index  index.php;
                fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
                include        fastcgi_params;
            }

    #注意这里
    #$document_root$fastcgi_script_name;
    #检测配置文件
    /usr/local/nginx/sbin/nginx -t

    #提示表示成功
    #nginx: the configuration file /usr/local/nginx/conf/nginx.conf syntax is ok
    #nginx: configuration file /usr/local/nginx/conf/nginx.conf test is successful

    #开启Nginx
    /usr/local/nginx/sbin/nginx &
    #平滑重启Nginx
    /usr/local/nginx/sbin/nginx -s reload

    #添加开机启动
    vi /etc/rc.d/rc.local
    #最后移行加入
    /usr/local/nginx/sbin/nginx

    #测试
    cd /usr/local/nginx/html/
    touch index.php
    vi /usr/local/nginx/html/index.php

    <?php
    phpinfo();
    ?>
