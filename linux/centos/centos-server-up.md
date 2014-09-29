# centos服务器配置流程



使用http用户服务： 

    user  : ngxin 
    group : nginx

普通用户根据项目来定： (举例用户 `jibun`)

    user  : jibun
    group : nginx,users

## 安装配置

### 关于源

* [阿里云镜像源](http://mirrors.aliyun.com/help/centos)
* [163镜像源](http://mirrors.163.com/.help/centos.html)
* [yum 命令](yum.md)

### PHP

**安装：** 

```bash    
yum install php php-fpm php-gd php-mcrypt php-pear php-mysql php-mysqlite
##或者 
yum install php php-*
```

检查版本： 

    $ php -v 

**服务：**

    # service php-fpm start|restart|stop

**开启启动：**

    chkconfig php-fpm on

**配置：**

> 文件
    
    /etc/php.ini       
    /etc/php.d/        
    /etc/php-fpm.conf
    /etc/php-fpm.d/  

`/etc/php.ini` 配置要点

    ; 时间区域
    date.timezone = "Asia/Shanghai"

必要的库，有的在对应的 `/etc/php.d/` 文件夹中

    ;extension=iconv.so
    ;extension=mcrypt.so    
    ;extension=sockets.so
    ;extension=soap.so

`/etc/php-fpm.conf`

`/etc/php-fpm.d/www.conf` fpm的配置文件修改要点

端口

    listen = 127.0.0.1:9000
    ; listen = /var/run/php-fpm/php-fpm.sock

用户写入文件的 user:group 和权限

    listen.owner = nginx
    listen.group = nginx
    listen.mode = 0666  
    user = nginx
    group = nginx

执行时限**

    request_terminate_timeout = 180s

**权限：**

对session.savepath 的组用户写权限

    # chown root:nginx -R /var/lib/php/

### Mysql

**安装：** 

    yum install mysql mysql-server

**检查：**

    mysql --version

**开启启动：**

    chkconfig mysqld on

**设定初始密码：**
    
    # service mysqld start
    # mysqladmin -u root password '123456'

**重新设定密码：**

开启安全模式：

    # service mysqld stop
    # mysqld_safe --skip-grant-tables&
    # mysql --user=root mysql
    > update user set Password=PASSWORD('new-password');
    > flush privileges;
    > exit;
    # killall mysqld_safe&
    # service mysqld start

* 停止服务
* root权限开启mysql安全模式
* 进入mysql数据库
* 更新表，刷新，退出
* 杀死mysqld 安全进程
* 开启mysql服务

**配置：** 

> 文件：

    /etc/my.cnf

### nginx 

添加内容到 `/etc/yum.repos.d/CentOS-Base.repo` 

    [nginx]
    name=nginx repo
    baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
    gpgcheck=0
    enabled=1

**安装：** 

    yum install nginx

**开启启动：**

    chkconfig nginxd on

**查看：**

    nginx -v 

**配置：**

    /etc/nginx/conf.d/

以 fb.geely.com 为模板

    server {
        listen       80;
        server_name  fb.geely.com;
        charset utf-8;
        root        /home/nginx/fb.geely.com;
        access_log  /home/nginx/logs/fb.geely.com.log  main;
        error_log   /home/nginx/logs/fb.geely.com_error.log;

        location /{
            index  index.php index.html index.htm;
            #if (!-e $request_filename){
            #        rewrite ^/(.*)$ /index.php/$1 last;
            #}
        }

        location ~ ^.+\.php{
            fastcgi_buffer_size 128k;
            fastcgi_buffers 32 32k;
            #fastcgi_pass  unix:/var/run/php-fpm/php-fpm.sock;
            fastcgi_pass  127.0.0.1:9000; 
            fastcgi_index index.php;
            fastcgi_split_path_info ^((?U).+\.php)(/?.+)$;
            fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
            fastcgi_param PATH_INFO $fastcgi_path_info;
            fastcgi_param PATH_TRANSLATED $document_root$fastcgi_path_info;
            include       fastcgi_params;
        }

        #error_page  404              /404.html;
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   /usr/share/nginx/html;
        }
    }


SSL HTTPS 配置模板

    server {
        listen       443 ssl;
        server_name  fb.geely.com;

        # ssl on;
        ssl_certificate     /home/nginx/fb.geely.com/ssl/fb.geely.com.crt;
        ssl_certificate_key /home/nginx/fb.geely.com/ssl/fb.geely.com.key;

        ssl_session_cache shared:SSL:1m;
        ssl_session_timeout  5m;

        ssl_ciphers  HIGH:!aNULL:!MD5;
        ssl_prefer_server_ciphers   on;

        root        /home/nginx/fb.geely.com;

        access_log  /home/nginx/logs/fb.geely.com.log  main;
        error_log   /home/nginx/logs/fb.geely.com_error.log;

        location / {
            index index.html index.php; 
        }
        
        location ~ ^.+\.php{
            fastcgi_pass  127.0.0.1:9000;
            #fastcgi_pass  unix:/var/run/php5-fpm.sock;
            fastcgi_index index.php;
            fastcgi_split_path_info ^((?U).+\.php)(/?.+)$;
            fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
            fastcgi_param PATH_INFO $fastcgi_path_info;
            fastcgi_param PATH_TRANSLATED $document_root$fastcgi_path_info;
            include       fastcgi_params;
        }

        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
                root   /usr/share/nginx/html;
        }

        location ~ /\.ht {
                deny  all;
        }
    }

> 参看 [SSL CA证书](/linux/server/nlb.md)

## 安全

自动检测服务并重启脚本：

    #!/bin/bash
    # file path /home/jibun/shell/checkservice.sh

    service_log=/home/jibun/logs/server.log

    ng=`service nginx status|awk '{print $5}'`
    php=`service php-fpm status|awk '{print $5}'`
    mysql=`service mysqld status|awk '{print $5}'`

    if [[ $ng != *"running"* ]]; then
        echo "---------------------------------" >> $service_log
        date '+%Y-%m-%d %H:%M' >> $service_log
        service nginx start >> $service_log
    fi

    if [[ $php != *"running"* ]]; then
        echo "---------------------------------" >> $service_log
        date '+%Y-%m-%d %H:%M' >> $service_log
        service php-fpm start >> $service_log
    fi

    if [[ $mysql != *"running"* ]]; then
        echo "---------------------------------" >> $service_log
        date '+%Y-%m-%d %H:%M' >> $service_log
        service mysqld start >> $service_log
    fi

> 注意 mysql的服务默认为 mysqld ,有部分为mysql，自行斟酌


## 创建普通用户

创建用户`jibun`并创建密码

    useradd -G nginx -g users jibun 
    passwd jibun 

## 服务器路径

创建服务器文件路径。

    /home/nginx         # nginx服务器所在文件夹
    /home/nginx/logs    # 存放日志
    /home/nginx/domain  # 对应的项目

执行过程

    # cd home
    # mkdir nginx 
    # chown nginx:nginx -R nginx
    # chmod g+w nginx
    # cd nginx 
    # mkdir logs
    # su jibun
    $ mkdir www.domain.com

上传对应的 `www.domain.com` 文件夹即可。  

用户`jibun`文件夹：

    /home/jibun            # home文件夹
    /home/jibun/shell      # 定时执行脚本
    /home/jibun/backdir    # 备份位置
    /home/jibun/gitsource  # git源库（防篡改之用）

## 参看服务器方案

* 参看 [linux 下同步方案以及站点文件的防篡改](/linux/server/rsync)
* 参看 [mysql master slave 主从同步](/linux/server/mysql-master-slave)
* 参看 [负载均衡技术](/linux/server/nlb.md)
* 参看 [压力测试 和 网速测试](/linux/server/stress-web-test-ab.md)