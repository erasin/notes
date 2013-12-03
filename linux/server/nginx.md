#Nginx


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
