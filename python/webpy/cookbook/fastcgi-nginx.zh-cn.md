# Webpy + Nginx with FastCGI搭建Web.py
**title:**在nginx中使用FastCGI搭建部署webpy  
**tags:**nginx,fastcgi,webpy  
**info:**  webpy 在 nginx 上搭建部署的实例 

这一节讲解的是如何使用Nginx和FastCGI搭建Web.py应用

## 环境依赖的软件包

* Nginx 0.8.\* or 0.7.\* (需要包含fastcgi和rewrite模块)。
* Webpy 0.32
* Spawn-fcgi 1.6.2
* Flup

注意：Flup是最常见的忘记装的软件，需要安装

更老的版本应该也可以工作，但是没有测试过，最新的是可以工作的

## 一些资源

* [Nginx wiki](http://wiki.nginx.org/NginxInstall)
* [Spawn-fcgi](http://redmine.lighttpd.net/projects/spawn-fcgi/news)
* [Flup](http://trac.saddi.com/flup)

## Notes

* 你可以重命名`index.py`为任何你想要的文件名。
* `/path/to/www` 为代码路径。
* `/path/to/www/index.py`为python代码的完整路径。

## Nginx 配置文件

    location / {
        fastcgi_param REQUEST_METHOD $request_method;
        fastcgi_param QUERY_STRING $query_string;
        fastcgi_param CONTENT_TYPE $content_type;
        fastcgi_param CONTENT_LENGTH $content_length;
        fastcgi_param GATEWAY_INTERFACE CGI/1.1;
        fastcgi_param SERVER_SOFTWARE nginx/$nginx_version;
        fastcgi_param REMOTE_ADDR $remote_addr;
        fastcgi_param REMOTE_PORT $remote_port;
        fastcgi_param SERVER_ADDR $server_addr;
        fastcgi_param SERVER_PORT $server_port;
        fastcgi_param SERVER_NAME $server_name;
        fastcgi_param SERVER_PROTOCOL $server_protocol;
        fastcgi_param SCRIPT_FILENAME $fastcgi_script_name;
        fastcgi_param PATH_INFO $fastcgi_script_name;
        fastcgi_pass 127.0.0.1:9002;
    }

对于静态文件可以添加如下配置:

	location /static/ {
	    if (-f $request_filename) {
		rewrite ^/static/(.*)$  /static/$1 break;
	    }
	}

__注意:__ 地址和端口号可能会是不同的。

## Spawn-fcgi

可以通过一下命令启动一个Spawn-fcgi进程:

	spawn-fcgi -d /path/to/www -f /path/to/www/index.py -a 127.0.0.1 -p 9002

## 启动和关闭的命令

启动:

	#!/bin/sh
	spawn-fcgi -d /path/to/www -f /path/to/www/index.py -a 127.0.0.1 -p 9002

关闭:

	#!/bin/sh
	kill `pgrep -f "python /path/to/www/index.py"`

__Note:__ 你可以随意填写地址和端口信息，但是一定需要和Nginx配置文件相匹配。

## Hello world!

讲下面的代码保存为index.py（或者任何你喜欢的），注意，使用Nginx配置的话，`web.wsgi.runwsgi = lambda func, addr=None: web.wsgi.runfcgi(func, addr)`这一行代码是必须的。

	#!/usr/bin/env python
	# -*- coding: utf-8 -*-

	import web

	urls = ("/.*", "hello")
	app = web.application(urls, globals())

	class hello:
		def GET(self):
			return 'Hello, world!'

	if __name__ == "__main__":
		web.wsgi.runwsgi = lambda func, addr=None: web.wsgi.runfcgi(func, addr)
		app.run()

注意: 同样需要给代码设置权限，代码如下chmod +x index.py。

## 运行

1. 打开一个 `spawn-fcgi` 进程.
2. 打开 Nginx.

如果需要检查应用程序是否运行，使用`ps aux|grep index.py`可以很容易的查看。

重启nginx配置:

	/path/to/nginx/sbin/nginx -s reload

停止nginx:

	/path/to/nginx/sbin/nginx -s stop

注意：运行后可访问http://localhost访问网站，更多信息可以去参考nginx官方文档。


## 示例配置

###1、 index.py 入口添加

    web.wsgi.runwsgi = lambda func, addr=None: web.wsgi.runfcgi(func, addr)

###2、 创建 start.sh 开始脚本

    #!/bin/bash

    if [ $# -ne 1 ];then
        read -p ' 输入参数 start OR stop! :' action
    else
        action=$1
    fi

    if [ "$action" == "start" ];then
        spawn-fcgi -d /home/http/blog -f /home/http/blog/blog.py -a 127.0.0.1 -p 9002
        echo "it started!"
    elif [ "$action" == "stop" ];then
        kill `pgrep -f "python /home/http/webpy/index.py"` 
        echo `ps -p $sid -f`
        read -p ' 确认杀死改进程!(y or n)' killit
        if [ "$killit" == y ];then
            kill $sid
            echo "kill the webpy!"
        else 
            echo 'why not?'
        fi
    else
        echo " 输入 start OR stop!"
    fi

使用 ./start.sh start 来开始  
如果系统有多个python版本，这时注意你的 py文件是否添加了 默认的python头 `#!/usr/bin/python2.7`

###3、建nginx config配置文件

    server {

        listen  80;
        server_name  blog.dom;
        root   /home/http/blog;

        access_log  /var/log/nginx/blog.log;

        location / {
            fastcgi_param REQUEST_METHOD $request_method;
            fastcgi_param QUERY_STRING $query_string;
            fastcgi_param CONTENT_TYPE $content_type;
            fastcgi_param CONTENT_LENGTH $content_length;
            fastcgi_param GATEWAY_INTERFACE CGI/1.1;
            fastcgi_param SERVER_SOFTWARE nginx/$nginx_version;
            fastcgi_param REMOTE_ADDR $remote_addr;
            fastcgi_param REMOTE_PORT $remote_port;
            fastcgi_param SERVER_ADDR $server_addr;
            fastcgi_param SERVER_PORT $server_port;
            fastcgi_param SERVER_NAME $server_name;
            fastcgi_param SERVER_PROTOCOL $server_protocol;
            fastcgi_param SCRIPT_FILENAME $fastcgi_script_name;
            fastcgi_param PATH_INFO $fastcgi_script_name;
            fastcgi_pass 127.0.0.1:9002;
        }

        location /static/ {
            root   /home/http/blog;
            if (-f $request_filename) {
               rewrite ^/static/(.*)$  /static/$1 break;
            }
        }

        error_page   500 502 503 504  /50x.html;

        location = /50x.html {
            root   /usr/share/nginx/html;
        }

        location ~ /\.ht {
            deny  all;
        }
    }

启动 nginx

###4、 hosts
上面用的是 blog.dom ,如果要正常访问必须要修改`/etc/hosts`了

    127.0.0.1 blog.dom



