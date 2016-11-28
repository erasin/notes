## CentOS 6 use  centos-release-SCL

First install the SCL repo:

```
# yum install centos-release-SCL
```

Then install PHP 5.4 and these modules:

```
# yum install php54 php54-php php54-php-gd php54-php-mbstring
```

You must also install the updated database module. This installs the new PHP 5.4 module for MySQL/MariaDB:

```
# yum install php54-php-mysqlnd
```

Disable loading the old PHP 5.3 Apache module:

```
# mv /etc/httpd/conf.d/php.conf /etc/httpd/conf.d/php53.off
```

You should now have a /etc/httpd/conf.d/php54-php.conf file, which loads the correct PHP 5.4 module for Apache.

Finally, restart Apache:

```
# service httpd restart
```



# centos6 use wetstatic

This article describes how to upgrade to PHP 5.4 or PHP 5.5 on a CentOS 5 or CentOS 6 server.
First, detect if any PHP packages are installed:
```
# yum list installed | grep php
```
If packages are installed remove them, for example:
```
# yum remove php.x86_64 php-cli.x86_64 php-common.x86_64 php-gd.x86_64 php-ldap.x86_64 php-mbstring.x86_64 php-mcrypt.x86_64 php-mysql.x86_64 php-pdo.x86_64
```

Add PHP 5.4 packages to yum using this command for CentOS 5.x
```
# rpm -Uvh http://mirror.webtatic.com/yum/el5/latest.rpm
```
Or, for CentOS 6.x:
```
# rpm -Uvh http://mirror.webtatic.com/yum/el6/latest.rpm
```
Now, you can check if the new PHP (5.4: php54w or 5.5: php55w) packages are available:
```
# yum list available | grep php
```
Or, version specific search:
```
# yum list available | grep php54
```
Next, install the new PHP 5.4 or 5.5 packages, for example when installing PHP 5.4 packages I used:
```
# yum install php54w.x86_64 php54w-cli.x86_64 php54w-common.x86_64 php54w-gd.x86_64 php54w-ldap.x86_64 php54w-mbstring.x86_64 php54w-mcrypt.x86_64 php54w-mysql.x86_64 php54w-pdo.x86_64
```
PHP should now be upgraded to the new version, you can verify with the command:
```
# php -v

PHP 5.4.17 (cli) (built: Jul 23 2013 00:02:04)
Copyright (c) 1997-2013 The PHP Group
Zend Engine v2.4.0, Copyright (c) 1998-2013 Zend Technologies
```
Finally, restart the Web server:
```
# service httpd restart
```



> [参考1](https://doc.owncloud.org/server/8.1/admin_manual/installation/php_54_installation.html)
> [参考2](http://www.shayanderson.com/linux/centos-5-or-centos-6-upgrade-php-to-php-54-or-php-55.htm)
