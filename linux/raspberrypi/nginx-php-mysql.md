#Raspberry nginx php5-fpm mysql服务器配置
**title:** Raspberry Pi nginx php服务配置  
**tags:** nginx,php,mysql,raspberrypi,raspbian,树莓派  

**OS:** Raspbian “wheezy”

## 下载服务
下载 nginx 和 Php相关服务

    # apt-get install nginx php5 php5-fpm php5-mysql php5-gd php5-curl 
    # apt-get install mysql 

在下载mysql服务的时候会提示设置密码,最好不要留空。

## 修改配置

    # vim /etc/nginx/nginx.conf

Raspbian 是 Debian的派生 ，所以默认的 user 为 www-data,这里的配置修改，请参考 [nginx手册][nginx-doc].  
可以建立站点配置文件在 site-enabled 中

### php5-fpm 在 nginx server 配置中

例子使用了 `path_info`

	location ~ ^.+\.php{
		root           /var/www
		# fastcgi_pass  127.0.0.1:9000;
		fastcgi_pass   unix:/var/run/php5-fpm.sock;
		fastcgi_index  index.php;
		fastcgi_split_path_info ^((?U).+\.php)(/?.+)$;
		fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
		fastcgi_param PATH_INFO $fastcgi_path_info;
		fastcgi_param PATH_TRANSLATED $document_root$fastcgi_path_info;
		include        fastcgi_params;
	}

其中的 `root /var/www;` 这里一定要使用绝对的路径，否则在访问php文件的时候有可能出现文件未找到错误(File not found)。

## 启动服务
系统默认已经在 `/etc/init.d/`中建立的启动脚本。  
使用脚本来对服务进行 start|restart|stop 操作

    # /etc/init.d/nginx start
    # /etc/init.d/php5-fpm start
    # /etc/init.d/mysql start

现在已经可以是哦那个服务了吧。

###默认将nginx等服务加入到init开机启动项

    # update-rc.d nginx defaults
    # update-rc.d php5-fpm defaults
    # update-rc.d mysql defaults

我这里的Raspbian 默认有安装apache2 所以要禁止其开机启动
    
    # update-rc.d apache2 remove 

查看是否有 apache2 的开机启动项

    $ ls /etc/rc2.d/|grep apache

如果有的话，执行前面的操作将其移除。

开机启动项目 可以参考 Debian 的 [update-rc.d](update-rc.d) 命令。

[nginx-doc]:http://nginx.org/cn/docs/
