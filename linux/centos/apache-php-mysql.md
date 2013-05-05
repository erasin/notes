#CentOS 5.5使用yum安装Apache+PHP+Mysql



## 一、安装centos5.5操作系统  

CentOS 开发社区已发布了新的 5.5 版本。CentOS 5.5 基于 Red Hat Enterpris Linux 5.5.0，其中包括 Kernel 2.6.18、Apache 2.2、PHP 5.1.6、MySQL 5.0、PostgreSQL 8、GNOME 2.16、KDE 3.5、OpenOffice.org 2.3、Firefox 3.0、Evolution 2.12 等等。此外，CentOS 5.3 更新了美工设计，并根据用户的请求恢复了 Contrib 仓库。  

CentOS 5.5 支持 i386 及 x86\_64 架构，其 ISO 映像可从以下地址获取。<http://www.centos.org/>

## 二、系统安装  

安装CentOS 5.5做服务器 

## 三、安装Apache+php+Mysql  

用yum安装前.先是替换为中国CentOS5.3镜像服务器!快速好用的yum更新源：http://www.wljcz.com/html/caozuoxitong/Linux/2009/0726/410.html 此文章有详细的介绍，按照提供的步骤修改好更新源后，就可以方便的用yum命令快速安装软件了.  
   
### 1、更新系统内核到最新. 

    yum -y update  

安装Apahce, PHP, Mysql, 以及php连接mysql库组件  
   
    yum -y install httpd php mysql mysql-server php-mysql  
   
### 2、安装mysql扩展  
   
    yum -y install mysql-connector-odbc mysql-devel libdbi-dbd-mysql  

或一次性粘贴安装:  
   
    yum -y install httpd php mysql mysql-server php-mysql httpd-manual mod_ssl mod_perl mod_auth_mysql php-mcrypt php-gd php-xml php-mbstring php-ldap php-pear php-xmlrpc mysql-connector-odbc mysql-devel libdbi-dbd-mysql  
   
### 3、设置mysql数据库root帐号密码。  
   
    mysqladmin -u root password ‘newpassword’  

"newpassword" 代表的是你想要设置的密码，新安装的mysql的root根用户密码默认为空，设置密码后可以让mysql数据库更安全  
   
    mysql -u root -p （此时会要求你输入刚刚设置的密码，输入后回车即可）  
    mysql> DROP DATABASE test; （删除test数据库）  
    mysql> DELETE FROM mysql.user WHERE user = ”; （删除匿名帐户）  
    mysql> FLUSH PRIVILEGES; （重载权限）  
   
4、按照以上的安装方式, 配置出来的默认站点目录为/var/www/html/新建一个php脚本:  
   
    <?php phpinfo();?>  

## 四、 配置防火墙  

添加允许访问HTTP、FTP端口  

    iptables -I RH-Firewall-1-INPUT -m state --state NEW -m tcp -p tcp --dport 21 -j ACCEPT  
    iptables -I RH-Firewall-1-INPUT -m state --state NEW -m tcp -p tcp --dport 80 -j ACCEPT  

重启 iptables: 

    service iptables restart  

## 五、安装phpMyAdmin  

进入phpMyAdmin官方下载phpMyAdmin,3.1以上需php 5.2以上，上传到你的网站目录下，然后进行配置。只需几步即可搞定。  
config.sample.inc.php更名为config.inc.php，然后打开config.inc.php文件，进行以下修改;  
代码:  

    // $cfg['Servers'][$i]['controluser'] = ‘pma’;  
    // $cfg['Servers'][$i]['controlpass'] = ‘pmapass’;  
    // $cfg['Servers'][$i]['pmadb'] = ‘phpmyadmin’;  
    // $cfg['Servers'][$i]['bookmarktable'] = ‘pma_bookmark’;  
    // $cfg['Servers'][$i]['relation'] = ‘pma_relation’;  
    // $cfg['Servers'][$i]['table_info'] = ‘pma_table_info’;  
    // $cfg['Servers'][$i]['table_coords'] = ‘pma_table_coords’;  
    // $cfg['Servers'][$i]['pdf_pages'] = ‘pma_pdf_pages’;  
    // $cfg['Servers'][$i]['column_info'] = ‘pma_column_info’;  
    // $cfg['Servers'][$i]['history'] = ‘pma_history’;  
    // $cfg['Servers'][$i]['designer_coords'] = ‘pma_designer_coords’;  

去掉每行前面的//注释  

    $cfg['blowfish_secret'] = ”; |修改为| $cfg['blowfish_secret'] = ‘http’;  
    $cfg['Servers'][$i]['controluser'] = ‘pma’; |把’pma’修改为你的帐号|$cfg['Servers'][$i]['controlpass'] = ‘pmapass’; |把’pmapass设置为你的mysql登录密码  
    $cfg['blowfish_secret'] = ”; | 添加短语密码例如：$cfg['blowfish_secret'] = ‘onohot’;  

六、安装php的扩展  

代码:  

    yum -y install php-gd php-xml php-mbstring php-ldap php-pear php-xmlrpc  

七、安装apache扩展  

代码:  

    yum -y install httpd-manual mod\_ssl mod\_perl mod\_auth\_mysql  

到此为止centos5.3下的php环境基本就配置完成了，用命令启动服务即可使用。  

## 八、如果想升级PHP的话  

提及php需要升级到5.2以上，而centos 5.x目前提供php版本为5.1.6，  
通过以下方法升级PHP到5.2比较方便，现推荐给大家。  
先将以下地址导入。  

    # rpm --import http://www.jasonlitka.com/media/RPM-GPG-KEY-jlitka  

修改 `/etc/yum.repos.d/CentOS-Base.repo` 增加下面信息  

    [utterramblings]  
    name=Jason's Utter Ramblings Repo  
    baseurl=http://www.jasonlitka.com/media/EL$releasever/$basearch/  
    enabled=1  
    gpgcheck=1  
    gpgkey=http://www.jasonlitka.com/media/RPM-GPG-KEY-jlitka  

执行命令，自动升级。  

    yum update php -y  
    yum install libmcrypt -y  

### 1. 更新系统内核到最新.  

    [root@linuxfei ~]#yum -y update  

系统更新后,如果yum安装时提示错误信息,请执行以下命令修复.  

    [root@linuxfei ~]#rpm –import /etc/pki/rpm-gpg/RPM-GPG-KEY*

### 2. 安装Apahce, PHP, Mysql, 以及php连接mysql库组件


    [root@linuxfei ~]#yum -y install httpd php mysql mysql-server php-mysql  
    //安装mysql扩展  
    [root@linuxfei ~]#yum -y install mysql-connector-odbc mysql-devel libdbi-dbd-mysql  
    //安装php的扩展  
    [root@linuxfei ~]# yum -y install php-gd php-xml php-mbstring php-ldap php-pear php-xmlrpc  
    //安装apache扩展  
    [root@linuxfei ~]#yum -y install httpd-manual mod\_ssl mod\_perl mod\_auth\_mysql  
    一次性粘贴安装:  
    [root@linuxfei ~]# yum -y install httpd php mysql mysql-server php-mysql httpd-manual mod\_ssl mod\_perl mod\_auth\_mysql php-mcrypt php-gd php-xml php-mbstring php-ldap php-pear php-xmlrpc mysql-connector-odbc mysql-devel libdbi-dbd-mysql  

### 3. 启动服务配置  

    [root@linuxfei ~]# /sbin/chkconfig httpd on [设置apache为自启动]  
    [root@linuxfei ~]# /sbin/chkconfig –-add mysqld [mysql服务]  
    [root@linuxfei ~]# /sbin/chkconfig mysqld on [mysqld服务]  
    [root@linuxfei ~]# /sbin/service httpd start [自启动 httpd 服务]  
    [root@linuxfei ~]# /sbin/service mysqld start [自启动mysqld服务]  

### 4.设置mysql数据库root帐号密码。  

    [root@linuxfei ~]# mysqladmin -u root password 'linuxfei'    [引号内填密码]  
    [root@linuxfei ~]# mysql -u root -p               ← 通过空密码用root登录  
    Enter password:linuxfei ← 在这里输入密码  
    Welcome to the MySQL monitor. Commands end with ; or \g. ← 确认用密码能够成功登录  
    Your MySQL connection id is 5 to server version: 4.1.20  
    Type 'help;' or '\h' for help. Type '\c' to clear the buffer.  

### 5.安装phpmyadmin  

    [root@linuxfei /]# wget http://gd2.down.chinaz.com:808/数据管理/phpMyAdmin-3.tar.gz  
    --2010-03-23 16:38:18-- http://gd2.down.chinaz.com:808/??????/phpMyAdmin-3.x.tar.gz  
    Resolving gd2.down.chinaz.com... 121.11.80.154  
    Connecting to gd2.down.chinaz.com|121.11.80.154|:808... connected.  
    HTTP request sent, awaiting response... 200 OK  
    Length: 4700100 (4.5M) [application/x-gzip]  
    Saving to: `phpMyAdmin-3.x.tar.gz'  
    100%[======================================>] 4,700,100    134K/s   in 37s     
    2010-03-23 16:38:56 (123 KB/s) - `phpMyAdmin-3.x.tar.gz' saved [4700100/4700100]  

解压phpmyadmin  

    #tar zxvf phpMyAdmin-3.x.tar.gz  
    #mv phpMyAdmin-3.3.1-all-languages /var/www/html/phpmyadmin  

将解压出来的目录移动到/var/www目录下并改名为phpmyadmin  

修改phpmyadmin根目录下的config.sample.inc.php 重命名为 config.inc.php  

打开并编辑config.inc.php  

    $cfg['blowfish_secret'] = '';  

然找到下边这两行 并把//去除  

    // $cfg['Servers'][$i]['controluser'] = 'pma';                 #mysql用户名  
    // $cfg['Servers'][$i]['controlpass'] = 'pmapass'              #mysql密码  

### apache php mysql配置  

|位置|说明
|----|----
|/etc/httpd/conf/httpd.conf     | 最主要的配置文件，不过很多其他的distribution都将这个文件拆成数个小文件，分别管理不同的参数。但是最主要配置文件还是以这个文件名为主。  
|/etc/httpd/conf.d/\*.conf      |这个事CentOS的特色之一，如果你不想修改原始配置文件httpd.conf的话，那么可以将你自己的额外参数独立出来，而启动apache时，这个文件就会被读入到主要配置文件。  
|/usr/lib/httpd/modules         |apache支持很多的模块，所以您想要使用的模块默认都放置在此目录  
|/var/www/html                  |这里是CentOS默认的“首页”所在目录。  
|/var/www/error                 |如果因为主机设置错误，或者是浏览器端要求的数据错误，在浏览器上出现的错误信息就已这个目录的默认信息为主。  
|/var/www/icons                 |提供apache的一些小图标  
|/var/www/cgi-bin               |默认给一些可执行的CGI程序放置的目录  
|/var/log/httpd                 |默认apache的日志文件都放在这里，对于流量大的网站来说，这个目录要很小心，因为这个文件很容易变的很大，您需要足够的空间哦  
|/usr/sbin/apachectl            |这是Apache的主要执行文件，这个执行文件其实是shell script,它可以主动检测系统上的一些设置值，好让您启动Apache时更简单  
|/usr/sbin/httpd                |这是主要的apache的二进制文件  
|/usr/bin/htpasswd              |当您想登陆某些网页时，需要输入账号与密码。那么Apache本身就提供一个最基本的密码保护方式。该密码的产生就是通过这个命令实现的  
|/etc/my.cnf                    |这是Mysql的配置文件，包括您想要进行mysql数据库的最佳化，或者是正对mysql进行一些额外的参数指定，都可以在这个文件里实现  
|/usr/lib/mysql                 |这个目录是MySQL数据库放置的位置，当启动任何MySQL的服务器时，请务必记得在备份时，将此目录完整的备份下来。  
|/usr/lib/httpd/modules/libphp4.so  |PHP提供给apache使用的模块，这个关系我们能否在apache网页上面设计php程序语言的最重要文件  
|/etc/httpd/conf.d/php.conf     |你要不要手动将该模块写入Httpd.conf中呢？不需要，因为系统已经主动将php设置参数写入到这个文件中了，而这个文件会在apache重新启动时被读入。  
|/etc/php.ini                   |这是PHP的主要配置文件，包括PHP能不能允许用户上传文件，能不能允许某些低安全性的标志等，都在这个配置文件中设置。  
|/etc/php.d/mysql.ini /usr/lib/php4/mysql.so    |PHP能否可以支持MySQL接口就看这两个文件了。这两个文件是由php-mysql软件提供的  
|/usr/bin/phpize /usr/include/php     |如果您以后想要安装类似PHP加速器可以让浏览速度加快的话，那么这个文件与目录就需要存在，否则加速器软件没法用。  

`httpd.conf`的基本设置  
首先，你需要在/etc/hosts内需要有个一个完整的主机名，否则在重启apache服务时，会提示找不到完整的主机名。  

针对主机环境的设置项目  

     #vi /etc/httpd/conf/httpd.conf  
     ServerTokens OS  
     # 这个项目在告诉客户端WWW服务器的版本和操作系统，不需要改编它  
     #如果你不想告诉太多的主机信息，将这个项目的OS改成Minor  
     ServerRoot "/etc/httpd"  
     #这个是设置文件的最顶层目录，通常使用绝对路径，下面某些数据设置使用相对路径时  
     #就是与这个目录设置值有关的下层目录，不需要更改它  
     ServerRoot  
     #设定Apache 安装的绝对路径  
     TimeOut  
     #设定 服务器接收至完成的最长等待时间  
     KeepAlive  
     #设定服务器是否开启连续请求功能,真实服务器一般都要开启  
     Port  
     #设定http服务的默认端口。  
     User/Group  
     #设定 服务器程序的执行者与属组,这个一般是apache  

下面我们就针对Apache做几个实验  

## Apache 服务

### 1：测试把默认网站目录改到root家目录下  

新建/root/website目录  

     #mkdir -p /root/website  
     #echo "website page" >> /root/website/index.html  
     #vi /etc/httpd/conf/httpd.conf  

找到 DocumentRoot "/var/www/html" 这一段   //apache的根目录  
把/var/www/html 这个目录改到 /root/website  
在找到  //定义apache /var/www/html这个区域  
把 /var/www/html改成/root/website  
这样我们就把apahce的默认路径改掉了  

然后重启服务  

    #service httpd restart    
    
这里在你重启服务的时候，可能会报错，说找不到目录，这个主要是由于***[selinux][selinux]***导致的  
那怎么解决呢？有2个办法，关掉selinux    

    #setenforce 0  

或者更改/root/website这个文件的selinux属性，让它匹配httpd这个服务器的要求  
怎么改？我们可以*复制/var/www/html这个目录的selinux属性*

    #chcon -R --reference /var/www/html /root/website  

然后在重启服务，之后你就看到它没有报错了  
不过你去访问localhost的时候，会发现访问拒绝 这是为什么呢？主要是因为你的/root的权限是750,ahache这个用户没有权限访问，你需要更改掉权限，可以这样改  ,这个最好修改为apache执行的用户权限

    #chmod -R 755 /root  

然后去访问 发现正常了  

使用 vhost 的时候 又是会出现端口错误

    [warn] _default_ VirtualHost overlap on port 80, the first has precedence

修改 conf/httpd.conf 

    #NameVirtualHost *:80  ==>  NameVirtualHost *:80  也就是说取消前面的注释

### 2：基于名称的虚拟主机  

需要两个域名解析到你的服务器，对应关系是  

    /var/www/server             server.example.com  
    /var/www/client             client.example.com  

当访问这两个域名时，可以分别显示出不同文件里面主页的内容  

    #echo "server page" >> /var/www/server/index.html  
    #echo "client page" >> /var/www/client/index.html  

然后我们编辑一个配置文件  

    #vi /etc/httpd/conf.d/virtual.conf //记住conf.d里面的内容也是apache的配置文件  

添加如下内容：  

    NameVirtualHost 192.168.76.133:80  

    ServerName service.example.com  
    DocumentRoot /var/www/server  

    ServerName client.example.com  
    DocumentRoot /var/www/client  

重启服务

    #service httpd restart  

这样基于名称的虚拟主机就配置好了  

如果你没有DNS你可以再你的机器上hosts文件里加记录 linux在/etc/hosts这个文件 windows在C:\windows\system32\drivers\etc\hosts文件  

加上这两行  

    127.0.0.1 server.example.com  
    127.0.0.1 client.example.com  

这样你在去测试，就会发现访问不同的域名显示不同的内容了 这样基于名称的虚拟主机就配置好了!  

### 3：基于IP地址的虚拟主机  

先添加一个临时网卡  

    #ifconfig eth0:0 192.168.76.132 //临时使用，重启后就会消失  

然后便捷virtual.conf文件  

    #vi /etc/httpd/conf.d/virtual.conf  

把内容修改为  

    #NameVirtualHost 192.168.76.133:80  
    ServerName service.example.com  
    DocumentRoot /var/www/server  

    ServerName client.example.com  
    DocumentRoot /var/www/client  

让后你在用ip访问，发现也能显示不同的内容,或者你编辑hosts文件，用域名访问也没问题  
这样基于IP地址的虚拟主机也成功了！  

### 4：别名  
在/etc/httpd/conf/httpd.conf里加入  
 
Alias /test "/root/website/"    // 别名 这样你用192.168.76.133/test访问 也会显示192.168.76.133的页面  

这个地方需要注意的就是/test 还是/test/ 这个是用区别的 你用/test 那么你访问的时候只能用192.168.76.133/test访问   如果你用/test/ 那么192.168.76.133/test/访问，而/test将不会放你访问  

忘了这里你的先把/etc/httpd/conf.d目录里面刚刚设置的虚拟目录注释掉 不然没法访问，是因为做了虚拟目录，而httpd.conf里面的设置就无法访问 当然可以用localhost来访问，其他的访问都不行  

### 5：实现网页的资源下载  
首先添加别名 修改`/etc/httpd/conf/httpd.conf` 在Alias /test "/root/website/" 后面加入 

    Alias /down "/var/ftp/pub"  

让后对/var/ftp/pub区域设置参数  

     Options Indexes MultiViews  
     AllowOverride None  
     Order allow,deny  
     Allow from all  

在Options 加入 MultiViews   //没有index时自动列出目录文档  

然后重启服务，这样http://192.168.76.133/down/里面就可以列出/var/ftp/pub里面的文件了，试着点一个另存为，是否可以下载？ 呵呵 成功！  

### 6：.htpasswd的实现  

修改 /etc/httpd/conf/httpd.conf  
我们针对刚刚做的/var/ftp/pub来做  
加入如下信息  

    Alias /down "/var/ftp/pub/"  
    Options Indexes MultiViews  
    AllowOverride AuthConfig  
    Order allow,deny  
    Allow from all  
    AuthType Basic  
    AuthName "this is test"  
    AuthUserFile /etc/httpd/htpasswd  
    Require User test  

然后重启httpd服务,  
让后生成.htpasswd用户密码  

    htpasswd -c /etc/httpd/htpasswd test  

让后去访问192.168.76.133/down会需要密码  
这样就成功了

本文来源<http://www.qkweb.net/showa/715.html>

[selinux]: http://baike.baidu.com.cn/view/487687.htm "Security-Enhanced Linux 强制控制访问系统"
