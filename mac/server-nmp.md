# server

## 安装nginx

用 brew 一键安装 nignx：


    brew install nginx

如果需要安装其他 nginx 的版本，可以 "brew edit nginx" 打开修改 nginx 的安装信息包 formula，默认会用 vi 打开，在文件开头处修改 nginx 相应版本的下载地址就行。

brew 执行完后，nginx 就安装好了。可以用以下指令对 nginx 进行操作：


    #打开 nginx
    sudo nginx
    #重新加载配置|重启|停止|退出 nginx
    nginx -s reload|reopen|stop|quit
    #测试配置是否有语法错误
    nginx -t

    详细 nginx 帮助信息：


nginx 版本: nginx/1.2.6
用法: nginx [-?hvVtq] [-s signal] [-c filename] [-p prefix] [-g directives]

    选项:
        -?,-h           : 打开帮助信息
        -v              : 显示版本信息并退出
        -V              : 显示版本和配置选项信息，然后退出
        -t              : 检测配置文件是否有语法错误，然后退出
        -q              : 在检测配置文件期间屏蔽非错误信息
        -s signal       : 给一个 nginx 主进程发送信号：stop（停止）, quit（退出）, reopen（重启）, reload（重新加载配置文件）
        -p prefix       : 设置前缀路径（默认是：/usr/local/Cellar/nginx/1.2.6/）
        -c filename     : 设置配置文件（默认是：/usr/local/etc/nginx/nginx.conf）
        -g directives   : 设置配置文件外的全局指令

打开 nginx 后，默认的访问端口 8080，如果要改为常用的 80 端口，则要修改 "/usr/local/etc/nginx/nginx.conf" 下监听(listen)端口值。

默认的文件访问目录(root)是 "/usr/local/Cellar/nginx/1.2.6/html"（这里的1.2.6是安装的nginx的版本，文件夹名以安装的nginx版本为准）。

把 nginx 设置为开机启动运行：

    mkdir -p ~/Library/LaunchAgents
    cp /usr/local/Cellar/nginx/1.2.6/homebrew.mxcl.nginx.plist ~/Library/LaunchAgents/
    launchctl load -w ~/Library/LaunchAgents/homebrew.mxcl.nginx.plist

不过试了一下，不是超级用户登陆，而是普通用户登陆，并且监听的端口在1024以下的（例如把默认的8080端口改为了80端口），nginx 开机是启动不了。因此，要 nginx 开机启动的话，需要给予它管理员权限：

    sudo chown root:wheel /usr/local/Cellar/nginx/1.2.6/sbin/nginx
    sudo chmod u+s /usr/local/Cellar/nginx/1.2.6/sbin/nginx

## 安装mysql

    brew install mysql
在上面 brew 安装完后，会在终端上显示一些初始配置 mysql 的信息，如下:

1）初始安装 mysql 的一些配置数据库（例如：information_schema、mysql）

    mysql_install_db --verbose --user=`whoami` --basedir="$(brew --prefix mysql)" --datadir=/usr/local/var/mysql --tmpdir=/tmp
执行完后，就可以在终端里运行 "mysql"，直接进入mysql数据库里。对，不用输入密码，可以直接连接，初始默认是可以匿名访问的。超级用户 "root" 也是没设密码，要设密码的话可以执行下面指令

    /usr/local/opt/mysql/bin/mysqladmin -u root password 'new-password'
现在访问 mysql 还是不用密码就可以连接，如果要设置一些登陆密码的安全访问限制，则需执行下面的 mysql安全安装指令

    /usr/local/opt/mysql/bin/mysql_secure_installation
主要是设置修改root密码（设置过了可以不用设置，略过）、删除匿名访问、删除root网络访问、删除test数据库。指令执行完后，登陆mysql就需要密码验证了

    mysql -u root -p
开机启动 mysql


    mkdir -p ~/Library/LaunchAgents/
    cp /usr/local/Cellar/mysql/5.5.28/homebrew.mxcl.mysql.plist ~/Library/LaunchAgents/
    launchctl load -w ~/Library/LaunchAgents/homebrew.mxcl.mysql.plist

如果要停止 mysql 服务则：


    launchctl unload ~/Library/LaunchAgents/homebrew.mxcl.mysql.plist

设置 MySQL 的配置文件 my.cnf（默认在 /usr/local/Cellar/mysql/5.5.28/my.cnf，最好把配置文件移到 /etc/my.cnf、/etc/mysql/my.cnf、/usr/local/etc/my.cnf、~/.my.cnf 这些位置之一，避免以后升级版本后 MySQL 配置被覆盖）：

```
[client]
port = 3306

[mysqld]
port = 3306
server_id = 1
character-set-server = utf8
default-storage-engine = INNODB
socket = /tmp/mysql.sock
skip-external-locking

key_buffer = 16K
query_cache_limit = 256K
query_cache_size = 4M
max_allowed_packet = 1M
table_open_cache = 128          # 表缓存的数目，一般是 max_connections 的倍数

max_connections = 32
thread_concurrency = 2

sort_buffer_size = 64K
read_buffer_size = 256K
read_rnd_buffer_size = 256K
net_buffer_length = 2K
thread_stack = 64K

#skip-bdb

sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES



[mysqldump]
quick
max_allowed_packet = 16M



[mysql]
no-auto-rehash
default-character-set = utf8


[isamchk]
key_buffer = 8M
sort_buffer_size = 8M



[myisamchk]
key_buffer = 8M
sort_buffer_size = 8M



[mysqlhotcopy]
interactive-timeout
```

至此，mysql安装完毕。

## 安装php-fpm

Mac是预装了php，不过很多扩展都没安装，目测最多只能在终端里执行下php指令，所以我选择重新安装php。
由于 brew 默认是没有 php 安装，所以要使用 “brew tap” 来安装 brew 的第三方程序包，
这里使用 josegonzalez 的php安装包，具体操作如下：

    brew tap homebrew/dupes
    brew tap josegonzalez/homebrew-php

执行完后，就可以用 brew 安装php了。这里php有几个版本可以安装，具体可以执行 "brew search php"
查看一下有什么php版本可以安装，一般会有“php52、php53、php54”版本，我安装的是最新的php5.4版本。
由于PHP5.4版本已经内嵌了 FPM（FastCGI Process Manager），在安装选项里标明就行，本人 php 的安装配置指令如下：

    brew install php54 --with-imap --with-tidy --with-debug --with-pgsql --with-mysql --with-fpm

更多的安装选项可以通过 "brew options php54" 查看。指令执行完后，php 跟 php-fpm 就安装好了。

由于是重装php，之前系统预装的php还没卸载，因此在终端调用php时，还是以之前系统的php版本做解析，
所以这里需要修改path，指定 php 的解析路径。在~/.bashrc（没有则创建）最后加入一行：

    export PATH="$(brew --prefix php54)/bin:$PATH"

添加之后再执行一下source，使之生效

    source ./.profile

OK，php-fpm安装完成。

要修改配置 php 或者 php-fpm 的话，可以修改 "/usr/local/etc/php/5.4/php.ini" 、 "/usr/local/etc/php/5.4/php-fpm.conf"。

启动 php-fpm 的话就直接在终端里执行 "php-fpm"，默认打开 php-fpm 会显示一个状态 shell 出来，也可以把 php-fpm 的配置文件里的 "daemonize = no" 改为 "daemonize = yes"，就会以后台守护进程的方式启动，对于刚修改的配置文件，可以执行 "php-fpm -t" 来检测配置有没有问题。

开机启动 php-fpm（下面的 5.4.9 是当前安装 php 的具体版本号）：

    mkdir -p ~/Library/LaunchAgents
    cp /usr/local/Cellar/php54/5.4.9/homebrew-php.josegonzalez.php54.plist ~/Library/LaunchAgents/
    launchctl load -w ~/Library/LaunchAgents/homebrew-php.josegonzalez.php54.plist

为了方便，写了个启动、关闭、重启 php-fpm 的 shell 脚本：

```
#!/bin/sh

param=$1

start()
{
    fpms=`ps aux | grep -i "php-fpm" | grep -v grep | awk '{print $2}'`
    if [ ! -n "$fpms" ]; then
        php-fpm
        echo "PHP-FPM Start"
    else
        echo "PHP-FPM Already Start"
    fi
}

stop()
{
    fpms=`ps aux | grep -i "php-fpm" | grep -v grep | awk '{print $2}'`
    echo $fpms | xargs kill -9

    for pid in $fpms; do
        if echo $pid | egrep -q '^[0-9]+$'; then
            echo "PHP-FPM Pid $pid Kill"
        else
            echo "$pid IS Not A PHP-FPM Pid"
        fi
    done
}

case $param in
    'start')
        start;;
    'stop')
        stop;;
    'restart')
        stop
        start;;
    *)
        echo "Usage: ./phpfpm.sh start|stop|restart";;
esac
```
设置 Nginx 的 PHP-FPM 配置

打开 nginx 默认注释掉的php location设置，修改如下（具体配置参数，例如路径，这里以我本地安装为准）：

    location ~ \.php$ {
        fastcgi_intercept_errors on;
        fastcgi_pass   127.0.0.1:9000;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  /usr/local/Cellar/nginx/1.2.6/html$fastcgi_script_name;
        include        /usr/local/etc/nginx/fastcgi_params;
    }

OK，这样就可以在访问目录下（默认是/usr/local/Cellar/nginx/1.2.6/html）执行 php 文件了。嗯，赶快输出一下 "phpinfo()" 吧：）

一些问题


    Error: SHA256 mismatch
    Expected: 3fe780e5179e90c4d37276e79acc0d0692f1bc0911985af694b92c664c0ef3c4
    Actual: f9dbbb4e5ecd98010a3c4686d0713dcda6a77223fb9d05537089b576ab8f7fdd
    Archive: /Library/Caches/Homebrew/php54-5.4.28
    To retry an incomplete download, remove the file above.

这里是 PHP 5.4.28 源代码地址出错的问题，修改文件 /usr/local/Library/Taps/josegonzalez/homebrew-php/Formula/abstract-php-version.rb 下 Php54Defs 的源码地址 PHP_SRC_TARBALL 的值，这里改用 http://mirrors.sohu.com/php/php-5.4.28.tar.bz2 这个地址。
