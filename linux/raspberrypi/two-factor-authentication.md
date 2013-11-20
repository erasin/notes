# 为Raspberry PI 的ssh添加 google 两步验证


**1. 下载 pam 部件**

	sudo apt-get install libpam-google-authenticator

**2. 创建 GA**

	google-authenticator

> 一路按 `y`即可。


中文版本：

	google-authenticator -l kevin@kbnix.com -t -d -r 3 -R 30 -w 3

参数

* -l 设置移动设备上的程序对应的 OTP 码的标签，便于区分不同的应用。
* -t 设置采用基于时间的验证。相应的也可以 -c 参数采用基于计数的验证。
* -d 禁止重用基于时间的验证码
* -r 限制登录频率
* -R 设置登录频率限制的时间间隔
* -w 设置时间窗的大小，主要在移动设备的时间不是准确同步的情况下比较有用


开始创建一个GA，他会自动生成一个QR供你的GA客户端读取，QR地址也会提供。并会输出 `~/.google_authenticator`文件内容为默认密钥和一些`紧急验证号`，这一些验证号在你没有GA客户端的情况下可以使用，但是每个号码只能使用一次。

> 最好将提供的QR的URL `https://www.google.com/chart?chs=200×200&chld=M|0&cht=qr&chl=otpauth://totp/名称%3Fsecret%3D...`(将`...`改为自己的密钥) 地址也一同保存到 `.google_authenticator`文件中去。

**3. 在你的手机上添加帐号**

Android ,黑莓手机，诺基亚都有对应的app。

在APP搜索`Google’s two-factor authentication`一类的安装把，然后扫描上面第二部中提供的QR即可。

**4. 设置相关的文件，配置GA到PAM中和SSH中**

打开 `/etc/pam.d/sshd` 文件，并且添加下面的句子

	auth required pam_google_authenticator.so

这 `.so` 在 `/lib/security/`中。

修改文件 `/etc/ssh/sshd_config` ,找到其中一项 `ChallengeResponseAuthentication` 相应验证将其值由默认的`no`修改为`yes`

	ChallengeResponseAuthentication yes

重启 sshd

	sudo /etc/init.d/sshd restart

> 测试的时候多开即可链接，避免错误问题

**5. 测试**

	ssh pi@pi 
	password: 


## 其他注意事项

1. Google Authenticator对SSH验证和使用商用的OTP系统有如下区别：
	商用OTP系统一般是C/S网络版方式，有一个统一的AuthenticationServer，为了保证高可用一般会有一主一备两台服务器。Google Authenticator是单机版的验证方式。
2. Google Authenticator是一个基于时间的产生验证码的程序，因此不管是服务器端还是手机客户端，对时间的要求都是非常严格。要时刻保证与NTP服务器同步。
3. Google Authenticator和条形码扫描器默认是不会产生任何GPRS和WIFI流量的。
4. Google Authenticator是一种验证的方式，可以扩展到其他的地方，例如wordpress的管理员登陆等等。我这里可以提供一个PHP验证Google Authenticator算法的链接：
	http://www.idontplaydarts.com/2011/07/google-totp-two-factor-authentication-for-php/
5. 如果不需要用户登录时输入OTP密码，而是在用户su到root时要求输入，可以把pam认证语句加入到/etc/pam.d/su当中。
6. 当服务器启用了pam认证之后，所有的用户都是要求输入TOTP密码，所以需要每个用户都在自己的目录下生成一个.google_authenticator文件。

## 参考 

* [Enable two-factor authentication for the SSH on your Raspberry PI](http://msorin.wordpress.com/2013/04/09/enable-two-factor-authentication-for-the-ssh-on-your-raspberry-pi/)
* [使用Google Authenticator对SSH进行验证](http://www.ipcpu.com/2012/07/google-auth-sshd/)
* [一步一步实现linux双因素两步认证](http://www.kbnix.com/2012/11/14/step_by_step_linux_two_factor)
