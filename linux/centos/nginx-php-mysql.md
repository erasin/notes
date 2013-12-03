# Centos 安装 nginx php mysql

os: centos 6.2

## nginx 

参考 <http://wiki.nginx.org/Install>

CentOS:

追加节点到 `/etc/yum.repos.d/CentOS-Base.repo`

	[nginx]
	name=nginx repo
	baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
	gpgcheck=0
	enabled=1

下载

	yum makecache
	yum install nginx

新建配置

	ls /etc/nginx/conf.d
	vim /etc/nginx/conf.d/site.conf

php 模块 支持 PATHINFO

	location ~ ^.+\.php{
		#fastcgi_pass  127.0.0.1:9000;
		fastcgi_pass  unix:/var/run/php-fpm/php-fpm.sock;
		fastcgi_index index.php;
		fastcgi_split_path_info ^((?U).+\.php)(/?.+)$;
		fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
		fastcgi_param PATH_INFO $fastcgi_path_info;
		fastcgi_param PATH_TRANSLATED $document_root$fastcgi_path_info;
		include       fastcgi_params;
	}

另外：

	location ~ \.php/?.*$ {
		#fastcgi_pass  127.0.0.1:9000;
		fastcgi_pass   unix:/var/run/php-fpm/php-fpm.sock;
		fastcgi_index  index.php;
		set $path_info "";
		set $real_script_name $fastcgi_script_name;
		if ($fastcgi_script_name ~ "^(.+?\.php)(/.+)$") {
			set $real_script_name $1;
			set $path_info $2;
		}
		fastcgi_param  SCRIPT_FILENAME $document_root$fastcgi_script_name;
		fastcgi_param  SCRIPT_NAME $real_script_name;
		fastcgi_param  PATH_INFO $path_info;
		include        fastcgi_params;
	}



建议：创建 `/wwwroot` 文件夹

	site   --->  /wwwroot/site
	log    --->  /wwwroot/log
	...

## php

下载

	yum install php php-fpm php-mysql php-mcrypt php-php-gd php-xml php-mbstring php-ldap php-pear php-xmlrpc   

修改配置

	vim /etc/php.ini 
	vim /etc/php-fpm.ini

## mysql

	yum install mysql mysql-devel mysql-connector-odbc mysql-embedded mysql-server

密码设置

	mysqladmin -u root password ‘newpassword’

配置

	vim /etc/my.cnf

## 开启服务

	chkconfig nginx on
	chkconfig php-fpm on
	chkconfig mysqld on

	service nginx start
	service php-fpm start
	service mysqld start

## 配置文件

	/etc/nginx/
	/etc/php.ini
	/etc/php-fpm.ini
	/etc/my.cnf



