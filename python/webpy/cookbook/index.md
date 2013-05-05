
# webpy cookbook

具体一些改变，还是去官网去看吧。

[webpy.org](http://webpy.org)

## 基本应用:

*   [Hello World][1]
*   [提供静态文件访问][2]
*   [理解URL控制][3]
*   [跳转与重定向][4]
*   [使用子应用][5]
*   [提供XML访问][6]
*   [从post读取原始数据][7]

## 高级应用:

*   [用web.ctx获得客户端信息][8]
*   [应用处理器，添加钩子和卸载钩子][9]
*   [如何使用web.background][10]
*   [自定义NotFound信息][11]
*   [如何流传输大文件][12]
*   [对自带的webserver日志进行操作][13]
*   [用cherrypy提供SSL支持][14]
*   [实时语言切换][15]

## Sessions and user state 会话和用户状态:

*   [如何使用Session][16]
*   [如何在调试模式下使用Session][17]
*   [在template中使用session][18]
*   [如何操作Cookie][19]
*   [用户认证][20]
*   [一个在postgreSQL数据库环境下的用户认证的例子][21]
*   [如何在子应用中操作Session][22]

## Utils 实用工具

*   [如何发送邮件][23]
*   [如何利用Gmail发送邮件][24]
*   [使用soaplib实现webservice][25]

## Templates 模板

*   [Templetor: web.py 模板系统][26]
*   [使用站点布局模板][27]
*   [交替式风格 (未译)][28]
*   [导入函数到模板中 (未译)][29]
*   [模板文件中的i18n支持][30]
*   [在web.py中使用Mako模板引擎 ][31]
*   [在web.py中使用Cheetah模板引擎][32]
*   [在web.py中使用Jinja2模板引擎][33]
*   [如何在谷歌应用程序引擎使用模板][34]

## Testing 测试:

*   [Testing with Paste and Nose (未译)][35]
*   [RESTful doctesting using an application's request method (未译)][36]

## User input 用户输入:

*   [文件上传][37]
*   [保存上传的文件][38]
*   [上传文件大小限定][39]
*   [通过 web.input 接受用户输入][40]
*   [怎样使用表单][41]
*   [显示个别表单字段][42]

## Database 数据库

*   [使用多数据库][43]
*   [Select: 查询数据][44]
*   [Update: 更新数据 ][45]
*   [Delete: 删除数据][46]
*   [Insert: 新增数据][47]
*   [Query: 高级数据库查询][48]
*   [怎样使用数据库事务][49]
*   [使用 sqlalchemy][50]
*   [整合 SQLite UDF (用户定义函数) 到 webpy 数据库层][51]
*   [使用字典动态构造where子句][52]

## Deployment 部署:

*   [通过Fastcgi和lighttpd部署][53]
*   [通过Webpy和Nginx with FastCGI搭建Web.py][54]
*   [CGI deployment through Apache (未译)][55]
*   `mod-python` deployment through Apache (requested)
*   [通过Apache和mod-wsgi部署][56]
*   [mod-wsgi deployment through Nginx (未译)][57]
*   [Fastcgi deployment through Nginx (未译)][54]

## Subdomains 子域名:

*   Subdomains and how to access the username (requested)

 [1]: helloworld.zh-cn
 [2]: staticfiles.zh-cn
 [3]: url-handling.zh-cn
 [4]: redirect-seeother.zh-cn
 [5]: subapp.zh-cn
 [6]: xmlfiles.zh-cn
 [7]: postbasic.zh-cn
 [8]: ctx.zh-cn
 [9]: application-processors.zh-cn
 [10]: background.zh-cn
 [11]: custom-notfound.zh-cn
 [12]: streaming-large-files.zh-cn
 [13]: logging.zh-cn
 [14]: ssl.zh-cn
 [15]: runtime-language-switch.zh-cn
 [16]: sessions.zh-cn
 [17]: session-with-reloader.zh-cn
 [18]: session-in-template.zh-cn
 [19]: cookies.zh-cn
 [20]: userauth.zh-cn
 [21]: userauthpgsql.zh-cn
 [22]: sessions-with-subapp.zh-cn
 [23]: sendmail.zh-cn
 [24]: sendmail-using-gmail.zh-cn
 [25]: webservice.zh-cn
 [26]: templetor.zh-cn
 [27]: layout-template.zh-cn
 [28]: alternating-style.zh-cn
 [29]: template-import.zh-cn
 [30]: i18n-support-in-template-file.zh-cn
 [31]: template-mako.zh-cn
 [32]: template-cheetah.zh-cn
 [33]: template-jinja.zh-cn
 [34]: templates-on-gae.zh-cn
 [35]: testing-with-paste-and-nose.zh-cn
 [36]: restful-doctesting-using-request.zh-cn
 [37]: fileupload.zh-cn
 [38]: storeupload.zh-cn
 [39]: limiting-upload-size.zh-cn
 [40]: input.zh-cn
 [41]: forms.zh-cn
 [42]: form-fields.zh-cn
 [43]: multidbs.zh-cn
 [44]: select.zh-cn
 [45]: update.zh-cn
 [46]: delete.zh-cn
 [47]: Insert.zh-cn
 [48]: query.zh-cn
 [49]: transactions.zh-cn
 [50]: sqlalchemy.zh-cn
 [51]: sqlite-udf.zh-cn
 [52]: where-dict.zh-cn
 [53]: fastcgi-lighttpd.zh-cn
 [54]: fastcgi-nginx.zh-cn
 [55]: cgi-apache.zh-cn
 [56]: mod-wsgi-apache.zh-cn
 [57]: mod-wsgi-nginx.zh-cn
