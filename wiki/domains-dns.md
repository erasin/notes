#域名的DNS解析指南
**title:**域名DNS解析详解指南
**tags:**DNS解析,域名解析,A记录,MX,CNAME,URL Redirect(301),URL frame,ns,TXT Record   
**info:**域名解析和dns设置指南

##DNS 
DNS 是** Domain Name System（域名系统）**的缩写，此系统用于管理和识别域名。DNS 的最基本功能是为域的一个或多个 IP 地址提供名称。

例如，可以将域名 wolf.example.com 转换为 198.102.434.8。这样可便于记忆网址和电子邮件地址。

DNS 还用于查明应将特定地址的电子邮件发送到何处。这是通过 MX 记录完成的。

##IP 地址	
互联网协议地址是可让设备在网络上找到信息的唯一数字。    
由于一个域名可有一个或多个关联的 IP 地址，因此 Google 企业应用套件不支持仅使用 IP 地址的电子邮件和网络发布配置。

##域名	
域名是与一个或多个 IP 地址关联的易于记忆的名称（网址和电子邮件地址）。

由于网页是由其网址定义的，因此可将网页移动到不同的 IP 地址而不会影响访问者。
    
_示例_：<www.sosxt.com>

* singlespeed.com 是域名
* com 是顶级域名
* singlespeed 是 com 的一个子部分，代表第二级域名
* www 是子域（也称为第三级域名或 CNAME）
  整个域名的总长度不能超过 255 个字符，但有些注册商给出的限额要更短一些。

###域名注册商	
域名注册商指销售互联网域名（如 blueshirt.com 或 organicfood.org）的商家。除注册外，大多数此类公司还提供托管服务。

如果您的域名注册商与域名托管服务商相互独立，则您需要将托管服务商的名称服务器添加到注册商的帐户中。

例如，如果您从 dot.tk（域名注册服务商）处购买了域名并在 namecheap.com（域名托管服务商）处托管您的域名，则需要将 namecheap.com 的名称服务器（freedns1.registrar-servers.com 和 freedns2.registrar-servers.com）添加到 dot.tk 处的帐户中。

###顶级域名
**顶级域**名是域名的最后一部分，也就是最后一个点号后面的字母。     
下面是一些示例：biz com org edu us ca fr de travel local es pl

###第二级域名	
第二级域名仅次于顶级域名。下面是一些常见的示例：

   第二级域名    域名
    Google        google.cn
    Wikipedia     wikipedia.org
    Ontariotravel ontariotravel.com
    Craigslist    craigslist.com
  
###第三级域名###
第三级域名也称为子域名和 CNAME。在网址中，子域名会写在域名前面。下面是一些示例：

    子域名      网址
    affiliates  http://affiliates.art.com
    www         http://www.rockfound.org
    men         http://men.style.com
    mail        http://mail.google.com
    bus         http://www.bus.umich.edu
  
###域名托管服务商	
域名托管服务商会运行您的域名的 DNS 服务器，其中包括 A 记录、MX 记录和 CNAME 记录等。大多数域名托管服务商还提供域名注册服务。

##A 记录	
A 记录（亦称主机记录）是最重要的 DNS 记录。这些记录会将域名或子域名与某个 IP 地址链接。

A 记录和 IP 地址不必一一对应。多个 A 记录可对应一个 IP 地址，即一台机器可提供多个网站服务。此外，一个 A 记录也可对应多个 IP 地址。这有助于提高容错性能和负载均衡，并允许网站改变其实际托管位置。

##NS 记录	
名称服务器记录可决定哪些服务器将用于传递域名的 DNS 信息。每个域名必须定义两条 NS 记录。通常，您会有主要和辅助名称服务器记录，NS 记录通过域名注册商加以更新，需要 24 至 72 小时才能生效。

如果域名注册商与域名托管服务商相互独立，那么您的托管服务商就会提供两个名称服务器，供您通过注册商更新 NS 记录。

##MX 记录
邮件交换记录 (MX) 会将电子邮件定向到某个域的服务器，并按照优先级进行排列。如果不能使用第一优先级记录发送邮件，则会使用第二优先级记录，依此类推。

##CNAME 记录	
规范名记录是 A 记录的别名。对于每条 CNAME 记录，您可选择别名和主机。

**优点**：您可以把www.hacknote.tk, mail.hacknote.tk, news.hacknote.tk这些子域名的CNAME记录全部指向hacknote.tk,当IP变化时，只需要修改hacknote.tk的A记录就可以了，不要手动修改每一个子域名的IP。

##URL Redirect	
URL重定向，访问您的域名hacknote.tk时，将自动跳转到您的另一个网络地址(URL) http://hi.baidu.com/regdt32 ，此时在浏览器地址栏显示的是http://hi.baidu.com/regdt32 。 URL转发可以转发到某一个目录下，甚至是一个文件下，而CNAME是不可以的，这就是URL Redirect和CNAME的主要区别。这种url转发方式是对seo不利的，对搜索引擎最友好的跳转方式是使用301转向(也叫301跳转，301重定向)。

##URL Redirect（301）	
为什么要使用**301转向**？首先第一条就是刚才我们提到的为了seo优化，在如何url网址规范化中也提到，设置301转向后，原来的页面将返回301 HTTP 状态码给浏览器或者搜索引擎，告知搜索引擎次页面已经永久重定向到了新的网页上，spider就不会索引原网页，同时也会将原网页的PR权重传递给跳转到的新网页。

如果返回的是**302状态码**，则告诉spider页面跳转只是临时的，spider仍然会索引原网页，这就造成了**重复页面问题**(即Google 补充材料)。如www.hacknote.tk 302 跳转到 hacknote.tk，它们就会被当做两个重复页面。

其次，由于各种原因要更换域名时，可以使用301重定向将老域名301转向到新的域名，这样老域名的流量不会流失，PR权重也会转移。

另外当有注册了多个不同域名，指向到同一网站时，也可以只设置一个主域名，将其他域名设置301转向到主域名，这样各个域名就不会分散权重。因此使用301转向很有必要。

##URL Frame	
Frame,框架的意思，即通过一个框架iframe来实现转发。翻译过来就叫**URL隐藏转发**，所有URL Redirect也叫显示转发。还是上面的例子，

URL Frame和URL Redirect的**区别**就在于浏览器显示的地址还是http://hacknote.tk 。

##TXT Record	
TXT记录，一般指为某个主机名或域名设置的说明，如：

    admin in TXT "管理员，电话：14444444444"
    mail in TXT "Email：master@baidu.com"
  
可以方便别人联系到您
