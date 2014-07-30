# Nginx Rewrite 规则相关指令


<http://qing.tiyee.net/post/2013-04-03/40048686737>

相关指令有if,rewrite,set,return,break等，其中最关键的就是rewrite.一个简单的Nginx Rewrite规则语法如下：

	rewrite ^/b/(.*)\.html /play.php?video=$1 break;

## 1.break指令

默认值：none ;使用环境：server,location,if ;

该指令的作用是完成当前的规则集，不再处理rewrite指令。

## 2.if指令

默认值：none ;使用环境：server,location

该指令用于检查一个条件是否符合，如果条件符合，则执行大括号内的语句。If指令不支持嵌套，不支持多个条件&&和||处理。

* A.变量名，错误的值包括：空字符串""或者任何以0开始的字符串
* B.变量比较可以使用"="(表示等于)和"!="(表示不等于)
* C.正则表达式模式匹配可以使用"~*"和"~"符号
* D."~"符号表示区分大小写字母的匹配
* E."~*"符号表示不区分大小写字母的匹配
* F."!~"和"!~*"符号的作用刚好和"~"、"~*"相反，表示不匹配
* G."-f"和"!-f"用来判断文件是否存在
* H."-d"和"!-d"用来判断目录是否存在
* I."-e"和"!-e"用来判断文件或目录是否存在
* J."-x"和"!-x"用来判断文件是否为可执行
* K.部分正则表达式可以在()内，用$1~$9来访问

## 3.return指令

语法：return code ;使用环境：server,location,if ;

该指令用于结束规则的执行并返回状态码给客户端。

示例：如果访问的URL以".sh"或".bash"结尾，则返回403状态码

	location ~ .*\.(sh|bash)?$
	{
		return 403;
	}

## 4.rewrite 指令

语法：rewrite regex replacement flag

默认值：none ; 使用环境：server,location,if

该指令根据表达式来重定向URI，或者修改字符串。指令根据配置文件中的顺序来执行。注意重写表达式只对相对路径有效。如果你想配对主机名，你应该使用if语句，示例如下：

	if( $host ~* www\.(.*) )
	{
		set $host_without_www $1;
		rewrite  ^(.*)$  http://$host_without_www$1 permanent;
	}
 

rewrite指令的最后一项参数为flag标记，支持flag标记有：

* last     相当于apache里面的[L]标记，表示rewrite。
* break本条规则匹配完成后，终止匹配，不再匹配后面的规则。
* redirect  返回302临时重定向，浏览器地址会显示跳转后的URL地址。
* permanent  返回301永久重定向， 浏览器地址会显示跳转后的URL地址。

使用last和break实现URI重写，浏览器地址栏不变。而且两者有细微差别，使用alias指令必须用last标记;使用proxy_pass指令时，需要使用break标记。Last标记在本条rewrite规则执行完毕后，会对其所在server{......}标签重新发起请求，而break标记则在本条规则匹配完成后，终止匹配。

一般在跟location中(location /{...})或直接在server标签中编写rewrite规则，推荐使用last标记；在非根location中(location /cms/{...})，则使用break。

如果URI中含有参数(/app/test.php?id=5)，默认情况下参数会被自动附加到替换串上，你可以通过在替换串的末尾加上?标记来解决这一问题。

例如：

	rewrite ^/test(.*)$ http://www.tiyee.net/home  permanent;

访问http://www.tiyee.net/test?id=5 会跳转到 http://www.tiyee.net/home?id=5

 

例如：如果我们将类似URL /photo/123456 重定向到 /path/to/photo/12/1234/123456.png

	Rewrite "/photo/([0-9]{2})([0-9]{2})([0-9]{2})" /path/to/photo/$1/$1$2/$1$2$3.png ;

注：如果正则表达式里面有花括号"{"或"}" ，应该使用双引号或单引号。

## 5.Set指令

语法：set variable value ; 默认值:none ; 使用环境：server,location,if;

该指令用于定义一个变量，并给变量赋值。变量的值可以为文本、变量以及文本变量的联合。

示例：set $varname "hello world";

## 6.Uninitialized_variable_warn指令

语法：uninitialized_variable_warn on|off

使用环境：http,server,location,if

该指令用于开启和关闭未初始化变量的警告信息，默认值为开启。

 

## 7.Nginx Rewrite可以用到的全局变量

$args ,$content_length ,$content_type ,$document_root ,$document_uri ,$host ,$http_user_agent ,$http_cookie ,$limit_rate ,$request_body_file ,$request_method ,$remote_addr ,$remote_port ,$remote_user ,$request_filename ,$request_uri ,$query_string ,$scheme ,$server_protocol ,$server_addr ,$server_name ,$server_port ,$uri 

====(详情见附录)====

 

## Nginx的Rewrite规则编写实例

1.当访问的文件和目录不存在时，重定向到某个php文件

	if( !-e $request_filename )
	{
		rewrite ^/(.*)$ index.php last;
	}

2.目录对换 /123456/xxxx  ====>   /xxxx?id=123456

	rewrite ^/(\d+)/(.+)/  /$2?id=$1 last;

3.如果客户端使用的是IE浏览器，则重定向到/ie目录下

	if( $http_user_agent  ~ MSIE)
	{
		rewrite ^(.*)$ /ie/$1 break;
	}
 
4.禁止访问多个目录

	location ~ ^/(cron|templates)/
	{
		deny all;
		break;
	}
 
5.禁止访问以/data开头的文件

	location ~ ^/data
	{
		deny all;
	}
 
6.禁止访问以.sh,.flv,.mp3为文件后缀名的文件

	location ~ .*\.(sh|flv|mp3)$
	{
		return 403;
	}

7.设置某些类型文件的浏览器缓存时间

	location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
	{
		expires 30d;
	}
	location ~ .*\.(js|css)$
	{
		expires 1h;
	}

8.给favicon.ico和robots.txt设置过期时间;
这里为favicon.ico为99天,robots.txt为7天并不记录404错误日志

	location ~(favicon.ico) {
		log_not_found off;
		expires 99d;
		break;
	}

                                                             
	location ~(robots.txt) {
		log_not_found off;
		expires 7d;
		break;
	}

9.设定某个文件的过期时间;这里为600秒，并不记录访问日志

	location ^~ /html/scripts/loadhead_1.js {
		access_log   off;
		root /opt/lampp/htdocs/web;
		expires 600;
		break;
	}

10.文件反盗链并设置过期时间

这里的return 412 为自定义的http状态码，默认为403，方便找出正确的盗链的请求 “rewrite ^/ http://img.tiyee.net/leech.gif;”显示一张防盗链图片 “access_log off;”不记录访问日志，减轻压力 “expires 3d”所有文件3天的浏览器缓存

	location ~* ^.+\.(jpg|jpeg|gif|png|swf|rar|zip|css|js)$ {
		valid_referers none blocked *.c1gstudio.com *.c1gstudio.net localhost 208.97.167.194;
		if ($invalid_referer) {
			rewrite ^/ http://img.tiyee.net/leech.gif;
			return 412;
			break;
		}
		access_log   off;
		root /opt/lampp/htdocs/web;
		expires 3d;
		break;
	}
 
11.只充许固定ip访问网站，并加上密码
 
	root  /opt/htdocs/www;
	allow   208.97.167.194;
	allow   222.33.1.2;
	allow   231.152.49.4;
	deny    all;
	auth_basic “C1G_ADMIN”;
	auth_basic_user_file htpasswd;

12将多级目录下的文件转成一个文件，增强seo效果 /job-123-456-789.html 指向/job/123/456/789.html

	rewrite ^/job-([0-9]+)-([0-9]+)-([0-9]+)\.html$ /job/$1/$2/jobshow_$3.html last;

13.将根目录下某个文件夹指向2级目录
如/shanghaijob/ 指向 /area/shanghai/
如果你将last改成permanent，那么浏览器地址栏显是/location/shanghai/

	rewrite ^/([0-9a-z]+)job/(.*)$ /area/$1/$2 last;

上面例子有个问题是访问/shanghai 时将不会匹配

	rewrite ^/([0-9a-z]+)job$ /area/$1/ last;
	rewrite ^/([0-9a-z]+)job/(.*)$ /area/$1/$2 last;

这样/shanghai 也可以访问了，但页面中的相对链接无法使用，
如./list_1.html真实地址是/area/shanghia/list_1.html会变成/list_1.html,导至无法访问。

那我加上自动跳转也是不行咯
(-d $request_filename)它有个条件是必需为真实目录，而我的rewrite不是的，所以没有效果

	if (-d $request_filename){
		rewrite ^/(.*)([^/])$ http://$host/$1$2/ permanent;
	}
 
知道原因后就好办了，让我手动跳转吧

	rewrite ^/([0-9a-z]+)job$ /$1job/ permanent;
	rewrite ^/([0-9a-z]+)job/(.*)$ /area/$1/$2 last;

14.文件和目录不存在的时候重定向：

	if (!-e $request_filename) {
		proxy_pass http://127.0.0.1;
	}
 
## Nginx和Apache的Rewrite规则实例对比

1.一般简单的Nginx和Apache规则的区别不大，基本能够完全兼容，例如：

	Apache: RewriteRule  ^/abc/$   /web/abc.php [L]
	Nginx:  rewrite  ^/abc/$  /web/abc.php last ;

我们可以看出来只要把Apache的RewriteRule改为Nginx的rewrite，Apache的[L]改为last 即可。

如果将Apache的规则改为Nginx规则后，用命令Nginx -t 检查发现错误，则我们可以尝试给条件加上引号，例如：

	rewrite “^/([0-9]{5}).html$”   /x.php?id=$1 last;

2.Apache和Nginx的Rewrite规则在URL跳转时有细微区别：

	Apache:  RewriteRule ^/html/([a-zA-Z]+)/.*$  /$1/  [R=301,L]
	Nginx:   rewrite ^/html/([a-zA-Z]+)/.*$  http://$host/$1/ premanent ;

我们可以看到在Nginx的跳转中，我们需要加上http://$host，这是在Nginx中强烈要求的。
 

3.下面是一些Apache和Nginx规则的对应关系

Apache    		| Nginx
----------------|----------
RewriteCond		| if
RewriteRule		| rewrite
[R] 			| redirect
[P] 			| last
[R,L] 			| redirect
[P,L]			| last
[PT,L]			| last

例如：允许指定的域名访问本站，其他的域名一律转向www.tiyee.net

Apache:

	RewriteCond %{HTTP_HOST} !^(.*?)\.aaa\.com$ [NC]
	RewriteCond %{HTTP_HOST} !^localhost$ 
	RewriteCond %{HTTP_HOST} !^192\.168\.0\.(.*?)$
	RewriteRule ^/(.*)$ http://www.tiyee.net [R,L]

Nginx:

	if( $host ~* ^(.*)\.aaa\.com$ )
	{
		set $allowHost ‘1’;
	}
	if( $host ~* ^localhost )
	{
		set $allowHost ‘1’;
	}
	if( $host ~* ^192\.168\.1\.(.*?)$ )
	{
		set $allowHost ‘1’;
	}
	if( $allowHost !~ ‘1’ )
	{
		rewrite ^/(.*)$ http://www.tiyee.net redirect ;
	}

##《附录：nginx全局变量》

经常需要配置Nginx ，其中有许多以  $ 开头的变量，经常需要查阅nginx 所支持的变量。Nginx支持的http变量实现在 ngx_http_variables.c 的 ngx_http_core_variables存储实现

	ngx_http_core_variables
	static ngx_http_variable_t  ngx_http_core_variables[] = {
                                
	     { ngx_string("http_host"), NULL, ngx_http_variable_header,
	       offsetof(ngx_http_request_t, headers_in.host), 0, 0 },
	                                
	     { ngx_string("http_user_agent"), NULL, ngx_http_variable_header,
	       offsetof(ngx_http_request_t, headers_in.user_agent), 0, 0 },
	                                
	     { ngx_string("http_referer"), NULL, ngx_http_variable_header,
	       offsetof(ngx_http_request_t, headers_in.referer), 0, 0 },
	                                
		 #if (NGX_HTTP_GZIP)
	     { ngx_string("http_via"), NULL, ngx_http_variable_header,
	       offsetof(ngx_http_request_t, headers_in.via), 0, 0 },
		 #endif
	                                
		 #if (NGX_HTTP_PROXY || NGX_HTTP_REALIP)
	     { ngx_string("http_x_forwarded_for"), NULL, ngx_http_variable_header,
	       offsetof(ngx_http_request_t, headers_in.x_forwarded_for), 0, 0 },
		 #endif
	                                
	     { ngx_string("http_cookie"), NULL, ngx_http_variable_headers,
	       offsetof(ngx_http_request_t, headers_in.cookies), 0, 0 },
	                                
	     { ngx_string("content_length"), NULL, ngx_http_variable_header,
	       offsetof(ngx_http_request_t, headers_in.content_length), 0, 0 },
	                                
	     { ngx_string("content_type"), NULL, ngx_http_variable_header,
	       offsetof(ngx_http_request_t, headers_in.content_type), 0, 0 },
	                                
	     { ngx_string("host"), NULL, ngx_http_variable_host, 0, 0, 0 },
	                                
	     { ngx_string("binary_remote_addr"), NULL,
	       ngx_http_variable_binary_remote_addr, 0, 0, 0 },
	                                
	     { ngx_string("remote_addr"), NULL, ngx_http_variable_remote_addr, 0, 0, 0 },
	                                
	     { ngx_string("remote_port"), NULL, ngx_http_variable_remote_port, 0, 0, 0 },
	                                
	     { ngx_string("server_addr"), NULL, ngx_http_variable_server_addr, 0, 0, 0 },
	                                
	     { ngx_string("server_port"), NULL, ngx_http_variable_server_port, 0, 0, 0 },
	                                
	     { ngx_string("server_protocol"), NULL, ngx_http_variable_request,
	       offsetof(ngx_http_request_t, http_protocol), 0, 0 },
	                                
	     { ngx_string("scheme"), NULL, ngx_http_variable_scheme, 0, 0, 0 },
	                                
	     { ngx_string("request_uri"), NULL, ngx_http_variable_request,
	       offsetof(ngx_http_request_t, unparsed_uri), 0, 0 },
	                                
	     { ngx_string("uri"), NULL, ngx_http_variable_request,
	       offsetof(ngx_http_request_t, uri),
	       NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("document_uri"), NULL, ngx_http_variable_request,
	       offsetof(ngx_http_request_t, uri),
	       NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("request"), NULL, ngx_http_variable_request_line, 0, 0, 0 },
	                                
	     { ngx_string("document_root"), NULL,
	       ngx_http_variable_document_root, 0, NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("realpath_root"), NULL,
	       ngx_http_variable_realpath_root, 0, NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("query_string"), NULL, ngx_http_variable_request,
	       offsetof(ngx_http_request_t, args),
	       NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("args"),
	       ngx_http_variable_request_set,
	       ngx_http_variable_request,
	       offsetof(ngx_http_request_t, args),
	       NGX_HTTP_VAR_CHANGEABLE|NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("is_args"), NULL, ngx_http_variable_is_args,
	       0, NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("request_filename"), NULL,
	       ngx_http_variable_request_filename, 0,
	       NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("server_name"), NULL, ngx_http_variable_server_name, 0, 0, 0 },
	                                
	     { ngx_string("request_method"), NULL,
	       ngx_http_variable_request_method, 0,
	       NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("remote_user"), NULL, ngx_http_variable_remote_user, 0, 0, 0 },
	                                
	     { ngx_string("body_bytes_sent"), NULL, ngx_http_variable_body_bytes_sent,
	       0, 0, 0 },
	                                
	     { ngx_string("request_completion"), NULL,
	       ngx_http_variable_request_completion,
	       0, 0, 0 },
	                                
	     { ngx_string("request_body"), NULL,
	       ngx_http_variable_request_body,
	       0, 0, 0 },
	                                
	     { ngx_string("request_body_file"), NULL,
	       ngx_http_variable_request_body_file,
	       0, 0, 0 },
	                                
	     { ngx_string("sent_http_content_type"), NULL,
	       ngx_http_variable_sent_content_type, 0, 0, 0 },
	                                
	     { ngx_string("sent_http_content_length"), NULL,
	       ngx_http_variable_sent_content_length, 0, 0, 0 },
	                                
	     { ngx_string("sent_http_location"), NULL,
	       ngx_http_variable_sent_location, 0, 0, 0 },
	                                
	     { ngx_string("sent_http_last_modified"), NULL,
	       ngx_http_variable_sent_last_modified, 0, 0, 0 },
	                                
	     { ngx_string("sent_http_connection"), NULL,
	       ngx_http_variable_sent_connection, 0, 0, 0 },
	                                
	     { ngx_string("sent_http_keep_alive"), NULL,
	       ngx_http_variable_sent_keep_alive, 0, 0, 0 },
	                                
	     { ngx_string("sent_http_transfer_encoding"), NULL,
	       ngx_http_variable_sent_transfer_encoding, 0, 0, 0 },
	                                
	     { ngx_string("sent_http_cache_control"), NULL, ngx_http_variable_headers,
	       offsetof(ngx_http_request_t, headers_out.cache_control), 0, 0 },
	                                
	     { ngx_string("limit_rate"), ngx_http_variable_request_set_size,
	       ngx_http_variable_request_get_size,
	       offsetof(ngx_http_request_t, limit_rate),
	       NGX_HTTP_VAR_CHANGEABLE|NGX_HTTP_VAR_NOCACHEABLE, 0 },
	                                
	     { ngx_string("nginx_version"), NULL, ngx_http_variable_nginx_version,
	       0, 0, 0 },
	                                
	     { ngx_string("hostname"), NULL, ngx_http_variable_hostname,
	       0, 0, 0 },
	                                
	     { ngx_string("pid"), NULL, ngx_http_variable_pid,
	       0, 0, 0 },
	                                
	     { ngx_null_string, NULL, NULL, 0, 0, 0 }
	};

把这些变量提取下，总结如下：

* arg_PARAMETER      #这个变量包含GET请求中，如果有变量PARAMETER时的值。
* args               #这个变量等于请求行中(GET请求)的参数，例如foo=123&bar=blahblah;
* binary_remote_addr #二进制的客户地址。
* body_bytes_sent    #响应时送出的body字节数数量。即使连接中断，这个数据也是精确的。
* content_length     #请求头中的Content-length字段。
* content_type       #请求头中的Content-Type字段。
* cookie_COOKIE      #cookie COOKIE变量的值
* document_root      #当前请求在root指令中指定的值。
* document_uri       #与
* uri相同。
* host               #请求主机头字段，否则为服务器名称。
* hostname           #Set to the machine’s hostname as returned by gethostname
* http_HEADER
* is_args            #如果有
* args参数，这个变量等于”?”，否则等于”"，空值。
* http_user_agent    #客户端agent信息
* http_cookie        #客户端cookie信息
* limit_rate         #这个变量可以限制连接速率。
* query_string       #与
* args相同。
* request_body_file  #客户端请求主体信息的临时文件名。
* request_method     #客户端请求的动作，通常为GET或POST。
* remote_addr        #客户端的IP地址。
* remote_port        #客户端的端口。
* remote_user        #已经经过Auth Basic Module验证的用户名。
* request_completion #如果请求结束，设置为OK. 当请求未结束或如果该请求不是请求链串的最后一个时，为空(Empty)。
* request_method     #GET或POST
* request_filename   #当前请求的文件路径，由root或alias指令与URI请求生成。
* request_uri        #包含请求参数的原始URI，不包含主机名，如：”/foo/bar.php?arg=baz”。不能修改。
* scheme             #HTTP方法（如http，https）。
* server_protocol    #请求使用的协议，通常是HTTP/1.0或HTTP/1.1。
* server_addr        #服务器地址，在完成一次系统调用后可以确定这个值。
* server_name        #服务器名称。
* server_port        #请求到达服务器的端口号。