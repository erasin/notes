#8-host#
**title:**8host  
**tags:**8host,捌号空间   
**info:**8host服务器的配置,虚拟空间的配置。

<http://bbs.8-host.net/read-htm-tid-278.html>

##绑定子域名##
修改根目录下 .htsub 文件；   
子域名     /www/面板登录用户名/域名加载点/htdocs/要绑定的子目录
>      demo.8-host.net     /www/8888.host/8-host_net/htdocs/demo

##流量时实##
[url]http://244519981.host.fb6f.8-host.com/my-real-speed/[/url]
http://bbs.8-host.net/attachment/Mon_0705/4_1_f4b7cb50b2d1e00.jpg

##目录##
>      backup 用于备份和存储非web用途文件  
>      htdocs 是存放web内容的目录，也就是http的根目录  
>      logs 里面是用户错误日志和访问日志  
>      sys 系统保留，用户没有权限的目录  
>      .htgroup .htpasswd 系统用，一般户不用动  
>      .htphp4ini 是php4的配置文件  
>      .htphp5ini 是php5的配置文件  
>      .htphpver 是用于选择php版本的  

## htaccess ##
在 htdocs 目录下建立 .htaccess

## django 部署 ##

注意：不要通过IE复制以下字符！  
首先在工程目录（通常是htdocs）建立.htaccess内容如下：  
Quote:

>      RewriteEngine On
>      RewriteCond %{REQUEST_FILENAME} !-f
>      RewriteRule ^(.*)$ mysite.fcgi/$1 [QSA,L]


然后再建立mysite.fcgi内容如下：  
Quote:

>     #!/usr/local/bin/python
>     import sys, os
>     
>     # 添加自定义Python路径
>     sys.path.insert(0, "/usr/local/bin/python")
>     
>     # 对于大多数eggs和模块，都可以上传到用户自己目录并加以调用
>     #sys.path.insert(0, "/www/user.root.dir/module.upload.dir")
>     
>     # 切换到工程目录（可选）
>     # os.chdir("/www/user.root.dir/site.mount.point/htdocs")
>     
>     # 设定DJANGO_SETTINGS_MODULE环境变量
>     os.environ['DJANGO_SETTINGS_MODULE'] # "myproject.settings"
>     
>     from django.core.servers.fastcgi import runfastcgi
>     runfastcgi(method#"threaded", daemonize#"false") 


##DNS 解析##
注意：如果设置了Email转发，就无法用MX记录，如果用MX记录，就无法用Email转发，cname和mx记录切记最后的点号，不带www的空域名hostname写@，泛解析hostname写*
基本格式如下

>      A记录： hostname A记录 IP地址 （例如： 74.117.56.75）
>     cname： hostname cname 要cname的域名地址和最后一个点. （例如： zou.lu.）
>     MX: hostname（一般写@） MX TTL数值+空格+点号 （例如： 10 mxdomain.qq.com.）
>     URL转发功能因为服务器被墙，暂时无法使用，见谅。
