#htaccess 用法

## 1.时区设置

有些时候，当你在PHP里使用date或mktime函数时，由于时区的不同，它会显示出一些很奇怪的信息。下面是解决这个问题的方法之一。就是设置你的服务器的时区。你可以在这里找到所有支持的时区的清单。
SetEnv TZ Australia/Melbourne  
 
## 2. 搜索引擎友好的301永久转向方法

为什么这是搜索引擎友好的呢？因为现在很多现代的搜索引擎都有能根据检查301永久转向来更新它现有的记录的功能。

Redirect 301 http://www.aqee.net/home http://www.aqee.net/  

## 3. 屏蔽下载对话框

通常，当你下载东西的时候，你会看到一个对话框询问你是保持这个文件还是直接打开它。如果你不想看到这个东西，你可以把下面的一段代码放到你的.htaccess文件里。

	AddType application/octet-stream .pdf  
	AddType application/octet-stream .zip  
	AddType application/octet-stream .mov  

## 4. 省去www前缀

SEO的一个原则是，确保你的网站只有一个URL。因此，你需要把所有的通过www的访问转向的非www，或者反这来。

	RewriteEngine On  
	RewriteBase /  
	RewriteCond %{HTTP_HOST} ^www.aqee.net [NC]  
	RewriteRule ^(.*)$ http://aqee.net/$1 [L,R=301]  

## 5. 个性化Error页面

对每个错误代码定制自己个性化的错误页面。

	ErrorDocument 401 /error/401.php  
	ErrorDocument 403 /error/403.php  
	ErrorDocument 404 /error/404.php  
	ErrorDocument 500 /error/500.php  

## 6. 压缩文件

通过压缩你的文件体积来优化网站的访问速度。

	# 压缩 text, html, javascript, css, xml:  
	AddOutputFilterByType DEFLATE text/plain  
	AddOutputFilterByType DEFLATE text/html  
	AddOutputFilterByType DEFLATE text/xml  
	AddOutputFilterByType DEFLATE text/css  
	AddOutputFilterByType DEFLATE application/xml  
	AddOutputFilterByType DEFLATE application/xhtml+xml  
	AddOutputFilterByType DEFLATE application/rss+xml  
	AddOutputFilterByType DEFLATE application/javascript  
	AddOutputFilterByType DEFLATE application/x-javascript  

## 7. 缓存文件

缓存文件是另外一个提高你的网站访问速度的好方法。

	<FilesMatch “.(flv|gif|jpg|jpeg|png|ico|swf|js|css|pdf)$”>  
	Header set Cache-Control “max-age=2592000″  
	</FilesMatch>  

## 8. 对某些文件类型禁止使用缓存

而另一方面，你也可以定制对某些文件类型禁止使用缓存。

	# 显式的规定对脚本和其它动态文件禁止使用缓存  
	<FilesMatch “.(pl|php|cgi|spl|scgi|fcgi)$”>  
	Header unset Cache-Control  
	</FilesMatch>  

# 安全问题

下面的htaccess代码能够提高你的web服务器的安全水平。图片链接盗用保护非常有用，它能防止其他人偷盗使用你的服务器上的图片资源。

## 1. 通过.htaccess放盗链

痛恨那些偷盗链接你的web服务器上的图片资源而耗尽了你的带宽的行为吗？试试这个，你可以防止这种事情的发生。

	RewriteBase /  
	RewriteCond %{HTTP_REFERER} !^$  
	RewriteCond %{HTTP_REFERER} !^http://(www.)?aqee.net/.*$ [NC]  
	RewriteRule .(gif|jpg|swf|flv|png)$ /feed/ [R=302,L]  

## 2. 防黑客

如果你想提高网站的安全等级，你可以去掉下面的几行代码，这样可以防止一些常见恶意URL匹配的黑客攻击技术。

	RewriteEngine On  
	# proc/self/environ? 没门！  
	RewriteCond %{QUERY_STRING} proc/self/environ [OR]  
  
	# 阻止脚本企图通过URL修改mosConfig值  
	RewriteCond %{QUERY_STRING} mosConfig_[a-zA-Z_]{1,21}(=|%3D) [OR]  
	  
	# 阻止脚本通过URL传递的base64_encode垃圾信息  
	RewriteCond %{QUERY_STRING} base64_encode.*(.*) [OR]  
	  
	# 阻止在URL含有<script>标记的脚本  
	RewriteCond %{QUERY_STRING} (<|%3C).*script.*(>|%3E) [NC,OR]  
	  
	# 阻止企图通过URL设置PHP的GLOBALS变量的脚本  
	RewriteCond %{QUERY_STRING} GLOBALS(=|[|%[0-9A-Z]{0,2}) [OR]  
	  
	# 阻止企图通过URL设置PHP的_REQUEST变量的脚本  
	RewriteCond %{QUERY_STRING} _REQUEST(=|[|%[0-9A-Z]{0,2})  
	  
	# 把所有被阻止的请求转向到403禁止提示页面！  
	RewriteRule ^(.*)$ index.php [F,L]  

## 3. 阻止访问你的 .htaccess 文件

下面的代码可以阻止别人访问你的.htaccess文件。同样，你也可以设定阻止多种文件类型。

	# 保护你的 htaccess 文件  
	<Files .htaccess>  
	order allow,deny  
	deny from all  
	</Files>  
  
	# 阻止查看指定的文件  
	<Files secretfile.jpg>  
	 order allow,deny  
	 deny from all  
	</Files>  
  
	# 多种文件类型  
	<FilesMatch “.(htaccess|htpasswd|ini|phps|fla|psd|log|sh)$”>  
	 Order Allow,Deny  
	 Deny from all  
	</FilesMatch>  

## 4. 重命名 htaccess 文件

你可以通过重命名htaccess文件来对其进行保护。

	AccessFileName htacc.ess  

## 5. 禁止目录浏览

禁止服务器对外显示目录结构，反之亦然。

	# 禁止目录浏览  
	Options All -Indexes  
  
	# 开放目录浏览
	Options All +Indexes  

## 6. 改变缺省的Index页面

你可以把缺省的 index.html, index.php 或 index.htm 改成其它页面。

	DirectoryIndex business.html  
## 7. 通过引用信息来阻止某些不欢迎的浏览者

	# 阻止来自某网站的用户  
	<IfModule mod_rewrite.c>  
	 RewriteEngine on  
	 RewriteCond %{HTTP_REFERER} scumbag.com [NC,OR]  
	 RewriteCond %{HTTP_REFERER} wormhole.com [NC,OR]  
	 RewriteRule .* - [F]  
	</ifModule>  

## 8. 通过判断浏览器头信息来阻止某些请求

这个方法可以通过阻止某些机器人或蜘蛛爬虫抓取你的网站来节省你的带宽流量。

	# 阻止来自某些特定网站的用户  
	<IfModule mod_rewrite.c>  
	SetEnvIfNoCase ^User-Agent$ .*(craftbot|download|extract|stripper|sucker|ninja|clshttp|webspider
	|leacher|collector|grabber|webpictures) HTTP_SAFE_BADBOT  
	SetEnvIfNoCase ^User-Agent$ .*(libwww-perl|aesop_com_spiderman) HTTP_SAFE_BADBOT  
	Deny from env=HTTP_SAFE_BADBOT  
	</ifModule>  

## 9. 禁止脚本执行，加强你的目录安全
	# 禁止某些目录里的脚本执行权限  
	AddHandler cgi-script .php .pl .py .jsp .asp .htm .shtml .sh .cgi  
	Options -ExecCGI 
