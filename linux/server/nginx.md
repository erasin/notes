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


    if ($http_host !~ "^www\.xingtouzi\.com$") {
        rewrite ^(.*) http://www.xingtouzi.com$1 permanent;
    }
    if (!-e $request_filename){
        rewrite ^/(.*)$ /index.php/$1 last;
    }

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

`fastcgi_*` 可以理解成nginx接受client请求时的响应使用的。proxy是nginx作为client转发时使用的，如果header过大，超出了默认的1k，就会引发上述的 upstream sent too big header。



##盗链

常见的防盗链设置如下:

    location ~* \.(gif|jpg|png|swf|flv)$ {
        valid_referers none blocked www.demo1.com www.demo2.net;
        if ($invalid_referer) {
            rewrite ^/ http://www.demo1.com/403.html;
            #return 404;
        }
        expires 3d;
    }

第一行表示对gif、jpg、png、swf、flv后缀的文件实行防盗链

第二行表示对两个域名这两个来路进行判断

if{}里面内容意思是：如果来入不是指定判断的来路时跳转到错误页面。

## Access-Control-Allow-Origin
```
add_header Access-Control-Allow-Origin *;
```

##  上传

    client_max_body_size 35m;        #客户端上传文件大小设为35M


## 连接数

worker_rlimit_nofile
注意：设置了这个后，你修改worker_connections值时，是不能超过worker_rlimit_nofile的这个值，不然又会有前面的那个warn提示。
保存配置文件，退出重启nginx。


如果nginx 中worker_connections 值设置是1024，worker_processes 值设置是4，按反向代理模式下最大连接数的理论计算公式：
   最大连接数 = worker_processes * worker_connections/4
查看相关资料，生产环境中worker_connections 建议值最好超过9000，计划将一台nginx 设置为10240，再观察一段时间。


## 隐藏版本号

 server_tokens off;

php 则 expose_php = Off 
