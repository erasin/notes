# nginx

Nginx 有2个模块用于控制访问“数量”和“速度”，简单的说，控制你最多同时有 多少个访问，并且控制你每秒钟最多访问多少次， 你的同时并发访问不能太多，也不能太快，不然就“杀无赦”。

    HttpLimitZoneModule    限制同时并发访问的数量
    HttpLimitReqModule      限制访问数据，每秒内最多几个请求

## 一、普通配置

普通配置就是针对【用户浏览器】→【网站服务器】这种常规模式的nginx配置。那么，如果我要对单IP做访问限制，绝大多数教程都是这样写的：

```
## 用户的 IP 地址 $binary_remote_addr 作为 Key，每个 IP 地址最多有 50 个并发连接
## 你想开 几千个连接 刷死我？ 超过 50 个连接，直接返回 503 错误给你，根本不处理你的请求了

limit_conn_zone $binary_remote_addr zone=TotalConnLimitZone:10m ;
limit_conn  TotalConnLimitZone  50;
limit_conn_log_level notice;

## 用户的 IP 地址 $binary_remote_addr 作为 Key，每个 IP 地址每秒处理 10 个请求
## 你想用程序每秒几百次的刷我，没戏，再快了就不处理了，直接返回 503 错误给你

limit_req_zone $binary_remote_addr zone=ConnLimitZone:10m  rate=10r/s;
limit_req_log_level notice;

## 具体服务器配置
server {
	listen   80;
	location ~ \.php$ {
                ## 最多 5 个排队， 由于每秒处理 10 个请求 + 5个排队，你一秒最多发送 15 个请求过来，再多就直接返回 503 错误给你了
		limit_req zone=ConnLimitZone burst=5 nodelay;
		fastcgi_pass   127.0.0.1:9000;
		fastcgi_index  index.php;
		include	fastcgi_params;
	}
}
```

这样一个最简单的服务器安全限制访问就完成了，这个基本上你 Google 一搜索能搜索到  90% 的网站都是这个例子，
而且还强调用“$binary_remote_addr”可以节省内存之类的云云。

## 二、CDN之后

目前国内已经争相出现了百度云加速、加速乐、360网站卫士以及安全宝等免费CDN。让我们这些小网站也能免费享受以前高大上的CDN加速服务。

于是，网站的访问模式就变为：
用户浏览器 → CDN节点 → 网站源服务器

甚至是更复杂的模式：
用户浏览器 → CDN节点（CDN入口、CC\DDoS攻击流量清洗等） → 阿里云盾 → 源服务器

可以看到，我们的网站中间经历了好几层的透明加速和安全过滤， 这种情况下，我们就不能用上面的“普通配置”。因为普通配置中基于【源IP的限制】的结果就是，我们把【CDN节点】或者【阿里云盾】给限制了，因为这里“源IP”地址不再是真实用户的IP，而是中间CDN节点的IP地址。

我们需要限制的是最前面的真实用户，而不是中间为我们做加速的加速服务器。

其实，当一个 CDN 或者透明代理服务器把用户的请求转到后面服务器的时候，这个 CDN 服务器会在 Http 的头中加入一个记录

X-Forwarded-For :  用户IP, 代理服务器IP

如果中间经历了不止一个代理服务器，这个记录会是这样

X-Forwarded-For :  用户IP, 代理服务器1-IP, 代理服务器2-IP, 代理服务器3-IP, ….

可以看到经过好多层代理之后， 用户的真实IP 在第一个位置， 后面会跟一串中间代理服务器的IP地址，从这里取到用户真实的IP地址，针对这个 IP 地址做限制就可以了。

那么针对CDN模式下的访问限制配置就应该这样写：

```
cd /usr/local/src
#下载echo模块并解压：
wget https://github.com/openresty/echo-nginx-module/archive/v0.57.tar.gz
tar zxvf v0.57.tar.gz

#下载nginx并解压
wget http://nginx.org/download/nginx-1.6.0.tar.gz
tar -xzvf nginx-1.6.0.tar.gz
cd nginx-1.6.0/

#查看在用nginx的编译参数（如果是全新安装则省略）
/usr/local/nginx/sbin/nginx -V
nginx version: nginx/1.6.0
built by gcc 4.4.7 20120313 (Red Hat 4.4.7-4) (GCC) #以下这行即为旧的编译参数：
configure arguments: --user=www --group=www --prefix=/usr/local/nginx --with-http_gzip_static_module
#在旧的编译参数基础上新增【--add-module=/echo模块的解压路径】参数，开始编译
./configure --prefix=/usr/local/nginx/nginx  --add-module=/usr/local/src/echo-nginx-module-0.57

#make编译
make -j2

#平滑升级nginx (如果是全新安装请执行：make install)
mv /usr/local/nginx/sbin/nginx /usr/local/nginx/sbin/nginx.old
cp -f objs/nginx /usr/local/nginx/sbin/
make upgrade
```
