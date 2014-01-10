#Nginx


## 配置

**运行用户**

    user www-data;   

**启动进程,通常设置成和cpu的数量相等**

    worker_processes  1;

**全局错误日志及PID文件**

    error_log  /var/log/nginx/error.log;  
    pid        /var/run/nginx.pid;  

**events #工作模式及连接数上限**

    events {  
        use   epoll;             
        # 单个后台 worker process 进程的最大并发链接数  
        worker_connections  1024;
        # multi_accept on;  
    }  

epoll是多路复用IO(I/O Multiplexing)中的一种方式,但是仅用于linux2.6以上内核,可以大大提高nginx的性能。


### http


### server

#### 404 

为php节点添加

    try_files $uri =404;



## htpasswd HTTP验证

**安装** `apache2-utils`

使用apache htpasswd来创建用户登陆访问

**创建**一个用户密码文件

	cd /var/www/mywebsite.com/
	htpasswd -c htpasswd.mywebsite username
	# username 密码输入
	# 继续创建
	htpasswd htpasswd.mywebsite user2
	...

**修改** `htpasswd.mywebsite`文件的对应权限,并保证nginx用户可以访问该文件的路径。

	sudo chown http:http htpasswd.mywebsite
	sudo chmod 640 htpasswd.mywebsite

上面的 `http:http` 为 nginx的进程用户，修改为自己系统的默认。 

**添加到配置**

	auth_basic "Restricted 限制注释";
    auth_basic_user_file /var/www/mywwebsite.com/htpasswd.mywebsite;

## error 

查看错误日志`/var/log/nginx/error.log`

错误502

**upstream sent too big header while reading response header from upstream**

    fastcgi_buffer_size 128k;
    fastcgi_buffers 32 32k;

`fastcgi_*` 可以理解成nginx接受client请求时的响应使用的。proxy是nginx作为client转发时使用的，如果header过大，超出了默认的1k，就会引发上述的upstream sent too big header。



